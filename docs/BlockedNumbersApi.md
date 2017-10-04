# \BlockedNumbersApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestapiV10AccountAccountIdExtensionExtensionIdBlockedNumberBlockedNumberIdDelete**](BlockedNumbersApi.md#RestapiV10AccountAccountIdExtensionExtensionIdBlockedNumberBlockedNumberIdDelete) | **Delete** /restapi/v1.0/account/{accountId}/extension/{extensionId}/blocked-number/{blockedNumberId} | 
[**RestapiV10AccountAccountIdExtensionExtensionIdBlockedNumberBlockedNumberIdGet**](BlockedNumbersApi.md#RestapiV10AccountAccountIdExtensionExtensionIdBlockedNumberBlockedNumberIdGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/blocked-number/{blockedNumberId} | 
[**RestapiV10AccountAccountIdExtensionExtensionIdBlockedNumberBlockedNumberIdPut**](BlockedNumbersApi.md#RestapiV10AccountAccountIdExtensionExtensionIdBlockedNumberBlockedNumberIdPut) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/blocked-number/{blockedNumberId} | 
[**RestapiV10AccountAccountIdExtensionExtensionIdBlockedNumberGet**](BlockedNumbersApi.md#RestapiV10AccountAccountIdExtensionExtensionIdBlockedNumberGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/blocked-number | 
[**RestapiV10AccountAccountIdExtensionExtensionIdBlockedNumberPost**](BlockedNumbersApi.md#RestapiV10AccountAccountIdExtensionExtensionIdBlockedNumberPost) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/blocked-number | 


# **RestapiV10AccountAccountIdExtensionExtensionIdBlockedNumberBlockedNumberIdDelete**
> RestapiV10AccountAccountIdExtensionExtensionIdBlockedNumberBlockedNumberIdDelete($accountId, $extensionId, $blockedNumberId)



Delete Blocked Number by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **blockedNumberId** | **string**| Internal identifier of a blocked number list entry | 

### Return type

void (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdBlockedNumberBlockedNumberIdGet**
> BlockedNumberInfo RestapiV10AccountAccountIdExtensionExtensionIdBlockedNumberBlockedNumberIdGet($accountId, $extensionId, $blockedNumberId)



Get Blocked Number by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **blockedNumberId** | **string**| Internal identifier of a blocked number list entry | 

### Return type

[**BlockedNumberInfo**](BlockedNumberInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdBlockedNumberBlockedNumberIdPut**
> BlockedNumberInfo RestapiV10AccountAccountIdExtensionExtensionIdBlockedNumberBlockedNumberIdPut($accountId, $extensionId, $blockedNumberId, $body)



Update Blocked Number Label


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **blockedNumberId** | **string**| Internal identifier of a blocked number list entry | 
 **body** | [**BlockedNumberInfo**](BlockedNumberInfo.md)|  | [optional] 

### Return type

[**BlockedNumberInfo**](BlockedNumberInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdBlockedNumberGet**
> InlineResponseDefault14 RestapiV10AccountAccountIdExtensionExtensionIdBlockedNumberGet($accountId, $extensionId)



Get Blocked Number List


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 

### Return type

[**InlineResponseDefault14**](inline_response_default_14.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdBlockedNumberPost**
> BlockedNumberInfo RestapiV10AccountAccountIdExtensionExtensionIdBlockedNumberPost($accountId, $extensionId, $body)



Add New Blocked Number


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **body** | [**BlockedNumberInfo**](BlockedNumberInfo.md)|  | [optional] 

### Return type

[**BlockedNumberInfo**](BlockedNumberInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

