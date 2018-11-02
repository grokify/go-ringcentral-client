package glipgroups

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	hum "github.com/grokify/gotilla/net/httputilmore"
	ro "github.com/grokify/oauth2more/ringcentral"
)

const (
	GroupTypeTeam = "Team"
)

type GroupsSet struct {
	GroupsMap map[string]Group
}

func (set *GroupsSet) AddGroups(groups []Group) {
	for _, grp := range groups {
		set.GroupsMap[grp.ID] = grp
	}
}

func (set *GroupsSet) FindGroupsByName(groupName string) []Group {
	groupName = strings.TrimSpace(groupName)
	groups := []Group{}
	for _, group := range set.GroupsMap {
		if groupName == group.Name {
			groups = append(groups, group)
		}
	}
	return groups
}

func (set *GroupsSet) FindGroupByName(groupName string) (Group, error) {
	groups := set.FindGroupsByName(groupName)
	if len(groups) != 1 {
		return Group{}, fmt.Errorf("Found [%v] groups for name [%v]", len(groups), groupName)
	}
	return groups[0], nil
}

type Group struct {
	ID               string    `json:"id,omitempty"`
	Name             string    `json:"name,omitempty"`
	Description      string    `json:"description,omitempty"`
	CreationTime     time.Time `json:"creationTime,omitempty"`
	LastModifiedTime time.Time `json:"lastModifiedTime,omitempty"`
	Members          []string  `json:"members,omitempty"`
}

func NewGroupsSetApiRequest(client *http.Client, serverUrl string, groupType string) (GroupsSet, error) {
	set := GroupsSet{GroupsMap: map[string]Group{}}

	query := url.Values{}
	query.Add("recordCount", "250")

	if len(groupType) > 0 {
		query.Add("type", groupType)
	}

	for {
		groupsURL := ro.BuildURL(serverUrl, "/glip/groups", true, query)
		resp, err := client.Get(groupsURL)
		if err != nil {
			return set, err
		}

		groupsResp, err := GetGroupsResponseFromHTTPResponse(resp)
		if err != nil {
			return set, err
		}
		set.AddGroups(groupsResp.Records)

		if len(groupsResp.Navigation.PrevPageToken) > 0 {
			query.Add("pageToken", groupsResp.Navigation.PrevPageToken)
		} else {
			break
		}
	}
	return set, nil
}

type GetGroupsResponse struct {
	Records    []Group    `json:"records,omitempty"`
	Navigation Navigation `json:"navigation,omitempty"`
}

func GetGroupsResponseFromHTTPResponse(resp *http.Response) (GetGroupsResponse, error) {
	bytes, err := hum.ResponseBody(resp)
	if err != nil {
		return GetGroupsResponse{}, err
	}
	var apiResp GetGroupsResponse
	err = json.Unmarshal(bytes, &apiResp)
	return apiResp, err
}

type Navigation struct {
	PrevPageToken string `json:"prevPageToken,omitempty"`
	NextPageToken string `json:"nextPageToken,omitempty"`
}
