package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	hum "github.com/grokify/gotilla/net/httputilmore"
	nhu "github.com/grokify/gotilla/net/nethttputil"
	"github.com/joho/godotenv"

	rc "github.com/grokify/go-ringcentral/client"
	cu "github.com/grokify/go-ringcentral/clientutil"
	ro "github.com/grokify/oauth2more/ringcentral"
)

type Handler struct {
	AppPort        int
	APIClient      *rc.APIClient
	AppCredentials *ro.ApplicationCredentials
}

func (h *Handler) RingOut(res http.ResponseWriter, req *http.Request) {
	ru := nhu.RequestUtil{Request: req}

	pwdCredentials := ro.PasswordCredentials{
		Username:        ru.QueryParamString("username"),
		Extension:       ru.QueryParamString("ext"),
		Password:        ru.QueryParamString("password"),
		RefreshTokenTTL: int64(-1),
	}

	ringOut := cu.RingOutRequest{
		To:       ru.QueryParamString("to"),
		From:     ru.QueryParamString("from"),
		CallerId: ru.QueryParamString("clid"),
	}

	log.Printf("%v\n", ringOut)

	apiClient, err := cu.NewApiClientPassword(*h.AppCredentials, pwdCredentials)
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

	res.Header().Set(hum.HeaderContentType, hum.HeaderContentTypeValueJSONUTF8)
	res.Write(bytes)
}

func serveNetHttp(handler Handler) {
	mux := http.NewServeMux()
	mux.HandleFunc("/ringout.asp", http.HandlerFunc(handler.RingOut))
	mux.HandleFunc("/ringout.asp/", http.HandlerFunc(handler.RingOut))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", handler.AppPort), mux))
}

func main() {
	err := godotenv.Load(os.Getenv("ENV_PATH"))
	if err != nil {
		panic(err)
	}

	handler := Handler{
		AppPort: 8080,
		AppCredentials: &ro.ApplicationCredentials{
			ServerURL:    os.Getenv("RINGCENTRAL_SERVER_URL"),
			ClientID:     os.Getenv("RINGCENTRAL_CLIENT_ID"),
			ClientSecret: os.Getenv("RINGCENTRAL_CLIENT_SECRET"),
		},
	}

	serveNetHttp(handler)
}
