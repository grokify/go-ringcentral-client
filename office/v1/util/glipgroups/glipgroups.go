package glipgroups

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/grokify/mogo/net/urlutil"
)

const (
	GroupTypeTeam         = "Team"
	groupNameIdJoin       = " - "
	ApiPathListUserGroups = "/restapi/v1.0/glip/groups"
)

type GroupsSet struct {
	GroupsMap map[string]Group // ID to Group Map
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

func (set *GroupsSet) FindGroupsByNameLower(groupName string) []Group {
	groupName = strings.ToLower(strings.TrimSpace(groupName))
	groups := []Group{}
	for _, group := range set.GroupsMap {
		if groupName == strings.ToLower(group.Name) {
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

func (set *GroupsSet) GroupNamesSorted(withIds bool) []string {
	names := []string{}
	for _, group := range set.GroupsMap {
		name := group.Name
		if withIds {
			name += groupNameIdJoin + group.ID
		}
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func (set *GroupsSet) GroupsSorted() []Group {
	names := set.GroupNamesSorted(true)
	groups := []Group{}
	for _, name := range names {
		parts := strings.Split(name, groupNameIdJoin)
		if len(parts) > 0 {
			id := parts[len(parts)-1]
			if g, ok := set.GroupsMap[id]; ok {
				groups = append(groups, g)
			}
		}
	}
	return groups
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
	groupType = strings.TrimSpace(groupType)
	if len(groupType) == 0 {
		groupType = "Team"
	}

	set := GroupsSet{GroupsMap: map[string]Group{}}

	query := url.Values{}
	query.Add("recordCount", "250")

	if len(groupType) > 0 {
		query.Add("type", groupType)
	}

	for {
		// apiUrlString := urlutil.JoinAbsolute(serverUrl, ApiPathListUserGroups)
		// groupsURL := ro.BuildURL(serverUrl, "/glip/groups", true, query)
		groupsURL, err := urlutil.URLAddQueryValuesString(
			urlutil.JoinAbsolute(serverUrl, ApiPathListUserGroups),
			query)
		if err != nil {
			return set, err
		}
		resp, err := client.Get(groupsURL.String())
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
	bytes, err := ioutil.ReadAll(resp.Body)
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
