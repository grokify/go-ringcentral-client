package clientutil

import (
	"os"
	"time"

	"github.com/joho/godotenv"

	rc "github.com/grokify/go-ringcentral/client"
	ro "github.com/grokify/oauth2more/ringcentral"
)

func LoadEnv() error {
	envPaths := []string{}
	if len(os.Getenv("ENV_PATH")) > 0 {
		envPaths = append(envPaths, os.Getenv("ENV_PATH"))
	}
	return godotenv.Load(envPaths...)
}

func NewApiClient(app ro.ApplicationCredentials, user ro.UserCredentials) (*rc.APIClient, error) {
	httpClient, err := ro.NewClientPassword(app, user)
	if err != nil {
		return nil, err
	}
	apiConfig := rc.NewConfiguration()
	apiConfig.BasePath = app.ServerURL
	apiConfig.HTTPClient = httpClient
	apiClient := rc.NewAPIClient(apiConfig)
	return apiClient, nil
}

func NewApiClientEnv() (*rc.APIClient, error) {
	return NewApiClient(
		ro.ApplicationCredentials{
			ServerURL:    os.Getenv("RINGCENTRAL_SERVER_URL"),
			ClientID:     os.Getenv("RINGCENTRAL_CLIENT_ID"),
			ClientSecret: os.Getenv("RINGCENTRAL_CLIENT_SECRET")},
		ro.UserCredentials{
			Username:  os.Getenv("RINGCENTRAL_USERNAME"),
			Extension: os.Getenv("RINGCENTRAL_EXTENSION"),
			Password:  os.Getenv("RINGCENTRAL_PASSWORD")})
}

type Event struct {
	UUID           string    `json:"uuid,omitempty"`
	Event          string    `json:"event,omitempty"`
	Timestamp      time.Time `json:"timestamp,omitempty"`
	SubscriptionId string    `json:"subscriptionId,omitempty"`
	OwnerId        string    `json:"ownerId,omitempty"`
	Body           EventBody `json:"body,omitempty"`
}

type EventBody struct {
	ExpiresIn            int       `json:"expiresIn,omitempty"`
	CreationTime         time.Time `json:"creationTime,omitempty"`
	Direction            string    `json:"direction,omitempty"`
	LastModificationTime time.Time `json:"lastModificationTime,omitempty"`
	Subject              string    `json:"subject,omitempty"`
	Type                 string    `json:"type,omitempty"`
}
