package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"

	ru "github.com/grokify/go-ringcentral/clientutil"
	ro "github.com/grokify/oauth2more/ringcentral"
)

// example: $ go run fax_send.go -to=+16505550100 -file=$GOPATH/src/github.com/grokify/go-ringcentral/examples/fax_send/test_file.pdf
func main() {
	var toPhoneNumber = flag.String("to", "", "Recipient fax number")
	var filename = flag.String("file", "", "Path to file, leave blank for empty")
	var coverPageText = flag.String("coverPageText", "Hello World!", "Cover page text")
	flag.Parse()

	if len(*toPhoneNumber) == 0 {
		fmt.Println("Usage: fax_send.go -to=+16505550100 -file=/path/to/file [-coverPageText='Hello World!']\nexiting...")
		return
	}
	if len(*filename) == 0 {
		fmt.Println("Usage: fax_send.go -to=+16505550100 -file=/path/to/file [-coverPageText='Hello World!']\nexiting...")
		return
	}

	err := config.LoadDotEnvSkipEmpty(os.Getenv("ENV_PATH"), "./.env")
	if err != nil {
		log.Fatal(err)
	}

	apiClient, err := ru.NewApiClientPassword(
		ro.ApplicationCredentials{
			ServerURL:    os.Getenv("RINGCENTRAL_SERVER_URL"),
			ClientID:     os.Getenv("RINGCENTRAL_CLIENT_ID"),
			ClientSecret: os.Getenv("RINGCENTRAL_CLIENT_SECRET")},
		ro.PasswordCredentials{
			Username:  os.Getenv("RINGCENTRAL_USERNAME"),
			Extension: os.Getenv("RINGCENTRAL_EXTENSION"),
			Password:  os.Getenv("RINGCENTRAL_PASSWORD")})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(filename)

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}

	params := map[string]interface{}{}
	if len(*coverPageText) > 0 {
		params["coverPageText"] = *coverPageText
	}

	info, resp, err := apiClient.MessagesApi.SendFaxMessage(
		context.Background(),
		"~",
		"~",
		file,
		"High",
		[]string{*toPhoneNumber},
		params,
	)

	if err != nil {
		panic(err)
	} else if resp.StatusCode > 299 {
		panic(fmt.Errorf("API Status %v", resp.StatusCode))
	}
	fmtutil.PrintJSON(info)
	fmt.Println("DONE")
}
