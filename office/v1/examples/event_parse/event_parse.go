package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/grokify/mogo/fmt/fmtutil"

	rc "github.com/grokify/go-ringcentral-client/office/v1/client"
	ru "github.com/grokify/go-ringcentral-client/office/v1/util"
)

type Body struct {
	ID   string
	Body interface{}
}

type BodyOne struct {
	Foo string
}

type BodyTwo struct {
	Bar string
}

func main() {
	str1 := `{"ID":"MY_ID","Foo":"Here"}`
	str2 := `{"ID":"MY_ID","Bar":"Here"}`

	fmt.Println(str1)
	fmt.Println(str2)

	body := &Body{}
	err := json.Unmarshal([]byte(str1), body)
	if err != nil {
		log.Fatal(err)
	}
	fmtutil.PrintJSON(body)

	if 1 == 1 {
		msg, err := ru.GetFileBytesForEventType("instant_message_sms")
		if err != nil {
			log.Fatal(err)
		}

		evt, err := ru.EventParseBytes(msg)
		if err != nil {
			log.Fatal(err)
		}

		if evt.IsEventType(ru.InstantMessageEvent) {
			body, err := evt.GetInstantMessageBody()
			if err != nil {
				log.Fatal(err)
			}
			fmtutil.PrintJSON(body)
		} else {
			log.Fatal("is not InstantMessageEvent")
		}
	}

	if 1 == 0 {
		msg, err := ru.GetFileBytesForEventType("glip_post")
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(string(msg))

		evt, err := ru.EventParseBytes(msg)
		if err != nil {
			log.Fatal(err)
		}
		//fmtutil.PrintJSON(evt)
		if evt.IsEventType(ru.GlipPostEvent) {
			body, err := evt.GetGlipPostEventBody()
			if err != nil {
				log.Fatal(err)
			}
			fmtutil.PrintJSON(body)
		}
	}
	if 1 == 1 {
		file := filepath.Join(
			os.Getenv("GOPATH"),
			"/src/github.com/grokify/go-ringcentral-client/office/v1/util/example_api-response_list-messages.json")
		resp, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(string(resp))
		info := &rc.GetMessageList{}
		err = json.Unmarshal(resp, info)
		if err != nil {
			log.Fatal(err)
		}
		fmtutil.PrintJSON(info)
	}

	fmt.Println("DONE")
}
