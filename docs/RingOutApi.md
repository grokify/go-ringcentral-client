# \RingOutApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestapiV10AccountAccountIdExtensionExtensionIdRingoutPost**](RingOutApi.md#RestapiV10AccountAccountIdExtensionExtensionIdRingoutPost) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/ringout | 
[**RestapiV10AccountAccountIdExtensionExtensionIdRingoutRingoutIdDelete**](RingOutApi.md#RestapiV10AccountAccountIdExtensionExtensionIdRingoutRingoutIdDelete) | **Delete** /restapi/v1.0/account/{accountId}/extension/{extensionId}/ringout/{ringoutId} | 
[**RestapiV10AccountAccountIdExtensionExtensionIdRingoutRingoutIdGet**](RingOutApi.md#RestapiV10AccountAccountIdExtensionExtensionIdRingoutRingoutIdGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/ringout/{ringoutId} | 


# **RestapiV10AccountAccountIdExtensionExtensionIdRingoutPost**
> RingOutInfo RestapiV10AccountAccountIdExtensionExtensionIdRingoutPost($accountId, $extensionId, $body)



Initiate RingOut Call


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **body** | [**Body14**](Body14.md)|  | [optional] 

### Return type

[**RingOutInfo**](RingOutInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdRingoutRingoutIdDelete**
> RestapiV10AccountAccountIdExtensionExtensionIdRingoutRingoutIdDelete($accountId, $extensionId, $ringoutId)



Cancel RingOut Call


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **ringoutId** | **string**| Internal identifier of a RingOut call | 

### Return type

void (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdRingoutRingoutIdGet**
> RingOutInfo RestapiV10AccountAccountIdExtensionExtensionIdRingoutRingoutIdGet($accountId, $extensionId, $ringoutId)



Get RingOut Call Status


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **ringoutId** | **string**| Internal identifier of a RingOut call | 

### Return type

[**RingOutInfo**](RingOutInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

