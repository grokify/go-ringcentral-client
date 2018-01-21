# \SCIMApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](SCIMApi.md#CreateUser) | **Post** /scim/v2/Users | Create User
[**DeleteUser**](SCIMApi.md#DeleteUser) | **Delete** /scim/v2/Users/{id} | Delete User
[**GetServiceProviderConfig**](SCIMApi.md#GetServiceProviderConfig) | **Get** /scim/v2/ServiceProviderConfig | Get Service Provider Config
[**GetUserById**](SCIMApi.md#GetUserById) | **Get** /scim/v2/Users/{id} | Get User
[**ListUsers**](SCIMApi.md#ListUsers) | **Get** /scim/v2/Users | Get User List
[**PatchUser**](SCIMApi.md#PatchUser) | **Patch** /scim/v2/Users/{id} | Partially update/patch a user
[**SearchUsersViaPost**](SCIMApi.md#SearchUsersViaPost) | **Post** /scim/v2/Users/.search | search or list users
[**UpdateUser**](SCIMApi.md#UpdateUser) | **Put** /scim/v2/Users/{id} | Update or replace user


# **CreateUser**
> UserInfo CreateUser(ctx, body)
Create User

<p style='font-style:italic;'>Since 1.0.31 (Release 9.2)</p><p>Creates a user.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditAccounts</td><td>Viewing and updating user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**UserCreationRequest**](UserCreationRequest.md)| a new user without &#39;id&#39; | 

### Return type

[**UserInfo**](UserInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json, application/scim+json
 - **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUser**
> DeleteUser(ctx, id)
Delete User

<p style='font-style:italic;'>Since 1.0.31 (Release 9.2)</p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditAccounts</td><td>Deleting User using scim</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| Internal identifier of a user | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json, application/scim+json
 - **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceProviderConfig**
> ServiceProviderConfig GetServiceProviderConfig(ctx, )
Get Service Provider Config

<p style='font-style:italic;'>Since 1.0.31 (Release 9.2)</p><p>Returns the list of users requested.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing Service Provider confiog</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ServiceProviderConfig**](ServiceProviderConfig.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json, application/scim+json
 - **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserById**
> UserInfo GetUserById(ctx, id)
Get User

<p style='font-style:italic;'>Since 1.0.31 (Release 9.2)</p><p>Returns a user by ID.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| Internal identifier of a user | 

### Return type

[**UserInfo**](UserInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json, application/scim+json
 - **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsers**
> GetUserListResponse ListUsers(ctx, optional)
Get User List

<p style='font-style:italic;'>Since 1.0.31 (Release 9.2)</p><p>Returns the list of users requested.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string**| only support &#39;userName&#39; or &#39;email&#39; filter expressions for now | 
 **count** | **int32**| page size | [default to 100]
 **startIndex** | **int32**| start index (1-based) | [default to 1]

### Return type

[**GetUserListResponse**](GetUserListResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json, application/scim+json
 - **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchUser**
> UserInfo PatchUser(ctx, id, optional)
Partially update/patch a user

<p style='font-style:italic;'>Since 1.0.31 (Release 9.2)</p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditAccounts</td><td>Partially update/patch a user</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| Internal identifier of a user | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Internal identifier of a user | 
 **body** | [**ScimUserPatch**](ScimUserPatch.md)| patch operations list | 

### Return type

[**UserInfo**](UserInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json, application/scim+json
 - **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchUsersViaPost**
> GetUserListResponse SearchUsersViaPost(ctx, optional)
search or list users

<p style='font-style:italic;'>Since 1.0.31 (Release 9.2)</p><p>Returns the list of users requested.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Searching SCIM Users</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ScimSearchRequestInfo**](ScimSearchRequestInfo.md)| search parameters | 

### Return type

[**GetUserListResponse**](GetUserListResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json, application/scim+json
 - **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser**
> UserInfo UpdateUser(ctx, id, body)
Update or replace user

<p style='font-style:italic;'>Since 1.0.31 (Release 9.2)</p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditExtensions</td><td>Updating User using SCIM</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| Internal identifier of a user | 
  **body** | [**UserUpdateRequest**](UserUpdateRequest.md)| An Exisiting User | 

### Return type

[**UserInfo**](UserInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json, application/scim+json
 - **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

