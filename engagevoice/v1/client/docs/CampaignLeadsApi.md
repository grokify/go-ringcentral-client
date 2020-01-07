# \CampaignLeadsApi

All URIs are relative to *https://portal.vacd.biz/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLeadStates**](CampaignLeadsApi.md#GetLeadStates) | **Get** /admin/accounts/{accountId}/campaignLeads/leadStates | Get Lead States
[**GetSystemDispositions**](CampaignLeadsApi.md#GetSystemDispositions) | **Get** /admin/accounts/{accountId}/campaignLeads/systemDispositions | Get Systems Dispositions
[**SearchLeads**](CampaignLeadsApi.md#SearchLeads) | **Post** /admin/accounts/{accountId}/campaignLeads/leadSearch | Search Leads



## GetLeadStates

> []string GetLeadStates(ctx, accountId)
Get Lead States

Returns a listing of all lead states for an account  Permissions: READ on Account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 

### Return type

**[]string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemDispositions

> []string GetSystemDispositions(ctx, accountId)
Get Systems Dispositions

Returns a listing of all system dispositions for an account  Permissions: READ on Account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 

### Return type

**[]string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchLeads

> []CampaignLeadSearchResultsView SearchLeads(ctx, accountId, campaignLeadSearchCriteria, optional)
Search Leads

Allows searching of campaign leads for a single campaign  Permissions: READ on Account (Permission Override)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**campaignLeadSearchCriteria** | [**CampaignLeadSearchCriteria**](CampaignLeadSearchCriteria.md)|  | 
 **optional** | ***SearchLeadsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchLeadsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**|  | 
 **page** | **optional.Int32**|  | 
 **maxRows** | **optional.Int32**|  | 

### Return type

[**[]CampaignLeadSearchResultsView**](CampaignLeadSearchResultsView.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

