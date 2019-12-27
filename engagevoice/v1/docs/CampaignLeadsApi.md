# \CampaignLeadsApi

All URIs are relative to *https://portal.vacd.biz/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCampaignLeadStates**](CampaignLeadsApi.md#GetCampaignLeadStates) | **Get** /admin/accounts/{accountId}/campaignLeads/leadStates | Returns a listing of all lead states for an account  Permissions: READ on Account



## GetCampaignLeadStates

> []string GetCampaignLeadStates(ctx, accountId)
Returns a listing of all lead states for an account  Permissions: READ on Account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

