# \DevicesApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestapiV10AccountAccountIdDeviceDeviceIdGet**](DevicesApi.md#RestapiV10AccountAccountIdDeviceDeviceIdGet) | **Get** /restapi/v1.0/account/{accountId}/device/{deviceId} | 
[**RestapiV10AccountAccountIdDeviceGet**](DevicesApi.md#RestapiV10AccountAccountIdDeviceGet) | **Get** /restapi/v1.0/account/{accountId}/device | 
[**RestapiV10AccountAccountIdExtensionExtensionIdDeviceGet**](DevicesApi.md#RestapiV10AccountAccountIdExtensionExtensionIdDeviceGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/device | 


# **RestapiV10AccountAccountIdDeviceDeviceIdGet**
> DeviceInfo RestapiV10AccountAccountIdDeviceDeviceIdGet($accountId, $deviceId)



Get Device by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **deviceId** | **string**| Internal identifier of a device | 

### Return type

[**DeviceInfo**](DeviceInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdDeviceGet**
> InlineResponseDefault5 RestapiV10AccountAccountIdDeviceGet($accountId)



Get Account Device List


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 

### Return type

[**InlineResponseDefault5**](inline_response_default_5.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdDeviceGet**
> InlineResponseDefault17 RestapiV10AccountAccountIdExtensionExtensionIdDeviceGet($accountId, $extensionId)



Get Extension Device List


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 

### Return type

[**InlineResponseDefault17**](inline_response_default_17.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

