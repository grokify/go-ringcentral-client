# \AddressBookApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactContactIdDelete**](AddressBookApi.md#RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactContactIdDelete) | **Delete** /restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact/{contactId} | 
[**RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactContactIdGet**](AddressBookApi.md#RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactContactIdGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact/{contactId} | 
[**RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactContactIdPut**](AddressBookApi.md#RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactContactIdPut) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact/{contactId} | 
[**RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactGet**](AddressBookApi.md#RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact | 
[**RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactPost**](AddressBookApi.md#RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactPost) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact | 
[**RestapiV10AccountAccountIdExtensionExtensionIdAddressBookGroupGet**](AddressBookApi.md#RestapiV10AccountAccountIdExtensionExtensionIdAddressBookGroupGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/group | 
[**RestapiV10AccountAccountIdExtensionExtensionIdAddressBookSyncGet**](AddressBookApi.md#RestapiV10AccountAccountIdExtensionExtensionIdAddressBookSyncGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book-sync | 


# **RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactContactIdDelete**
> RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactContactIdDelete($accountId, $extensionId, $contactId)



Delete Contact by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **contactId** | **string**| Internal identifier of a contact record in the RingCentral database | 

### Return type

void (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactContactIdGet**
> PersonalContactInfo RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactContactIdGet($accountId, $extensionId, $contactId)



Get Contact by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **contactId** | **string**| Internal identifier of a contact record in the RingCentral database | 

### Return type

[**PersonalContactInfo**](PersonalContactInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactContactIdPut**
> PersonalContactInfo RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactContactIdPut($accountId, $extensionId, $contactId, $body)



Update Contact by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **contactId** | **string**| Internal identifier of a contact record in the RingCentral database | 
 **body** | [**PersonalContactInfo**](PersonalContactInfo.md)|  | [optional] 

### Return type

[**PersonalContactInfo**](PersonalContactInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactGet**
> InlineResponseDefault9 RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactGet($accountId, $extensionId, $phoneNumber, $startsWith, $sortBy, $page, $perPage)



Get Contact List


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **phoneNumber** | **string**| Phone number in E.164 (11-digits) format with or without plus &#39;+&#39;. Multiple values are supported | [optional] 
 **startsWith** | **string**| If specified, only contacts whose First name or Last name start with the mentioned substring are returned. Case-insensitive | [optional] 
 **sortBy** | **string**| Sorts results by the specified property. The default is &#39;First Name&#39; | [optional] 
 **page** | **int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | [optional] 
 **perPage** | **int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | [optional] 

### Return type

[**InlineResponseDefault9**](inline_response_default_9.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactPost**
> PersonalContactInfo RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactPost($accountId, $extensionId, $body)



Create New Contact


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **body** | [**PersonalContactInfo**](PersonalContactInfo.md)|  | [optional] 

### Return type

[**PersonalContactInfo**](PersonalContactInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdAddressBookGroupGet**
> InlineResponseDefault10 RestapiV10AccountAccountIdExtensionExtensionIdAddressBookGroupGet($accountId, $extensionId)



Get Contact Group List


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 

### Return type

[**InlineResponseDefault10**](inline_response_default_10.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdAddressBookSyncGet**
> InlineResponseDefault8 RestapiV10AccountAccountIdExtensionExtensionIdAddressBookSyncGet($accountId, $extensionId, $syncType, $syncToken, $perPage, $pageId)



Contacts Synchronization


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **syncType** | **string**| Type of synchronization. The default value is &#39;FSync&#39; | [optional] 
 **syncToken** | **string**| Value of syncToken property of the last sync request response | [optional] 
 **perPage** | **int32**| Number of records per page to be returned. The max number of records is 250, which is also the default. For FSync — if the number of records exceeds the parameter value (either specified or default), all of the pages can be retrieved in several requests. For ISync — if the number of records exceeds the page size, the number of incoming changes to this number is limited | [optional] 
 **pageId** | **int32**| Internal identifier of a page. It can be obtained from the &#39;nextPageId&#39; parameter passed in response body | [optional] 

### Return type

[**InlineResponseDefault8**](inline_response_default_8.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

