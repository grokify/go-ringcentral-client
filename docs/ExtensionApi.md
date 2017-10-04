# \ExtensionApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestapiV10AccountAccountIdDepartmentDepartmentIdMembersGet**](ExtensionApi.md#RestapiV10AccountAccountIdDepartmentDepartmentIdMembersGet) | **Get** /restapi/v1.0/account/{accountId}/department/{departmentId}/members | 
[**RestapiV10AccountAccountIdExtensionExtensionIdGet**](ExtensionApi.md#RestapiV10AccountAccountIdExtensionExtensionIdGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId} | 
[**RestapiV10AccountAccountIdExtensionExtensionIdGrantGet**](ExtensionApi.md#RestapiV10AccountAccountIdExtensionExtensionIdGrantGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/grant | 
[**RestapiV10AccountAccountIdExtensionExtensionIdPhoneNumberGet**](ExtensionApi.md#RestapiV10AccountAccountIdExtensionExtensionIdPhoneNumberGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/phone-number | 
[**RestapiV10AccountAccountIdExtensionExtensionIdProfileImageGet**](ExtensionApi.md#RestapiV10AccountAccountIdExtensionExtensionIdProfileImageGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/profile-image | 
[**RestapiV10AccountAccountIdExtensionExtensionIdProfileImagePost**](ExtensionApi.md#RestapiV10AccountAccountIdExtensionExtensionIdProfileImagePost) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/profile-image | 
[**RestapiV10AccountAccountIdExtensionExtensionIdProfileImagePut**](ExtensionApi.md#RestapiV10AccountAccountIdExtensionExtensionIdProfileImagePut) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/profile-image | 
[**RestapiV10AccountAccountIdExtensionExtensionIdProfileImageScaleSizeGet**](ExtensionApi.md#RestapiV10AccountAccountIdExtensionExtensionIdProfileImageScaleSizeGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/profile-image/{scaleSize} | 
[**RestapiV10AccountAccountIdExtensionExtensionIdPut**](ExtensionApi.md#RestapiV10AccountAccountIdExtensionExtensionIdPut) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId} | 
[**RestapiV10AccountAccountIdExtensionGet**](ExtensionApi.md#RestapiV10AccountAccountIdExtensionGet) | **Get** /restapi/v1.0/account/{accountId}/extension | 


# **RestapiV10AccountAccountIdDepartmentDepartmentIdMembersGet**
> InlineResponseDefault4 RestapiV10AccountAccountIdDepartmentDepartmentIdMembersGet($accountId, $departmentId, $page, $perPage)



Get Department Members


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **departmentId** | **string**| Internal identifier of a Department extension (same as extensionId but only the ID of a department extension is valid) | 
 **page** | **int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | [optional] 
 **perPage** | **int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | [optional] 

### Return type

[**InlineResponseDefault4**](inline_response_default_4.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdGet**
> ExtensionInfo RestapiV10AccountAccountIdExtensionExtensionIdGet($accountId, $extensionId)



Get Extension Info by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 

### Return type

[**ExtensionInfo**](ExtensionInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdGrantGet**
> InlineResponseDefault19 RestapiV10AccountAccountIdExtensionExtensionIdGrantGet($accountId, $extensionId, $page, $perPage)



Get Extension Grants


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **page** | **int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | [optional] 
 **perPage** | **int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | [optional] 

### Return type

[**InlineResponseDefault19**](inline_response_default_19.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdPhoneNumberGet**
> InlineResponseDefault23 RestapiV10AccountAccountIdExtensionExtensionIdPhoneNumberGet($accountId, $extensionId, $usageType, $page, $perPage)



Get Extension Phone Numbers


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **usageType** | **string**| Usage type of the phone number | [optional] 
 **page** | **int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | [optional] 
 **perPage** | **int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | [optional] 

### Return type

[**InlineResponseDefault23**](inline_response_default_23.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdProfileImageGet**
> Binary RestapiV10AccountAccountIdExtensionExtensionIdProfileImageGet($accountId, $extensionId)



Get Profile Image


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 

### Return type

[**Binary**](Binary.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdProfileImagePost**
> RestapiV10AccountAccountIdExtensionExtensionIdProfileImagePost($accountId, $extensionId, $body)



Update Profile Image (same as PUT)


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **body** | [**Binary**](Binary.md)|  | [optional] 

### Return type

void (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdProfileImagePut**
> RestapiV10AccountAccountIdExtensionExtensionIdProfileImagePut($accountId, $extensionId, $body)



Update Profile Image


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **body** | [**Binary**](Binary.md)|  | [optional] 

### Return type

void (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdProfileImageScaleSizeGet**
> Binary RestapiV10AccountAccountIdExtensionExtensionIdProfileImageScaleSizeGet($accountId, $extensionId, $scaleSize)



Get Scaled Profile Image


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **scaleSize** | **string**| Dimensions of a profile image which will be returned in response. | 

### Return type

[**Binary**](Binary.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdPut**
> ExtensionInfo RestapiV10AccountAccountIdExtensionExtensionIdPut($accountId, $extensionId, $body)



Update Extension by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **body** | [**interface{}**](interface{}.md)|  | [optional] 

### Return type

[**ExtensionInfo**](ExtensionInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionGet**
> InlineResponseDefault7 RestapiV10AccountAccountIdExtensionGet($accountId, $page, $perPage, $status, $type_)



Get Extension List


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **page** | **int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | [optional] 
 **perPage** | **int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default. | [optional] 
 **status** | **string**| Extension current state. Multiple values are supported. If &#39;Unassigned&#39; is specified, then extensions without extensionNumber are returned. If not specified, then all extensions are returned | [optional] 
 **type_** | **string**| Extension type. Multiple values are supported | [optional] 

### Return type

[**InlineResponseDefault7**](inline_response_default_7.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

