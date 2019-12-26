package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/jessevdk/go-flags"

	utils "github.com/grokify/go-ringcentral-engage/engagedigital/engagedigitalutil"
)

type options struct {
	Site   string `short:"s" long:"site" description:"A site" required:"true"`
	Token  string `short:"t" long:"token" description:"A token" required:"true"`
	Object string `short:"o" long:"object" description:"An object" required:"true"`
	Id     string `short:"i" long:"id" description:"An object id" required:"false"`
}

func main() {
	opts := options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	client := utils.NewApiClient(opts.Site, opts.Token)

	opts.Id = strings.TrimSpace(opts.Id)
	opts.Object = strings.ToLower(strings.TrimSpace(opts.Object))

	switch opts.Object {

	case "locale":
		info, resp, err := client.LocalesApi.GetAllLocales(context.Background())
		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode != 200 {
			log.Fatal(resp.StatusCode)
		}
		fmtutil.PrintJSON(info)
	}

	fmt.Println("DONE")
}
