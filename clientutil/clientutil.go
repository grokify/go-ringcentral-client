package clientutil

import (
	"os"

	"github.com/grokify/gotilla/net/urlutil"
	"github.com/joho/godotenv"

	rc "github.com/grokify/go-ringcentral/client"
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

func NewApiClient(app ro.ApplicationCredentials, user ro.UserCredentials) (*rc.APIClient, error) {
	httpClient, err := ro.NewClientPassword(app, user)
	if err != nil {
		return nil, err
	}
	apiConfig := rc.NewConfiguration()
	apiConfig.BasePath = app.ServerURL
	apiConfig.HTTPClient = httpClient
	apiClient := rc.NewAPIClient(apiConfig)
	return apiClient, nil
}

func NewApiClientEnv() (*rc.APIClient, error) {
	return NewApiClient(
		ro.NewApplicationCredentialsEnv(),
		ro.NewUserCredentialsEnv())
}

func NewScimApiClient(app ro.ApplicationCredentials, user ro.UserCredentials) (*rs.APIClient, error) {
	httpClient, err := ro.NewClientPassword(app, user)
	if err != nil {
		return nil, err
	}
	apiConfig := rs.NewConfiguration()
	apiConfig.BasePath = urlutil.JoinAbsolute(app.ServerURL, "/scim/v2")
	apiConfig.HTTPClient = httpClient
	apiClient := rs.NewAPIClient(apiConfig)
	return apiClient, nil
}
