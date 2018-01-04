package main

import (
	"context"
	"fmt"
	"os"

	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/joho/godotenv"

	ro "github.com/grokify/oauth2more/ringcentral"
	rc "github.com/grokify/ringcentral-sdk-go-scg/client"
)

func loadEnv() error {
	envPaths := []string{}
	if len(os.Getenv("ENV_PATH")) > 0 {
		envPaths = append(envPaths, os.Getenv("ENV_PATH"))
	}
	return godotenv.Load(envPaths...)
}

func newApiClient() (*rc.APIClient, error) {
	httpClient, err := ro.NewClientPassword(
		ro.ApplicationCredentials{
			ServerURL:    os.Getenv("RINGCENTRAL_SERVER_URL"),
			ClientID:     os.Getenv("RINGCENTRAL_CLIENT_ID"),
			ClientSecret: os.Getenv("RINGCENTRAL_CLIENT_SECRET"),
		},
		ro.UserCredentials{
			Username:  os.Getenv("RINGCENTRAL_USERNAME"),
			Extension: os.Getenv("RINGCENTRAL_EXTENSION"),
			Password:  os.Getenv("RINGCENTRAL_PASSWORD"),
		},
	)
	if err != nil {
		return nil, err
	}

	apiConfig := rc.NewConfiguration()
	apiConfig.HTTPClient = httpClient
	apiClient := rc.NewAPIClient(apiConfig)
	return apiClient, nil
}

func main() {
	err := loadEnv()
	if err != nil {
		panic(err)
	}

	apiClient, err := newApiClient()
	if err != nil {
		panic(err)
	}

	info, resp, err := apiClient.UserSettingsApi.LoadExtensionInfo(context.Background(), "~", "~")
	if err != nil {
		panic(err)
	} else if resp.StatusCode > 299 {
		panic(fmt.Errorf("API Status %v", resp.StatusCode))
	}
	fmtutil.PrintJSON(info)
	fmt.Println("DONE")
}
