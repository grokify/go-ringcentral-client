package main

import (
	"fmt"

	"github.com/grokify/ringcentral-sdk-go-scg"
)

const (
	RINGCENTRAL_SERVER_URL_PRODUCTION = "https://platform.ringcentral.com"
)

func main() {
	accessToken := os.Getenv(EnvAccessToken)
	mediaId := os.Getenv(EnvMediaId)

	mediaApi := ringcentral.NewExtensionApi()
	mediaApi.Configuration.SetAccessToken(accessToken)

	fmt.Println("DONE")
}
