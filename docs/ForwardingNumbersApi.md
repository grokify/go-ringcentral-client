# \ForwardingNumbersApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestapiV10AccountAccountIdExtensionExtensionIdForwardingNumberGet**](ForwardingNumbersApi.md#RestapiV10AccountAccountIdExtensionExtensionIdForwardingNumberGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/forwarding-number | 
[**RestapiV10AccountAccountIdExtensionExtensionIdForwardingNumberPost**](ForwardingNumbersApi.md#RestapiV10AccountAccountIdExtensionExtensionIdForwardingNumberPost) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/forwarding-number | 


# **RestapiV10AccountAccountIdExtensionExtensionIdForwardingNumberGet**
> InlineResponseDefault18 RestapiV10AccountAccountIdExtensionExtensionIdForwardingNumberGet($accountId, $extensionId, $page, $perPage)



Get Forwarding Numbers


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **page** | **int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | [optional] 
 **perPage** | **int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | [optional] 

### Return type

[**InlineResponseDefault18**](inline_response_default_18.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdForwardingNumberPost**
> ForwardingNumberInfo RestapiV10AccountAccountIdExtensionExtensionIdForwardingNumberPost($accountId, $extensionId, $body)



Add New Forwarding Number


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **body** | [**Body9**](Body9.md)|  | [optional] 

### Return type

[**ForwardingNumberInfo**](ForwardingNumberInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

