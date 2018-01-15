package main

import (
	"context"
	"fmt"
	"os"

	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/joho/godotenv"

	rc "github.com/grokify/go-ringcentral/client"
	rcu "github.com/grokify/go-ringcentral/clientutil"
)

func loadEnv() error {
	envPaths := []string{}
	if len(os.Getenv("ENV_PATH")) > 0 {
		envPaths = append(envPaths, os.Getenv("ENV_PATH"))
	}
	return godotenv.Load(envPaths...)
}

func testRetrieveUsers() {

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

func (su *ScimApiUtil) ListScimUsers(params map[string]interface{}) (rc.GetUserListResponse, error) {
	info, resp, err := su.ApiClient.SCIMApi.ListScimUsers(su.Context, params)
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
	apiClient, err := rcu.NewApiClientEnv()
	if err != nil {
		panic(err)
	}

	scimUtil := ScimApiUtil{
		ApiClient: apiClient,
		Context:   context.Background(),
	}

	// Retrieve a list of users
	info, err := scimUtil.ListScimUsers(map[string]interface{}{})
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
	fmt.Println("DONE")
}
