package main

import (
	"context"
	"fmt"
	"os"

	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"

	rc "github.com/grokify/go-ringcentral/office/client"
	ru "github.com/grokify/go-ringcentral/office/clientutil"
	ro "github.com/grokify/oauth2more/ringcentral"
)

func main() {
	err := config.LoadDotEnvSkipEmpty(os.Getenv("ENV_PATH"), "./.env")
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
	info, resp, err := apiClient.GlipApi.LoadGroupList(
		context.Background(), &rc.LoadGroupListOpts{})

	if err != nil {
		panic(err)
	} else if resp.StatusCode >= 300 {
		panic(fmt.Errorf("API Status %v", resp.StatusCode))
	}
	fmtutil.PrintJSON(info)

	fmt.Println("DONE")
}
