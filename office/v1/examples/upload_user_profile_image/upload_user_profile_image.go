package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/grokify/mogo/config"
	"github.com/grokify/mogo/errors/errorsutil"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/net/http/httputilmore"
	"github.com/jessevdk/go-flags"

	ru "github.com/grokify/go-ringcentral-client/office/v1/util"
	ro "github.com/grokify/goauth/ringcentral"
)

type cliOptions struct {
	File string `short:"f" long:"file" description:"path to image file"`
}

func main() {
	opts := cliOptions{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	err = config.LoadDotEnvSkipEmpty(os.Getenv("ENV_PATH"), "./.env")
	if err != nil {
		panic(err)
	}

	httpClient, err := ro.NewHTTPClientEnvFlexStatic("")
	if err != nil {
		log.Fatal(errorsutil.Wrap(err, "getHttpClientEnv"))
	}
	serverURL := os.Getenv("RINGCENTRAL_SERVER_URL")
	apiClient, err := ru.NewApiClientHttpClientBaseURL(httpClient, serverURL)
	if err != nil {
		log.Fatal(err)
	}

	if 1 == 0 {
		f, err := os.Open(opts.File)
		if err != nil {
			log.Fatal(err)
		}

		info, resp, err := apiClient.UserSettingsApi.UploadImageByPostForm(
			context.Background(), "~", "~", f)

		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode >= 300 {
			log.Fatal(fmt.Sprintf("Status [%v]", resp.StatusCode))
		}
		fmtutil.PrintJSON(info)
	}

	if 1 == 1 {
		resp, err := ru.UploadUserProfileImage(httpClient,
			os.Getenv("RINGCENTRAL_SERVER_URL"),
			"~",
			"~",
			opts.File)
		if err != nil {
			log.Fatal(err)
		}

		httputilmore.PrintResponse(resp, true)
	}

	fmt.Println("DONE")
}
