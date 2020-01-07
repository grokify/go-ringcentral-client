# \UsersApi

All URIs are relative to *https://DOMAIN.api.engagement.dimelo.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](UsersApi.md#CreateUser) | **Post** /users | Creating a user
[**DeleteUser**](UsersApi.md#DeleteUser) | **Delete** /users/{userId} | Deleting a user
[**GetAllUsers**](UsersApi.md#GetAllUsers) | **Get** /users | Getting all users
[**GetUser**](UsersApi.md#GetUser) | **Get** /users/{userId} | Getting a user from its id
[**InviteUser**](UsersApi.md#InviteUser) | **Post** /users/invite | Inviting a user
[**UpdateUser**](UsersApi.md#UpdateUser) | **Put** /users/{userId} | Updating a user



## CreateUser

> User CreateUser(ctx, email, firstname, lastname, password, roleId, optional)
Creating a user

This method creates a new user. In case of success it renders the created user, otherwise, it renders an error (422 HTTP code).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string**| User email (mandatory). | 
**firstname** | **string**| User firstname (mandatory). | 
**lastname** | **string**| User lastname (mandatory). | 
**password** | **string**| User plain password (mandatory). | 
**roleId** | **string**| User role id (mandatory). | 
 **optional** | ***CreateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **categoryIds** | [**optional.Interface of []string**](string.md)| User list of category ids (multiple). | 
 **enabled** | **optional.Bool**| Whether the user is enabled or not (boolean). | 
 **externalId** | **optional.String**| User external id, used for SSO. | 
 **gender** | **optional.String**| User gender (\&quot;man\&quot; or \&quot;woman\&quot;). | 
 **identityIds** | [**optional.Interface of []string**](string.md)| User list of identity ids (multiple). | 
 **locale** | **optional.String**| Language for the user interface. | 
 **nickname** | **optional.String**| User nickname. | 
 **teamIds** | [**optional.Interface of []string**](string.md)| User list of team ids (multiple). | 
 **timezone** | **optional.String**| Use the timezone endpoint to get the timezone name (String), default is empty for domain timezone. | 
 **spokenLanguages** | [**optional.Interface of []string**](string.md)| List of locales corresponding to the languages spoken by the user (multiple). | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> User DeleteUser(ctx, userId)
Deleting a user

This method deletes the given user. In case of success it renders the deleted user, otherwise, it renders an error (422 HTTP code).  Authorization​: only users that can update users. The user affiliated to the token must have at least all the permissions of the other user. If the user affiliated to the token has the manage_users_of_my_teams permission, the deleted user will need to belong to at least one of the teams he’s the leader of.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**|  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllUsers

> GetAllUsersResponse GetAllUsers(ctx, optional)
Getting all users

This method renders users ordered by creation date (descending).  Authorization​: only users that can view users. If the user affiliated to the token has the manage_users_of_my_teams permission, only the users belonging to at least one of the teams he’s the leader of will be returned.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllUsersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **optional.String**| To filter users on given email. | 
 **categoryId** | **optional.String**| To filter users on given category id. | 
 **identityId** | **optional.String**| To filter users on given identity id. | 
 **externalId** | **optional.String**| To filter users on given external id. | 
 **roleId** | **optional.String**| To filter users on given role id. | 
 **teamId** | **optional.String**| To filter users on given team id. | 
 **offset** | **optional.Int32**| The record index to start. Default value is 0. | 
 **limit** | **optional.Int32**| The max number of records to return. Default value is 30, max value is 150. | 

### Return type

[**GetAllUsersResponse**](GetAllUsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> User GetUser(ctx, userId)
Getting a user from its id

This method renders a user from given id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**|  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteUser

> User InviteUser(ctx, email, firstname, lastname, roleId, optional)
Inviting a user

This method invites a new user. In case of success it renders the created user, otherwise, it renders an error (422 HTTP code).  Authorization​: only users that can invite other users. If the user affiliated to the token has the manage_users_of_my_teams permission, the invited user will need to belong to at least one of the teams he’s the leader of. It will not be possible to assign the user to other teams.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string**| User email (mandatory). | 
**firstname** | **string**| User firstname (mandatory). | 
**lastname** | **string**| User lastname (mandatory). | 
**roleId** | **string**| User role id (mandatory). | 
 **optional** | ***InviteUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InviteUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **categoryIds** | [**optional.Interface of []string**](string.md)| User list of category ids (multiple). | 
 **enabled** | **optional.Bool**| Whether the user is enabled or not (boolean). | 
 **externalId** | **optional.String**| User external id. | 
 **gender** | **optional.String**| User gender (\&quot;man\&quot; or \&quot;woman\&quot;). | 
 **identityIds** | [**optional.Interface of []string**](string.md)| User list of identity ids (multiple). | 
 **locale** | **optional.String**| Language for the user interface. | 
 **nickname** | **optional.String**| User nickname. | 
 **teamIds** | [**optional.Interface of []string**](string.md)| User list of team ids (multiple). | 
 **timezone** | **optional.String**| Use the timezone endpoint to get the timezone name (String), default is empty for | 
 **spokenLanguages** | [**optional.Interface of []string**](string.md)| List of locales corresponding to the languages spoken by the user (multiple). | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> User UpdateUser(ctx, userId, optional)
Updating a user

This method updates users from given attributes and renders it in case of success.  Authorization​: only users that can update users. If the user affiliated to the token has the `manage_users_of_my_teams` permission, the updated user will need to belong to at least one of the teams he’s the leader of. The teams the user affiliated to the token is the leader of will be the only ones which can be added or removed.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**|  | 
 **optional** | ***UpdateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **categoryIds** | [**optional.Interface of []string**](string.md)| User list of category ids (multiple). | 
 **email** | **optional.String**| User email. | 
 **enabled** | **optional.Bool**| Whether the user is enabled or not (boolean). | 
 **externalId** | **optional.String**| User external id, used for SSO. | 
 **firstname** | **optional.String**| User firstname. | 
 **gender** | **optional.String**| User gender (\&quot;man\&quot; or \&quot;woman\&quot;). | 
 **identityIds** | [**optional.Interface of []string**](string.md)| User list of identity ids (multiple). | 
 **lastname** | **optional.String**| User lastname. | 
 **locale** | **optional.String**| Language for the user interface. | 
 **nickname** | **optional.String**| User nickname. | 
 **password** | **optional.String**| User plain password. | 
 **roleId** | **optional.String**| User role id. | 
 **teamIds** | [**optional.Interface of []string**](string.md)| User list of team ids (multiple). | 
 **timezone** | **optional.String**| Use the timezone endpoint to get the timezone name (String), default is empty for domain timezone. | 
 **spokenLanguages** | [**optional.Interface of []string**](string.md)| List of locales corresponding to the languages spoken by the user (multiple). | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

