package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/grokify/go-ringcentral-engage/engagevoice/engagevoice"
	"github.com/grokify/go-ringcentral-engage/engagevoice/engagevoiceutil"
	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	EnvPath   string `short:"e" long:"envPath" description:"Environment File Path"`
	Token     string `short:"t" long:"token" description:"Token"`
	AccountID string `short:"a" long:"accountID" description:"AccountID"`
	Object    string `short:"o" long:"object" description:"Object to retrieve" required:"true"`
	ID        string `short:"i" long:"id" description:"id to get"`
}

func main() {
	opts := Options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	err = config.LoadDotEnvFirst(opts.EnvPath, os.Getenv("ENV_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	if len(opts.Token) == 0 {
		opts.Token = os.Getenv("ENGAGE_VOICE_API_TOKEN")
	}
	if len(opts.AccountID) == 0 {
		opts.AccountID = os.Getenv("ENGAGE_VOICE_ACCOUNT_ID")
	}

	clientApis := engagevoiceutil.NewClientAPIs(opts.Token)
	apiClient := clientApis.APIClient

	object := strings.ToLower(opts.Object)
	switch object {
	case "agent":
		info, resp, err := apiClient.AgentsApi.GetAgents(
			context.Background(), opts.AccountID, opts.ID)
		handleErrors(resp, err)
		fmtutil.PrintJSON(info)
	case "agentgroup":
		info, resp, err := apiClient.AgentsApi.GetAgentGroups(
			context.Background(), opts.AccountID)
		handleErrors(resp, err)
		fmtutil.PrintJSON(info)
	case "campaign":
		info, resp, err := apiClient.DialGroupsApi.GetDialGroups(
			context.Background(), opts.AccountID,
		)
		handleErrors(resp, err)
		fmtutil.PrintJSON(info)
		dialGroupIDs := []int64{}
		for _, dg := range info {
			dialGroupIDs = append(dialGroupIDs, dg.DialGroupId)
		}
		fmtutil.PrintJSON(dialGroupIDs)
		for _, dialGroupID := range dialGroupIDs {
			info, resp, err := apiClient.CampaignsApi.GetDialGroupCampaigns(
				context.Background(),
				opts.AccountID,
				strconv.Itoa(int(dialGroupID)),
			)
			handleErrors(resp, err)
			fmtutil.PrintJSON(info)
			for _, campaign := range info {
				campaignID := campaign.CampaignId
				fmt.Printf("CAMPAIGN [%v]\n", campaignID)
				uploadLeads(apiClient, opts.AccountID, strconv.Itoa(int(campaignID)))

			}
		}
	case "country":
		info, resp, err := apiClient.CountriesApi.GetAvailableCountries(
			context.Background(), opts.AccountID)
		handleErrors(resp, err)
		fmtutil.PrintJSON(info)
	}

	if 1 == 0 {
		info, resp, err := apiClient.DialGroupsApi.GetDialGroups(
			context.Background(), opts.AccountID,
		)
		handleErrors(resp, err)
		fmtutil.PrintJSON(info)
		dialGroupIDs := []int64{}
		for _, dg := range info {
			dialGroupIDs = append(dialGroupIDs, dg.DialGroupId)
		}
		fmtutil.PrintJSON(dialGroupIDs)
		for _, dialGroupID := range dialGroupIDs {
			info, resp, err := apiClient.CampaignsApi.GetDialGroupCampaigns(
				context.Background(),
				opts.AccountID,
				strconv.Itoa(int(dialGroupID)),
			)
			handleErrors(resp, err)
			fmtutil.PrintJSON(info)
			for _, campaign := range info {
				campaignID := campaign.CampaignId
				fmt.Printf("CAMPAIGN [%v]\n", campaignID)
				uploadLeads(apiClient, opts.AccountID, strconv.Itoa(int(campaignID)))

			}
		}
	}
	fmt.Println("DONE")
}

func uploadLeads(client *engagevoice.APIClient, accountId, campaignId string) {
	dt := time.Now()
	leadsRequest := engagevoice.UploadLeadsRequest{
		Description:       "Test from Go " + dt.Format(time.RFC3339),
		DialPriority:      "IMMEDIATE",
		DuplicateHandling: "RETAIN_ALL",
		ListState:         "ACTIVE",
		TimeZoneOption:    "NPA_NXX",
		UploadLeads: []engagevoice.Lead{
			{
				LeadPhone: "6505550100",
				ExternId:  1,
				FirstName: "Jon",
				LastName:  "Snow",
				ExtendedLeadData: engagevoice.ExtendedLeadData{
					Important:  "Priority 1: needs help right away",
					Interested: true,
				},
			},
		},
	}
	info, resp, err := client.LeadsApi.UploadLeads(
		context.Background(),
		accountId,
		campaignId,
		leadsRequest,
	)
	handleErrors(resp, err)
	fmtutil.PrintJSON(info)
}

func handleErrors(resp *http.Response, err error) {
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode >= 300 {
		log.Fatal("E_STATUS_CODE_GTE_300")
	}
}
