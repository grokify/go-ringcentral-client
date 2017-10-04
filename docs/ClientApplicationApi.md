# \ClientApplicationApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestapiV10ClientInfoCustomDataCustomDataKeyPut**](ClientApplicationApi.md#RestapiV10ClientInfoCustomDataCustomDataKeyPut) | **Put** /restapi/v1.0/client-info/custom-data/{customDataKey} | 


# **RestapiV10ClientInfoCustomDataCustomDataKeyPut**
> InlineResponseDefault28 RestapiV10ClientInfoCustomDataCustomDataKeyPut($customDataKey, $body)



Update Custom Data by Key


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customDataKey** | **string**| Custom data access key. The number of unique custom data keys is limited to 100 keys per extension, summarized for all the applications. For example, if you have created 50 custom data keys under the Android mobile client application for the particular extension, then logged in the iOS application and created another 50 keys, the web client application won&#39;t be allowed to create any custom data key for that extension | 
 **body** | [**Body17**](Body17.md)|  | [optional] 

### Return type

[**InlineResponseDefault28**](inline_response_default_28.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

