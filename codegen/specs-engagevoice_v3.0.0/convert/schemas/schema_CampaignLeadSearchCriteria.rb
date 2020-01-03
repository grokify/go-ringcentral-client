#!ruby

str = 'private List<Integer> leadIds = new ArrayList<>();
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
    private Integer agentId;'
    
    extra = '
    private List<Timezone> leadTimezones;
	private DateTime loadedDtsStart;
    private DateTime loadedDtsEnd;
	private SuppressedType suppressed = SuppressedType.ALL;
	private List<DispositionType> systemDispositions;
	private List<LeadState> leadStates;
	private DateTimeComparableSearchField nextDialTimeCriteria;
	private DateTimeComparableSearchField lastPassTimeCriteria;
	private ComparableSearchField leadPassesCriteria;

```
SuppressedType {name}
ComparableSearchField
DateTimeComparableSearchField
DateTime


api/services/outbound/leadManagerActions/CampaignLeadSearchCriteria.java:public class CampaignLeadSearchCriteria implements JsonObject {
LMRC9100:connectfirst john.wang$ grep -R 'class ComparableSearchField' *
api/domain/ComparableSearchField.java:public class ComparableSearchField implements JsonObject {


ComparableSearchField {
    "object": {}, ??
    "operator":{
        "symbol": "string"
    }
symbol: string // GREATER_THAN(">"), LESS_THAN("<"), EQUALS("=");



}
```