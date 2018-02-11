package main

import (
	"fmt"

	"github.com/grokify/gotilla/fmt/fmtutil"

	rcu "github.com/grokify/go-ringcentral/clientutil"
)

func main() {
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
	body, err := evt.GetInstantMessageBody()
	if err != nil {
		panic(err)
	}
	fmtutil.PrintJSON(body)

	fmt.Println("DONE")
}
