package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/antihax/optional"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/jessevdk/go-flags"

	"github.com/grokify/go-ringcentral-engage/engagedigital/engagedigital"
	utils "github.com/grokify/go-ringcentral-engage/engagedigital/engagedigitalutil"
)

type options struct {
	Site   string `short:"s" long:"site" description:"A site" required:"true"`
	Token  string `short:"t" long:"token" description:"A token" required:"true"`
	Action string `short:"a" long:"action" description:"A action (create|update|delete)" required:"true"`
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

	switch opts.Action {
	case "create":
		fmt.Println("I_CREATING_CATEGORY")
		opts.Name = strings.TrimSpace(opts.Name)
		if len(opts.Name) == 0 {
			log.Fatal("E_CREATE_NO_CATEGORY_NAME")
		}
		apiOpts := &engagedigital.CreateCategoryOpts{
			Name: optional.NewString(opts.Name)}
		info, resp, err := client.CategoriesApi.CreateCategory(
			context.Background(),
			apiOpts)
		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode != 200 {
			log.Fatal(resp.StatusCode)
		}
		fmtutil.PrintJSON(info)
	case "update":
		opts.Id = strings.TrimSpace(opts.Id)

		info, resp, err := client.CategoriesApi.GetCategory(context.Background(), opts.Id)
		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode != 200 {
			log.Fatal(resp.StatusCode)
		}
		apiOpts := &engagedigital.UpdateCategoryOpts{
			Name:         optional.NewString(opts.Name),
			Unselectable: optional.NewBool(false)}
		if !info.Unselectable {
			apiOpts.Unselectable = optional.NewBool(true)
		}
		info, resp, err = client.CategoriesApi.UpdateCategory(
			context.Background(), opts.Id, apiOpts)
		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode != 200 {
			log.Fatal(resp.StatusCode)
		}
		fmtutil.PrintJSON(info)
	case "delete":
		opts.Id = strings.TrimSpace(opts.Id)
		if len(opts.Id) == 0 {
			log.Fatal("E_DELETE_NO_CATEGORY_ID")
		}
		info, resp, err := client.CategoriesApi.DeleteCategory(
			context.Background(), opts.Id, nil)
		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode != 200 {
			log.Fatal(resp.StatusCode)
		}
		fmtutil.PrintJSON(info)
	}

	fmt.Println("DONE")
}
