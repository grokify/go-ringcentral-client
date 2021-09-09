package clientutil

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/grokify/oauth2more/credentials"
	"github.com/grokify/simplego/net/urlutil"
	"github.com/joho/godotenv"

	rc "github.com/grokify/go-ringcentral-client/office/v1/client"
	rs "github.com/grokify/go-scim-client"
	ro "github.com/grokify/oauth2more/ringcentral"
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
	httpClient, err := creds.NewClient()
	if err != nil {
		return nil, err
	}
	return NewApiClientHttpClientBaseURL(httpClient, creds.Application.ServerURL)
}

func NewApiClientPassword(app credentials.ApplicationCredentials, pwd credentials.PasswordCredentials) (*rc.APIClient, error) {
	httpClient, err := ro.NewClientPassword(app, pwd)
	if err != nil {
		return nil, err
	}
	return NewApiClientHttpClientBaseURL(httpClient, app.ServerURL)
}

func NewApiClientPasswordSimple(app credentials.ApplicationCredentials, user credentials.PasswordCredentials) (*rc.APIClient, error) {
	httpClient, err := ro.NewClientPasswordSimple(app, user)
	if err != nil {
		return nil, err
	}
	return NewApiClientHttpClientBaseURL(httpClient, app.ServerURL)
}

func NewApiClientPasswordEnv() (*rc.APIClient, error) {
	return NewApiClientPassword(
		ro.NewApplicationCredentialsEnv(),
		ro.NewPasswordCredentialsEnv())
}

func NewScimApiClient(app credentials.ApplicationCredentials, pwd credentials.PasswordCredentials) (*rs.APIClient, error) {
	httpClient, err := ro.NewClientPassword(app, pwd)
	if err != nil {
		return nil, err
	}
	apiConfig := rs.NewConfiguration()
	apiConfig.BasePath = urlutil.JoinAbsolute(app.ServerURL, "/scim/v2")
	apiConfig.HTTPClient = httpClient
	apiClient := rs.NewAPIClient(apiConfig)
	return apiClient, nil
}

func ApiResponseErrorBody(err error) []byte {
	openAPIErr := err.(rc.GenericOpenAPIError)
	return openAPIErr.Body()
}
