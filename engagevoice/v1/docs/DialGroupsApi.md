# \DialGroupsApi

All URIs are relative to *https://portal.vacd.biz/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDialGroups**](DialGroupsApi.md#GetDialGroups) | **Get** /admin/accounts/{accountId}/dialGroups | Returns a listing of dial groups for an account



## GetDialGroups

> []DialGroup GetDialGroups(ctx, accountId)
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

