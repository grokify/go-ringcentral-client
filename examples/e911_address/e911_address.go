package main

import (
	"context"
	"fmt"
	"os"

	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"

	rc "github.com/grokify/go-ringcentral/client"
	ru "github.com/grokify/go-ringcentral/clientutil"
	ro "github.com/grokify/oauth2more/ringcentral"
)

func getAccountDeviceList(apiClient *rc.APIClient, accountId string) (rc.GetAccountDevicesResponse, error) {
	info, resp, err := apiClient.AccountProvisioningApi.ListAccountDevices(context.Background(), accountId)
	if err != nil {
		return info, err
	}
	if resp.StatusCode >= 300 {
		return info, fmt.Errorf("API Status Code %v", resp.StatusCode)
	}
	return info, err
}

func getPrintDeviceE911Address(apiClient *rc.APIClient, deviceId string) {
	device, resp, err := apiClient.AccountProvisioningApi.LoadAccountDevice(
		context.Background(), "~", deviceId)
	if err != nil {
		panic(err)
	} else if resp.StatusCode >= 300 {
		panic("BAD_STATUS")
	}
	fmtutil.PrintJSON(device)
}

func main() {
	err := config.LoadDotEnvSkipEmpty(os.Getenv("ENV_PATH"), "./.env")
	if err != nil {
		panic(err)
	}

	apiClient, err := ru.NewApiClientPassword(
		ro.ApplicationCredentials{
			ServerURL:    os.Getenv("RINGCENTRAL_SERVER_URL"),
			ClientID:     os.Getenv("RINGCENTRAL_CLIENT_ID"),
			ClientSecret: os.Getenv("RINGCENTRAL_CLIENT_SECRET")},
		ro.PasswordCredentials{
			Username:  os.Getenv("RINGCENTRAL_USERNAME"),
			Extension: os.Getenv("RINGCENTRAL_EXTENSION"),
			Password:  os.Getenv("RINGCENTRAL_PASSWORD")})
	if err != nil {
		panic(err)
	}

	deviceList, err := getAccountDeviceList(apiClient, "~")
	if err != nil {
		panic(err)
	}
	fmtutil.PrintJSON(deviceList)

	// UpdateDevice
	if 1 == 0 {
		deviceId := "12345678"
		getPrintDeviceE911Address(apiClient, deviceId)

		body := rc.AccountDeviceUpdate{
			EmergencyServiceAddress: &rc.EmergencyAddressInfoRequest{
				CustomerName: "RingForce",
				Street:       "20 DAVIS DR",
				Street2:      "",
				City:         "BELMONT",
				State:        "CA",
				Zip:          "94402-3002",
				Country:      "US",
			},
		}
		fmtutil.PrintJSON(body)

		info, resp, err := apiClient.AccountProvisioningApi.UpdateDevice(
			context.Background(), "~", deviceId, body,
		)
		if err != nil {
			panic(err)
		} else if resp.StatusCode >= 300 {
			panic(fmt.Sprintf("Status %v", resp.StatusCode))
		}
		fmtutil.PrintJSON(info)

		getPrintDeviceE911Address(apiClient, deviceId)
	}

	fmt.Println("DONE")
}
