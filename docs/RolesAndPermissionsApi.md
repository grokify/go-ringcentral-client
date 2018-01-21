# \RolesAndPermissionsApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckPermission**](RolesAndPermissionsApi.md#CheckPermission) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/authz-profile/check | Check User Permissions
[**GetPermission**](RolesAndPermissionsApi.md#GetPermission) | **Get** /restapi/v1.0/dictionary/permission/{permissionId} | Get Permission Info
[**GetPermissionCategory**](RolesAndPermissionsApi.md#GetPermissionCategory) | **Get** /restapi/v1.0/dictionary/permission-category/{permissionCategoryId} | Get Permission Category
[**GetProfile**](RolesAndPermissionsApi.md#GetProfile) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/authz-profile | Get User Permissions
[**GetSystemRole**](RolesAndPermissionsApi.md#GetSystemRole) | **Get** /restapi/v1.0/dictionary/user-role/{roleId} | Get User Role
[**ListPermissionCategories**](RolesAndPermissionsApi.md#ListPermissionCategories) | **Get** /restapi/v1.0/dictionary/permission-category | Get Permissions Categories
[**ListPermissions**](RolesAndPermissionsApi.md#ListPermissions) | **Get** /restapi/v1.0/dictionary/permission | Get Permissions
[**ListSystemRoles**](RolesAndPermissionsApi.md#ListSystemRoles) | **Get** /restapi/v1.0/dictionary/user-role | Get Standard User Roles
[**LoadUserRole**](RolesAndPermissionsApi.md#LoadUserRole) | **Get** /restapi/v1.0/account/{accountId}/user-role/default | Get Default User Role
[**UpdateUserRole**](RolesAndPermissionsApi.md#UpdateUserRole) | **Put** /restapi/v1.0/account/{accountId}/user-role/default | Set Default User Role


# **CheckPermission**
> AuthProfileCheckResource CheckPermission(ctx, extensionId, accountId, optional)
Check User Permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **extensionId** | **string**|  | 
  **accountId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **extensionId** | **string**|  | 
 **accountId** | **string**|  | 
 **permissionId** | **string**|  | 
 **targetExtensionId** | **string**|  | 

### Return type

[**AuthProfileCheckResource**](AuthProfileCheckResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPermission**
> PermissionResource GetPermission(ctx, permissionId)
Get Permission Info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **permissionId** | **string**|  | 

### Return type

[**PermissionResource**](PermissionResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPermissionCategory**
> PermissionCategoryResource GetPermissionCategory(ctx, permissionCategoryId)
Get Permission Category

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **permissionCategoryId** | **string**|  | 

### Return type

[**PermissionCategoryResource**](PermissionCategoryResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProfile**
> AuthProfileResource GetProfile(ctx, extensionId, accountId)
Get User Permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **extensionId** | **string**|  | 
  **accountId** | **string**|  | 

### Return type

[**AuthProfileResource**](AuthProfileResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSystemRole**
> RoleResource GetSystemRole(ctx, roleId)
Get User Role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **roleId** | **string**|  | 

### Return type

[**RoleResource**](RoleResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPermissionCategories**
> PermissionCategoryCollectionResource ListPermissionCategories(ctx, optional)
Get Permissions Categories

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string**|  | [default to 1]
 **perPage** | **string**|  | [default to 100]
 **servicePlanId** | **string**|  | 

### Return type

[**PermissionCategoryCollectionResource**](PermissionCategoryCollectionResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPermissions**
> PermissionCollectionResource ListPermissions(ctx, optional)
Get Permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string**|  | [default to 1]
 **perPage** | **string**|  | [default to 100]
 **assignable** | **bool**|  | 
 **servicePlanId** | **string**|  | 

### Return type

[**PermissionCollectionResource**](PermissionCollectionResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSystemRoles**
> RolesCollectionResource ListSystemRoles(ctx, optional)
Get Standard User Roles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string**|  | [default to 1]
 **perPage** | **string**|  | [default to 100]
 **servicePlanId** | **string**|  | 

### Return type

[**RolesCollectionResource**](RolesCollectionResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadUserRole**
> LoadUserRole(ctx, accountId)
Get Default User Role

<p style='font-style:italic;'>Since 1.0.28 (Release 8.4)</p><p>Returns account default user role.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>RoleManagement</td><td>Editing and assignment of user roles</td></tr></tbody></table><h4>API Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of an account | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserRole**
> UpdateUserRole(ctx, accountId, body)
Set Default User Role

<p style='font-style:italic;'>Since 1.0.30 (Release 9.1)</p><p>Updates account default user role.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>RoleManagement</td><td>Editing and assignment of user roles</td></tr></tbody></table><h4>API Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of an account | 
  **body** | [**DefaultUserRoleRequest**](DefaultUserRoleRequest.md)| JSON body | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

