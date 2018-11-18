# \CallLogApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAccountCallLog**](CallLogApi.md#DeleteAccountCallLog) | **Delete** /restapi/v1.0/account/{accountId}/call-log | Delete Call Log Records by Filter
[**DeleteAccountCallLogRecord**](CallLogApi.md#DeleteAccountCallLogRecord) | **Delete** /restapi/v1.0/account/{accountId}/call-log/{callRecordId} | Delete Account Call Log Record(s) by ID
[**DeleteExtensionCallLog**](CallLogApi.md#DeleteExtensionCallLog) | **Delete** /restapi/v1.0/account/{accountId}/extension/{extensionId}/call-log | Delete Call Log Records by Filter
[**GetCallRecords**](CallLogApi.md#GetCallRecords) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/call-log/{callRecordId} | Get Call Records by ID
[**ListCallRecordingData**](CallLogApi.md#ListCallRecordingData) | **Get** /restapi/v1.0/account/{accountId}/recording/{recordingId}/content | Get Call Recordings Data
[**ListCallRecordings**](CallLogApi.md#ListCallRecordings) | **Get** /restapi/v1.0/account/{accountId}/recording/{recordingId} | Get Call Recordings
[**ListCompanyActiveCalls**](CallLogApi.md#ListCompanyActiveCalls) | **Get** /restapi/v1.0/account/{accountId}/active-calls | Get User Active Calls
[**ListExtensionActiveCalls**](CallLogApi.md#ListExtensionActiveCalls) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/active-calls | Get User Active Calls
[**LoadAccountCallLog**](CallLogApi.md#LoadAccountCallLog) | **Get** /restapi/v1.0/account/{accountId}/call-log | Get Call Log Records by Filter
[**LoadAccountCallLogRecord**](CallLogApi.md#LoadAccountCallLogRecord) | **Get** /restapi/v1.0/account/{accountId}/call-log/{callRecordId} | Get Account Call Log Record(s) by ID
[**LoadExtensionCallLog**](CallLogApi.md#LoadExtensionCallLog) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/call-log | Get Call Log Records by Filter
[**SyncAccountCallLog**](CallLogApi.md#SyncAccountCallLog) | **Get** /restapi/v1.0/account/{accountId}/call-log-sync | Call Log Synchronization
[**SyncExtensionCallLog**](CallLogApi.md#SyncExtensionCallLog) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/call-log-sync | Call Log Synchronization


# **DeleteAccountCallLog**
> DeleteAccountCallLog(ctx, accountId, optional)
Delete Call Log Records by Filter

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditCallLog</td><td>Viewing and updating user call logs</td></tr><tr><td class='code'>ReadCallLog</td><td>Viewing user call logs</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
 **optional** | ***DeleteAccountCallLogOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteAccountCallLogOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dateTo** | **optional.String**| The end datetime for records deletion in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is current time | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccountCallLogRecord**
> DeleteAccountCallLogRecord(ctx, accountId, callRecordId)
Delete Account Call Log Record(s) by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **callRecordId** | **int32**| Internal identifier of a call log record | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteExtensionCallLog**
> DeleteExtensionCallLog(ctx, accountId, extensionId, optional)
Delete Call Log Records by Filter

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditCallLog</td><td>Viewing and updating user call logs</td></tr><tr><td class='code'>ReadCallLog</td><td>Viewing user call logs</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***DeleteExtensionCallLogOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteExtensionCallLogOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dateTo** | **optional.String**| The end datetime for records deletion in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is current time | 
 **phoneNumber** | **optional.String**|  | 
 **extensionNumber** | **optional.String**|  | 
 **type_** | [**optional.Interface of []string**](string.md)|  | 
 **direction** | [**optional.Interface of []string**](string.md)|  | 
 **dateFrom** | **optional.Time**|  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCallRecords**
> UserCallLogRecord GetCallRecords(ctx, callRecordId, extensionId, accountId, optional)
Get Call Records by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **callRecordId** | [**[]string**](string.md)|  | 
  **extensionId** | **string**|  | 
  **accountId** | **string**|  | 
 **optional** | ***GetCallRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCallRecordsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **view** | **optional.String**|  | 

### Return type

[**UserCallLogRecord**](UserCallLogRecord.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCallRecordingData**
> Binary ListCallRecordingData(ctx, accountId, recordingId)
Get Call Recordings Data

<p style='font-style:italic;'>Since 1.0.16 (Release 7.1)</p><p>Returns call recording metadata.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadCallRecording</td><td>Downloading call recording content</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **recordingId** | **string**| Internal identifier of a recording (returned in Call Log) | 

### Return type

[**Binary**](Binary.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCallRecordings**
> GetCallRecordingResponse ListCallRecordings(ctx, accountId, recordingId)
Get Call Recordings

<p style='font-style:italic;'>Since 1.0.18 (Release 6.5)</p><p>Returns call recording metadata.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadCallRecording</td><td>Downloading call recording content</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **recordingId** | **string**| Internal identifier of a recording (returned in Call Log) | 

### Return type

[**GetCallRecordingResponse**](GetCallRecordingResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCompanyActiveCalls**
> ExtensionActiveCallsResponse ListCompanyActiveCalls(ctx, accountId, optional)
Get User Active Calls

<p style='font-style:italic;'>Since 1.0.13 (Release 6.5)</p><p>Returns records of all calls that are in progress, ordered by start time in descending order</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadCallLog</td><td>Viewing user call logs</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
 **optional** | ***ListCompanyActiveCallsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListCompanyActiveCallsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **direction** | [**optional.Interface of []string**](string.md)| The direction for the result records. It is allowed to specify more than one direction. If not specified, both inbound and outbound records are returned. Multiple values are accepted | 
 **type_** | [**optional.Interface of []string**](string.md)| Call type of a record. It is allowed to specify more than one type. If not specified, all call types are returned. Multiple values are accepted | 
 **transport** | [**optional.Interface of []string**](string.md)| Call transport type. &#39;PSTN&#39; specifies that a call leg is initiated from the PSTN network provider; &#39;VoIP&#39; - from an RC phone. By default this filter is disabled | 
 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | 
 **perPage** | **optional.Int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | 

### Return type

[**ExtensionActiveCallsResponse**](ExtensionActiveCallsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListExtensionActiveCalls**
> ExtensionActiveCallsResponse ListExtensionActiveCalls(ctx, accountId, extensionId, optional)
Get User Active Calls

<p style='font-style:italic;'>Since 1.0.13 (Release 6.5)</p><p>Returns records of all extension calls that are in progress, ordered by start time in descending order.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadCallLog</td><td>Viewing user call logs</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***ListExtensionActiveCallsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListExtensionActiveCallsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **direction** | [**optional.Interface of []string**](string.md)| The direction for the result records. It is allowed to specify more than one direction. If not specified, both inbound and outbound records are returned. Multiple values are accepted | 
 **type_** | [**optional.Interface of []string**](string.md)| Call type of a record. It is allowed to specify more than one type. If not specified, all call types are returned. Multiple values are accepted | 
 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | 
 **perPage** | **optional.Int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | 

### Return type

[**ExtensionActiveCallsResponse**](ExtensionActiveCallsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadAccountCallLog**
> AccountCallLogResponse LoadAccountCallLog(ctx, accountId, optional)
Get Call Log Records by Filter

<p style='font-style:italic;'>Since 1.0.3 (Release 5.11)</p><p>Returns call log records filtered by the specified parameters.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadCallLog</td><td>Viewing user call logs</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
 **optional** | ***LoadAccountCallLogOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoadAccountCallLogOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **extensionNumber** | **optional.String**| Extension number of a user. If specified, returns call log for a particular extension only. Cannot be specified together with the phoneNumber filter | 
 **phoneNumber** | **optional.String**| Phone number of a caller/call recipient. If specified, returns all calls (both incoming and outcoming) with the mentioned phone number. Cannot be specified together with the extensionNumber filter | 
 **direction** | [**optional.Interface of []string**](string.md)| The direction for the result records. It is allowed to specify more than one direction. If not specified, both inbound and outbound records are returned. Multiple values are accepted | 
 **type_** | [**optional.Interface of []string**](string.md)| Call type of a record. It is allowed to specify more than one type. If not specified, all call types are returned. Multiple values are accepted | 
 **view** | [**optional.Interface of []string**](string.md)| The default value is &#39;Simple&#39; for both account and extension call log | 
 **withRecording** | **optional.Bool**| &#39;True&#39; if only recorded calls have to be returned | 
 **dateFrom** | **optional.Time**| The start datetime for resulting records in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is dateTo minus 24 hours | 
 **dateTo** | **optional.Time**| The end datetime for resulting records in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is current time | 
 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are allowed. The default value is &#39;1&#39; | 
 **perPage** | **optional.Int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default. | 
 **sessionId** | **optional.String**|  | 

### Return type

[**AccountCallLogResponse**](AccountCallLogResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadAccountCallLogRecord**
> GetAccountCallLogRecordResponse LoadAccountCallLogRecord(ctx, accountId, callRecordId)
Get Account Call Log Record(s) by ID

<p style='font-style:italic;'>Since 1.0.3 (Release 5.11)</p><p>Returns individual call log record(s) by ID(s). Batch request is supported, see Batch Requests for details.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadCallLog</td><td>Viewing user call logs</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **callRecordId** | **int32**| Internal identifier of a call log record | 

### Return type

[**GetAccountCallLogRecordResponse**](GetAccountCallLogRecordResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadExtensionCallLog**
> ExtensionCallLogResponse LoadExtensionCallLog(ctx, accountId, extensionId, optional)
Get Call Log Records by Filter

<p style='font-style:italic;'>Since 1.0.3 (Release 5.11)</p><p>Returns call log records filtered by the specified parameters.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadCallLog</td><td>Viewing user call logs</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***LoadExtensionCallLogOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoadExtensionCallLogOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **extensionNumber** | **optional.String**| Extension number of a user. If specified, returns call log for a particular extension only. Cannot be specified together with the phoneNumber filter | 
 **showBlocked** | **optional.Bool**| If &#39;True&#39; then calls from/to blocked numbers are returned. The default value is &#39;True&#39; | 
 **phoneNumber** | **optional.String**| Phone number of a caller/call recipient. If specified, returns all calls (both incoming and outcoming) with the mentioned phone number. Cannot be specified together with the extensionNumber filter | 
 **direction** | [**optional.Interface of []string**](string.md)| The direction for the result records. It is allowed to specify more than one direction. If not specified, both inbound and outbound records are returned. Multiple values are accepted | 
 **sessionId** | **optional.String**|  | 
 **type_** | [**optional.Interface of []string**](string.md)| Call type of a record. It is allowed to specify more than one type. If not specified, all call types are returned. Multiple values are accepted | 
 **transport** | [**optional.Interface of []string**](string.md)| Call transport type. &#39;PSTN&#39; specifies that a call leg is initiated from the PSTN network provider; &#39;VoIP&#39; - from an RC phone. By default this filter is disabled | 
 **view** | [**optional.Interface of []string**](string.md)| The default value is &#39;Simple&#39; for both account and extension call log | 
 **withRecording** | **optional.Bool**| &#39;True&#39; if only recorded calls have to be returned | 
 **dateTo** | **optional.Time**| The end datetime for resulting records in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is current time | 
 **dateFrom** | **optional.Time**| The start datetime for resulting records in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is dateTo minus 24 hours | 
 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | 
 **perPage** | **optional.Int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | 

### Return type

[**ExtensionCallLogResponse**](ExtensionCallLogResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncAccountCallLog**
> SyncAccountCallLog(ctx, accountId, optional)
Call Log Synchronization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
 **optional** | ***SyncAccountCallLogOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SyncAccountCallLogOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **syncType** | [**optional.Interface of []string**](string.md)| Type of synchronization. &#39;FSync&#39; is a default value | 
 **syncToken** | **optional.String**| Value of syncToken property of last sync request response | 
 **dateFrom** | **optional.String**| The start datetime for resulting records in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is the current moment | 
 **recordCount** | **optional.Int32**| ForT?FSync the parameter is mandatory, it limits the number of records to be returned in response. For ISync it specifies with how many records to extend sync Frame to the past, the maximum number of records is 250 | 
 **statusGroup** | [**optional.Interface of []string**](string.md)| Type of calls to be returned. The default value is &#39;All&#39; | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncExtensionCallLog**
> CallLogSync SyncExtensionCallLog(ctx, accountId, extensionId, optional)
Call Log Synchronization

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadCallLog</td><td>Viewing user call logs</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***SyncExtensionCallLogOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SyncExtensionCallLogOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **syncType** | [**optional.Interface of []string**](string.md)| Type of synchronization. &#39;FSync&#39; is a default value | 
 **syncToken** | **optional.String**| Value of syncToken property of last sync request response | 
 **dateFrom** | **optional.String**| The start datetime for resulting records in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is the current moment | 
 **recordCount** | **optional.Int32**| ForT?FSync the parameter is mandatory, it limits the number of records to be returned in response. For ISync it specifies with how many records to extend sync Frame to the past, the maximum number of records is 250 | 
 **statusGroup** | [**optional.Interface of []string**](string.md)| Type of calls to be returned. The default value is &#39;All&#39; | 

### Return type

[**CallLogSync**](CallLogSync.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

