package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env"
	om "github.com/grokify/oauth2more"
	"github.com/grokify/simplego/config"
	"github.com/grokify/simplego/fmt/fmtutil"
	"github.com/jessevdk/go-flags"

	rc "github.com/grokify/go-ringcentral-client/office/v1/client"
	ru "github.com/grokify/go-ringcentral-client/office/v1/util"
)

type Options struct {
	FirstName string `short:"f" long:"firstname" description:"First Name" required:"true"`
	LastName  string `short:"l" long:"lastname" description:"Last Name" required:"true"`
}

type RingCentralConfig struct {
	TokenJSON string `env:"RINGCENTRAL_TOKEN_JSON"`
	ServerURL string `env:"RINGCENTRAL_SERVER_URL"`
}

func Initialize() (*rc.APIClient, error) {
	appCfg := RingCentralConfig{}
	err := env.Parse(&appCfg)
	if err != nil {
		return nil, err
	}

	httpClient, err := om.NewClientTokenJSON(
		context.Background(), []byte(appCfg.TokenJSON))
	if err != nil {
		return nil, err
	}

	apiClient, err := ru.NewApiClientHttpClientBaseURL(
		httpClient, appCfg.ServerURL)
	return apiClient, err
}

// This code takes a bot token and creates a permanent webhook.
func main() {
	opts := &Options{}
	_, err := flags.Parse(opts)
	if err != nil {
		log.Fatal(err)
	}

	err = config.LoadDotEnvSkipEmpty(os.Getenv("ENV_PATH"), "./.env")
	if err != nil {
		log.Fatal(err)
	}

	apiClient, err := Initialize()
	if err != nil {
		log.Fatal(err)
	}

	req := rc.ExtensionUpdateRequest{
		Contact: rc.ContactInfoUpdateRequest{
			FirstName: opts.FirstName,
			LastName:  opts.LastName}}

	info, resp, err := apiClient.UserSettingsApi.UpdateExtension(
		context.Background(), "~", "~", req)
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode >= 300 {
		log.Fatal(fmt.Sprintf("API Status %v", resp.StatusCode))
	}

	fmt.Printf("Status: %v\n", resp.StatusCode)

	fmtutil.PrintJSON(info)

	fmt.Println("DONE")
}
