package clientutil

import (
	"strings"
)

const (
	InstantMessageSMSExample = "/restapi/v1.0/account/~/extension/823476228762/message-store/instant?type=SMS"
	InstantMessageSMSPattern = `message-store/instant?type=SMS`
)

func IsInstantMessageSMS(s string) bool {
	if strings.Index(s, InstantMessageSMSPattern) == -1 {
		return false
	}
	return true
}
