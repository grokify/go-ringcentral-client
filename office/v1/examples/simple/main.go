package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/grokify/oauth2more"
	"github.com/grokify/oauth2more/ringcentral"
	"github.com/grokify/simplego/config"
	"github.com/grokify/simplego/fmt/fmtutil"
	"github.com/grokify/simplego/mime/multipartutil"
	"github.com/grokify/simplego/net/httputilmore"
	"github.com/grokify/simplego/net/urlutil"
	"github.com/jessevdk/go-flags"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

type Options struct {
	EnvPath string `short:"e" long:"envPath" description:"Environment File Path"`
	EnvVar  string `short:"v" long:"envVar" description:"Environment Varaible Name"`
	Token   string `short:"t" long:"token" description:"Token"`
}

func loadEnv() (Options, error) {
	opts := Options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		return opts, err
	}

	err = config.LoadDotEnvSkipEmpty(opts.EnvPath, os.Getenv("ENV_PATH"), "./.env")
	if err != nil {
		return opts, errors.Wrap(err, "E_LOAD_DOT_ENV")
	}
	return opts, nil
}

func getApplicationConfig(cfg []byte) (ringcentral.Credentials, error) {
	ac := ringcentral.Credentials{}
	err := json.Unmarshal(cfg, &ac)
	if err != nil {
		return ac, errors.Wrap(
			err, fmt.Sprintf("E_JSON_UNMARSHAL [%v]", string(cfg)))
	}
	return ac, nil
}

func getTokenFromApplicationConfig(ac ringcentral.Credentials) (*oauth2.Token, error) {
	token, err := ringcentral.NewTokenPassword(
		ac.Application,
		ac.PasswordCredentials)
	if err != nil {
		return nil, err
	}
	token.Expiry = token.Expiry.UTC()

	fmtutil.PrintJSON(token)
	return token, nil
}

const (
	ExtensionAnsweringRuleURL = "/restapi/v1.0/account/~/extension/~/answering-rule"
	ExtensionGreetingURL      = "/restapi/v1.0/account/~/extension/~/greeting"
)

func main() {
	opts, err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}

	cfgJSON := ""

	if len(opts.EnvVar) > 0 {
		cfgJSON = os.Getenv(opts.EnvVar)
	}
	if len(cfgJSON) == 0 {
		log.Fatal("E_NO_APP_CONFIG")
	}

	ac, err := getApplicationConfig([]byte(cfgJSON))
	if err != nil {
		log.Fatal(err)
	}
	token, err := getTokenFromApplicationConfig(ac)
	if err != nil {
		log.Fatal(err)
	}
	client := oauth2more.NewClientTokenOAuth2(token)

	resp, err := client.Get(urlutil.JoinAbsolute(ac.Application.ServerURL, ExtensionAnsweringRuleURL))
	if err != nil {
		log.Fatal(err)
	}
	err = httputilmore.PrintResponse(resp, true)
	if err != nil {
		log.Fatal(err)
	}

	params := url.Values{}
	params.Set("type", "HoldMusic")
	params.Set("answeringRuleId", "business-hours-rule")

	req, err := multipartutil.NewRequest(
		http.MethodPost,
		urlutil.JoinAbsolute(ac.Application.ServerURL, ExtensionGreetingURL),
		params,
		[]multipartutil.FileInfo{
			{
				MIMEPartName: "binary",
				Filepath:     "star-wars_decoded.wav",
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	if 1 == 0 {
		err = httputilmore.PrintRequestOut(req, true)
		if err != nil {
			log.Fatal(err)
		}
	}
	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	err = httputilmore.PrintResponse(resp, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DONE")
}
