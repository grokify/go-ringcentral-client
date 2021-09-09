package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/antihax/optional"
	"github.com/caarlos0/env"
	"github.com/grokify/oauth2more/credentials"
	"github.com/grokify/simplego/config"
	"github.com/grokify/simplego/fmt/fmtutil"
	hum "github.com/grokify/simplego/net/httputilmore"
	uu "github.com/grokify/simplego/net/urlutil"
	"github.com/jessevdk/go-flags"

	rc "github.com/grokify/go-ringcentral-client/office/v1/client"
	ru "github.com/grokify/go-ringcentral-client/office/v1/util"
)

type CliOptions struct {
	EnvFile       string   `short:"e" long:"env" description:"Env filepath"`
	To            []string `short:"t" long:"to" description:"Recipients"`
	Files         []string `short:"f" long:"file" description:"Files to send"`
	CoverPageText string   `short:"c" long:"coverpagetext" description:"Cover Page Text"`
}

type RingCentralConfig struct {
	ServerURL    string `env:"RINGCENTRAL_SERVER_URL"`
	ClientID     string `env:"RINGCENTRAL_CLIENT_ID"`
	ClientSecret string `env:"RINGCENTRAL_CLIENT_SECRET"`
	Username     string `env:"RINGCENTRAL_USERNAME"`
	Extension    string `env:"RINGCENTRAL_EXTENSION"`
	Password     string `env:"RINGCENTRAL_PASSWORD"`
}

func NewRingCentralConfigEnv() (*RingCentralConfig, error) {
	appCfg := &RingCentralConfig{}
	return appCfg, env.Parse(appCfg)
}

func (cfg *RingCentralConfig) ApplicationCredentials() credentials.ApplicationCredentials {
	return credentials.ApplicationCredentials{
		ServerURL:    cfg.ServerURL,
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret}
}

func (cfg *RingCentralConfig) PasswordCredentials() credentials.PasswordCredentials {
	return credentials.PasswordCredentials{
		Username: cfg.Username,
		Password: cfg.Password}
}

func sendFaxRaw(opts CliOptions, httpClient *http.Client) {
	fax := ru.FaxRequest{
		To:            opts.To,
		CoverPageText: opts.CoverPageText,
		Resolution:    "High",
		FilePaths:     opts.Files,
	}

	url := uu.JoinAbsolute(os.Getenv("RINGCENTRAL_SERVER_URL"), "/restapi/v1.0/account/~/extension/~/fax")

	resp, err := fax.Post(httpClient, url)
	if err != nil {
		panic(err)
	}
	err = hum.PrintResponse(resp, true)
	if err != nil {
		panic(err)
	}
}

// example: $ go run fax_send.go -to=+16505550100 -file=$GOPATH/src/github.com/grokify/go-ringcentral-client/office/v1/examples/fax_send/test_file.pdf
func main() {
	opts := CliOptions{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}
	err = config.LoadDotEnvFirst(opts.EnvFile, os.Getenv("ENV_PATH"), "./.env")
	if err != nil {
		log.Fatal(err)
	}
	rcCfg, err := NewRingCentralConfigEnv()
	if err != nil {
		log.Fatal(err)
	}

	fmtutil.PrintJSON(opts)
	fmtutil.PrintJSON(rcCfg)

	apiClient, err := ru.NewApiClientPassword(
		rcCfg.ApplicationCredentials(),
		rcCfg.PasswordCredentials())
	if err != nil {
		log.Fatal(err)
	}

	httpClient := apiClient.HTTPClient()

	if 1 == 0 {
		sendFaxRaw(opts, httpClient)
	}

	if 1 == 1 {
		fmt.Println(opts.Files[0])

		file, err := os.Open(opts.Files[0])
		if err != nil {
			log.Fatal(err)
		}

		params := rc.SendFaxMessageOpts{}
		if len(opts.CoverPageText) > 0 {
			//params.FaxResolution = optional.NewString("High")
			params.CoverPageText = optional.NewString(opts.CoverPageText)
		}
		fmtutil.PrintJSON(opts)
		if 1 == 1 {
			params.Attachment = optional.NewInterface(file)
		}
		info, resp, err := apiClient.MessagesApi.SendFaxMessage(
			context.Background(),
			"~",
			"~",
			opts.To,
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
