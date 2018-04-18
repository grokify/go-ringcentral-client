package clientutil

import (
	"fmt"
	"strings"
)

const (
	InstantMessageSMSExample   = "/restapi/v1.0/account/~/extension/12345678/message-store/instant?type=SMS"
	InstantMessageSMSPattern   = `message-store/instant?type=SMS`
	GlipPostEventFilterPattern = `/restapi/v1.0/glip/posts`
)

type EventType int

const (
	AccountPresenceEvent EventType = iota
	ContactDirectoryEvent
	DetailedExtensionPresenceEvent
	DetailedExtensionPresenceWithSIPEvent
	ExtensionFavoritesEvent
	ExtensionFavoritesPresenceEvent
	ExtensionGrantListEvent
	ExtensionListEvent
	ExtensionInfoEvent
	ExtensionPresenceEvent
	ExtensionPresenceLineEvent
	GlipGroupsEvent
	GlipPostEvent
	GlipUnreadMessageCountEvent
	InboundMessageEvent
	IncomingCallEvent
	InstantMessageEvent
	MessageEvent
	MissedCallEvent
	RCVideoNotificationsEvent
)

// Events is an array of event structs for reference.
var Events = []string{
	"AccountPresenceEvent",
	"ContactDirectoryEvent",
	"DetailedExtensionPresenceEvent",
	"DetailedExtensionPresenceWithSIPEvent",
	"ExtensionFavoritesEvent",
	"ExtensionFavoritesPresenceEvent",
	"ExtensionGrantListEvent",
	"ExtensionListEvent",
	"ExtensionInfoEvent",
	"ExtensionPresenceEvent",
	"ExtensionPresenceLineEvent",
	"GlipGroupsEvent",
	"GlipPostEvent",
	"GlipUnreadMessageCountEvent",
	"InboundMessageEvent",
	"IncomingCallEvent",
	"InstantMessageEvent",
	"MessageEvent",
	"MissedCallEvent",
	"RCVideoNotificationsEvent",
}

func (d EventType) String() string { return Events[d] }

func IsInstantMessageSMS(s string) bool {
	if strings.Index(s, InstantMessageSMSPattern) == -1 {
		return false
	}
	return true
}

func ParseEventTypeForFilter(eventFilter string) (EventType, error) {
	if strings.Index(eventFilter, GlipPostEventFilterPattern) > -1 {
		return GlipPostEvent, nil
	}
	return GlipPostEvent, fmt.Errorf("No Event found for filter %v", eventFilter)
}
