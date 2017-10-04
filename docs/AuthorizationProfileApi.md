# \AuthorizationProfileApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestapiV10AccountAccountIdExtensionExtensionIdAuthzProfileCheckGet**](AuthorizationProfileApi.md#RestapiV10AccountAccountIdExtensionExtensionIdAuthzProfileCheckGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/authz-profile/check | 
[**RestapiV10AccountAccountIdExtensionExtensionIdAuthzProfileGet**](AuthorizationProfileApi.md#RestapiV10AccountAccountIdExtensionExtensionIdAuthzProfileGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/authz-profile | 


# **RestapiV10AccountAccountIdExtensionExtensionIdAuthzProfileCheckGet**
> InlineResponseDefault13 RestapiV10AccountAccountIdExtensionExtensionIdAuthzProfileCheckGet($accountId, $extensionId, $permissionId, $targetExtensionId)



Check User Permissions


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **permissionId** | **string**| Permission to check | [optional] 
 **targetExtensionId** | **string**| Optional. Internal identifier of an extension for which user permissions are to be checked. The default value is the currently logged-in extension | [optional] 

### Return type

[**InlineResponseDefault13**](inline_response_default_13.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdAuthzProfileGet**
> InlineResponseDefault12 RestapiV10AccountAccountIdExtensionExtensionIdAuthzProfileGet($accountId, $extensionId)



Get User Permissions


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 

### Return type

[**InlineResponseDefault12**](inline_response_default_12.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

