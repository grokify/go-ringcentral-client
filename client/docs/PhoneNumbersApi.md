# \PhoneNumbersApi

All URIs are relative to *https://api.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListExtensionPhoneNumbers**](PhoneNumbersApi.md#ListExtensionPhoneNumbers) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/phone-number | Get Extension Phone Number List


# **ListExtensionPhoneNumbers**
> GetExtensionPhoneNumbersResponse ListExtensionPhoneNumbers(ctx, accountId, extensionId, optional)
Get Extension Phone Number List

<p style='font-style:italic;'>Since 1.0.2</p><p>Returns the list of phone numbers that are used by a particular extension, and can be filtered by the phone number type. The returned list contains all numbers which are directly mapped to a given extension plus the features and also company-level numbers which may be used when performing different operations on behalf of this extension.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **usageType** | [**[]string**](string.md)| Usage type of the phone number | 
 **page** | **int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | 
 **perPage** | **int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | 

### Return type

[**GetExtensionPhoneNumbersResponse**](GetExtensionPhoneNumbersResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

