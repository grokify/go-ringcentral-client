package main

import (
	"context"
	"fmt"
	"os"

	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/joho/godotenv"

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

func main() {
	err := loadEnv()
	if err != nil {
		panic(err)
	}
	apiClient, err := rcu.NewApiClientEnv()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	// Retrieve a list of users
	info, resp, err := apiClient.SCIMApi.ListScimUsers(ctx, map[string]interface{}{})
	if err != nil {
		panic(err)
	} else if resp.StatusCode >= 300 {
		panic(fmt.Errorf("API Status %v", resp.StatusCode))
	}
	fmtutil.PrintJSON(info)

	for _, user := range info.Resources {
		fmt.Println(user.Id)

		// Retrieve an individual user
		uinfo, resp, err := apiClient.SCIMApi.GetUserById(ctx, user.Id)
		if err != nil {
			panic(err)
		} else if resp.StatusCode >= 300 {
			panic(fmt.Errorf("API Status %v", resp.StatusCode))
		}
		fmtutil.PrintJSON(uinfo)
	}
	fmt.Println("DONE")
}
