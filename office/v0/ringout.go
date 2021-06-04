package legacy

import (
	"io/ioutil"
	"net/http"
	"strings"

	uu "github.com/grokify/simplego/net/urlutil"
	scu "github.com/grokify/simplego/strconv/strconvutil"
)

const (
	DocumentationURL = "https://success.ringcentral.com/articles/RC_Knowledge_Article/2080"
	CmdCall          = "call"
)

var RingOutURL = "https://service.ringcentral.com/ringout.asp"

func Call(params CallRequestInfo) (*CallResponseInfo, *http.Response, error) {
	params.Command = CmdCall

	roURL, err := uu.URLAddQueryInterfaceString(RingOutURL, params)
	if err != nil {
		return nil, nil, err
	}
	resp, err := http.Get(roURL.String())
	if err != nil {
		return nil, resp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp, err
	}

	resInfo := &CallResponseInfo{}
	parts := strings.Split(string(body), " ")
	if len(parts) == 3 {
		resInfo.Status = strings.TrimSpace(parts[0])
		resInfo.SessionID = strings.TrimSpace(parts[1])
		resInfo.WS = strings.TrimSpace(parts[2])
	}
	return resInfo, resp, nil
}

type CallRequestInfoStrings struct {
	Command   string `url:"cmd,omitempty"`
	Username  string `url:"username,omitempty"`
	Extension string `url:"ext,omitempty"`
	Password  string `url:"password,omitempty"`
	To        string `url:"to,omitempty"`
	From      string `url:"from,omitempty"`
	CallerID  string `url:"clid,omitempty"`
	Prompt    string `url:"prompt,omitempty"`
}

func (s *CallRequestInfoStrings) ToCanonical() CallRequestInfo {
	can := CallRequestInfo{
		Command:  s.Command,
		Password: s.Password,
	}
	if len(strings.TrimSpace(s.Username)) > 0 {
		can.Username = scu.MustParseE164ToInt(s.Username)
	}
	if len(strings.TrimSpace(s.To)) > 0 {
		can.To = scu.MustParseE164ToInt(s.To)
	}
	if len(strings.TrimSpace(s.From)) > 0 {
		can.From = scu.MustParseE164ToInt(s.From)
	}
	if len(strings.TrimSpace(s.CallerID)) > 0 {
		can.CallerID = scu.MustParseE164ToInt(s.CallerID)
	}
	if len(strings.TrimSpace(s.Extension)) > 0 {
		can.Extension = scu.MustParseInt(s.Extension)
	}
	s.Prompt = strings.ToLower(strings.TrimSpace(s.Prompt))
	if s.Prompt == "true" || s.Prompt == "1" {
		can.Prompt = 1
	}
	return can
}

type CallRequestInfo struct {
	Command   string `url:"cmd,omitempty"`
	Username  int    `url:"username,omitempty"`
	Extension int    `url:"ext,omitempty"`
	Password  string `url:"password,omitempty"`
	To        int    `url:"to,omitempty"`
	From      int    `url:"from,omitempty"`
	CallerID  int    `url:"clid,omitempty"`
	Prompt    int    `url:"prompt,omitempty"`
}

type CallResponseInfo struct {
	Status    string
	SessionID string
	WS        string // can be .11
}
