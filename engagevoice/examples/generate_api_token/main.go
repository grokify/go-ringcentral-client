package main

import (
	"fmt"
	"log"

	"github.com/grokify/go-ringcentral-engage/engagevoice/engagevoiceutil"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	Username string `short:"u" long:"username" description:"Username"`
	Password string `short:"p" long:"password" description:"Password"`
	Token    string `short:"t" long:"token" description:"Token"`
}

func main() {
	opts := Options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	if len(opts.Username) > 0 && len(opts.Password) > 0 {
		apiToken, userInfo, err := engagevoiceutil.GenerateAPIToken(opts.Username, opts.Password)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("API_TOKEN: [%v]\n", apiToken)
		fmt.Printf("ACCOUNT_ID: [%v]\n", userInfo.Accounts[0].AccountID)
	}

	tokens, err := engagevoiceutil.ListTokens(opts.Token)
	if err != nil {
		log.Fatal(err)
	}
	fmtutil.PrintJSON(tokens)

	fmt.Println("DONE")
}
