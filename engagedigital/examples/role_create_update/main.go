package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/antihax/optional"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/jessevdk/go-flags"
	"github.com/pkg/errors"

	"github.com/grokify/go-ringcentral-engage/engagedigital/engagedigital"
	utils "github.com/grokify/go-ringcentral-engage/engagedigital/engagedigitalutil"
)

type options struct {
	Site   string `short:"s" long:"site" description:"A site" required:"true"`
	Token  string `short:"t" long:"token" description:"A token" required:"true"`
	Action string `short:"a" long:"action" description:"A action(create|update)" required:"true"`
}

func GetRole(client *engagedigital.APIClient, roleName string) (engagedigital.Role, error) {
	info, resp, err := client.RolesApi.GetAllRoles(context.Background(), nil)
	if err != nil {
		err = errors.Wrap(err, "E_GET_ALL_ROLES")
		return engagedigital.Role{}, err
	} else if resp.StatusCode != 200 {
		return engagedigital.Role{}, fmt.Errorf("E_STATUS_CODE [%v]", resp.StatusCode)
	}
	for _, role := range info.Records {
		if role.Label == roleName {
			return role, nil
		}
	}
	return engagedigital.Role{}, fmt.Errorf("E_ROLE_NOT_FOUND [%v]", roleName)
}

func main() {
	opts := options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	client := utils.NewApiClient(opts.Site, opts.Token)

	switch opts.Action {
	case "create":
		dt := time.Now()
		roleLabel := "TestRole+" + dt.Format(time.RFC3339)
		apiOpts := &engagedigital.CreateRoleOpts{
			UseEmoji: optional.NewBool(true)}
		info, resp, err := client.RolesApi.CreateRole(
			context.Background(), roleLabel, apiOpts)
		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode != 200 {
			log.Fatal(resp.StatusCode)
		}
		fmtutil.PrintJSON(info)
	case "update":
		roleName := "Agent"
		role, err := GetRole(client, roleName)
		if err != nil {
			log.Fatal(err)
		}

		apiOpts := &engagedigital.UpdateRoleOpts{}
		if role.UseEmoji {
			apiOpts.UseEmoji = optional.NewBool(false)
		} else {
			apiOpts.UseEmoji = optional.NewBool(true)
		}

		info, resp, err := client.RolesApi.UpdateRole(
			context.Background(), role.Id, apiOpts)
		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode != 200 {
			log.Fatal(resp.StatusCode)
		}
		fmtutil.PrintJSON(info)
	}
	fmt.Println("DONE")
}
