package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/path/filepathutil"
	"github.com/grokify/spectrum/openapi2"
	"github.com/grokify/spectrum/openapi3"
	"github.com/jessevdk/go-flags"
)

type options struct {
	Directory string `short:"d" long:"directory" description:"OAS Directory" required:"true"`
	Version   int    `short:"v" long:"version" description:"OAS Version" required:"true"`
}

func main() {
	if 1 == 0 {
		// file := "specs-voice_v3.0.0/openapi-spec_campaigns.json"
		// file = "specs-voice_v3.0.0/openapi-spec_countries.json"
		// file = "specs-example_v3.0.0.json"
		// file = "specs-voice_v3.0.0/openapi-spec_agent-groups.json"
		file := "specs-voice_v3.0.0/openapi-spec_agents.json"
		if 1 == 1 {
			spec, err := openapi3.ReadFile(file, false)
			if err != nil {
				fmt.Println("TEST_UNMARSHAL")
				log.Fatal(err)
			}
			fmtutil.MustPrintJSON(spec)
			name := "Country"
			fmtutil.MustPrintJSON(spec.Components.Schemas[name])

			moreSpec := openapi3.SpecMore{Spec: spec}

			fmt.Printf("HAS [%v][%v]\n", name, moreSpec.SchemaNameExists(name, false))
			err = spec.Validate(context.Background())
			if err != nil {
				log.Fatal(err)
			}
			panic("READ_FILE_NEW")
		}
		spec, err := openapi3.ReadFile(file, true)
		if err != nil {
			fmt.Println("TEST_LOADER")
			log.Fatal(err)
		}
		fmtutil.MustPrintJSON(spec)
		err = spec.Validate(context.Background())
		if err != nil {
			fmt.Println("TEST")
			log.Fatal(err)
		}
		fmtutil.MustPrintJSON(spec)

		fmt.Println("DONE")
		panic("Z")
	}
	opts := options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	version := 2
	if opts.Version == 2 || opts.Version == 3 {
		version = opts.Version
	}
	dir := opts.Directory

	dir = filepathutil.TrimRight(dir)

	_, leaf := filepath.Split(dir)

	outfile := leaf + ".json"

	switch version {
	case 2:
		err = openapi2.WriteFileDirMerge(outfile, dir, 0644)
	case 3:
		_, err = openapi3.WriteFileDirMerge(outfile, dir, 0644, nil)
	}
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DONE")
}
