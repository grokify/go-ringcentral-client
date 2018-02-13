package main

import (
	"encoding/json"
	"fmt"

	"github.com/grokify/gotilla/fmt/fmtutil"

	rcu "github.com/grokify/go-ringcentral/clientutil"
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
		panic(err)
	}
	fmtutil.PrintJSON(body)

	if 1 == 0 {
		msg, err := rcu.GetFileBytesForEventType("instant_message_sms")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(msg))

		evt, err := rcu.EventParseBytes(msg)
		if err != nil {
			panic(err)
		}
		fmtutil.PrintJSON(evt)
		if evt.IsEventType(rcu.InstantMessageEvent) {
			body, err := evt.GetInstantMessageBody()
			if err != nil {
				panic(err)
			}
			fmtutil.PrintJSON(body)
		}
	}

	fmt.Println("DONE")
}
