package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jessevdk/go-flags"

	utils "github.com/grokify/go-ringcentral-engage/engagedigital/engagedigitalutil"
	ex "github.com/grokify/go-ringcentral-engage/engagedigital/examples"
)

type options struct {
	Site   string `short:"s" long:"site" description:"A site" required:"true"`
	Token  string `short:"t" long:"token" description:"A token" required:"true"`
	Action string `short:"a" long:"action" description:"Action" required:"true"`
	Update []bool `short:"u" long:"update" description:"Update team"`
	Id     string `short:"i" long:"id" description:"An object id" required:"false"`
	Id2    string `long:"id2" description:"A source object id" required:"false"`
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
	case "create":
		if len(opts.Id) > 0 && len(opts.Id2) > 0 {
			body := "Let's intervene here!"
			ex.HandleApiResponse(client.InterventionCommentsApi.CreateInterventionComment(
				context.Background(), body, opts.Id, opts.Id2))
		} else {
			log.Fatal("E_CREATE_INTERVENTION_ID_USER_ID")
		}
	case "delete":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.InterventionCommentsApi.DeleteInterventionComment(
				context.Background(), opts.Id))
		} else {
			log.Fatal("E_CREATE_NO_INTERVENTION_COMMENT_ID")
		}
	case "read":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.InterventionCommentsApi.GetInterventionComment(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.InterventionCommentsApi.GetAllInterventionComments(context.Background(), nil))
		}
	}

	fmt.Println("DONE")
}
