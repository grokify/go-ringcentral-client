package main

import (
	"log"

	"github.com/grokify/go-ringcentral-client/office/v1/util/mergedusers"
	"github.com/grokify/mogo/fmt/fmtutil"
)

var (
	exampleBoundary = `glboundary_f3VcIoEs3XcwvCsNOm1b2TkCmCQViub1w3oF`
	exampleBody     = `--glboundary_f3VcIoEs3XcwvCsNOm1b2TkCmCQViub1w3oF
Content-type: application/json; charset=utf-8
Content-length: 44

{"response":[{"status":200},{"status":200}]}
--glboundary_f3VcIoEs3XcwvCsNOm1b2TkCmCQViub1w3oF
Content-type: application/json; charset=utf-8
Content-length: 519

{"id":"557601020","firstName":"John","lastName":"Wang","gender":"male","email":"john.wang@ringcentral.com","location":"Belmont, CA","avatar":"https://example.com/557601020.jpg","companyId":"123","creationTime":"2015-06-22T22:08:33.303Z","lastModifiedTime":"2018-10-22T02:26:32.212Z","employeeSince":null,"jobTitle":"Platform Products","birthday":null,"webPage":null}
--glboundary_f3VcIoEs3XcwvCsNOm1b2TkCmCQViub1w3oF
Content-type: application/json; charset=utf-8
Content-length: 519

{"id":"557601020","firstName":"John","lastName":"Wang","gender":"male","email":"john.wang@ringcentral.com","location":"Belmont, CA","avatar":"https://examplecom/557601020.jpg","companyId":"123","creationTime":"2015-06-22T22:08:33.303Z","lastModifiedTime":"2018-10-22T02:26:32.212Z","employeeSince":null,"jobTitle":"Platform Products","birthday":null,"webPage":null}
--glboundary_f3VcIoEs3XcwvCsNOm1b2TkCmCQViub1w3oF--`
)

func main() {
	mergedUserSet := mergedusers.NewMergedUserSet()
	mergedUserSet, err := mergedusers.AddBatchGlipPersonInfosBodyBoundary(mergedUserSet, []byte(exampleBody), exampleBoundary)
	if err != nil {
		log.Fatal(err)
	}
	fmtutil.PrintJSON(mergedUserSet)
}
