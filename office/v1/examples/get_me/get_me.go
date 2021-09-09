package main

import (
	"context"
	"fmt"
	"os"

	"github.com/grokify/oauth2more/credentials"
	"github.com/grokify/simplego/config"
	"github.com/grokify/simplego/fmt/fmtutil"

	ru "github.com/grokify/go-ringcentral-client/office/v1/util"
)

func main() {
	err := config.LoadDotEnvSkipEmpty(os.Getenv("ENV_PATH"), "./.env")
	if err != nil {
		panic(err)
	}

	apiClient, err := ru.NewApiClientPassword(
		credentials.ApplicationCredentials{
			ServerURL:    os.Getenv("RINGCENTRAL_SERVER_URL"),
			ClientID:     os.Getenv("RINGCENTRAL_CLIENT_ID"),
			ClientSecret: os.Getenv("RINGCENTRAL_CLIENT_SECRET")},
		credentials.PasswordCredentials{
			Username: os.Getenv("RINGCENTRAL_USERNAME"),
			Password: os.Getenv("RINGCENTRAL_PASSWORD")})
	if err != nil {
		panic(err)
	}

	info, resp, err := apiClient.UserSettingsApi.LoadExtensionInfo(context.Background(), "~", "~")
	if err != nil {
		panic(err)
	} else if resp.StatusCode >= 300 {
		panic(fmt.Errorf("API Status %v", resp.StatusCode))
	}
	fmtutil.PrintJSON(info)

	fmt.Println("DONE")
}
