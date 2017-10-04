# \PresenceApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestapiV10AccountAccountIdExtensionExtensionIdPresenceGet**](PresenceApi.md#RestapiV10AccountAccountIdExtensionExtensionIdPresenceGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/presence | 


# **RestapiV10AccountAccountIdExtensionExtensionIdPresenceGet**
> PresenceInfo RestapiV10AccountAccountIdExtensionExtensionIdPresenceGet($accountId, $extensionId)



Get Extension Presence


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 

### Return type

[**PresenceInfo**](PresenceInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

