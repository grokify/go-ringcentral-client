package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/grokify/goauth/authutil"
	"github.com/grokify/mogo/config"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/jessevdk/go-flags"

	rc "github.com/grokify/go-ringcentral-client/office/v1/client"
	ru "github.com/grokify/go-ringcentral-client/office/v1/util"
)

type Options struct {
	SubscriptionId string `short:"s" long:"subscriptionId" description:"SubscriptionId to delete" required:"true"`
}

type RingCentralConfig struct {
	TokenJSON  string `env:"RINGCENTRAL_TOKEN_JSON"`
	ServerURL  string `env:"RINGCENTRAL_SERVER_URL"`
	WebhookURL string `env:"RINGCENTRAL_WEBHOOK_URL"`
}

// This code takes a bot token and creates a permanent webhook.
func main() {
	opts := &Options{}
	_, err := flags.Parse(opts)
	if err != nil {
		panic(err)
	}

	_, err = config.LoadDotEnv([]string{os.Getenv("ENV_PATH"), "./.env"}, 1)
	if err != nil {
		log.Fatal(err)
	}

	appCfg := RingCentralConfig{}
	err = env.Parse(&appCfg)
	if err != nil {
		log.Fatal(err)
	}

	httpClient, err := authutil.NewClientTokenJSON(
		context.Background(), []byte(appCfg.TokenJSON))
	if err != nil {
		log.Fatal(err)
	}

	apiClient, err := ru.NewApiClientHttpClientBaseURL(
		httpClient, appCfg.ServerURL)
	if err != nil {
		log.Fatal(err)
	}

	//msi := map[string]interface{}{}
	groups, resp, err := apiClient.GlipApi.LoadGroupList(context.Background(),
		&rc.LoadGroupListOpts{})
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode >= 300 {
		log.Fatal(fmt.Sprintf("API Status %v", resp.StatusCode))
	}

	fmt.Printf("Status: %v\n", resp.StatusCode)

	fmtutil.PrintJSON(groups)

	fmt.Println("DONE")
}
