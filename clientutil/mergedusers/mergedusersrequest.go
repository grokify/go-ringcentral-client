package mergedusers

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	rc "github.com/grokify/go-ringcentral/client"
	ru "github.com/grokify/go-ringcentral/clientutil"
	"github.com/grokify/gotilla/mime/multipartutil"
)

const (
	glipPersonInfosUrlPart = "/restapi/v1.0/glip/persons/"
	extensionInfosUrlPart  = "/restapi/v1.0/account/~/extension/"
)

var (
	rxAccountIdUrl = regexp.MustCompile(`^(.+restapi/v1.0/account/\d+).*`)
	rxAccountId    = regexp.MustCompile(`/account/(\d+)`)
)

func NewMergedUsersApiIds(client *http.Client, serverUrl string, userIds []string) (MergedUserSet, error) {
	set := NewMergedUserSet()
	glipPersonInfosUrl := GlipPersonInfosUrl(serverUrl, userIds)

	resp, err := client.Get(glipPersonInfosUrl)
	if err != nil {
		return set, err
	}
	mr, err := multipartutil.NewMultipartReaderForHttpResponse(resp)
	if err != nil {
		return set, err
	}
	set, err = AddBatchGlipPersonInfosMultipartReader(set, mr)
	if err != nil {
		return set, err
	}

	extensionInfosUrl := ExtensionInfosUrl(serverUrl, userIds)

	resp2, err := client.Get(extensionInfosUrl)
	if err != nil {
		return set, err
	}
	set, err = AddBatchExtensionInfosHttpResponse(set, resp2)
	if err != nil {
		return set, err
	}
	return AddMainCompanyPhoneNumber(client, serverUrl, set)
}

func GlipPersonInfosUrl(serverUrl string, userIds []string) string {
	return fmt.Sprintf("%s%s%s", serverUrl, glipPersonInfosUrlPart, strings.Join(userIds, ","))
}

func ExtensionInfosUrl(serverUrl string, userIds []string) string {
	return fmt.Sprintf("%s%s%s", serverUrl, extensionInfosUrlPart, strings.Join(userIds, ","))
}

func ApiUrlToAccountIdUrl(apiUrl string) string {
	return rxAccountIdUrl.ReplaceAllString(apiUrl, "$1")
}

func AddMainCompanyPhoneNumber(client *http.Client, serverUrl string, set MergedUserSet) (MergedUserSet, error) {
	accountIdToExtIds := map[string]map[string]int{}
	for extId, mu := range set.MergedUserMap {
		if len(mu.ExtensionInfo.Uri) > 0 {
			m := rxAccountId.FindStringSubmatch(mu.ExtensionInfo.Uri)
			if len(m) > 1 {
				actId := m[1]
				if _, ok := accountIdToExtIds[actId]; !ok {
					accountIdToExtIds[actId] = map[string]int{}
				}
				accountIdToExtIds[actId][extId] = 1
			}
		}
	}

	if len(accountIdToExtIds) > 0 {
		apiClient, err := ru.NewApiClientHttpClientBaseURL(client, serverUrl)
		if err != nil {
			return set, err
		}

		for actId, extIds := range accountIdToExtIds {
			mainNumber, err := GetCompanyMainNumber(apiClient, actId)
			if err != nil {
				return set, err
			}
			if len(mainNumber) > 0 {
				for extId := range extIds {
					if user, ok := set.MergedUserMap[extId]; ok {
						user.MainNumber = mainNumber
						set.MergedUserMap[extId] = user
					}
				}
			}
		}
	}
	return set, nil
}

func GetCompanyMainNumber(apiClient *rc.APIClient, accountId string) (string, error) {
	info, resp, err := apiClient.CompanySettingsApi.LoadAccount(context.Background(), accountId)
	if err != nil {
		return "", err
	} else if resp.StatusCode >= 300 {
		return "", fmt.Errorf("API Error Status [%v]", resp.StatusCode)
	}
	return info.MainNumber, nil
}
