# \UserSettingsApi

All URIs are relative to *https://platform.devtest.ringcentral.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteExtension**](UserSettingsApi.md#DeleteExtension) | **Delete** /restapi/v1.0/account/{accountId}/extension/{extensionId} | Delete Extension
[**DownloadImage**](UserSettingsApi.md#DownloadImage) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/profile-image | Get User Profile Image
[**GetNotificationSettings**](UserSettingsApi.md#GetNotificationSettings) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/notification-settings | Get Notification Settings
[**ListExtensionGrants**](UserSettingsApi.md#ListExtensionGrants) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/grant | Get Extension Grants
[**ListSecretQuestions**](UserSettingsApi.md#ListSecretQuestions) | **Get** /restapi/v1.0/dictionary/secret-question | Get Secret Questions
[**LoadConferencingInfo**](UserSettingsApi.md#LoadConferencingInfo) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/conferencing | Get User Conferencing Settings
[**LoadExtensionCallerId**](UserSettingsApi.md#LoadExtensionCallerId) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/caller-id | Get Extension Caller ID
[**LoadExtensionInfo**](UserSettingsApi.md#LoadExtensionInfo) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId} | Get Extension Info
[**LoadExtensionProfileImage**](UserSettingsApi.md#LoadExtensionProfileImage) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/profile-image/{scaleSize} | Get Extension Profile Image (Scaled)
[**LoadExtensionUserCredentials**](UserSettingsApi.md#LoadExtensionUserCredentials) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/credentials | Get User Credentials
[**LoadSecretQuestion**](UserSettingsApi.md#LoadSecretQuestion) | **Get** /restapi/v1.0/dictionary/secret-question/{questionId} | Get Secret Question
[**UpdateConferencingInfo**](UserSettingsApi.md#UpdateConferencingInfo) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/conferencing | Update User Conferencing Settings
[**UpdateExtension**](UserSettingsApi.md#UpdateExtension) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId} | Update Extension
[**UpdateExtensionCallerId**](UserSettingsApi.md#UpdateExtensionCallerId) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/caller-id | Update Extension Caller ID
[**UpdateExtensionUserCredentials**](UserSettingsApi.md#UpdateExtensionUserCredentials) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/credentials | Update User Credentials
[**UpdateNotificationSettings**](UserSettingsApi.md#UpdateNotificationSettings) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/notification-settings | Update Notification Settings
[**UploadImageByPostForm**](UserSettingsApi.md#UploadImageByPostForm) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/profile-image | Upload User Profile Image
[**UploadImageByPutForm**](UserSettingsApi.md#UploadImageByPutForm) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/profile-image | Update User Profile Image
[**ValidateExtensionUserCredentials**](UserSettingsApi.md#ValidateExtensionUserCredentials) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/credentials/validate | Validate User Credentials
[**VerifyExtensionUserCredentials**](UserSettingsApi.md#VerifyExtensionUserCredentials) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/credentials/verify | Verify Extension User Credentials


# **DeleteExtension**
> DeleteExtension(ctx, extensionId, accountId)
Delete Extension

<p style='font-style:italic;'>Since 1.0.10 (Release 6.2)</p><p>Deletes extension(s) by ID(s).</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditAccounts</td><td>Viewing and updating user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **extensionId** | **string**|  | 
  **accountId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadImage**
> Binary DownloadImage(ctx, extensionId, accountId)
Get User Profile Image

<p style='font-style:italic;'>Since 1.0.20 (Release 7.4)</p><p>Returns the extension profile image.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **extensionId** | **string**|  | 
  **accountId** | **string**|  | 

### Return type

[**Binary**](Binary.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, image/png, image/jpeg, image/gif

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNotificationSettings**
> NotificationSettings GetNotificationSettings(ctx, accountId, extensionId)
Get Notification Settings

<p style='font-style:italic;'>Since 1.0.26 (Release 8.2)</p><p>Returns notification settings for the current extension.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 

### Return type

[**NotificationSettings**](NotificationSettings.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListExtensionGrants**
> GetExtensionGrantListResponse ListExtensionGrants(ctx, accountId, extensionId, optional)
Get Extension Grants

<p style='font-style:italic;'></p><p>Returns the list of extension grants.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***ListExtensionGrantsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListExtensionGrantsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | 
 **perPage** | **optional.Int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | 

### Return type

[**GetExtensionGrantListResponse**](GetExtensionGrantListResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSecretQuestions**
> GetSecretQuestionListResponse ListSecretQuestions(ctx, optional)
Get Secret Questions

<p style='font-style:italic;'>Since 1.0.20 (Release 7.4)</p><p>Returns the list of secret questions for a specific language.</p><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListSecretQuestionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListSecretQuestionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signupOnly** | **optional.Bool**|  | [default to false]
 **page** | **optional.String**|  | [default to 1]
 **perPage** | **optional.String**|  | [default to 100]

### Return type

[**GetSecretQuestionListResponse**](GetSecretQuestionListResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadConferencingInfo**
> GetConferencingInfoResponse LoadConferencingInfo(ctx, accountId, extensionId, optional)
Get User Conferencing Settings

<p style='font-style:italic;'>Since 1.0.4 (Release 5.13)</p><p>Returns the information on the Free Conference Calling (FCC) feature for a given extension.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***LoadConferencingInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoadConferencingInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **countryId** | **optional.String**| Internal identifier of a country. If not specified, the response is returned for the brand country | 

### Return type

[**GetConferencingInfoResponse**](GetConferencingInfoResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadExtensionCallerId**
> ExtensionCallerIdInfo LoadExtensionCallerId(ctx, accountId, extensionId)
Get Extension Caller ID

<p style='font-style:italic;'>Since 1.0.28 (Release 8.4)</p><p>Returns information on an outbound caller ID of an extension.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]

### Return type

[**ExtensionCallerIdInfo**](ExtensionCallerIdInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadExtensionInfo**
> GetExtensionInfoResponse LoadExtensionInfo(ctx, accountId, extensionId)
Get Extension Info

<p style='font-style:italic;'>Since 1.0.0</p><p>Returns basic information about a particular extension of an account.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]

### Return type

[**GetExtensionInfoResponse**](GetExtensionInfoResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadExtensionProfileImage**
> Binary LoadExtensionProfileImage(ctx, accountId, extensionId, scaleSize)
Get Extension Profile Image (Scaled)

<p style='font-style:italic;'></p><p>Returns the extension profile image.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
  **scaleSize** | **string**| Dimensions of a profile image which will be returned in response. If this path parameter is not specified in request URI then | 

### Return type

[**Binary**](Binary.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadExtensionUserCredentials**
> ExtensionUserCredentials LoadExtensionUserCredentials(ctx, accountId, extensionId)
Get User Credentials

<p style='font-style:italic;'>Since 1.0.27 (Release 8.3)</p><p>Returns extension user credentials.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]

### Return type

[**ExtensionUserCredentials**](ExtensionUserCredentials.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadSecretQuestion**
> SecretQuestionInfo LoadSecretQuestion(ctx, questionId)
Get Secret Question

<p style='font-style:italic;'>Since 1.0.20 (Release 7.4)</p><p>Returns a particular secret question in specific language by question ID.</p><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **questionId** | **int32**| Internal identifier of a question | 

### Return type

[**SecretQuestionInfo**](SecretQuestionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConferencingInfo**
> GetConferencingInfoResponse UpdateConferencingInfo(ctx, accountId, extensionId, updateConferencingInfoRequest)
Update User Conferencing Settings

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditExtensions</td><td>Viewing and updating my extension info (includes extension name, number, email and phone number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
  **updateConferencingInfoRequest** | [**UpdateConferencingInfoRequest**](UpdateConferencingInfoRequest.md)| JSON body | 

### Return type

[**GetConferencingInfoResponse**](GetConferencingInfoResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateExtension**
> GetExtensionInfoResponse UpdateExtension(ctx, accountId, extensionId, extensionUpdateRequest)
Update Extension

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditExtensions</td><td>Viewing and updating my extension info (includes extension name, number, email and phone number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
  **extensionUpdateRequest** | [**ExtensionUpdateRequest**](ExtensionUpdateRequest.md)| JSON body | 

### Return type

[**GetExtensionInfoResponse**](GetExtensionInfoResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateExtensionCallerId**
> ExtensionCallerIdInfo UpdateExtensionCallerId(ctx, accountId, extensionId, extensionCallerIdInfo)
Update Extension Caller ID

<p style='font-style:italic;'>Since 1.0.28 (Release 8.4)</p><p>Updates outbound caller ID information of an extension.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditExtensions</td><td>Viewing and updating my extension info (includes extension name, number, email and phone number)</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
  **extensionCallerIdInfo** | [**ExtensionCallerIdInfo**](ExtensionCallerIdInfo.md)| JSON body | 

### Return type

[**ExtensionCallerIdInfo**](ExtensionCallerIdInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateExtensionUserCredentials**
> ExtensionUserCredentials UpdateExtensionUserCredentials(ctx, accountId, extensionId, extensionUserCredentailsRequest)
Update User Credentials

<p style='font-style:italic;'>Since 1.0.27 (Release 8.3)</p><p>Updates extension user credentials.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
  **extensionUserCredentailsRequest** | [**ExtensionUserCredentailsRequest**](ExtensionUserCredentailsRequest.md)| JSON body | 

### Return type

[**ExtensionUserCredentials**](ExtensionUserCredentials.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNotificationSettings**
> NotificationSettings UpdateNotificationSettings(ctx, accountId, extensionId, notificationSettingsUpdateRequest)
Update Notification Settings

<p style='font-style:italic;'>Since 1.0.26 (Release 8.2)</p><p>Updates notification settings for the current extension.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditExtensions</td><td>Viewing and updating my extension info (includes extension name, number, email and phone number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **int32**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
  **notificationSettingsUpdateRequest** | [**NotificationSettingsUpdateRequest**](NotificationSettingsUpdateRequest.md)|  | 

### Return type

[**NotificationSettings**](NotificationSettings.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadImageByPostForm**
> Binary UploadImageByPostForm(ctx, extensionId, accountId, image)
Upload User Profile Image

<p style='font-style:italic;'>Since 1.0.26 (Release 8.2)</p><p>Returns the extension profile image.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditExtensions</td><td>Viewing and updating my extension info (includes extension name, number, email and phone number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **extensionId** | **string**|  | 
  **accountId** | **string**|  | 
  **image** | ***os.File*****os.File**|  | 

### Return type

[**Binary**](Binary.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadImageByPutForm**
> Binary UploadImageByPutForm(ctx, extensionId, accountId, optional)
Update User Profile Image

<p style='font-style:italic;'>Since 1.0.26 (Release 8.2)</p><p>Updates the extension profile image..</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditExtensions</td><td>Viewing and updating my extension info (includes extension name, number, email and phone number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **extensionId** | **string**|  | 
  **accountId** | **string**|  | 
 **optional** | ***UploadImageByPutFormOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UploadImageByPutFormOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **image** | **optional.Interface of *os.File****optional.*os.File**|  | 

### Return type

[**Binary**](Binary.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateExtensionUserCredentials**
> ValidateExtensionUserCredentials(ctx, accountId, extensionId, validateExtensionUserCredentials)
Validate User Credentials

<p style='font-style:italic;'>Since 1.0.27 (Release 8.3)</p><p>Validates if extension user credentials specified can be applied.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditExtensions</td><td>Viewing and updating my extension info (includes extension name, number, email and phone number)</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
  **validateExtensionUserCredentials** | [**ValidateExtensionUserCredentials**](ValidateExtensionUserCredentials.md)| JSON body | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyExtensionUserCredentials**
> VerifyExtensionUserCredentials(ctx, accountId, extensionId, verifyExtensionUserCredentials)
Verify Extension User Credentials

<p style='font-style:italic;'>Since 1.0.27 (Release 8.3)</p><p>Verifies current extension user credentials.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditExtensions</td><td>Viewing and updating my extension info (includes extension name, number, email and phone number)</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
  **verifyExtensionUserCredentials** | [**VerifyExtensionUserCredentials**](VerifyExtensionUserCredentials.md)| JSON body | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

