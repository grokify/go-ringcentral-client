package main

import (
	"context"
	"fmt"
	"os"

	"github.com/grokify/mogo/fmt/fmtutil"
	scu "github.com/grokify/mogo/strconv/strconvutil"

	rc "github.com/grokify/go-ringcentral-client/office/v1/client"
	"github.com/grokify/go-ringcentral-client/office/v1/examples"
)

func ringoutBodyEnv() *rc.MakeRingOutRequest {
	return &rc.MakeRingOutRequest{
		From: rc.MakeRingOutCallerInfoRequestFrom{
			PhoneNumber: os.Getenv("RINGCENTRAL_DEMO_RINGOUT_FROM"),
		},
		To: rc.MakeRingOutCallerInfoRequestTo{
			PhoneNumber: os.Getenv("RINGCENTRAL_DEMO_RINGOUT_TO"),
		},
		PlayPrompt: scu.MustParseBool(os.Getenv("RINGCENTRAL_DEMO_RINGOUT_PROMPT")),
	}
}

func main() {
	err := examples.LoadEnv()
	if err != nil {
		panic(err)
	}

	apiClient, err := examples.NewApiClient()
	if err != nil {
		panic(err)
	}

	body := ringoutBodyEnv()
	fmtutil.PrintJSON(body)

	info, resp, err := apiClient.RingOutApi.MakeRingOutCallNew(
		context.Background(), "~", "~", *body,
	)
	if err != nil {
		panic(err)
	} else if resp.StatusCode > 299 {
		panic(fmt.Errorf("API Status %v", resp.StatusCode))
	}
	fmtutil.PrintJSON(info)
	fmt.Println("DONE")
}
