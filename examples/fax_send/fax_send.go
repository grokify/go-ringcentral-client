package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/joho/godotenv"

	rc "github.com/grokify/go-ringcentral/client"
	ro "github.com/grokify/oauth2more/ringcentral"
)

func loadEnv() error {
	envPaths := []string{}
	if len(os.Getenv("ENV_PATH")) > 0 {
		envPaths = append(envPaths, os.Getenv("ENV_PATH"))
	}
	return godotenv.Load(envPaths...)
}

func newApiClient() (*rc.APIClient, error) {
	httpClient, err := ro.NewClientPassword(
		ro.ApplicationCredentials{
			ServerURL:    os.Getenv("RINGCENTRAL_SERVER_URL"),
			ClientID:     os.Getenv("RINGCENTRAL_CLIENT_ID"),
			ClientSecret: os.Getenv("RINGCENTRAL_CLIENT_SECRET")},
		ro.UserCredentials{
			Username:  os.Getenv("RINGCENTRAL_USERNAME"),
			Extension: os.Getenv("RINGCENTRAL_EXTENSION"),
			Password:  os.Getenv("RINGCENTRAL_PASSWORD")})
	if err != nil {
		return nil, err
	}

	apiConfig := rc.NewConfiguration()
	apiConfig.BasePath = os.Getenv("RINGCENTRAL_SERVER_URL")
	apiConfig.HTTPClient = httpClient
	apiClient := rc.NewAPIClient(apiConfig)
	return apiClient, nil
}

func main() {
	var to = flag.Int("to", 0, "Recipient fax number")
	flag.Parse()

	if *to == 0 {
		fmt.Println("Usage: fax_send.go -to=16505551212\nexiting...")
		return
	}

	err := loadEnv()
	if err != nil {
		panic(err)
	}
	apiClient, err := newApiClient()
	if err != nil {
		panic(err)
	}

	filename := "test_file.pdf"

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	info, resp, err := apiClient.MessagesApi.SendFaxMessage(
		context.Background(),
		"~",
		"~",
		file,
		"High",
		[]string{fmt.Sprintf("+%d", *to)},
		map[string]interface{}{
			"coverPageText": "HelloWorld!",
		},
	)

	if err != nil {
		panic(err)
	} else if resp.StatusCode > 299 {
		panic(fmt.Errorf("API Status %v", resp.StatusCode))
	}
	fmtutil.PrintJSON(info)
	fmt.Println("DONE")
}
