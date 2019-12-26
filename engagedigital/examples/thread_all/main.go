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
	Site   string `short:"s" long:"site" description:"A site" required:"true"`
	Token  string `short:"t" long:"token" description:"A token" required:"true"`
	Action string `short:"a" long:"action" description:"Action" required:"true"`
	Id     string `short:"i" long:"id" description:"An object id" required:"false"`
	Object string `short:"o" long:"object" description:"An object" required:"false"`
}

func main() {
	opts := options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	client := utils.NewApiClient(opts.Site, opts.Token)

	switch opts.Action {
	case "read":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.ThreadsApi.GetThread(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.ThreadsApi.GetAllThreads(context.Background(), nil))
		}
	case "categorize":
		if len(opts.Id) > 0 {
			apiOpts := &engagedigital.CategorizeThreadOpts{
				ThreadCategoryIds: optional.NewInterface([]string{"5c8bf22b14bf8a84c44c774c"})}
			ex.HandleApiResponse(client.ThreadsApi.CategorizeThread(
				context.Background(), opts.Id, apiOpts))
		} else {
			log.Fatal("E_CATEGORIZE_NO_THREAD_ID")
		}
	case "archive":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.ThreadsApi.ArchiveThread(context.Background(), opts.Id))
		} else {
			log.Fatal("E_ARCHIVE_NO_THREAD_ID")
		}
	case "close":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.ThreadsApi.CloseThread(context.Background(), opts.Id))
		} else {
			log.Fatal("E_CLOSE_NO_THREAD_ID")
		}
	case "open":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.ThreadsApi.OpenThread(context.Background(), opts.Id))
		} else {
			log.Fatal("E_OPEN_NO_THREAD_ID")
		}
	default:
		log.Fatal(fmt.Sprintf("E_BAD_ACTION [%s]", opts.Action))
	}

	fmt.Println("DONE")
}
