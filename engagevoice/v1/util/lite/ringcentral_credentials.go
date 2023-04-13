package lite

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/grokify/goauth"
	"github.com/grokify/goauth/authutil"
	"github.com/grokify/mogo/encoding/jsonutil"
	"github.com/grokify/mogo/errors/errorsutil"
)

type EngageCredentials struct {
	Credentials goauth.Credentials
	EngageToken EngageToken
}

func NewEngageCredentialsJSON(rcCredentialsJSON []byte) (EngageCredentials, error) {
	ec := EngageCredentials{}
	rcCredentials, err := goauth.NewCredentialsJSON(rcCredentialsJSON, nil)
	if err != nil {
		return ec, errorsutil.Wrap(err, "NewEngageCredentialsJSON>>ringcentral.NewCredentialsJSON")
	}
	ec.Credentials = rcCredentials
	return ec, nil
}

func (ec *EngageCredentials) LoadNewTokens() error {
	rcToken, err := ec.Credentials.NewToken()
	if err != nil {
		return errorsutil.Wrap(err, "EngageCredentials>>ec.Credentials.NewToken()")
	}
	evToken, err := RcToEvToken(rcToken.AccessToken)
	if err != nil {
		return err
	}
	ec.EngageToken = evToken
	return nil
}

func (ec *EngageCredentials) NewClient() (*http.Client, error) {
	err := ec.LoadNewTokens()
	if err != nil {
		return nil, err
	}
	engageClient := authutil.NewClientToken(
		authutil.TokenBearer, ec.EngageToken.AccessToken, false)
	return engageClient, nil
}

type EngageToken struct {
	AccessToken string `json:"accessToken"`
	TokenType   string `json:"tokenType"`
}

func RcToEvToken(rctoken string) (EngageToken, error) {
	evToken := EngageToken{}

	res, err := http.PostForm(
		"https://engage.ringcentral.com/api/auth/login/rc/accesstoken",
		url.Values{"rcAccessToken": {rctoken}, "rcTokenType": {"Bearer"}})
	if err != nil {
		return evToken, errorsutil.Wrap(err, "RcToEvToken.PostForm")
	}
	if res.StatusCode >= 300 {
		return evToken, fmt.Errorf("RcToEvToken.ApiResponse.StatusCode [%v]", res.StatusCode)
	}

	_, err = jsonutil.UnmarshalReader(res.Body, &evToken)
	if err != nil {
		err = errorsutil.Wrap(err, "RcToEvToken.jsonutil.UnmarshalIoReader")
	}
	return evToken, err
}
