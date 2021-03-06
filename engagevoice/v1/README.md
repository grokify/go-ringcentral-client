# RingCentral Engage Voice

This is a Go Client SDK for the RingCentral Engage Voice API.

## Usage

See the sample code here:

* [Generate API Token](v1examples/generate_api_token/main.go)
* [API Calls](https://github.com/grokify/go-ringcentral-client/blob/master/engagevoice/v1/examples/get_simple/main.go)

## API Overview

Get an overview of the API here:

* [API Reference](https://grokify.github.io/go-ringcentral-client/engagevoice/)
* [OpenAPI 3.0 Specification](https://raw.githubusercontent.com/grokify/go-ringcentral-client/master/codegen/specs-engagevoice_v3.0.0.json)
* [Client SDK Docs](https://github.com/grokify/go-ringcentral-client/blob/master/engagevoice/v1/README.md)

## API Coverage & Documentation

The following endpoints are covered by the OpenAPI 3.0 spec, the API Reference and the SDK.

- [ ] Auth
  - [x] [/v1/admin/token GET](client/docs/AuthApi.md#gettokens)

- [ ] Active Calls
  - [ ] /v1/admin/accounts/{accountId}/activeCalls/list?product={productName}&productId={productId} 

- [ ] Agent Groups
  - [x] [/v1/admin/accounts/{accountId}/agentGroups GET](client/docs/AgentsApi.md#getagentgroups)
  - [x] [/v1/admin/accounts/{accountId}/agentGroups/{agentGroupId}/agents GET](client/docs/AgentsApi.md#getagents)

- [ ] Campaign Leads
  - [x] [/v1/admin/accounts/{accountId}/campaignLeads/leadSearch POST](client/docs/CampaignLeadsApi.md#searchcampaignleads)
  - [x] [/v1/admin/accounts/{accountId}/campaignLeads/leadStates GET](client/docs/CampaignLeadsApi.md#getcampaignleadstates)
  - [x] [/v1/admin/accounts/{accountId}/campaignLeads/systemDispositions GET](client/docs/CampaignLeadsApi.md#getsystemdispositions)

- [ ] Campaigns
  - [x] [/v1/admin/accounts/{accountId}/campaigns/{campaignId}/leadLoader/direct POST](client/docs/CampaignsApi.md#uploadleads)

- [ ] Countries
  - [x] [/v1/admin/accounts/{accountId}/countries/available GET](client/docs/CountriesApi.md#getavailablecountries)

- [ ] Dial Groups
  - [x] [/v1/admin/accounts/{accountId}/dialGroups GET](client/docs/DialGroupsApi.md#getdialgroups)
  - [x] [/v1/admin/accounts/{accountId}/dialGroups/{dialGroupId}/campaigns GET](client/docs/DialGroupsApi.md#getdialgroupcampaigns)
  - [x] [/v1/admin/accounts/{accountId}/dialGroups/{dialGroupId}/campaigns/{campaignId}/clearCache POST](client/docs/DialGroupsApi.md#clearcampaigncache)

- [ ] Users
  - [x] [/v1/admin/users GET](client/docs/UsersApi.md#getusers)
