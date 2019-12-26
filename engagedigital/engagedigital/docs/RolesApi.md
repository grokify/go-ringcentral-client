# \RolesApi

All URIs are relative to *https://DOMAIN.api.engagement.dimelo.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRole**](RolesApi.md#CreateRole) | **Post** /roles | Creating a role
[**GetAllRoles**](RolesApi.md#GetAllRoles) | **Get** /roles | Getting all roles
[**GetRole**](RolesApi.md#GetRole) | **Get** /roles/{roleId} | Getting a role from its id
[**UpdateRole**](RolesApi.md#UpdateRole) | **Put** /roles/{roleId} | Updating a role



## CreateRole

> Role CreateRole(ctx, label, optional)
Creating a role

This method creates a new role. In case of success it renders the created role, otherwise, it renders an error (422 HTTP code).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**label** | **string**|  | 
 **optional** | ***CreateRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateRoleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessHelpCenter** | **optional.Bool**|  | 
 **accessPreviousMessages** | **optional.Bool**|  | 
 **accessPullMode** | **optional.Bool**|  | 
 **adminStampAnswer** | **optional.Bool**|  | 
 **approveContent** | **optional.Bool**|  | 
 **assignIntervention** | **optional.Bool**|  | 
 **authorBlockContent** | **optional.Bool**|  | 
 **closeContentThread** | **optional.Bool**|  | 
 **createAndDestroyExtension** | **optional.Bool**|  | 
 **createCommunity** | **optional.Bool**|  | 
 **createContentSource** | **optional.Bool**|  | 
 **createUser** | **optional.Bool**|  | 
 **delayExportContent** | **optional.Bool**|  | 
 **deleteContentThread** | **optional.Bool**|  | 
 **impersonateUser** | **optional.Bool**|  | 
 **inviteUser** | **optional.Bool**|  | 
 **manageApiAccessTokens** | **optional.Bool**|  | 
 **manageAppSdkApplications** | **optional.Bool**|  | 
 **manageAutomaticExportsTasks** | **optional.Bool**| permission only available with the corresponding extension enabled | 
 **manageCategories** | **optional.Bool**|  | 
 **manageChat** | **optional.Bool**|  | 
 **manageCustomFields** | **optional.Bool**|  | 
 **manageCustomNotifications** | **optional.Bool**|  | 
 **manageEmailsTemplates** | **optional.Bool**|  | 
 **manageFolders** | **optional.Bool**|  | 
 **manageIce** | **optional.Bool**|  | 
 **manageIdentities** | **optional.Bool**|  | 
 **manageOwnNotifications** | **optional.Bool**|  | 
 **manageReplyAssistant** | **optional.Bool**| permission only available with the corresponding extension enabled | 
 **manageRoles** | **optional.Bool**|  | 
 **manageRulesEngineRules** | **optional.Bool**| permission only available with the corresponding extension enabled | 
 **manageSurveys** | **optional.Bool**| permission only available with the corresponding extension enabled | 
 **manageTags** | **optional.Bool**|  | 
 **manageTeams** | **optional.Bool**|  | 
 **manageTopologies** | **optional.Bool**|  | 
 **manageUsersOfMyTeams** | **optional.Bool**|  | 
 **monitorTasks** | **optional.Bool**|  | 
 **monitorTeamTasks** | **optional.Bool**|  | 
 **muteContent** | **optional.Bool**|  | 
 **openContentThread** | **optional.Bool**|  | 
 **publishContent** | **optional.Bool**|  | 
 **readCommunity** | **optional.Bool**|  | 
 **readContentSource** | **optional.Bool**|  | 
 **readEvent** | **optional.Bool**|  | 
 **readExport** | **optional.Bool**|  | 
 **readIdentity** | **optional.Bool**|  | 
 **readOwnStats** | **optional.Bool**|  | 
 **readPresence** | **optional.Bool**|  | 
 **readStats** | **optional.Bool**|  | 
 **readSurveys** | **optional.Bool**| permission only available with the corresponding extension enabled | 
 **readUser** | **optional.Bool**|  | 
 **receiveTasks** | **optional.Bool**|  | 
 **replyWithAssistant** | **optional.Bool**| permission only available with the corresponding extension enabled | 
 **searchContents** | **optional.Bool**|  | 
 **searchEvent** | **optional.Bool**|  | 
 **updateCommunity** | **optional.Bool**|  | 
 **updateContentSource** | **optional.Bool**|  | 
 **updateExtension** | **optional.Bool**|  | 
 **updateIdentity** | **optional.Bool**|  | 
 **updateIntervention** | **optional.Bool**|  | 
 **updateOwnIntervention** | **optional.Bool**|  | 
 **updateSettings** | **optional.Bool**|  | 
 **updateTimeSheet** | **optional.Bool**|  | 
 **updateUser** | **optional.Bool**|  | 
 **useEmoji** | **optional.Bool**|  | 

### Return type

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllRoles

> GetAllRolesResponse GetAllRoles(ctx, optional)
Getting all roles

This method renders roles ordered by creation date (ascending).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllRolesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| The record index to start. Default value is 0. | 
 **limit** | **optional.Int32**| The max number of records to return. Default value is 30, max value is 150. | 

### Return type

[**GetAllRolesResponse**](GetAllRolesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRole

> Role GetRole(ctx, roleId)
Getting a role from its id

This method renders a role from given id.  Authorization​: only users that can manage roles.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string**|  | 

### Return type

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> Role UpdateRole(ctx, roleId, optional)
Updating a role

This method updates an existing role from given attributes and renders it in case of success.  Authorization​: A user can’t update roles with more permissions than himself and can’t give a role a permission he doesn’t have.  Any permission updated with a user that does not have this permission will be ignored (The update is done, just not the unallowed permission)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string**|  | 
 **optional** | ***UpdateRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateRoleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessHelpCenter** | **optional.Bool**|  | 
 **accessPreviousMessages** | **optional.Bool**|  | 
 **accessPullMode** | **optional.Bool**|  | 
 **adminStampAnswer** | **optional.Bool**|  | 
 **approveContent** | **optional.Bool**|  | 
 **assignIntervention** | **optional.Bool**|  | 
 **authorBlockContent** | **optional.Bool**|  | 
 **closeContentThread** | **optional.Bool**|  | 
 **createAndDestroyExtension** | **optional.Bool**|  | 
 **createCommunity** | **optional.Bool**|  | 
 **createContentSource** | **optional.Bool**|  | 
 **createUser** | **optional.Bool**|  | 
 **delayExportContent** | **optional.Bool**|  | 
 **deleteContentThread** | **optional.Bool**|  | 
 **impersonateUser** | **optional.Bool**|  | 
 **inviteUser** | **optional.Bool**|  | 
 **manageApiAccessTokens** | **optional.Bool**|  | 
 **manageAppSdkApplications** | **optional.Bool**|  | 
 **manageAutomaticExportsTasks** | **optional.Bool**| permission only available with the corresponding extension enabled | 
 **manageCategories** | **optional.Bool**|  | 
 **manageChat** | **optional.Bool**|  | 
 **manageCustomFields** | **optional.Bool**|  | 
 **manageCustomNotifications** | **optional.Bool**|  | 
 **manageEmailsTemplates** | **optional.Bool**|  | 
 **manageFolders** | **optional.Bool**|  | 
 **manageIce** | **optional.Bool**|  | 
 **manageIdentities** | **optional.Bool**|  | 
 **manageOwnNotifications** | **optional.Bool**|  | 
 **manageReplyAssistant** | **optional.Bool**| permission only available with the corresponding extension enabled | 
 **manageRoles** | **optional.Bool**|  | 
 **manageRulesEngineRules** | **optional.Bool**| permission only available with the corresponding extension enabled | 
 **manageSurveys** | **optional.Bool**| permission only available with the corresponding extension enabled | 
 **manageTags** | **optional.Bool**|  | 
 **manageTeams** | **optional.Bool**|  | 
 **manageTopologies** | **optional.Bool**|  | 
 **manageUsersOfMyTeams** | **optional.Bool**|  | 
 **monitorTasks** | **optional.Bool**|  | 
 **monitorTeamTasks** | **optional.Bool**|  | 
 **muteContent** | **optional.Bool**|  | 
 **openContentThread** | **optional.Bool**|  | 
 **publishContent** | **optional.Bool**|  | 
 **readCommunity** | **optional.Bool**|  | 
 **readContentSource** | **optional.Bool**|  | 
 **readEvent** | **optional.Bool**|  | 
 **readExport** | **optional.Bool**|  | 
 **readIdentity** | **optional.Bool**|  | 
 **readOwnStats** | **optional.Bool**|  | 
 **readPresence** | **optional.Bool**|  | 
 **readStats** | **optional.Bool**|  | 
 **readSurveys** | **optional.Bool**| permission only available with the corresponding extension enabled | 
 **readUser** | **optional.Bool**|  | 
 **receiveTasks** | **optional.Bool**|  | 
 **replyWithAssistant** | **optional.Bool**| permission only available with the corresponding extension enabled | 
 **searchContents** | **optional.Bool**|  | 
 **searchEvent** | **optional.Bool**|  | 
 **updateCommunity** | **optional.Bool**|  | 
 **updateContentSource** | **optional.Bool**|  | 
 **updateExtension** | **optional.Bool**|  | 
 **updateIdentity** | **optional.Bool**|  | 
 **updateIntervention** | **optional.Bool**|  | 
 **updateOwnIntervention** | **optional.Bool**|  | 
 **updateSettings** | **optional.Bool**|  | 
 **updateTimeSheet** | **optional.Bool**|  | 
 **updateUser** | **optional.Bool**|  | 
 **useEmoji** | **optional.Bool**|  | 

### Return type

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

