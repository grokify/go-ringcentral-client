# \CompanyContactsApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBotExtension**](CompanyContactsApi.md#CreateBotExtension) | **Post** /restapi/v1.0/account/{accountId}/bot | Create Bot Extension
[**CreateExtension**](CompanyContactsApi.md#CreateExtension) | **Post** /restapi/v1.0/account/{accountId}/extension | Create Extension
[**ListCompanyDirectoryContacts**](CompanyContactsApi.md#ListCompanyDirectoryContacts) | **Get** /restapi/v1.0/account/{accountId}/directory/contacts | Get Company Directory
[**ListExtensions**](CompanyContactsApi.md#ListExtensions) | **Get** /restapi/v1.0/account/{accountId}/extension | Get Extensions
[**LoadCompanyDirectoryContact**](CompanyContactsApi.md#LoadCompanyDirectoryContact) | **Get** /restapi/v1.0/account/{accountId}/directory/contacts/{contactId} | Get Company Directory Contact
[**LoadCompanyDirectoryExtensionNumbers**](CompanyContactsApi.md#LoadCompanyDirectoryExtensionNumbers) | **Get** /restapi/v1.0/account/{accountId}/directory/extension-numbers | Get Existing Extension Numbers
[**LoadCompanyDirectoryFederation**](CompanyContactsApi.md#LoadCompanyDirectoryFederation) | **Get** /restapi/v1.0/account/{accountId}/directory/federation | Get Account Federation
[**LoadCompanyDirectoryFederationConflicts**](CompanyContactsApi.md#LoadCompanyDirectoryFederationConflicts) | **Get** /restapi/v1.0/account/{accountId}/directory/federation-conflicts | Get Account Federation Conflicts


# **CreateBotExtension**
> BotExtensionCreation CreateBotExtension(ctx, accountId, botExtensionCreationRequest)
Create Bot Extension

<p style='font-style:italic;'>Since 1.0.32 (Release 9.3)</p><p>Creates a bot extension. Please note: Bot extension is always created in Enabled status, no welcome email is sent.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditAccounts</td><td>Viewing and updating user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **botExtensionCreationRequest** | [**BotExtensionCreationRequest**](BotExtensionCreationRequest.md)| JSON body | 

### Return type

[**BotExtensionCreation**](BotExtensionCreation.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateExtension**
> ExtensionInfo CreateExtension(ctx, accountId, extensionCreationRequest)
Create Extension

<p style='font-style:italic;'>Since 1.0.10 (Release 6.2)</p><p>Creates an extension.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditAccounts</td><td>Viewing and updating user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionCreationRequest** | [**ExtensionCreationRequest**](ExtensionCreationRequest.md)| JSON body | 

### Return type

[**ExtensionInfo**](ExtensionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCompanyDirectoryContacts**
> CompanyDirectoryContacts ListCompanyDirectoryContacts(ctx, accountId, optional)
Get Company Directory

<p style='font-style:italic;'></p><p>Returns contact information on corporate users of federated accounts. Please note: 1. User, DigitalUser, VirtualUser and FaxUser types are returned as User type. 2.ApplicationExtension type is not returned. 3. Only extensions in Enabled, Disabled and NotActivated state are returned.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
 **optional** | ***ListCompanyDirectoryContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListCompanyDirectoryContactsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int64**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | 
 **perPage** | **optional.Int64**| Indicates the page size (number of items). The possible values are: Max, all or a numeric value. If not specified, all records are returned on one page | 
 **excludeFederatedContacts** | **optional.Bool**| If &#39;True&#39; then only contacts of current account are returned, if &#39;False&#39; then all contacts of all federation accounts are returned | 

### Return type

[**CompanyDirectoryContacts**](CompanyDirectoryContacts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListExtensions**
> GetExtensionListResponse ListExtensions(ctx, accountId, optional)
Get Extensions

<p style='font-style:italic;'>Since 1.0.0</p><p>Returns the list of extensions created for a particular account. All types of extensions are included in this list.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
 **optional** | ***ListExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListExtensionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int64**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | 
 **perPage** | **optional.Int64**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default. | 
 **status** | [**optional.Interface of []string**](string.md)| Extension current state. Multiple values are supported. If &#39;Unassigned&#39; is specified, then extensions without extensionNumber are returned. If not specified, then all extensions are returned | 
 **type_** | [**optional.Interface of []string**](string.md)| Extension type. Multiple values are supported | 

### Return type

[**GetExtensionListResponse**](GetExtensionListResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadCompanyDirectoryContact**
> ContactInfo LoadCompanyDirectoryContact(ctx, accountId, contactId)
Get Company Directory Contact

<p style='font-style:italic;'></p><p>Returns contact information on a particular corporate user of a federated account.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **contactId** | **string**| Internal identifier of an extension | 

### Return type

[**ContactInfo**](ContactInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadCompanyDirectoryExtensionNumbers**
> FederationExtensionNumbers LoadCompanyDirectoryExtensionNumbers(ctx, accountId)
Get Existing Extension Numbers

<p style='font-style:italic;'></p><p>Returns the list of account federation extensions.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]

### Return type

[**FederationExtensionNumbers**](FederationExtensionNumbers.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadCompanyDirectoryFederation**
> FederationInfo LoadCompanyDirectoryFederation(ctx, accountId)
Get Account Federation

<p style='font-style:italic;'></p><p>Returns information on a federation and associated accounts.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]

### Return type

[**FederationInfo**](FederationInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadCompanyDirectoryFederationConflicts**
> FederationConflicts LoadCompanyDirectoryFederationConflicts(ctx, accountId)
Get Account Federation Conflicts

<p style='font-style:italic;'></p><p>Returns the list of federation extensions which numbers are conflicting with extension numbers of the requested account.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]

### Return type

[**FederationConflicts**](FederationConflicts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

