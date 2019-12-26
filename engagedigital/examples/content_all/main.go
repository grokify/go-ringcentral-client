package main

import (
	"context"
	"fmt"
	"log"

	"github.com/antihax/optional"
	"github.com/jessevdk/go-flags"

	"github.com/grokify/go-ringcentral-engage/engagedigital/engagedigital"
	utils "github.com/grokify/go-ringcentral-engage/engagedigital/engagedigitalutil"
	ex "github.com/grokify/go-ringcentral-engage/engagedigital/examples"
)

type options struct {
	Site     string `short:"s" long:"site" description:"A site" required:"true"`
	Token    string `short:"t" long:"token" description:"A token" required:"true"`
	Action   string `short:"a" long:"action" description:"Action" required:"true"`
	Update   []bool `short:"u" long:"update" description:"Update team"`
	Id       string `short:"i" long:"id" description:"An object id" required:"false"`
	SourceId string `long:"id2" description:"A source object id" required:"false"`
	Object   string `short:"o" long:"object" description:"An object" required:"false"`
}

func main() {
	opts := options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	client := utils.NewApiClient(opts.Site, opts.Token)

	switch opts.Action {
	case "create":
		if len(opts.SourceId) == 0 {
			log.Fatal("E_CREATE_NO_SOURCE_ID")
		}
		apiOpts := &engagedigital.CreateContentOpts{
			SourceId: optional.NewString(opts.SourceId)}
		apiOpts.To = optional.NewInterface([]string{"demo@example.com"})
		apiOpts.Title = optional.NewString("Demo Email")
		body := "Test Content Body"
		ex.HandleApiResponse(client.ContentsApi.CreateContent(
			context.Background(), body, apiOpts))
	case "read":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.ContentsApi.GetContent(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.ContentsApi.GetAllContents(context.Background(), nil))
		}
	case "categorizecontent":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.ContentsApi.CategorizeContent(
				context.Background(), opts.Id, []string{"5c8bf22b14bf8a84c44c774c"}))
		} else {
			log.Fatal("E_IGNORE_NO_ID")
		}
	case "ignore":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.ContentsApi.IgnoreContent(context.Background(), opts.Id))
		} else {
			log.Fatal("E_IGNORE_NO_ID")
		}
	}

	fmt.Println("DONE")
}
