package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env"
	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"
	om "github.com/grokify/oauth2more"

	rc "github.com/grokify/go-ringcentral/client"
	ru "github.com/grokify/go-ringcentral/clientutil"
)

type RingCentralConfig struct {
	TokenJSON  string `env:"RINGCENTRAL_TOKEN_JSON"`
	ServerURL  string `env:"RINGCENTRAL_SERVER_URL"`
	WebhookURL string `env:"RINGCENTRAL_WEBHOOK_URL"`
}

// This code takes a bot token and creates a permanent webhook.
func main() {
	err := config.LoadDotEnvSkipEmpty(os.Getenv("ENV_PATH"), "./.env")
	if err != nil {
		log.Fatal(err)
	}

	appCfg := RingCentralConfig{}
	err = env.Parse(&appCfg)
	if err != nil {
		log.Fatal(err)
	}

	httpClient, err := om.NewClientTokenJSON(
		context.Background(), []byte(appCfg.TokenJSON))
	if err != nil {
		log.Fatal(err)
	}

	apiClient, err := ru.NewApiClientHttpClientBaseURL(
		httpClient, appCfg.ServerURL)
	if err != nil {
		log.Fatal(err)
	}

	req := rc.CreateSubscriptionRequest{
		EventFilters: []string{"/restapi/v1.0/glip/posts"},
		DeliveryMode: &rc.NotificationDeliveryModeRequest{
			TransportType: "WebHook",
			Address:       appCfg.WebhookURL},
		ExpiresIn: int32(500000000)}

	info, resp, err := apiClient.PushNotificationsApi.CreateSubscription(
		context.Background(), req)

	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode >= 300 {
		log.Fatal(fmt.Errorf("API Status %v", resp.StatusCode))
	}
	fmtutil.PrintJSON(info)

	fmt.Println("DONE")
}
