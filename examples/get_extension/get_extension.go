package main

import (
	"fmt"
	"os"

	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"
	rcOAuth "github.com/grokify/oauth2util-go/services/ringcentral"
	"github.com/grokify/ringcentral-sdk-go-scg"
	"golang.org/x/oauth2"
)

func main() {
	err := config.LoadDotEnv()
	if err != nil {
		panic(err)
	}

	client, err := rcOAuth.NewClientPasswordConf(
		oauth2.Config{
			ClientID:     os.Getenv("RC_CLIENT_ID"),
			ClientSecret: os.Getenv("RC_CLIENT_SECRET"),
			Endpoint:     rcOAuth.NewEndpoint(os.Getenv("RC_SERVER_URL"))},
		os.Getenv("RC_USER_USERNAME"),
		os.Getenv("RC_USER_PASSWORD"))

	if err != nil {
		panic(err)
	}

	extApi := ringcentral.NewExtensionApiWithBasePath(os.Getenv("RC_SERVER_URL"))
	extApi.Configuration.Transport = client.Transport

	extInfo, resp, err := extApi.RestapiV10AccountAccountIdExtensionExtensionIdGet("~", "~")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", resp.StatusCode)
	fmtutil.PrintJSON(extInfo)

	fmt.Println("DONE")
}
