package engagevoiceutil

import (
	"context"
	"fmt"
	"strconv"

	engagevoice "github.com/grokify/go-ringcentral/engagevoice/v1/client"
)

func GetAllCampaigns(ctx context.Context, apiClient *engagevoice.APIClient, accountID string) ([]engagevoice.Campaign, error) {
	campaigns := []engagevoice.Campaign{}
	info, resp, err := apiClient.DialGroupsApi.GetDialGroups(
		ctx, accountID,
	)
	if err != nil {
		return campaigns, err
	} else if resp.StatusCode >= 300 {
		return campaigns, fmt.Errorf("GetDialGroups Response Status Code [%v]", resp.StatusCode)
	}

	dialGroupIDs := []int64{}
	for _, dg := range info {
		dialGroupIDs = append(dialGroupIDs, dg.DialGroupId)
	}

	for _, dialGroupID := range dialGroupIDs {
		info, resp, err := apiClient.DialGroupsApi.GetCampaigns(
			ctx,
			accountID,
			strconv.Itoa(int(dialGroupID)),
		)
		if err != nil {
			return campaigns, err
		} else if resp.StatusCode >= 300 {
			return campaigns, fmt.Errorf("GetDialGroups Response Status Code [%v]", resp.StatusCode)
		}
		campaigns = append(campaigns, info...)
	}
	return campaigns, nil
}
