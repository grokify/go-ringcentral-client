package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/antihax/optional"
	"github.com/jessevdk/go-flags"

	"github.com/grokify/go-ringcentral-engage/engagedigital/engagedigital"
	utils "github.com/grokify/go-ringcentral-engage/engagedigital/engagedigitalutil"
	ex "github.com/grokify/go-ringcentral-engage/engagedigital/examples"
	"github.com/grokify/gotilla/time/timeutil"
)

type options struct {
	Site   string `short:"s" long:"site" description:"A site" required:"true"`
	Token  string `short:"t" long:"token" description:"A token" required:"true"`
	Object string `short:"o" long:"obect" description:"An object" required:"true"`
	Action string `short:"a" long:"action" description:"An action (create|update|delete)" required:"true"`
	Name   string `short:"n" long:"name" description:"A tag name" required:"false"`
	Id     string `short:"i" long:"id" description:"A tag id" required:"false"`
}

func main() {
	opts := options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	client := utils.NewApiClient(opts.Site, opts.Token)

	switch strings.ToLower(opts.Object) {
	case "customfield":
		handleCustomField(client, opts)
	case "presencestatus":
		handlePresenceStatus(client, opts)
	case "tag":
		handleTag(client, opts)
	case "replyassistantentry":
		handleReplyAssistantEntry(client, opts)
	case "replyassistantgroup":
		handleReplyAssistantGroup(client, opts)
	case "replyassistantversion":
		handleReplyAssistantVersion(client, opts)
	case "settings":
		handleSettings(client, opts)
	case "timesheet":
		handleTimeSheet(client, opts)
	default:
		log.Fatal(fmt.Sprintf("E_OBJECT_NOT_SUPPORTED [%v]", opts.Object))
	}

	fmt.Println("DONE")
}

func formatRespStatusCodeError(statusCode int) string {
	return fmt.Sprintf("E_API_ERROR [%v]", statusCode)
}

func handleReplyAssistantGroup(client *engagedigital.APIClient, opts options) {
	switch opts.Action {
	case "create":
		dt := time.Now()
		name := "New Reply Assistant Group Name " + dt.Format(time.RFC3339)
		ex.HandleApiResponse(client.ReplyAssistantGroupsApi.CreateReplyAssistantGroup(context.Background(), name, nil))
	case "read":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.ReplyAssistantGroupsApi.GetReplyAssistantGroup(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.ReplyAssistantGroupsApi.GetAllReplyAssistantGroups(context.Background(), nil))
		}
	case "update":
		if len(opts.Id) > 0 {
			dt := time.Now()
			name := "Updated Reply Assistant Group Name " + dt.Format(time.RFC3339)
			apiOpts := &engagedigital.UpdateReplyAssistantGroupOpts{
				Name: optional.NewString(name)}
			ex.HandleApiResponse(client.ReplyAssistantGroupsApi.UpdateReplyAssistantGroup(context.Background(), opts.Id, apiOpts))
		} else {
			log.Fatal("E_REPLYASSISTANTENTRY_UPDATE_ID_NOT_PRESENT")
		}
	case "delete":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.ReplyAssistantGroupsApi.DeleteReplyAssistantGroup(context.Background(), opts.Id))
		} else {
			log.Fatal("E_REPLYASSISTANTGROUP_DELETE_ID_NOT_PRESENT")
		}
	}
}

func handleReplyAssistantVersion(client *engagedigital.APIClient, opts options) {
	switch opts.Action {
	case "create":
		if len(opts.Id) > 0 {
			dt := time.Now()
			body := "New Reply Assistant Version Body " + dt.Format(time.RFC3339)
			apiOpts := &engagedigital.CreateReplyAssistantVersionOpts{}
			ex.HandleApiResponse(client.ReplyAssistantVersionsApi.CreateReplyAssistantVersion(context.Background(), body, opts.Id, apiOpts))
		} else {
			log.Fatal("E_REPLYASSISTANTVERSION_CREATE_NO_ENTRY_ID")
		}
	case "read":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.ReplyAssistantVersionsApi.GetReplyAssistantVersion(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.ReplyAssistantVersionsApi.GetAllReplyAssistantVersions(context.Background(), nil))
		}
	case "update":
		if len(opts.Id) > 0 {
			dt := time.Now()
			body := "Updated Reply Assistant Version Body " + dt.Format(time.RFC3339)
			apiOpts := &engagedigital.UpdateReplyAssistantVersionOpts{
				Body: optional.NewString(body)}
			ex.HandleApiResponse(client.ReplyAssistantVersionsApi.UpdateReplyAssistantVersion(context.Background(), opts.Id, apiOpts))
		} else {
			log.Fatal("E_REPLYASSISTANTENTRY_UPDATE_ID_NOT_PRESENT")
		}
	case "delete":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.ReplyAssistantVersionsApi.DeleteReplyAssistantVersion(context.Background(), opts.Id))
		} else {
			log.Fatal("E_REPLYASSISTANTVERSION_DELETE_ID_NOT_PRESENT")
		}
	}
}

func handleReplyAssistantEntry(client *engagedigital.APIClient, opts options) {
	switch opts.Action {
	case "create":
		dt := time.Now()
		label := "New Reply Assistant Entry Label " + dt.Format(time.RFC3339)
		ex.HandleApiResponse(client.ReplyAssistantEntriesApi.CreateReplyAssistantEntry(context.Background(), label))
	case "read":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.ReplyAssistantEntriesApi.GetReplyAssistantEntry(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.ReplyAssistantEntriesApi.GetAllReplyAssistantEntries(context.Background(), nil))
		}
	case "update":
		if len(opts.Id) > 0 {
			dt := time.Now()
			label := "Updated Reply Assistant Label " + dt.Format(time.RFC3339)
			apiOpts := &engagedigital.UpdateReplyAssistantEntryOpts{
				Label: optional.NewString(label)}
			ex.HandleApiResponse(client.ReplyAssistantEntriesApi.UpdateReplyAssistantEntry(context.Background(), opts.Id, apiOpts))
		} else {
			log.Fatal("E_REPLYASSISTANTENTRY_UPDATE_ID_NOT_PRESENT")
		}
	case "delete":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.ReplyAssistantEntriesApi.DeleteReplyAssistantEntry(context.Background(), opts.Id))
		} else {
			log.Fatal("E_REPLYASSISTANTENTRY_DELETE_ID_NOT_PRESENT")
		}
	}
}

func handleTimeSheet(client *engagedigital.APIClient, opts options) {
	switch opts.Action {
	case "create":
		dt := time.Now()
		label := "TestTimeSheet2-" + dt.Format(timeutil.DT14)
		opts.Id = strings.TrimSpace(opts.Id)
		/*if len(opts.Id) == 0 {
			log.Fatal("E_CREATE_TIMESHEET_NO_SOURCE_ID")
		}*/
		fmt.Println(label)
		apiOpts := &engagedigital.CreateTimeSheetOpts{
			Active:         optional.NewBool(false),
			HolidaysRegion: optional.NewString("us"),
			SourceIds:      optional.NewInterface([]string{opts.Id}),
			MondayHours:    optional.NewString("28800-72000")}
		ex.HandleApiResponse(client.TimeSheetsApi.CreateTimeSheet(context.Background(), label, apiOpts))
	case "read":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.TimeSheetsApi.GetTimeSheet(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.TimeSheetsApi.GetAllTimeSheets(context.Background(), nil))
		}
	case "update":
		if len(opts.Id) == 0 {
			log.Fatal("E_UPDATE_TIMESHEET_NO_ID")
		}
		info, resp, err := client.TimeSheetsApi.GetTimeSheet(context.Background(), opts.Id)
		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode != 200 {
			log.Fatal(fmt.Sprintf("STATUS_CODE [%v]", resp.StatusCode))
		}
		apiOpts := &engagedigital.UpdateTimeSheetOpts{}
		if info.MondayHours == "28800-72000" {
			apiOpts.MondayHours = optional.NewString("0-86400")
		} else {
			apiOpts.MondayHours = optional.NewString("28800-72000")
		}
		ex.HandleApiResponse(client.TimeSheetsApi.UpdateTimeSheet(context.Background(), opts.Id, apiOpts))
	case "delete":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.TimeSheetsApi.DeleteTimeSheet(context.Background(), opts.Id))
		} else {
			log.Fatal("E_DELETE_TIMESHEET_NO_ID")
		}
	}
}

func handleSettings(client *engagedigital.APIClient, opts options) {
	switch opts.Action {
	case "read":
		ex.HandleApiResponse(client.SettingsApi.GetAllSettings(context.Background()))
	case "update":
		info, resp, err := client.SettingsApi.GetAllSettings(context.Background())
		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode != 200 {
			log.Fatal(fmt.Sprintf("Status Code [%v]", resp.StatusCode))
		}
		apiOpts := &engagedigital.UpdateSettingsOpts{}
		if info.Locale == "en" {
			apiOpts.Locale = optional.NewString("fr")
		} else {
			apiOpts.Locale = optional.NewString("en")
		}
		ex.HandleApiResponse(client.SettingsApi.UpdateSettings(context.Background(), apiOpts))
	}
}

func handleCustomField(client *engagedigital.APIClient, opts options) {
	switch opts.Action {
	case "create":
		fmt.Println("I_CREATING_CUSTOM_FIELD")
		opts.Name = strings.TrimSpace(opts.Name)
		if len(opts.Name) == 0 {
			log.Fatal("E_CREATE_CUSTOM_FIELD_NO_NAME")
		}
		ex.HandleApiResponse(client.CustomFieldsApi.CreateCustomField(
			context.Background(), "Intervention", opts.Name, nil))
	case "read":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.CustomFieldsApi.GetCustomField(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.CustomFieldsApi.GetAllCustomFields(context.Background(), nil))
		}
	case "update":
		opts.Id = strings.TrimSpace(opts.Id)
		opts.Name = strings.TrimSpace(opts.Name)
		apiOpts := &engagedigital.UpdateCustomFieldOpts{
			Label: optional.NewString(opts.Name)}
		if len(opts.Name) == 0 || len(opts.Id) == 0 {
			log.Fatal("E_UPDATE_CUSTOM_FIELD_NO_ID_OR_NAME")
		}
		ex.HandleApiResponse(client.CustomFieldsApi.UpdateCustomField(
			context.Background(), opts.Id, apiOpts))
	case "delete":
		opts.Id = strings.TrimSpace(opts.Id)
		if len(opts.Id) == 0 {
			log.Fatal("E_DELETE_CUSTOM_FIELD_NO_ID")
		}
		ex.HandleApiResponse(client.CustomFieldsApi.DeleteCustomField(
			context.Background(), opts.Id))
	}
}

func handlePresenceStatus(client *engagedigital.APIClient, opts options) {
	switch opts.Action {
	case "create":
		fmt.Println("I_CREATING_PRESENCE_STATUS")
		opts.Name = strings.TrimSpace(opts.Name)
		if len(opts.Name) == 0 {
			log.Fatal("E_CREATE_PRESENCE_STATUS_NO_NAME")
		}
		ex.HandleApiResponse(client.PresenceStatusApi.CreatePresenceStatus(
			context.Background(),
			opts.Name))
	case "read":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.PresenceStatusApi.GetPresenceStatus(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.PresenceStatusApi.GetAllPresenceStatus(context.Background(), nil))
		}
	case "update":
		opts.Id = strings.TrimSpace(opts.Id)
		opts.Name = strings.TrimSpace(opts.Name)
		if len(opts.Name) == 0 || len(opts.Id) == 0 {
			log.Fatal("E_UPDATE_PRESENCE_STATUS_NO_ID_OR_NAME")
		}
		ex.HandleApiResponse(client.PresenceStatusApi.UpdatePresenceStatus(
			context.Background(), opts.Id, opts.Name))
	case "delete":
		opts.Id = strings.TrimSpace(opts.Id)
		if len(opts.Id) == 0 {
			log.Fatal("E_DELETE_PRESENCE_STATUS_NO_ID")
		}
		ex.HandleApiResponse(client.PresenceStatusApi.DeletePresenceStatus(
			context.Background(), opts.Id))
	}
}

func handleTag(client *engagedigital.APIClient, opts options) {
	switch opts.Action {
	case "create":
		fmt.Println("I_CREATING_TAG")
		opts.Name = strings.TrimSpace(opts.Name)
		if len(opts.Name) == 0 {
			log.Fatal("E_CREATE_TAG_NO_NAME")
		}
		ex.HandleApiResponse(client.TagsApi.CreateTag(
			context.Background(),
			opts.Name))
	case "read":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.TagsApi.GetTag(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.TagsApi.GetAllTags(context.Background(), nil))
		}
	case "update":
		opts.Id = strings.TrimSpace(opts.Id)
		opts.Name = strings.TrimSpace(opts.Name)
		if len(opts.Name) == 0 || len(opts.Id) == 0 {
			log.Fatal("E_UPDATE_TAG_NO_ID_OR_NAME")
		}
		ex.HandleApiResponse(client.TagsApi.UpdateTag(
			context.Background(), opts.Id, opts.Name))
	case "delete":
		opts.Id = strings.TrimSpace(opts.Id)
		if len(opts.Id) == 0 {
			log.Fatal("E_DELETE_TAG_NO_ID")
		}
		ex.HandleApiResponse(client.TagsApi.DeleteTag(
			context.Background(), opts.Id))
	}
}

/*
func handleTopology(client *engagedigital.APIClient, opts options) {
	logSlug := "TOPOLOGY"
	switch opts.Action {
	case "create":
		all, resp, err := client.TopologiesApi.GetAllTopologies(context.Background(), nil)
		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode != 200 {
			log.Fatal(formatRespStatusCodeError(resp.StatusCode))
		}
		defTop, err := utils.FindDefaultTopology(all.Records)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("I_CREATING_%s", logSlug)
		opts.Name = strings.TrimSpace(opts.Name)
		if len(opts.Name) == 0 {
			log.Fatal(fmt.Sprintf("E_CREATE_%s_NO_NAME", logSlug))
		}
		newTop := engagedigital.Topology{Config: defTop.Config}
		newTopString, err := json.Marshal(newTop)
		if err != nil {
			log.Fatal(err)
		}
		apiOpts := &engagedigital.CreateTopologyOpts{
			JsonConfig: optional.NewString(string(newTopString))}

		ex.HandleApiResponse(client.TopologiesApi.CreateTopology(
			context.Background(), opts.Name, apiOpts))
	case "read":

		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.TopologiesApi.GetTopology(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.TopologiesApi.GetAllTopologies(context.Background(), nil))
		}

			case "update":
				opts.Id = strings.TrimSpace(opts.Id)
				opts.Name = strings.TrimSpace(opts.Name)
				apiOpts := &engagedigital.UpdateCustomFieldOpts{
					Label: optional.NewString(opts.Name)}
				if len(opts.Name) == 0 || len(opts.Id) == 0 {
					log.Fatal(fmt.Sprintf("E_UPDATE_%s_NO_ID_OR_NAME", logSlug))
				}
				ex.HandleApiResponse(client.CustomFieldsApi.UpdateCustomField(
					context.Background(), opts.Id, apiOpts))
			case "delete":
				opts.Id = strings.TrimSpace(opts.Id)
				if len(opts.Id) == 0 {
					log.Fatal(fmt.Sprintf("E_DELETE_%s_NO_ID", logSlug))
				}
				ex.HandleApiResponse(client.CustomFieldsApi.DeleteCustomField(
					context.Background(), opts.Id))

	}
}
*/
