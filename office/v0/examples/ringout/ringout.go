package main

import (
	"fmt"
	"os"

	legacy "github.com/grokify/go-ringcentral-client/office/v0"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(os.Getenv("ENV_PATH"))
	if err != nil {
		panic(err)
	}

	legacy.RingOutURL = "http://localhost:8080/ringout.asp"

	reqS := legacy.CallRequestInfoStrings{
		Username:  os.Getenv("RINGCENTRAL_USERNAME"),
		Extension: os.Getenv("RINGCENTRAL_EXTENSION"),
		Password:  os.Getenv("RINGCENTRAL_PASSWORD"),
		To:        os.Getenv("RINGCENTRAL_DEMO_RINGOUT_TO"),
		From:      os.Getenv("RINGCENTRAL_DEMO_RINGOUT_FROM"),
		Prompt:    os.Getenv("RINGCENTRAL_DEMO_RINGOUT_PROMPT"),
	}

	info, resp, err := legacy.Call(reqS.ToCanonical())
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	fmtutil.PrintJSON(info)
}
