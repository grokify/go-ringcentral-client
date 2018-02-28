package main

import (
	"context"
	"fmt"
	"os"

	"github.com/grokify/gotilla/fmt/fmtutil"
	iom "github.com/grokify/gotilla/io/ioutilmore"
	"github.com/joho/godotenv"

	ru "github.com/grokify/go-ringcentral/clientutil"
	ro "github.com/grokify/oauth2more/ringcentral"
)

var LocalEnvFile = "./.env"

func loadEnv() error {
	envPaths := []string{}
	if len(os.Getenv("ENV_PATH")) > 0 {
		envPaths = append(envPaths, os.Getenv("ENV_PATH"))
	} else {
		isGood, err := iom.IsFileWithSizeGtZero(LocalEnvFile)
		if err == nil && isGood {
			envPaths = append(envPaths, LocalEnvFile)
		}
	}
	return godotenv.Load(envPaths...)
}

func main() {
	err := loadEnv()
	if err != nil {
		panic(err)
	}

	apiClient, err := ru.NewApiClientPassword(
		ro.ApplicationCredentials{
			ServerURL:    os.Getenv("RINGCENTRAL_SERVER_URL"),
			ClientID:     os.Getenv("RINGCENTRAL_CLIENT_ID"),
			ClientSecret: os.Getenv("RINGCENTRAL_CLIENT_SECRET")},
		ro.PasswordCredentials{
			Username:  os.Getenv("RINGCENTRAL_USERNAME"),
			Extension: os.Getenv("RINGCENTRAL_EXTENSION"),
			Password:  os.Getenv("RINGCENTRAL_PASSWORD")})
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
