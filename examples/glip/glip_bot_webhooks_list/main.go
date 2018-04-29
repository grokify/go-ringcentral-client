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
	"github.com/jessevdk/go-flags"

	rc "github.com/grokify/go-ringcentral/client"
	ru "github.com/grokify/go-ringcentral/clientutil"
)

type Options struct {
	List     []bool `short:"l" long:"list" description:"List subscriptions"`
	Delete   string `short:"d" long:"delete" description:"Delete subscriptions"`
	Recreate string `short:"r" long:"recreate" description:"Recreate subscription"`
}

type RingCentralConfig struct {
	TokenJSON  string `env:"RINGCENTRAL_TOKEN_JSON"`
	ServerURL  string `env:"RINGCENTRAL_SERVER_URL"`
	WebhookURL string `env:"RINGCENTRAL_WEBHOOK_URL"`
}

func ReplaceSubscription(apiClient *rc.APIClient, subscription rc.SubscriptionResponse) error {
	if subscription.Status == "Blacklisted" {
		fmt.Printf("Blacklisted %v %v\n",
			subscription.Id,
			subscription.DeliveryMode.Address)

		resp, err := apiClient.PushNotificationsApi.DeleteSubscription(
			context.Background(), subscription.Id)
		if err != nil {
			return err
		} else if resp.StatusCode >= 300 {
			return fmt.Errorf("API Status %v", resp.StatusCode)
		}

		req := rc.CreateSubscriptionRequest{
			EventFilters: []string{"/restapi/v1.0/glip/posts"},
			DeliveryMode: &rc.NotificationDeliveryModeRequest{
				TransportType: "WebHook",
				Address:       subscription.DeliveryMode.Address},
			ExpiresIn: int32(500000000)}

		info, resp, err := apiClient.PushNotificationsApi.CreateSubscription(
			context.Background(), req)

		if err != nil {
			return err
		} else if resp.StatusCode >= 300 {
			return fmt.Errorf("API Status %v", resp.StatusCode)
		}
		fmtutil.PrintJSON(info)
	}
	return nil
}

// This code takes a bot token and creates a permanent webhook.
func main() {
	opts := Options{}

	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	err = config.LoadDotEnvSkipEmpty(os.Getenv("ENV_PATH"), "./.env")
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

	if len(opts.List) > 0 || len(opts.Recreate) > 0 {
		info, resp, err := apiClient.PushNotificationsApi.GetSubscriptions(
			context.Background())

		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode >= 300 {
			log.Fatal(fmt.Errorf("API Status %v", resp.StatusCode))
		}
		fmtutil.PrintJSON(info)

		fmt.Println(opts.Recreate)

		if len(opts.Recreate) > 0 {
			for _, subscription := range info.Records {
				if subscription.Id == opts.Recreate {
					err := ReplaceSubscription(apiClient, subscription)
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	}

	fmt.Println("DONE")
}
