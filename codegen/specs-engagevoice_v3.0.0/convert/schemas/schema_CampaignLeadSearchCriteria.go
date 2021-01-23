package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	oas3 "github.com/getkin/kin-openapi/openapi3"
	"github.com/grokify/simplego/fmt/fmtutil"
	"github.com/grokify/swaggman/openapi3/fromspring"
	"github.com/rs/zerolog/log"
)

type SchemaInfo struct {
	Name                string
	JavaCode            string
	ExplicitCustomTypes []string
	Notes               string
}

var infos = map[string]SchemaInfo{
	"CampaignLeadSearchCriteria": SchemaInfo{
		Name: "CampaignLeadSearchCriteria",
		ExplicitCustomTypes: []string{
			"ComparableSearchField",
			"DateTimeComparableSearchField",
			"SuppressedType",
			"Timezone"},
		JavaCode: `private List<Integer> leadIds = new ArrayList<>();
			private List<Integer> listIds = new ArrayList<>();
			private List<String> externIds = new ArrayList<>();
			private List<String> physicalStates;
			private List<String> agentDispositions;
			private List<String> leadPhoneNumbers = new ArrayList<>();
			private boolean orphanedLeadsOnly;
			private String callerId;
			private String leadPhoneNum;
			private List<Integer> campaignIds = new ArrayList<>();
			private String firstName;
			private String lastName;
			private String address1;
			private String address2;
			private String city;
			private String zip;
			private String emailAddress;
			private String auxData1;
			private String auxData2;
			private String auxData3;
			private String auxData4;
			private String auxData5;
			private Integer pendingAgentId;
			private Integer agentId;
			private DateTime loadedDtsStart;
			private DateTime loadedDtsEnd;
			private List<Timezone> leadTimezones;
			private DateTimeComparableSearchField nextDialTimeCriteria;
			private DateTimeComparableSearchField lastPassTimeCriteria;
			private ComparableSearchField leadPassesCriteria;
			`,
	},
	"CampaignLeadSearchResultsView": {
		Name:                "CampaignLeadSearchResultsView",
		ExplicitCustomTypes: []string{"CampaignLead"},
		JavaCode: `
private Integer leadId;
private Integer dialGroupId;
private Integer campaignId;
private String campaignName;
private String pendingAgentName;
private String lastPassDisposition;
private String lastPassAgentName;
private String listDesc;
private Integer listId;
private String uploadedBy;

DateTime loadedDts;
private DateTime lastPassDate;
DateTime lastPassDts;
DateTime nextDialTime;
DateTime stateDts;
private DateTime uploadDate;

private CampaignLead campaignLead;

Long leadId;
String id;
String name;

String lastPassDispo;
String agentDispostion;

String address1;
String address2;
String auxData1;
String auxData2;
String auxData3;
String auxData4;
String auxData5;

String auxExternalUrl;
String auxGreeting;
String auxPhone;

String callerId;
String city;
Long dupeKeyOverride;
String email;
String externId;
String extraData;
String firstName;
String gateKeeper;
String lastName;
Integer leadPasses;
String leadPhone;
String leadState;
String leadTimezone;

String liveAnswerMessage;
String machAnswerMessage;
Integer maxPasses;
String midName;

String sortCol;

String state;
String suffix;
String title;
String zip;
Boolean suppressed;

Integer speedToLeadAgentConn;
Integer speedToLeadFirstPass;
`,
	},
	"CampaignLead": SchemaInfo{
		Name:  "CampaignLead",
		Notes: "No `leadPriority` (null). campaign (null), campaignRedials (null),quotaTargets (empty array), quotaTargetIds (empty array)",
		JavaCode: `
		Long leadId;
		Long id;
		String name;
		
		String lastPassDispo;
		String agentDispostion;
	
		String address1;
		String address2;
		String auxData1;
		String auxData2;
		String auxData3;
		String auxData4;
		String auxData5;

		String auxExternalUrl;
		String auxGreeting;
		String auxPhone;
	
		String callerId;
		String city;
		Long dupeKeyOverride;
		String email;
		String externId;
		String extraData;
		String firstName;
		String gateKeeper;
		String lastName;
		Integer leadPasses;
		String leadPhone;
		String leadState;
		String leadTimezone;

		String liveAnswerMessage;
		String machAnswerMessage;
		Integer maxPasses;
		String midName;

		String sortCol;
	
		String state;
		String suffix;
		String title;
		String zip;
		Boolean suppressed;
	
		Integer speedToLeadAgentConn;
		Integer speedToLeadFirstPass;

		DateTime loadedDts;
		DateTime nextDialTime;
	`,
	},
}

func ProcInfo(schemaInfo SchemaInfo) {
	lines := strings.Split(schemaInfo.JavaCode, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	fmtutil.PrintJSON(lines)

	mss, err := fromspring.ParseSpringLinesToMapStringSchemaRefs(
		lines, schemaInfo.ExplicitCustomTypes)
	if err != nil {
		log.Info().Msg("S1")
		log.Fatal().Err(err)
	}
	fmtutil.PrintJSON(mss)

	oas := oas3.Swagger{
		Components: oas3.Components{
			Schemas: map[string]*oas3.SchemaRef{
				schemaInfo.Name: &oas3.SchemaRef{
					Value: &oas3.Schema{
						Type:       fromspring.TypeObject,
						Properties: mss,
					},
				},
			},
		},
	}
	joas, err := oas.MarshalJSON()
	if err != nil {
		log.Fatal().Err(err)
	}

	file := "openapi-spec_schema_" + schemaInfo.Name + ".json"
	err = ioutil.WriteFile(file, joas, 0644)
	if err != nil {
		log.Fatal().Err(err)
	}
	fmt.Printf("WROTE [%v]\n", file)

}

func main() {
	for _, info := range infos {
		ProcInfo(info)
	}

	fmt.Println("DONE")
}
