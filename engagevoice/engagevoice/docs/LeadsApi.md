# \LeadsApi

All URIs are relative to *https://portal.vacd.biz/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UploadLeads**](LeadsApi.md#UploadLeads) | **Post** /admin/accounts/{accountId}/campaigns/{campaignId}/leadLoader/direct | Upload Leads



## UploadLeads

> UploadLeadsResponse UploadLeads(ctx, accountId, campaignId, uploadLeadsRequest)
Upload Leads

Uploads a single lead or list of leads to a new or existing list

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**campaignId** | **string**|  | 
**uploadLeadsRequest** | [**UploadLeadsRequest**](UploadLeadsRequest.md)|  | 

### Return type

[**UploadLeadsResponse**](UploadLeadsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

