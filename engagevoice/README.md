# RingCentral Engage Voice

## APIs

- [ ] Active Calls
  - [ ] /admin/accounts/{accountId}/activeCalls/list?product={productName}&productId={productId} 

- [ ] Agent Groups
  - [x] [/v1/admin/accounts/{accountId}/agentGroups GET](v1/docs/AgentsApi.md#getagentgroups)
  - [x] [/v1/admin/accounts/{accountId}/agentGroups/{agentGroupId}/agents GET](v1/docs/AgentsApi.md#getagents)

- [ ] Campaign Leads
  - [x] [/v1/admin/accounts/{accountId}/campaignLeads/leadStates GET](v1/docs/CampaignLeadsApi.md#getcampaignleadstates)
  - [x] [/v1/admin/accounts/{accountId}/campaignLeads/systemDispositions GET](v1/docs/CampaignLeadsApi.md#getsystemdispositions)

- [ ] Campaigns
  - [x] [/admin/accounts/{accountId}/campaigns/{campaignId}/leadLoader/direct POST](v1/docs/LeadsApi.md#uploadleads)

- [ ] Countries
  - [x] [/v1/admin/accounts/{accountId}/countries/available GET](v1/docs/CountriesApi.md#getavailablecountries)

- [ ] Dial Groups
  - [x] [/v1/admin/accounts/{accountId}/dialGroups GET](v1/docs/DialGroupsApi.md#getdialgroups)
  - [x] [/admin/accounts/{accountId}/dialGroups/{dialGroupId}/campaigns GET](v1/docs/CampaignsApi.md#getdialgroupcampaigns)
  
  
