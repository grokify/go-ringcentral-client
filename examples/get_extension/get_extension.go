package main

import (
	"fmt"
	"os"

	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"
	rcOAuth "github.com/grokify/oauth2util-go/services/ringcentral"
	"github.com/grokify/ringcentral-sdk-go-scg"
)

func main() {
	err := config.LoadDotEnv()
	if err != nil {
		panic(err)
	}

	client, err := rcOAuth.NewClientPassword(
		rcOAuth.ApplicationCredentials{
			ClientID:     os.Getenv("RC_CLIENT_ID"),
			ClientSecret: os.Getenv("RC_CLIENT_SECRET"),
			ServerURL:    os.Getenv("RC_SERVER_URL")},
		rcOAuth.UserCredentials{
			Username:  os.Getenv("RC_USER_USERNAME"),
			Extension: os.Getenv("RC_USER_EXTENSION"),
			Password:  os.Getenv("RC_USER_PASSWORD")})
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
