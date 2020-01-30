# \CampaignsApi

All URIs are relative to *https://portal.vacd.biz/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PatchCampaignLead**](CampaignsApi.md#PatchCampaignLead) | **Patch** /admin/accounts/{accountId}/campaignLeads/{leadId} | Patch Campaign Lead
[**UpdateCampaignLead**](CampaignsApi.md#UpdateCampaignLead) | **Put** /admin/accounts/{accountId}/campaignLeads/{leadId} | Update Campaign Lead
[**UploadLeads**](CampaignsApi.md#UploadLeads) | **Post** /admin/accounts/{accountId}/campaigns/{campaignId}/leadLoader/direct | Upload Leads



## PatchCampaignLead

> CampaignLead PatchCampaignLead(ctx, accountId, leadId, campaignId, campaignLead, optional)
Patch Campaign Lead

Allows updating of a campaign lead, only updating those fields passed in  Permissions: READ on Account (Permission Override), UPDATE on Campaign

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**leadId** | **int32**|  | 
**campaignId** | **int32**|  | 
**campaignLead** | [**CampaignLead**](CampaignLead.md)|  | 
 **optional** | ***PatchCampaignLeadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PatchCampaignLeadOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **listId** | **optional.Int32**|  | 
 **timezoneOption** | **optional.String**|  | 
 **duplicateHandling** | **optional.String**| &#x60;RETAIN_ALL&#x60;: Retain all records, &#x60;REMOVE_ALL_EXISTING&#x60;: Remove duplicates from all existing lists, &#x60;REMOVE_FROM_LIST&#x60;: Remove duplicates from this list | [default to RETAIN_ALL]

### Return type

[**CampaignLead**](CampaignLead.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCampaignLead

> CampaignLead UpdateCampaignLead(ctx, accountId, leadId, campaignId, campaignLead, optional)
Update Campaign Lead

Allows updating of a campaign lead, updating entire lead including fields not passed in.  Permissions: READ on Account (Permission Override), UPDATE on Campaign

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**leadId** | **int32**|  | 
**campaignId** | **int32**|  | 
**campaignLead** | [**CampaignLead**](CampaignLead.md)|  | 
 **optional** | ***UpdateCampaignLeadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCampaignLeadOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **listId** | **optional.Int32**|  | 
 **timezoneOption** | **optional.String**|  | 
 **duplicateHandling** | **optional.String**| &#x60;RETAIN_ALL&#x60;: Retain all records, &#x60;REMOVE_ALL_EXISTING&#x60;: Remove duplicates from all existing lists, &#x60;REMOVE_FROM_LIST&#x60;: Remove duplicates from this list | [default to RETAIN_ALL]

### Return type

[**CampaignLead**](CampaignLead.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

