package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/grokify/gotilla/fmt/fmtutil"

	rc "github.com/grokify/go-ringcentral/client"
	rcu "github.com/grokify/go-ringcentral/clientutil"
)

func getDemoMessage() rc.GlipCreatePost {
	return rc.GlipCreatePost{
		Text: "Body of the post",
		Attachments: []rc.GlipMessageAttachmentInfoRequest{
			{
				Type_: "Card",
				Title: "Attachment Title",
				Color: "#00ff2a",
				Author: &rc.GlipMessageAttachmentAuthorInfo{
					Name:    "Author Name",
					Uri:     "https://example.com/author_link",
					IconUri: "https://upload.wikimedia.org/wikipedia/commons/thumb/f/fd/000080_Navy_Blue_Square.svg/1200px-000080_Navy_Blue_Square.svg.png",
				},
				Text:         "Attachment text",
				Intro:        "Attachment intro appears before the attachment block",
				ImageUri:     "https://media3.giphy.com/media/l4FssTixISsPStXRC/giphy.gif",
				ThumbnailUri: "http://clipart-library.com/img/1415947.jpg",
				Fallback:     "Attachment fallback text",
				Fields: []rc.GlipMessageAttachmentFieldsInfo{
					{
						Title: "Field 1",
						Value: "A short field",
						Style: "Short",
					},
					{
						Title: "Field 2",
						Value: "This is [a linked short field](https://example.com)",
						Style: "Short",
					},
					{
						Title: "Field 3",
						Value: "A long, full-width field with *formatting* and [a link](https://example.com)",
						Style: "Long",
					},
				},
				Footnote: &rc.GlipMessageAttachmentFootnoteInfo{
					Text:    "Attachment footer and timestamp",
					IconUri: "http://www.iconsdb.com/icons/preview/red/square-ios-app-xxl.png",
					Time:    time.Now(),
				},
			},
		},
	}
}

func getDemoMessageSalesforce() rc.GlipCreatePost {
	return rc.GlipCreatePost{
		Text: "**Top Opportunities**\n\nFull report: https://my.salesforce.com/00O80000007MOgS",
		Attachments: []rc.GlipMessageAttachmentInfoRequest{
			{
				Type_:    "Card",
				Color:    "#00ff2a",
				Fallback: "Attachment fallback text",
				Fields: []rc.GlipMessageAttachmentFieldsInfo{
					{
						Title: "Opportunity",
						Value: "[Electric Company of America (0038000004OrS7y)](https://example.com)",
						Style: "Short"},
					{
						Title: "Owner",
						Value: "Matthew West",
						Style: "Short"},
					{
						Title: "Stage",
						Value: "Contract Negotiation",
						Style: "Short"},
					{
						Title: "Close Date",
						Value: "2017-12-20",
						Style: "Short"},
					{
						Title: "Amount",
						Value: "$150,000",
						Style: "Short"},
					{
						Title: "Probability",
						Value: "95%",
						Style: "Short"},
				},
			},
			{
				Type_:    "Card",
				Color:    "#dfdd13",
				Fallback: "Attachment fallback text",
				Fields: []rc.GlipMessageAttachmentFieldsInfo{
					{
						Title: "Opportunity",
						Value: "[Eureka Oil and Gas (0038000001MgG2z)](https://example.com)",
						Style: "Short"},
					{
						Title: "Owner",
						Value: "Alice Collins",
						Style: "Short"},
					{
						Title: "Stage",
						Value: "Proposal/Quote",
						Style: "Short"},
					{
						Title: "Close Date",
						Value: "2017-09-20",
						Style: "Short"},
					{
						Title: "Amount",
						Value: "$750,000",
						Style: "Short"},
					{
						Title: "Probability",
						Value: "70%",
						Style: "Short"},
					{
						Title: "Field 1",
						Value: "A short field",
						Style: "Short",
					},
				},
			},
			{
				Type_:    "Card",
				Color:    "#ff0000",
				Fallback: "Attachment fallback text",
				Fields: []rc.GlipMessageAttachmentFieldsInfo{
					{
						Title: "Opportunity",
						Value: "[Pacfic Restaurants (0038000004WhM4a)](https://example.com)",
						Style: "Short"},
					{
						Title: "Owner",
						Value: "Justin Lyons",
						Style: "Short"},
					{
						Title: "Stage",
						Value: "Discovery",
						Style: "Short"},
					{
						Title: "Close Date",
						Value: "2017-09-20",
						Style: "Short"},
					{
						Title: "Amount",
						Value: "$500,000",
						Style: "Short"},
					{
						Title: "Probability",
						Value: "30%",
						Style: "Short"},
				},
			},
		},
	}
}

func main() {
	err := rcu.LoadEnv()
	if err != nil {
		panic(err)
	}
	groupId := os.Getenv("GLIP_POST_GROUP_ID")

	apiClient, err := rcu.NewApiClientPasswordEnv()
	if err != nil {
		panic(err)
	}

	msgs := []rc.GlipCreatePost{}
	msgs = append(msgs, getDemoMessage())
	msgs = append(msgs, getDemoMessageSalesforce())

	fmtutil.PrintJSON(msgs)

	for _, msg := range msgs {
		info, resp, err := apiClient.GlipApi.CreatePost(context.Background(), groupId, msg)
		if err != nil {
			panic(err)
		} else if resp.StatusCode > 299 {
			panic(fmt.Errorf("API Status %v", resp.StatusCode))
		}
		fmtutil.PrintJSON(info)
	}
	fmt.Println("DONE")
}
