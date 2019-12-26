# \CampaignsApi

All URIs are relative to *https://portal.vacd.biz/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDialGroupCampaigns**](CampaignsApi.md#GetDialGroupCampaigns) | **Get** /admin/accounts/{accountId}/dialGroups/{dialGroupId}/campaigns | Returns a listing of campaigns for a dial group



## GetDialGroupCampaigns

> []Campaign GetDialGroupCampaigns(ctx, accountId, dialGroupId)
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

