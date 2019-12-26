package examples

import (
	"fmt"
	"log"
	"net/http"

	"github.com/grokify/go-ringcentral-engage/engagedigital/engagedigital"
	"github.com/grokify/gotilla/fmt/fmtutil"
)

func HandleApiResponse(info interface{}, resp *http.Response, err error) {
	if err != nil {
		fmt.Println("E_ERR")
		if resp.StatusCode != 200 {
			fmt.Printf("StatusCode [%v]\n", resp.StatusCode)
			err2 := err.(engagedigital.GenericOpenAPIError)
			fmt.Println(string(err2.Body()))
		}
		log.Fatal(err)
	} else if resp.StatusCode != 200 {
		log.Fatal(resp.StatusCode)
	}
	fmtutil.PrintJSON(info)
}
