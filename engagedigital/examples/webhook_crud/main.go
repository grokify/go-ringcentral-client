package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/antihax/optional"
	"github.com/jessevdk/go-flags"

	"github.com/grokify/go-ringcentral-engage/engagedigital/engagedigital"
	utils "github.com/grokify/go-ringcentral-engage/engagedigital/engagedigitalutil"
	ex "github.com/grokify/go-ringcentral-engage/engagedigital/examples"
)

type options struct {
	Site   string `short:"s" long:"site" description:"A site" required:"true"`
	Token  string `short:"t" long:"token" description:"A token" required:"true"`
	Object string `short:"o" long:"object" description:"A object" required:"false"`
	Action string `short:"a" long:"action" description:"An action (create|update|delete)" required:"true"`
	Id     string `short:"i" long:"id" description:"An object id" required:"false"`
	Name   string `short:"n" long:"name" description:"A tag name" required:"false"`
	URL    string `short:"u" long:"url" description:"A webhook URL" required:"false"`
}

func main() {
	opts := options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	client := utils.NewApiClient(opts.Site, opts.Token)

	opts.Object = "webhook"

	switch opts.Object {
	case "webhook":
		handleWebhooks(client, opts)
	default:
		log.Fatal(fmt.Sprintf("E_OBJECT_NOT_SUPPORTED [%v]", opts.Object))
	}

	fmt.Println("DONE")
}

func formatRespStatusCodeError(statusCode int) string {
	return fmt.Sprintf("E_API_ERROR [%v]", statusCode)
}

func handleWebhooks(client *engagedigital.APIClient, opts options) {
	switch opts.Action {
	case "create":
		fmt.Println("I_CREATING_CUSTOM_FIELD")
		opts.Name = strings.TrimSpace(opts.Name)
		if len(opts.Name) == 0 {
			log.Fatal("E_CREATE_CUSTOM_FIELD_NO_NAME")
		}
		opts.URL = strings.TrimSpace(opts.URL)
		if len(opts.URL) == 0 {
			log.Fatal("E_CREATE_CUSTOM_FIELD_NO_URL")
		}
		apiOpts := &engagedigital.CreateWebhookOpts{
			Active:      optional.NewBool(true),
			StagingUse:  optional.NewBool(false),
			VerifyToken: optional.NewString(opts.Token),
			Secret:      optional.NewString("ABC")}
		ex.HandleApiResponse(client.WebhooksApi.CreateWebhook(
			context.Background(), opts.Token, opts.Name, opts.URL, []string{
				"content.imported",
				"intervention.assigned",
				"intervention.canceled",
				"intervention.closed",
				"intervention.custom_fields_updated",
				"intervention.deferred",
				"intervention.opened",
				"intervention.reactivated",
				"intervention.recategorized",
				"intervention.reopened",
				"intervention.user_updated",
				"push_agent.availability_change"}, apiOpts))
	case "read":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.WebhooksApi.GetWebhook(context.Background(), opts.Id, opts.Token))
		} else {
			ex.HandleApiResponse(client.WebhooksApi.GetAllWebhooks(context.Background(), opts.Token, nil))
		}
	case "update":
		opts.Id = strings.TrimSpace(opts.Id)
		if len(opts.Id) == 0 {
			log.Fatal("E_UPDATE_CUSTOM_FIELD_NO_ID")
		}
		apiOpts := &engagedigital.UpdateWebhookOpts{
			RegisteredEvents: optional.NewInterface(
				[]string{
					"content.imported",
					"intervention.assigned"})}
		ex.HandleApiResponse(client.WebhooksApi.UpdateWebhook(
			context.Background(), opts.Id, opts.Token, apiOpts))
	case "delete":
		opts.Id = strings.TrimSpace(opts.Id)
		if len(opts.Id) == 0 {
			log.Fatal("E_DELETE_CUSTOM_FIELD_NO_ID")
		}
		ex.HandleApiResponse(client.WebhooksApi.DeleteWebhook(
			context.Background(), opts.Id, opts.Token))
	}
}
