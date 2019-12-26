package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/antihax/optional"
	"github.com/jessevdk/go-flags"

	"github.com/grokify/go-ringcentral-engage/engagedigital/engagedigital"
	utils "github.com/grokify/go-ringcentral-engage/engagedigital/engagedigitalutil"
	"github.com/grokify/gotilla/fmt/fmtutil"
)

// curl -XPOST https://DOMAIN.api.engagement.dimelo.com/1.0/attachments -H 'Authorization: Bearer TOKEN' -F 'file=@ringcentral-engage-logo.svg'

type options struct {
	Site  string `short:"s" long:"site" description:"A site" required:"true"`
	Token string `short:"t" long:"token" description:"A token" required:"true"`
	File  string `short:"f" long:"file" description:"A file" required:"true"`
}

func main() {
	opts := options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	client := utils.NewApiClient(opts.Site, opts.Token)

	file, err := os.Open(opts.File)
	if err != nil {
		log.Fatal(err)
	}

	apiOpts := &engagedigital.CreateAttachmentOpts{
		File: optional.NewInterface(file)}

	info, resp, err := client.AttachmentsApi.CreateAttachment(context.Background(), apiOpts)

	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode != 200 {
		log.Fatal(resp.StatusCode)
	}
	fmtutil.PrintJSON(info)

	fmt.Println("DONE")
}
