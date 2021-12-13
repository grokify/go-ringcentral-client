package lite

import (
	"net/http"

	"github.com/grokify/mogo/net/httputilmore"
)

const (
	EnvEngageVoiceServerURL = "ENGAGE_VOICE_SERVER_URL"
	EnvEngageVoiceAccountID = "ENGAGE_VOICE_ACCOUNT_ID"
	EnvEngageVoiceToken     = "ENGAGE_VOICE_TOKEN"
	APIUsersURLPath         = "/api/v1/admin/users"
	APITokensURLPath        = "/api/v1/admin/token"
)

func NewClientToken(token string) *http.Client {
	header := http.Header{}
	header.Add(EngageVoiceTokenHeader, token)

	client := &http.Client{}
	client.Transport = httputilmore.TransportWithHeaders{
		Transport: client.Transport,
		Header:    header}
	return client
}

func NewClientRingCentralPassword(rcCredentials []byte) (*http.Client, error) {
	appCredentials, err := NewEngageCredentialsJSON(rcCredentials)
	if err != nil {
		return nil, err
	}
	return appCredentials.NewClient()

	/*
		// 1. Get RingCentral Token
		ringcentralCredentials, err := ringcentral.NewCredentialsJSON(
			appCredentialsJSON)
		if err != nil {
			return nil, err
		}

		ringcentralAccessToken, err := ringcentralCredentials.NewToken()
		if err != nil {
			return nil, err
		}

		// 2. Get Engage Token
		engageAccessToken, err := RcToEvToken(ringcentralAccessToken.AccessToken)
		if err != nil {
			return nil, err
		}

		// 3. Get Engage Client
		engageClient := goauth.NewClientToken(
			goauth.TokenBearer, engageAccessToken.AccessToken, false)
		return engageClient, nil*/
}
