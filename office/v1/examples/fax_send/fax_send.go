package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/antihax/optional"
	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"
	hum "github.com/grokify/gotilla/net/httputilmore"
	"github.com/jessevdk/go-flags"
	"github.com/pkg/errors"

	rc "github.com/grokify/go-ringcentral/office/v1/client"
	ru "github.com/grokify/go-ringcentral/office/v1/util"
	"github.com/grokify/oauth2more/ringcentral"
)

type CliOptions struct {
	EnvFile string `short:"e" long:"env" description:"Env filepath"`
	EnvVar  string `short:"v" long:"envVar" description:"Environment Variable Name"`
	//Token         string   `short:"t" long:"token" description:"Token"`
	To            []string `short:"t" long:"to" description:"Recipients"`
	Files         []string `short:"f" long:"file" description:"Files to send"`
	CoverPageText string   `short:"c" long:"coverpagetext" description:"Cover Page Text"`
}

func getCredentials() (ringcentral.Credentials, error) {
	credsEmpty := ringcentral.Credentials{}
	opts := CliOptions{}
	_, err := flags.Parse(&opts)
	if err != nil {
		return credsEmpty, err
	}

	files, err := config.LoadDotEnv(opts.EnvFile)
	if err != nil {
		return credsEmpty, errors.Wrap(err, "E_LOAD_DOT_ENV")
	}
	fmtutil.PrintJSON(files)

	if len(opts.EnvVar) == 0 {
		return credsEmpty, errors.New("E_NO_ENV_VAR")
	}
	return ringcentral.NewCredentialsJSON([]byte(os.Getenv(opts.EnvVar)))
}

func getApiClient() (*rc.APIClient, error) {
	creds, err := getCredentials()
	if err != nil {
		return nil, err
	}
	fmtutil.PrintJSON(creds)
	return ru.NewApiClientCredentials(creds)
}

// example: $ go run fax_send.go -to=+16505550100 -file=$GOPATH/src/github.com/grokify/go-ringcentral/examples/fax_send/test_file.pdf
func main() {
	creds, err := getCredentials()
	if err != nil {
		log.Fatal(err)
	}
	httpClient, err := creds.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	cu := ringcentral.ClientUtil{
		Client:    httpClient,
		ServerURL: creds.Application.ServerURL}

	usr, err := cu.GetSCIMUser()
	if err != nil {
		log.Fatal(err)
	}
	fmtutil.PrintJSON(usr)

	demoData := ru.FaxRequest{
		To:            []string{"+16505626570"},
		CoverPageText: "Hello World Go",
		Resolution:    "High",
		FilePaths:     []string{"test_file.pdf"}}

	if 1 == 0 {
		resp, err := demoData.Post(httpClient, creds.Application.ServerURL)
		if err != nil {
			log.Fatal(err)
		}
		err = hum.PrintResponse(resp, true)
		if err != nil {
			log.Fatal(err)
		}
	}

	if 1 == 1 {
		apiClient, err := ru.NewApiClientHttpClientBaseURL(httpClient, creds.Application.ServerURL)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(demoData.FilePaths[0])

		file, err := os.Open(demoData.FilePaths[0])
		if err != nil {
			log.Fatal(err)
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
			panic(err)
		} else if resp.StatusCode > 299 {
			panic(fmt.Errorf("API Status %v", resp.StatusCode))
		}
		fmtutil.PrintJSON(info)
	}

	fmt.Println("DONE")
}
