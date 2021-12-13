package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/joho/godotenv"

	rc "github.com/grokify/go-ringcentral-client/office/v1/client"
	rcu "github.com/grokify/go-ringcentral-client/office/v1/util"
)

const (
	envPrefix = "RINGCENTRAL_"
)

func loadEnv() error {
	envPaths := []string{}
	if len(os.Getenv("ENV_PATH")) > 0 {
		envPaths = append(envPaths, os.Getenv("ENV_PATH"))
	}
	return godotenv.Load(envPaths...)
}

type ScimApiUtil struct {
	ApiClient *rc.APIClient
	Context   context.Context
}

func (su *ScimApiUtil) GetSCIMUserById(userId string) (rc.UserInfo, error) {
	uinfo, resp, err := su.ApiClient.SCIMApi.GetUserById(su.Context, userId)
	if err != nil {
		return uinfo, err
	} else if resp.StatusCode >= 300 {
		return uinfo, fmt.Errorf("API Status %v", resp.StatusCode)
	}
	return uinfo, nil
}

func (su *ScimApiUtil) ListScimUsers(params rc.ListUsersOpts) (rc.GetUserListResponse, error) {
	info, resp, err := su.ApiClient.SCIMApi.ListUsers(su.Context, &params)
	if err != nil {
		return info, err
	} else if resp.StatusCode >= 300 {
		return info, fmt.Errorf("API Status %v", resp.StatusCode)
	}
	return info, nil
}

func (su *ScimApiUtil) LoadExtensionInfo(accountId, extensionId string) (rc.GetExtensionInfoResponse, error) {
	info, resp, err := su.ApiClient.UserSettingsApi.LoadExtensionInfo(su.Context, accountId, extensionId)
	if err != nil {
		return info, err
	} else if resp.StatusCode >= 300 {
		return info, fmt.Errorf("API Status %v", resp.StatusCode)
	}
	return info, nil
}

func main() {
	err := loadEnv()
	if err != nil {
		panic(err)
	}
	apiClient, err := rcu.NewApiClientPasswordEnv(envPrefix)
	if err != nil {
		panic(err)
	}

	scimUtil := ScimApiUtil{
		ApiClient: apiClient,
		Context:   context.Background(),
	}

	userId := "1"

	if 1 == 1 {
		extInfo, err := scimUtil.GetSCIMUserById(userId)
		if err != nil {
			log.Fatal(err)
		}
		fmtutil.PrintJSON(extInfo)
	}

	if 1 == 0 {
		extInfo, err := scimUtil.LoadExtensionInfo("~", userId)
		if err != nil {
			log.Fatal(err)
		}
		fmtutil.PrintJSON(extInfo)
	}

	params := rc.UserUpdateRequest{
		Schemas:  []string{"urn:ietf:params:scim:schemas:core:2.0:User"},
		UserName: "JonSnow",
		Emails: []rc.EmailInfo{
			{
				Value: "jon.snow@winterfell.gov",
			},
		},
		Active: true,
		Addresses: []rc.AddressInfo{
			{
				Locality: "Winterfell",
				Region:   "The North",
				Country:  "Seven Kingdoms",
			},
		},
	}

	if 1 == 0 {
		info, res, err := apiClient.SCIMApi.UpdateUser(scimUtil.Context, userId, params)
		if err != nil {
			panic(err)
		} else if res.StatusCode >= 300 {
			panic(fmt.Sprintf("API Response %v", res.StatusCode))
		}
		fmtutil.PrintJSON(info)
	}
	if 1 == 0 {
		// Retrieve a list of users
		info, err := scimUtil.ListScimUsers(rc.ListUsersOpts{})
		if err != nil {
			panic(err)
		}
		fmtutil.PrintJSON(info)

		for _, user := range info.Resources {
			fmt.Println(user.Id)

			uinfo, err := scimUtil.GetSCIMUserById(user.Id)
			if err != nil {
				panic(err)
			}
			fmtutil.PrintJSON(uinfo)
		}
	}
	fmt.Println("DONE")
}
