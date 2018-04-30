package clientutil

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	rc "github.com/grokify/go-ringcentral/client"
)

type GlipApiUtil struct {
	ApiClient *rc.APIClient
}

func (apiUtil *GlipApiUtil) GlipGroupMemberCount(groupId string) (int64, error) {
	if apiUtil.ApiClient == nil {
		return int64(-1), fmt.Errorf("GlipApiUtil is missing RingCentral ApiClient")
	}
	groupId = strings.ToLower(strings.TrimSpace(groupId))
	grp, resp, err := apiUtil.ApiClient.GlipApi.LoadGroup(context.Background(), groupId)
	if err != nil {
		return int64(-1), err
	} else if resp.StatusCode >= 300 {
		return int64(-1), fmt.Errorf("Glip API Response Code [%v]", resp.StatusCode)
	}
	return int64(len(grp.Members)), nil
}

// DirectMessage means a group of 2 or a team of 2
func (apiUtil *GlipApiUtil) AtMentionedOrGroupOfTwo(userId, groupId string, mentions []rc.GlipMentionsInfo) (bool, error) {
	if IsAtMentioned(userId, mentions) {
		return true, nil
	}

	count, err := apiUtil.GlipGroupMemberCount(groupId)
	if err != nil {
		return false, err
	}
	if count == int64(2) {
		return true, nil
	}
	return false, nil
}

func IsAtMentioned(userId string, mentions []rc.GlipMentionsInfo) bool {
	for _, mention := range mentions {
		if userId == mention.Id {
			return true
		}
	}
	return false
}

func GlipCreatePostIsEmpty(post rc.GlipCreatePost) bool {
	if len(strings.TrimSpace(post.Text)) == 0 && len(post.Attachments) == 0 {
		return false
	}
	return true
}

func StripAtMention(id, text string) string {
	rx := regexp.MustCompile(fmt.Sprintf("!\\[:Person\\]\\(%v\\)", id))
	noAtMention := rx.ReplaceAllString(text, " ")
	noAtMention = regexp.MustCompile(`\\s+`).ReplaceAllString(noAtMention, " ")
	return strings.TrimSpace(noAtMention)
}
