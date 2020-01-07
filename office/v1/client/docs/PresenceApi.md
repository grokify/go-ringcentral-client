# \PresenceApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountPresence**](PresenceApi.md#AccountPresence) | **Get** /restapi/v1.0/account/{accountId}/presence | Get all user statuses
[**GetMonitoringExtensions**](PresenceApi.md#GetMonitoringExtensions) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/presence/permission | Get User Presence Permissions
[**GetPresenceLine**](PresenceApi.md#GetPresenceLine) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/presence/line/{lineId} | Get Monitored Extensions by Id
[**GetPresenceStatus**](PresenceApi.md#GetPresenceStatus) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/presence | Get User Status
[**ListMonitoredExtensions**](PresenceApi.md#ListMonitoredExtensions) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/presence/line | Get Monitored Extensions
[**PutMonitoringExtensions**](PresenceApi.md#PutMonitoringExtensions) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/presence/permission | Update User Presence Permissions
[**UpdatePresenceLines**](PresenceApi.md#UpdatePresenceLines) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/presence/line | Update Monitored Extensions
[**UpdatePresenceStatus**](PresenceApi.md#UpdatePresenceStatus) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/presence | Update User Status


# **AccountPresence**
> AccountPresenceInfo AccountPresence(ctx, accountId)
Get all user statuses

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 

### Return type

[**AccountPresenceInfo**](AccountPresenceInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMonitoringExtensions**
> MonitoringExtensionsResource GetMonitoringExtensions(ctx, extensionId, accountId)
Get User Presence Permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
  **accountId** | **string**|  | 

### Return type

[**MonitoringExtensionsResource**](MonitoringExtensionsResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPresenceLine**
> PresenceLineResource GetPresenceLine(ctx, lineId, extensionId, accountId)
Get Monitored Extensions by Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lineId** | **string**|  | 
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]

### Return type

[**PresenceLineResource**](PresenceLineResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPresenceStatus**
> GetPresenceInfo GetPresenceStatus(ctx, accountId, extensionId)
Get User Status

<p style='font-style:italic;'>Since 1.0.2</p><p>Returns presence status of an extension or several extensions by their ID(s). Batch request is supported, see Batch Requests for details.The presenceStatus is returned as Offline (the parameters telephonyStatus, message, userStatus and dndStatus are not returned at all) for the following extension types: Department/Announcement Only/Take Messages Only (Voicemail)/Fax User/Paging Only Group/Shared Lines Group/IVR Menu/Application Extension/Park Location.If the user requests his/her own presence status, the response contains actual presence status even if the status publication is turned off.Batch request is supported. For batch requests the number of extensions in one request is limited to 30. If more extensions are included in the request, the error code 400 Bad Request is returned with the logical error code InvalidMultipartRequest and the corresponding message 'Extension Presence Info multipart request is limited to 30 extensions'.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadPresence</td><td>Getting user presence information</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]

### Return type

[**GetPresenceInfo**](GetPresenceInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMonitoredExtensions**
> GetMonitoredExtensionsResponse ListMonitoredExtensions(ctx, accountId, extensionId)
Get Monitored Extensions

<p style='font-style:italic;'>Since 1.0.13 (Release 6.5)</p><p>Returns list of lines - extensions which presence status can be indicated and monitored on BLF-enabled (Busy Lamp Field) devices.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadPresence</td><td>Getting user presence information</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]

### Return type

[**GetMonitoredExtensionsResponse**](GetMonitoredExtensionsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutMonitoringExtensions**
> MonitoringExtensionsResource PutMonitoringExtensions(ctx, extensionId, accountId, userPresencePermissionsUpdateRequest)
Update User Presence Permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
  **accountId** | **string**|  | 
  **userPresencePermissionsUpdateRequest** | [**UserPresencePermissionsUpdateRequest**](UserPresencePermissionsUpdateRequest.md)|  | 

### Return type

[**MonitoringExtensionsResource**](MonitoringExtensionsResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePresenceLines**
> UpdatePresenceLinesResponse UpdatePresenceLines(ctx, accountId, extensionId, updatePresenceLinesRequest)
Update Monitored Extensions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
  **updatePresenceLinesRequest** | [**UpdatePresenceLinesRequest**](UpdatePresenceLinesRequest.md)|  | 

### Return type

[**UpdatePresenceLinesResponse**](UpdatePresenceLinesResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePresenceStatus**
> PresenceInfoResource UpdatePresenceStatus(ctx, accountId, extensionId, presenceInfoResource)
Update User Status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
  **presenceInfoResource** | [**PresenceInfoResource**](PresenceInfoResource.md)|  | 

### Return type

[**PresenceInfoResource**](PresenceInfoResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: multipart/mixed, application/json
 - **Accept**: multipart/mixed, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

