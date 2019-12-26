package main

import (
	"context"
	"fmt"
	"log"

	"github.com/antihax/optional"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/jessevdk/go-flags"

	"github.com/grokify/go-ringcentral-engage/engagedigital/engagedigital"
	utils "github.com/grokify/go-ringcentral-engage/engagedigital/engagedigitalutil"
)

type options struct {
	Site  string `short:"s" long:"site" description:"A site" required:"true"`
	Token string `short:"t" long:"token" description:"A token" required:"true"`
}

func main() {
	opts := options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	client := utils.NewApiClient(opts.Site, opts.Token)

	info, resp, err := client.ChannelsApi.GetAllChannels(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode != 200 {
		log.Fatal(resp.StatusCode)
	}
	channelLabel := "realtime"
	channel := engagedigital.Channel{}
	match := false
	for _, chanTry := range info.Records {
		if chanTry.Label == channelLabel {
			channel = chanTry
			match = true
			break
		}
	}
	if !match {
		log.Fatal("Didn't find channel")
	}
	apiOpts := &engagedigital.UpdateChannelOpts{}
	if channel.HardCapability == 5 {
		apiOpts.HardCapability = optional.NewInt32(7)
	} else {
		apiOpts.HardCapability = optional.NewInt32(5)
	}

	info2, resp, err := client.ChannelsApi.UpdateChannel(
		context.Background(), channel.Id, apiOpts)
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode != 200 {
		log.Fatal(resp.StatusCode)
	}
	fmtutil.PrintJSON(info2)

	fmt.Println("DONE")
}
