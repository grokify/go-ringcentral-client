# \UserContactsApi

All URIs are relative to *https://platform.devtest.ringcentral.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContactFavorite**](UserContactsApi.md#ContactFavorite) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/favorite | Get Favorite Contacts
[**CreateContact**](UserContactsApi.md#CreateContact) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact | Create Contact
[**CreateContacts**](UserContactsApi.md#CreateContacts) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/favorite | Update Favorite Contacts
[**DeleteContact**](UserContactsApi.md#DeleteContact) | **Delete** /restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact/{contactId} | Delete Contact(s) by ID
[**ListContacts**](UserContactsApi.md#ListContacts) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact | Get Contacts
[**LoadContact**](UserContactsApi.md#LoadContact) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact/{contactId} | Get Contact(s) by ID
[**SyncAddressBook**](UserContactsApi.md#SyncAddressBook) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book-sync | Address Book Synchronization
[**UpdateContact**](UserContactsApi.md#UpdateContact) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact/{contactId} | Update Contact(s) by ID


# **ContactFavorite**
> ContactFavorite(ctx, accountId, extensionId)
Get Favorite Contacts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateContact**
> PersonalContactResource CreateContact(ctx, accountId, extensionId, optional)
Create Contact

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>Contacts</td><td>Creating, viewing, editing and deleting user personal contacts</td></tr><tr><td class='code'>ReadContacts</td><td>Viewing user personal contacts</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***CreateContactOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateContactOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **personalContactResource** | [**optional.Interface of PersonalContactResource**](PersonalContactResource.md)|  | 

### Return type

[**PersonalContactResource**](PersonalContactResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateContacts**
> CreateContacts(ctx, accountId, extensionId, optional)
Update Favorite Contacts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***CreateContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateContactsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **favoriteCollection** | [**optional.Interface of FavoriteCollection**](FavoriteCollection.md)|  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteContact**
> DeleteContact(ctx, accountId, extensionId, contactId)
Delete Contact(s) by ID

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>Contacts</td><td>Creating, viewing, editing and deleting user personal contacts</td></tr><tr><td class='code'>ReadContacts</td><td>Viewing user personal contacts</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
  **contactId** | **int32**| Internal identifier of a contact record in the RingCentral database | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListContacts**
> ContactList ListContacts(ctx, accountId, extensionId, optional)
Get Contacts

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadContacts</td><td>Viewing user personal contacts</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***ListContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListContactsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startsWith** | **optional.String**| If specified, only contacts whose First name or Last name start with the mentioned substring are returned. Case-insensitive | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sorts results by the specified property. The default is &#39;First Name&#39; | 
 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | 
 **perPage** | **optional.Int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | 
 **phoneNumber** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**ContactList**](ContactList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadContact**
> PersonalContactResource LoadContact(ctx, accountId, extensionId, contactId)
Get Contact(s) by ID

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadContacts</td><td>Viewing user personal contacts</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
  **contactId** | **int32**| Internal identifier of a contact record in the RingCentral database | 

### Return type

[**PersonalContactResource**](PersonalContactResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncAddressBook**
> AddressBookSync SyncAddressBook(ctx, accountId, extensionId, optional)
Address Book Synchronization

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadContacts</td><td>Viewing user personal contacts</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***SyncAddressBookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SyncAddressBookOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **syncType** | [**optional.Interface of []string**](string.md)| Type of synchronization. The default value is &#39;FSync&#39; | 
 **syncToken** | **optional.String**| Value of syncToken property of the last sync request response | 
 **perPage** | **optional.Int32**| Number of records per page to be returned. The max number of records is 250, which is also the default. For FSync ??? if the number of records exceeds the parameter value (either specified or default), all of the pages can be retrieved in several requests. For ISync ??? if the number of records exceeds the page size, the number of incoming changes to this number is limited | 
 **pageId** | **optional.Int32**| Internal identifier of a page. It can be obtained from the &#39;nextPageId&#39; parameter passed in response body | 

### Return type

[**AddressBookSync**](AddressBookSync.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateContact**
> PersonalContactResource UpdateContact(ctx, accountId, extensionId, contactId, optional)
Update Contact(s) by ID

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>Contacts</td><td>Creating, viewing, editing and deleting user personal contacts</td></tr><tr><td class='code'>ReadContacts</td><td>Viewing user personal contacts</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
  **contactId** | **int32**| Internal identifier of a contact record in the RingCentral database | 
 **optional** | ***UpdateContactOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateContactOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **personalContactResource** | [**optional.Interface of PersonalContactResource**](PersonalContactResource.md)|  | 

### Return type

[**PersonalContactResource**](PersonalContactResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

