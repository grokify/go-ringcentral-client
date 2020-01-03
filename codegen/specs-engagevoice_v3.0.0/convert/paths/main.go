package main

import "fmt"

var example1 = `
@RequestMapping(path = "campaignLeads/leadStates", method = RequestMethod.GET, produces = JSON_MEDIA_TYPE)
@ApiOperation(value = "Returns a listing of all lead states for an account", notes = "Permissions: READ on Account")
@JsonView(JsonViews.ListView.class)
@PreAuthorize("hasPermission('ACCT', 'ACCT')")
public LeadState[] getLeadStates(@PathVariable final String accountId) {
    `

func main() {
	fmt.Println("DONE")
}
