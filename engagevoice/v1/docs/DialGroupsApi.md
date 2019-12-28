# \DialGroupsApi

All URIs are relative to *https://portal.vacd.biz/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearCampaignCache**](DialGroupsApi.md#ClearCampaignCache) | **Post** /admin/accounts/{accountId}/dialGroups/{dialGroupId}/campaigns/{campaignId}/clearCache | Clear Campaign Cache
[**GetCampaigns**](DialGroupsApi.md#GetCampaigns) | **Get** /admin/accounts/{accountId}/dialGroups/{dialGroupId}/campaigns | Get Campaigns
[**GetDialGroups**](DialGroupsApi.md#GetDialGroups) | **Get** /admin/accounts/{accountId}/dialGroups | Get Dial Groups



## ClearCampaignCache

> ClearCampaignCache(ctx, accountId, dialGroupId, campaignId, body)
Clear Campaign Cache

lears the cache for a campaign  Permissions: READ on Campaign (Permission Override)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int64**|  | 
**dialGroupId** | **int64**|  | 
**campaignId** | **int64**|  | 
**body** | **map[string]interface{}**| An &#x60;application/json&#x60; Content-Type header is required. Submit an empty body to trigger the header. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaigns

> []Campaign GetCampaigns(ctx, accountId, dialGroupId)
Get Campaigns

Returns a listing of campaigns for a dial group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**dialGroupId** | **string**|  | 

### Return type

[**[]Campaign**](Campaign.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDialGroups

> []DialGroup GetDialGroups(ctx, accountId)
Get Dial Groups

Returns a listing of dial groups for an account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 

### Return type

[**[]DialGroup**](DialGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

