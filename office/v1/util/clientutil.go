package clientutil

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/grokify/goauth/credentials"
	"github.com/grokify/mogo/net/urlutil"
	"github.com/joho/godotenv"

	rc "github.com/grokify/go-ringcentral-client/office/v1/client"
	rs "github.com/grokify/go-scim-client"
	ro "github.com/grokify/goauth/ringcentral"
)

func LoadEnv() error {
	envPaths := []string{}
	if len(os.Getenv("ENV_PATH")) > 0 {
		envPaths = append(envPaths, os.Getenv("ENV_PATH"))
	}
	return godotenv.Load(envPaths...)
}

func NewApiClientHttpClientBaseURL(httpClient *http.Client, serverUrl string) (*rc.APIClient, error) {
	if len(strings.TrimSpace(serverUrl)) == 0 {
		return nil, fmt.Errorf("No RingCentral API Server URL provided")
	}
	apiConfig := rc.NewConfiguration()
	apiConfig.BasePath = strings.TrimSpace(serverUrl)
	apiConfig.HTTPClient = httpClient
	apiClient := rc.NewAPIClient(apiConfig)
	return apiClient, nil
}

func NewApiClientCredentials(creds credentials.Credentials) (*rc.APIClient, error) {
	httpClient, err := creds.NewClient(context.Background())
	if err != nil {
		return nil, err
	}
	return NewApiClientHttpClientBaseURL(httpClient, creds.OAuth2.ServerURL)
}

func NewApiClientPassword(oc credentials.OAuth2Credentials) (*rc.APIClient, error) {
	httpClient, err := ro.NewClientPassword(oc)
	if err != nil {
		return nil, err
	}
	return NewApiClientHttpClientBaseURL(httpClient, oc.ServerURL)
}

func NewApiClientPasswordSimple(oc credentials.OAuth2Credentials) (*rc.APIClient, error) {
	httpClient, err := ro.NewClientPasswordSimple(oc)
	if err != nil {
		return nil, err
	}
	return NewApiClientHttpClientBaseURL(httpClient, oc.ServerURL)
}

func NewApiClientPasswordEnv(envPrefix string) (*rc.APIClient, error) {
	return NewApiClientPassword(credentials.NewOAuth2CredentialsEnv(envPrefix))
}

func NewScimApiClient(oc credentials.OAuth2Credentials) (*rs.APIClient, error) {
	httpClient, err := ro.NewClientPassword(oc)
	if err != nil {
		return nil, err
	}
	apiConfig := rs.NewConfiguration()
	apiConfig.BasePath = urlutil.JoinAbsolute(oc.ServerURL, "/scim/v2")
	apiConfig.HTTPClient = httpClient
	apiClient := rs.NewAPIClient(apiConfig)
	return apiClient, nil
}

func ApiResponseErrorBody(err error) []byte {
	openAPIErr := err.(rc.GenericOpenAPIError)
	return openAPIErr.Body()
}
