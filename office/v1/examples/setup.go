package examples

import (
	"os"

	"github.com/joho/godotenv"

	rc "github.com/grokify/go-ringcentral-client/office/v1/client"
	"github.com/grokify/goauth"
	ro "github.com/grokify/goauth/ringcentral"
)

func LoadEnv() error {
	envPaths := []string{}
	if len(os.Getenv("ENV_PATH")) > 0 {
		envPaths = append(envPaths, os.Getenv("ENV_PATH"))
	}
	return godotenv.Load(envPaths...)
}

func NewApiClient() (*rc.APIClient, error) {
	httpClient, err := ro.NewClientPassword(
		goauth.NewCredentialsOAuth2Env("RINGCENTRAL_"))
	if err != nil {
		return nil, err
	}

	apiConfig := rc.NewConfiguration()
	apiConfig.BasePath = os.Getenv("RINGCENTRAL_SERVER_URL")
	apiConfig.HTTPClient = httpClient
	apiClient := rc.NewAPIClient(apiConfig)
	return apiClient, nil
}
