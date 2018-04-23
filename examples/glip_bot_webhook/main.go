package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"
	om "github.com/grokify/oauth2more"

	rc "github.com/grokify/go-ringcentral/client"
	ru "github.com/grokify/go-ringcentral/clientutil"
)

// This code takes a bot token and creates a permanent webhook.
func main() {
	err := config.LoadDotEnvSkipEmpty(os.Getenv("ENV_PATH"), "./.env")
	if err != nil {
		log.Fatal(err)
	}

	httpClient, err := om.NewClientTokenJSON(
		context.Background(), []byte(os.Getenv("RINGCENTRAL_TOKEN_JSON")))
	if err != nil {
		log.Fatal(err)
	}

	apiClient, err := ru.NewApiClientHttpClientBaseURL(
		httpClient, os.Getenv("RINGCENTRAL_SERVER_URL"))
	if err != nil {
		log.Fatal(err)
	}

	req := rc.CreateSubscriptionRequest{
		EventFilters: []string{"/restapi/v1.0/glip/posts"},
		DeliveryMode: &rc.NotificationDeliveryModeRequest{
			TransportType: "WebHook",
			Address:       os.Getenv("RINGCENTRAL_WEBHOOK_URL")},
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
