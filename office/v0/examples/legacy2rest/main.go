package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	hum "github.com/grokify/simplego/net/httputilmore"
	nhu "github.com/grokify/simplego/net/nethttputil"
	"github.com/joho/godotenv"

	rc "github.com/grokify/go-ringcentral-client/office/v1/client"
	ru "github.com/grokify/go-ringcentral-client/office/v1/util"
	"github.com/grokify/goauth"
	"github.com/grokify/goauth/credentials"
)

type Handler struct {
	AppPort        int
	APIClient      *rc.APIClient
	AppCredentials *credentials.OAuth2Credentials
}

func (h *Handler) RingOut(res http.ResponseWriter, req *http.Request) {
	reqUtil := nhu.RequestUtil{Request: req}

	pwdCredentials := credentials.OAuth2Credentials{
		ClientID:        h.AppCredentials.ClientID,
		ClientSecret:    h.AppCredentials.ClientSecret,
		ServerURL:       h.AppCredentials.ServerURL,
		OAuth2Endpoint:  h.AppCredentials.OAuth2Endpoint,
		GrantType:       goauth.GrantTypePassword,
		Username:        reqUtil.QueryParamString("username"),
		Password:        reqUtil.QueryParamString("password"),
		RefreshTokenTTL: int64(-1),
	}

	ringOut := ru.RingOutRequest{
		To:       reqUtil.QueryParamString("to"),
		From:     reqUtil.QueryParamString("from"),
		CallerId: reqUtil.QueryParamString("clid"),
	}

	log.Printf("%v\n", ringOut)

	apiClient, err := ru.NewApiClientPassword(pwdCredentials)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	info, resp, err := apiClient.RingOutApi.MakeRingOutCallNew(
		context.Background(), "~", "~", *ringOut.Body(),
	)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	} else if resp.StatusCode >= 500 {
		res.WriteHeader(http.StatusInternalServerError)
		return
	} else if resp.StatusCode >= 400 {
		res.WriteHeader(http.StatusBadRequest)
		return
	} else if resp.StatusCode >= 300 {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(info)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Header().Set(hum.HeaderContentType, hum.ContentTypeAppJsonUtf8)
	res.Write(bytes)
}

func serveNetHttp(handler Handler) {
	mux := http.NewServeMux()
	mux.HandleFunc("/ringout.asp", http.HandlerFunc(handler.RingOut))
	mux.HandleFunc("/ringout.asp/", http.HandlerFunc(handler.RingOut))

	done := make(chan bool)
	go http.ListenAndServe(fmt.Sprintf(":%v", handler.AppPort), mux)
	log.Printf("Server listening on port %v", handler.AppPort)
	<-done
}

func main() {
	err := godotenv.Load(os.Getenv("ENV_PATH"))
	if err != nil {
		panic(err)
	}

	handler := Handler{
		AppPort: 8080,
		AppCredentials: &credentials.OAuth2Credentials{
			ServerURL:    os.Getenv("RINGCENTRAL_SERVER_URL"),
			ClientID:     os.Getenv("RINGCENTRAL_CLIENT_ID"),
			ClientSecret: os.Getenv("RINGCENTRAL_CLIENT_SECRET"),
		},
	}

	serveNetHttp(handler)
}
