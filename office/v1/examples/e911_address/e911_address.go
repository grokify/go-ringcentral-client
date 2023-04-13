package main

import (
	"context"
	"fmt"
	"os"

	"github.com/grokify/goauth"
	"github.com/grokify/mogo/config"
	"github.com/grokify/mogo/fmt/fmtutil"

	rc "github.com/grokify/go-ringcentral-client/office/v1/client"
	ru "github.com/grokify/go-ringcentral-client/office/v1/util"
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
	_, err := config.LoadDotEnv([]string{os.Getenv("ENV_PATH"), "./.env"}, 1)
	if err != nil {
		panic(err)
	}

	apiClient, err := ru.NewApiClientPassword(
		goauth.NewCredentialsOAuth2Env("RINGCENTRAL_"))
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
			EmergencyServiceAddress: rc.EmergencyAddressInfoRequest{
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
