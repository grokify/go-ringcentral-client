package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/antihax/optional"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/jessevdk/go-flags"

	"github.com/grokify/go-ringcentral-engage/engagedigital/engagedigital"
	utils "github.com/grokify/go-ringcentral-engage/engagedigital/engagedigitalutil"
)

type options struct {
	Site   string `short:"s" long:"site" description:"A site" required:"true"`
	Token  string `short:"t" long:"token" description:"A token" required:"true"`
	Update []bool `short:"u" long:"update" description:"Update team"`
	//Object string `short:"o" long:"object" description:"An object" required:"true"`
	//Id     string `short:"i" long:"id" description:"An object id" required:"false"`
}

func getUsers(client *engagedigital.APIClient) (engagedigital.GetAllUsersResponse, error) {
	apiResp := engagedigital.GetAllUsersResponse{}
	info, resp, err := client.UsersApi.GetAllUsers(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return apiResp, err
	} else if resp.StatusCode != 200 {
		return apiResp, fmt.Errorf("API Response [%v]", resp.StatusCode)
	}
	return info, nil
}

func getTeams(client *engagedigital.APIClient) (engagedigital.GetAllTeamsResponse, error) {
	apiResp := engagedigital.GetAllTeamsResponse{}
	info, resp, err := client.TeamsApi.GetAllTeams(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return apiResp, err
	} else if resp.StatusCode != 200 {
		return apiResp, fmt.Errorf("API Response [%v]", resp.StatusCode)
	}
	return info, nil
}

func main() {
	opts := options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	client := utils.NewApiClient(opts.Site, opts.Token)

	usersApiResp, err := getUsers(client)
	if err != nil {
		log.Fatal(err)
	}

	t := time.Now()
	teamName := fmt.Sprintf("TestTeam:%s", t.Format(time.RFC3339))

	leaderIds := []string{}
	userIds := []string{}

	for i, user := range usersApiResp.Records {
		if i == 0 {
			leaderIds = append(leaderIds, user.Id)
		} else {
			userIds = append(userIds, user.Id)
		}
	}

	if len(opts.Update) == 0 {
		apiOpts := engagedigital.CreateTeamOpts{
			Name:      optional.NewString(teamName),
			LeaderIds: optional.NewInterface(leaderIds)}
		if len(userIds) > 0 {
			apiOpts.UserIds = optional.NewInterface(userIds)
		}

		info, resp, err := client.TeamsApi.CreateTeam(context.Background(), &apiOpts)
		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode != 200 {
			log.Fatal(resp.StatusCode)
		}
		fmt.Println("CREATE_TEAM")
		fmtutil.PrintJSON(info)
	} else {
		teamsApiResp, err := getTeams(client)
		if err != nil {
			log.Fatal(err)
		}
		if len(teamsApiResp.Records) == 0 {
			log.Fatal("E_NO_TEAMS_TO_UPDATE")
		}
		team1 := teamsApiResp.Records[0]
		teamId := team1.Id
		apiOpts := engagedigital.UpdateTeamOpts{
			Name:      optional.NewString(teamName),
			LeaderIds: optional.NewInterface(leaderIds)}
		if len(userIds) > 0 {
			apiOpts.UserIds = optional.NewInterface(userIds)
		}

		info, resp, err := client.TeamsApi.UpdateTeam(context.Background(), teamId, &apiOpts)
		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode != 200 {
			log.Fatal(resp.StatusCode)
		}
		fmt.Println("UPDATE_TEAM")
		fmtutil.PrintJSON(info)
	}

	fmt.Println("DONE")
}
