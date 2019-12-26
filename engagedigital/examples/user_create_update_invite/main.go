package main

import (
	"context"
	"fmt"
	"log"

	"github.com/antihax/optional"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/gotilla/net/smtputil"
	"github.com/jessevdk/go-flags"
	"github.com/pkg/errors"

	"github.com/grokify/go-ringcentral-engage/engagedigital/engagedigital"
	utils "github.com/grokify/go-ringcentral-engage/engagedigital/engagedigitalutil"
)

type options struct {
	Site   string `short:"s" long:"site" description:"A site" required:"true"`
	Token  string `short:"t" long:"token" description:"A token" required:"true"`
	Action string `short:"a" long:"action" description:"A action (create|update|invite)" required:"true"`
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

	roleName := "Agent"
	userEmail, err := smtputil.GmailAddressPlusTime("bb8.ringforce")
	if err != nil {
		log.Fatal(err)
	}

	switch opts.Action {
	case "create":
		//userEmail = "joe@example.com"
		userFirstName := "Test"
		userLastName := "User"

		role, err := GetRole(client, roleName)
		if err != nil {
			log.Fatal(err)
		}

		info, resp, err := client.UsersApi.CreateUser(
			context.Background(),
			userEmail,
			userFirstName,
			userLastName,
			"MyPasswordU*OQ@#4uo",
			role.Id, nil)
		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode != 200 {
			log.Fatal(resp.StatusCode)
		}
		fmtutil.PrintJSON(info)
	case "update":
		users, resp, err := client.UsersApi.GetAllUsers(context.Background(), nil)
		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode != 200 {
			log.Fatal(resp.StatusCode)
		}

		user := users.Records[0]

		apiOpts := &engagedigital.UpdateUserOpts{}
		if user.Locale == "en" {
			apiOpts.Locale = optional.NewString("fr")
		} else {
			apiOpts.Locale = optional.NewString("en")
		}

		info, resp, err := client.UsersApi.UpdateUser(
			context.Background(), user.Id, apiOpts)
		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode != 200 {
			log.Fatal(resp.StatusCode)
		}
		fmtutil.PrintJSON(info)
	case "invite":
		userEmail = "joe@example.com"
		userFirstName := "Test"
		userLastName := "User"

		role, err := GetRole(client, roleName)
		if err != nil {
			log.Fatal(err)
		}

		info, resp, err := client.UsersApi.InviteUser(
			context.Background(),
			userEmail,
			userFirstName,
			userLastName,
			role.Id, nil)
		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode != 200 {
			log.Fatal(resp.StatusCode)
		}
		fmtutil.PrintJSON(info)
	}

	fmt.Println("DONE")
}
