package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/grokify/simplego/config"
	"github.com/grokify/simplego/fmt/fmtutil"
	"github.com/jessevdk/go-flags"

	rc "github.com/grokify/go-ringcentral-client/office/v1/client"
	ru "github.com/grokify/go-ringcentral-client/office/v1/util"
)

type CliOptions struct {
	EnvFile string `short:"e" long:"env" description:"Env filepath"`
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

	apiClient, err := ru.NewApiClientPasswordEnv("RINGCENTRAL_")
	if err != nil {
		log.Fatal(err)
	}

	params := rc.CreateSmsMessage{
		From: rc.MessageStoreCallerInfoRequest{
			PhoneNumber: os.Getenv("RINGCENTRAL_DEMO_SMS_FROM")},
		To: []rc.MessageStoreCallerInfoRequest{
			{
				PhoneNumber: os.Getenv("RINGCENTRAL_DEMO_SMS_TO")}},
		Text: os.Getenv("RINGCENTRAL_DEMO_SMS_TEXT")}

	fmtutil.PrintJSON(params)

	info, resp, err := apiClient.MessagesApi.SendSMS(context.Background(), "~", "~", params)

	if err != nil {
		panic(err)
	} else if resp.StatusCode > 299 {
		panic(fmt.Errorf("API Status %v", resp.StatusCode))
	}
	fmtutil.PrintJSON(info)

	fmt.Println("DONE")
}
