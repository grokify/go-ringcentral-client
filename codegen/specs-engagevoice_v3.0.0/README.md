# Engage Voice API Spec

## Known Issues

`GET /admin/accounts/{accountId}/campaigns/{campaignId}/leadLoader/direct`

* `auxData1` should be `oneOf` but OpenAPI Generator does not support OneOf. It generates a `Oneoflongstring` which isn't defined in `models_leads.go`. Workaround is to support `string` only.

`POST /admin/accounts/{accountId}/campaigns/{campaignId}/leadLoader/direct`

* `dncTags` isn't defined in `LeadsLoader` schema right now because no definition is available.

### Campaign Object

* `customDialZoneGroup` in response is not included. In test data, it is `null`.
* `quotaGroup` in response is not included. In test data, it is `null`.

This is used in GetDialGroupCampaigns API.