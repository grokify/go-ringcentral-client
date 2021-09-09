package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/antihax/optional"
	"github.com/grokify/simplego/fmt/fmtutil"
	hum "github.com/grokify/simplego/net/httputilmore"
	"github.com/jessevdk/go-flags"
	"github.com/rs/zerolog/log"

	rc "github.com/grokify/go-ringcentral-client/office/v1/client"
	ru "github.com/grokify/go-ringcentral-client/office/v1/util"
	"github.com/grokify/oauth2more/credentials"
	"github.com/grokify/oauth2more/ringcentral"
)

type CliOptions struct {
	CredsPath     string   `short:"c" long:"credspath" description:"Environment File Path"`
	Account       string   `short:"a" long:"account" description:"Environment Variable Name"`
	To            []string `short:"t" long:"to" description:"Recipients" required:"true"`
	Files         []string `short:"f" long:"file" description:"Files to send"`
	CoverPageText string   `short:"p" long:"coverpagetext" description:"Cover Page Text"`
}

func main() {
	opts := CliOptions{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal().Err(err)
		panic(err)
	}
	fmtutil.PrintJSON(opts)

	cset, err := credentials.ReadFileCredentialsSet(opts.CredsPath, true)
	if err != nil {
		log.Fatal().Err(err).
			Str("credentials_filepath", opts.CredsPath).
			Msg("cannot read credentials file")
	}

	creds, err := cset.Get(opts.Account)
	if err != nil {
		log.Fatal().Err(err).
			Str("available_accounts", strings.Join(cset.Accounts(), ", ")).
			Str("credentials_account", opts.Account).
			Msg("cannot find credentials account")
	}

	httpClient, err := creds.NewClient()
	if err != nil {
		log.Fatal().Err(err)
	}

	cu := ringcentral.ClientUtil{
		Client:    httpClient,
		ServerURL: creds.Application.ServerURL}

	usr, err := cu.GetSCIMUser()
	if err != nil {
		log.Fatal().Err(err)
	}
	fmtutil.PrintJSON(usr)

	demoData := ru.FaxRequest{
		To:            opts.To,
		CoverPageText: "Hello World Go",
		Resolution:    "High",
		FilePaths:     []string{"test_file.pdf"}}

	if 1 == 0 {
		resp, err := demoData.Post(httpClient, creds.Application.ServerURL)
		if err != nil {
			log.Fatal().Err(err)
		}
		err = hum.PrintResponse(resp, true)
		if err != nil {
			log.Fatal().Err(err)
		}
	}

	if 1 == 1 {
		apiClient, err := ru.NewApiClientHttpClientBaseURL(httpClient, creds.Application.ServerURL)
		if err != nil {
			log.Fatal().Err(err)
		}
		fmt.Println(demoData.FilePaths[0])

		file, err := os.Open(demoData.FilePaths[0])
		if err != nil {
			log.Fatal().Err(err)
		}

		params := rc.SendFaxMessageOpts{}
		if len(demoData.CoverPageText) > 0 {
			//params.FaxResolution = optional.NewString("High")
			params.CoverPageText = optional.NewString(demoData.CoverPageText)
		}
		if 1 == 1 {
			params.Attachment = optional.NewInterface(file)
		}
		info, resp, err := apiClient.MessagesApi.SendFaxMessage(
			context.Background(),
			"~",
			"~",
			demoData.To,
			&params)

		if err != nil {
			log.Fatal().Err(err)
		} else if resp.StatusCode > 299 {
			panic(fmt.Errorf("API Status %v", resp.StatusCode))
		}
		fmtutil.PrintJSON(info)
	}

	fmt.Println("DONE")
}
