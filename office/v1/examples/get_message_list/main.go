package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/antihax/optional"
	"github.com/grokify/goauth"
	"github.com/grokify/goauth/authutil"
	"github.com/grokify/mogo/config"
	"github.com/grokify/mogo/fmt/fmtutil"

	rc "github.com/grokify/go-ringcentral-client/office/v1/client"
	ru "github.com/grokify/go-ringcentral-client/office/v1/util"
)

func main() {
	_, err := config.LoadDotEnv([]string{os.Getenv("ENV_PATH"), "./.env"}, 1)
	if err != nil {
		panic(err)
	}

	apiClient, err := ru.NewApiClientPassword(
		goauth.CredentialsOAuth2{
			ServerURL:    os.Getenv("RINGCENTRAL_SERVER_URL"),
			ClientID:     os.Getenv("RINGCENTRAL_CLIENT_ID"),
			ClientSecret: os.Getenv("RINGCENTRAL_CLIENT_SECRET"),
			GrantType:    authutil.GrantTypePassword,
			Username:     os.Getenv("RINGCENTRAL_USERNAME"),
			Password:     os.Getenv("RINGCENTRAL_PASSWORD")})
	if err != nil {
		panic(err)
	}

	dt, err := time.Parse(time.RFC3339, "2018-01-01T00:00:00Z")
	if err != nil {
		log.Fatal(err)
	}

	opts := &rc.ListMessagesOpts{
		DateFrom: optional.NewTime(dt)}

	info, resp, err := apiClient.MessagesApi.ListMessages(
		context.Background(), "~", "~", opts)
	if err != nil {
		panic(err)
	} else if resp.StatusCode >= 300 {
		panic(fmt.Errorf("API Status %v", resp.StatusCode))
	}
	fmtutil.PrintJSON(info)

	fmt.Println("DONE")
}
