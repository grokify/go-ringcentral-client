# \ConferencingApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestapiV10AccountAccountIdExtensionExtensionIdConferencingGet**](ConferencingApi.md#RestapiV10AccountAccountIdExtensionExtensionIdConferencingGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/conferencing | 
[**RestapiV10AccountAccountIdExtensionExtensionIdConferencingPut**](ConferencingApi.md#RestapiV10AccountAccountIdExtensionExtensionIdConferencingPut) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/conferencing | 


# **RestapiV10AccountAccountIdExtensionExtensionIdConferencingGet**
> ConferencingInfo RestapiV10AccountAccountIdExtensionExtensionIdConferencingGet($accountId, $extensionId, $countryId)



Get Conferencing info


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **countryId** | **string**| Internal identifier of a country. If not specified, the response is returned for the brand country | [optional] 

### Return type

[**ConferencingInfo**](ConferencingInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdConferencingPut**
> ConferencingInfo RestapiV10AccountAccountIdExtensionExtensionIdConferencingPut($accountId, $extensionId, $body)



Update Conferencing info


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **body** | [**Body7**](Body7.md)|  | [optional] 

### Return type

[**ConferencingInfo**](ConferencingInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

