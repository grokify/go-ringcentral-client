# \UserContactsApi

All URIs are relative to *https://api.ringcentral.com*

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
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateContact**
> PersonalContactResource CreateContact(ctx, accountId, extensionId, optional)
Create Contact

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>Contacts</td><td>Creating, viewing, editing and deleting user personal contacts</td></tr><tr><td class='code'>ReadContacts</td><td>Viewing user personal contacts</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

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
 **body** | [**PersonalContactResource**](PersonalContactResource.md)|  | 

### Return type

[**PersonalContactResource**](PersonalContactResource.md)

### Authorization

No authorization required

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
 **body** | [**FavoriteCollection**](FavoriteCollection.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteContact**
> DeleteContact(ctx, accountId, extensionId, contactId)
Delete Contact(s) by ID

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>Contacts</td><td>Creating, viewing, editing and deleting user personal contacts</td></tr><tr><td class='code'>ReadContacts</td><td>Viewing user personal contacts</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
  **contactId** | **int32**| Internal identifier of a contact record in the RingCentral database | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListContacts**
> ContactList ListContacts(ctx, accountId, extensionId, optional)
Get Contacts

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadContacts</td><td>Viewing user personal contacts</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

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
 **startsWith** | **string**| If specified, only contacts whose First name or Last name start with the mentioned substring are returned. Case-insensitive | 
 **sortBy** | [**[]string**](string.md)| Sorts results by the specified property. The default is &#39;First Name&#39; | 
 **page** | **int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | 
 **perPage** | **int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | 
 **phoneNumber** | [**[]string**](string.md)|  | 

### Return type

[**ContactList**](ContactList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadContact**
> PersonalContactResource LoadContact(ctx, accountId, extensionId, contactId)
Get Contact(s) by ID

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadContacts</td><td>Viewing user personal contacts</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
  **contactId** | **int32**| Internal identifier of a contact record in the RingCentral database | 

### Return type

[**PersonalContactResource**](PersonalContactResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncAddressBook**
> AddressBookSync SyncAddressBook(ctx, accountId, extensionId, optional)
Address Book Synchronization

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadContacts</td><td>Viewing user personal contacts</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

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
 **syncType** | [**[]string**](string.md)| Type of synchronization. The default value is &#39;FSync&#39; | 
 **syncToken** | **string**| Value of syncToken property of the last sync request response | 
 **perPage** | **int32**| Number of records per page to be returned. The max number of records is 250, which is also the default. For FSync ??? if the number of records exceeds the parameter value (either specified or default), all of the pages can be retrieved in several requests. For ISync ??? if the number of records exceeds the page size, the number of incoming changes to this number is limited | 
 **pageId** | **int32**| Internal identifier of a page. It can be obtained from the &#39;nextPageId&#39; parameter passed in response body | 

### Return type

[**AddressBookSync**](AddressBookSync.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateContact**
> PersonalContactResource UpdateContact(ctx, accountId, extensionId, contactId, optional)
Update Contact(s) by ID

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>Contacts</td><td>Creating, viewing, editing and deleting user personal contacts</td></tr><tr><td class='code'>ReadContacts</td><td>Viewing user personal contacts</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
  **contactId** | **int32**| Internal identifier of a contact record in the RingCentral database | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **contactId** | **int32**| Internal identifier of a contact record in the RingCentral database | 
 **body** | [**PersonalContactResource**](PersonalContactResource.md)|  | 

### Return type

[**PersonalContactResource**](PersonalContactResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

