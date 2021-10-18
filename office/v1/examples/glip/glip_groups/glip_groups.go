package main

import (
	"context"
	"fmt"
	"os"

	"github.com/grokify/goauth"
	"github.com/grokify/goauth/credentials"
	"github.com/grokify/simplego/config"
	"github.com/grokify/simplego/fmt/fmtutil"

	rc "github.com/grokify/go-ringcentral-client/office/v1/client"
	ru "github.com/grokify/go-ringcentral-client/office/v1/util"
)

func main() {
	err := config.LoadDotEnvSkipEmpty(os.Getenv("ENV_PATH"), "./.env")
	if err != nil {
		panic(err)
	}

	apiClient, err := ru.NewApiClientPassword(
		credentials.OAuth2Credentials{
			ServerURL:    os.Getenv("RINGCENTRAL_SERVER_URL"),
			ClientID:     os.Getenv("RINGCENTRAL_CLIENT_ID"),
			ClientSecret: os.Getenv("RINGCENTRAL_CLIENT_SECRET"),
			GrantType:    goauth.GrantTypePassword,
			Username:     os.Getenv("RINGCENTRAL_USERNAME"),
			Password:     os.Getenv("RINGCENTRAL_PASSWORD")})
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
