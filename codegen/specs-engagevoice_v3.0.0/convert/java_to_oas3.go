package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	oas3 "github.com/getkin/kin-openapi/openapi3"
	"github.com/grokify/mogo/errors/errorsutil"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/os/osutil"
	"github.com/grokify/mogo/type/stringsutil"
	"github.com/grokify/spectrum/openapi3"
	"github.com/grokify/spectrum/openapi3/springopenapi3"
	"github.com/rs/zerolog/log"
)

func ReadOas3Spec(file string, validate bool) (*openapi3.Spec, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errorsutil.Wrap(err, fmt.Sprintf("Filename [%v]", file))
	}
	swag := &openapi3.Spec{}
	err = swag.UnmarshalJSON(bytes)
	if err != nil {
		return nil, errorsutil.Wrap(err, fmt.Sprintf("Filename [%v]", file))
	}
	if validate {
		err = swag.Validate(context.Background())
		return swag, errorsutil.Wrap(err, fmt.Sprintf("Filename [%v]", file))
	}
	return swag, nil
}

func main() {
	if 1 == 1 {
		file := "../openapi-spec_agents.json"

		swag, err := ReadOas3Spec(file, true)
		if err != nil {
			log.Fatal().Err(err)
		}
		fmtutil.PrintJSON(swag)
		panic("Z")
	}

	bytes, err := ioutil.ReadFile("_class_agent.java")
	if err != nil {
		log.Fatal().Err(err)
	}
	fmt.Println(string(bytes))
	lines := strings.Split(string(bytes), "\n")
	fmtutil.PrintJSON(lines)
	lines = stringsutil.SliceTrim(lines, " \t", true)
	fmtutil.PrintJSON(lines)
	//panic("Z")

	columnsRaw := springopenapi3.ParseSpringCodeColumnsRaw(lines)
	fmtutil.PrintJSON(columnsRaw)

	schema := oas3.Schema{Properties: map[string]*oas3.SchemaRef{}}
	fmtutil.PrintJSON(schema)

	if 1 == 0 {
		// line := "private Boolean userManagedByRC;"
		line := "private Boolean userManagedByRC = false;"
		name, prop, err := springopenapi3.ParseSpringLineToSchema(line)
		if err != nil {
			log.Fatal().Err(err)
		}
		fmt.Printf("NAME [%v]\n", name)
		fmtutil.PrintJSON(prop)
		panic("Z")
	}

	mss, err := springopenapi3.ParseSpringPropertyLinesSliceToSchema(columnsRaw)
	if err != nil {
		log.Info().Msg("S1")
		log.Fatal().Err(err)
	}
	fmtutil.PrintJSON(mss)

	err = osutil.WriteFileJSON("_schema_agent.json", mss, 644, "", "  ")
	if err != nil {
		log.Fatal().Err(err)
	}
	fmt.Printf("WROTE [%v]\n", "_schema_agent.json")

	swag2 := openapi3.Spec{
		Components: &oas3.Components{
			Schemas: map[string]*oas3.SchemaRef{
				"Agent": {
					Value: &oas3.Schema{
						Properties: mss,
					},
				},
			},
		},
	}
	fmtutil.PrintJSON(swag2)
	swagFile := "_openapi-spec_agents_models.json"
	err = osutil.WriteFileJSON(swagFile, swag2, 0644, "", "  ")
	if err != nil {
		log.Fatal().Err(err)
	}
	fmt.Printf("WROTE [%v]\n", swagFile)

	fmt.Println("DONE")
}
