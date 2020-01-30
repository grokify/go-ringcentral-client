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

	engagevoice "github.com/grokify/go-ringcentral/engagevoice/v1/client"
	engagevoiceutil "github.com/grokify/go-ringcentral/engagevoice/v1/util"
	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	EnvPath        string `short:"e" long:"envPath" description:"Environment File Path"`
	Token          string `short:"t" long:"token" description:"Token"`
	AccountID      string `short:"a" long:"accountID" description:"AccountID"`
	Object         string `short:"o" long:"object" description:"Object to retrieve" required:"true"`
	ID             string `short:"i" long:"id" description:"id to get"`
	AccountIdInt64 int64
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
		accountIdInt, err := strconv.Atoi(opts.AccountID)
		if err != nil {
			log.Fatal(err)
		}
		opts.AccountIdInt64 = int64(accountIdInt)
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
		campaigns, err := engagevoiceutil.GetAllCampaigns(
			context.Background(), apiClient, opts.AccountID)
		if err != nil {
			log.Fatal(err)
		}
		fmtutil.PrintJSON(campaigns)
	case "campaignclearcache":
		campaigns, err := engagevoiceutil.GetAllCampaigns(
			context.Background(), apiClient, opts.AccountID)
		if err != nil {
			log.Fatal(err)
		}
		if len(campaigns) == 0 {
			log.Fatal("E_NO_CAMPAIGNS")
		}
		for _, campaign := range campaigns {
			fmt.Println("CLEARING...")
			fmtutil.PrintJSON(campaign)
			resp, err := apiClient.DialGroupsApi.ClearCampaignCache(
				context.Background(),
				opts.AccountIdInt64,
				campaign.DialGroup.Id,
				campaign.CampaignId,
				map[string]interface{}{},
			)
			handleErrors(resp, err)
			fmt.Printf("SUCCESS [%v]\n", resp.StatusCode)
			break
		}
	case "leadstate":
		info, resp, err := apiClient.CampaignLeadsApi.GetLeadStates(
			context.Background(), opts.AccountID)
		handleErrors(resp, err)
		fmtutil.PrintJSON(info)
	case "country":
		info, resp, err := apiClient.CountriesApi.GetAvailableCountries(
			context.Background(), opts.AccountID)
		handleErrors(resp, err)
		fmtutil.PrintJSON(info)
	case "dialgroup":
		info, resp, err := apiClient.DialGroupsApi.GetDialGroups(
			context.Background(), opts.AccountID)
		handleErrors(resp, err)
		fmtutil.PrintJSON(info)
	case "systemdispositions":
		info, resp, err := apiClient.CampaignLeadsApi.GetSystemDispositions(
			context.Background(), opts.AccountID)
		handleErrors(resp, err)
		fmtutil.PrintJSON(info)
	case "leadsearch":
		//qry := engagevoice.CampaignLeadSearchCriteria{FirstName: "Jon"}
		qry := engagevoice.CampaignLeadSearchCriteria{}
		info, resp, err := apiClient.CampaignLeadsApi.SearchLeads(
			context.Background(), opts.AccountID,
			qry, nil,
		)
		handleErrors(resp, err)
		fmtutil.PrintJSON(info)
	case "token":
		info, resp, err := apiClient.AuthApi.GetTokens(context.Background())
		handleErrors(resp, err)
		fmtutil.PrintJSON(info)
	case "user":
		info, resp, err := apiClient.UsersApi.GetUsers(context.Background())
		handleErrors(resp, err)
		fmtutil.PrintJSON(info)
	default:
		log.Fatal(fmt.Sprintf("Object not found [%v]", object))
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
			info, resp, err := apiClient.DialGroupsApi.GetCampaigns(
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
		UploadLeads: []engagevoice.CampaignLead{
			{
				LeadPhone: "6505550100",
				ExternId:  "1",
				FirstName: "Jon",
				LastName:  "Snow",
				/*ExtendedLeadData: engagevoice.ExtendedLeadData{
					Important:  "Priority 1: needs help right away",
					Interested: true,
				},*/
			},
		},
	}
	info, resp, err := client.CampaignsApi.UploadLeads(
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
