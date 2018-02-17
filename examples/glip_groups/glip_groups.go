package main

import (
	"context"
	"fmt"

	"github.com/grokify/gotilla/fmt/fmtutil"

	rcu "github.com/grokify/go-ringcentral/clientutil"
)

func main() {
	err := rcu.LoadEnv()
	if err != nil {
		panic(err)
	}

	apiClient, err := rcu.NewApiClientPasswordEnv()
	if err != nil {
		panic(err)
	}

	info, resp, err := apiClient.GlipApi.LoadGroupList(
		context.Background(), map[string]interface{}{},
	)

	if err != nil {
		panic(err)
	} else if resp.StatusCode >= 300 {
		panic(fmt.Errorf("API Status %v", resp.StatusCode))
	}
	fmtutil.PrintJSON(info)

	fmt.Println("DONE")
}
