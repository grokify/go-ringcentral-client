# \CallLogApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestapiV10AccountAccountIdActiveCallsGet**](CallLogApi.md#RestapiV10AccountAccountIdActiveCallsGet) | **Get** /restapi/v1.0/account/{accountId}/active-calls | 
[**RestapiV10AccountAccountIdCallLogCallLogIdGet**](CallLogApi.md#RestapiV10AccountAccountIdCallLogCallLogIdGet) | **Get** /restapi/v1.0/account/{accountId}/call-log/{callLogId} | 
[**RestapiV10AccountAccountIdCallLogGet**](CallLogApi.md#RestapiV10AccountAccountIdCallLogGet) | **Get** /restapi/v1.0/account/{accountId}/call-log | 
[**RestapiV10AccountAccountIdExtensionExtensionIdActiveCallsGet**](CallLogApi.md#RestapiV10AccountAccountIdExtensionExtensionIdActiveCallsGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/active-calls | 
[**RestapiV10AccountAccountIdExtensionExtensionIdCallLogCallLogIdGet**](CallLogApi.md#RestapiV10AccountAccountIdExtensionExtensionIdCallLogCallLogIdGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/call-log/{callLogId} | 
[**RestapiV10AccountAccountIdExtensionExtensionIdCallLogDelete**](CallLogApi.md#RestapiV10AccountAccountIdExtensionExtensionIdCallLogDelete) | **Delete** /restapi/v1.0/account/{accountId}/extension/{extensionId}/call-log | 
[**RestapiV10AccountAccountIdExtensionExtensionIdCallLogGet**](CallLogApi.md#RestapiV10AccountAccountIdExtensionExtensionIdCallLogGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/call-log | 
[**RestapiV10AccountAccountIdExtensionExtensionIdCallLogSyncGet**](CallLogApi.md#RestapiV10AccountAccountIdExtensionExtensionIdCallLogSyncGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/call-log-sync | 
[**RestapiV10AccountAccountIdRecordingRecordingIdContentGet**](CallLogApi.md#RestapiV10AccountAccountIdRecordingRecordingIdContentGet) | **Get** /restapi/v1.0/account/{accountId}/recording/{recordingId}/content | 
[**RestapiV10AccountAccountIdRecordingRecordingIdGet**](CallLogApi.md#RestapiV10AccountAccountIdRecordingRecordingIdGet) | **Get** /restapi/v1.0/account/{accountId}/recording/{recordingId} | 


# **RestapiV10AccountAccountIdActiveCallsGet**
> InlineResponseDefault1 RestapiV10AccountAccountIdActiveCallsGet($accountId, $direction, $type_, $page, $perPage)



Get Account Active (Recent) Calls


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **direction** | **string**| The direction for the result records. It is allowed to specify more than one direction. If not specified, both inbound and outbound records are returned. Multiple values are accepted | [optional] 
 **type_** | **string**| Call type of a record. It is allowed to specify more than one type. If not specified, all call types are returned. Multiple values are accepted | [optional] 
 **page** | **int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39;. | [optional] 
 **perPage** | **int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default. | [optional] 

### Return type

[**InlineResponseDefault1**](inline_response_default_1.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdCallLogCallLogIdGet**
> CallLogInfo RestapiV10AccountAccountIdCallLogCallLogIdGet($accountId, $callLogId)



Get Account Call Log Record by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **callLogId** | **string**| Internal identifier of a call log record | 

### Return type

[**CallLogInfo**](CallLogInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdCallLogGet**
> InlineResponseDefault3 RestapiV10AccountAccountIdCallLogGet($accountId, $extensionNumber, $phoneNumber, $direction, $type_, $view, $withRecording, $dateFrom, $dateTo, $page, $perPage)



Get Account Call Log


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionNumber** | **string**| Extension number of a user. If specified, returns call log for a particular extension only. Cannot be specified together with the phoneNumber filter | [optional] 
 **phoneNumber** | **string**| Phone number of a caller/call recipient. If specified, returns all calls (both incoming and outcoming) with the mentioned phone number. Cannot be specified together with the extensionNumber filter | [optional] 
 **direction** | **string**| The direction for the result records. It is allowed to specify more than one direction. If not specified, both inbound and outbound records are returned. Multiple values are accepted | [optional] 
 **type_** | **string**| Call type of a record. It is allowed to specify more than one type. If not specified, all call types are returned. Multiple values are accepted | [optional] 
 **view** | **string**| The default value is &#39;Simple&#39; for both account and extension call log | [optional] 
 **withRecording** | **bool**| &#39;True&#39; if only recorded calls have to be returned | [optional] 
 **dateFrom** | **time.Time**| The start datetime for resulting records in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is dateTo minus 24 hours | [optional] 
 **dateTo** | **time.Time**| The end datetime for resulting records in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is current time | [optional] 
 **page** | **int32**| Indicates the page number to retrieve. Only positive number values are allowed. The default value is &#39;1&#39; | [optional] 
 **perPage** | **int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default. | [optional] 

### Return type

[**InlineResponseDefault3**](inline_response_default_3.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdActiveCallsGet**
> InlineResponseDefault1 RestapiV10AccountAccountIdExtensionExtensionIdActiveCallsGet($accountId, $extensionId, $direction, $type_, $page, $perPage)



Get Extension Active (Recent) Calls


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **direction** | **string**| The direction for the result records. It is allowed to specify more than one direction. If not specified, both inbound and outbound records are returned. Multiple values are accepted | [optional] 
 **type_** | **string**| Call type of a record. It is allowed to specify more than one type. If not specified, all call types are returned. Multiple values are accepted | [optional] 
 **page** | **int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | [optional] 
 **perPage** | **int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | [optional] 

### Return type

[**InlineResponseDefault1**](inline_response_default_1.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdCallLogCallLogIdGet**
> CallLogInfo RestapiV10AccountAccountIdExtensionExtensionIdCallLogCallLogIdGet($accountId, $extensionId, $callLogId)



Get Extension Call Log Record by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **callLogId** | **string**| Internal identifier of a call log record | 

### Return type

[**CallLogInfo**](CallLogInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdCallLogDelete**
> RestapiV10AccountAccountIdExtensionExtensionIdCallLogDelete($accountId, $extensionId, $dateTo)



Delete Extension Call Log


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **dateTo** | **time.Time**| The end datetime for records deletion in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is current time | [optional] 

### Return type

void (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdCallLogGet**
> InlineResponseDefault3 RestapiV10AccountAccountIdExtensionExtensionIdCallLogGet($accountId, $extensionId, $extensionNumber, $phoneNumber, $direction, $type_, $view, $withRecording, $dateTo, $dateFrom, $page, $perPage)



Get Extension Call Log


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **extensionNumber** | **string**| Extension number of a user. If specified, returns call log for a particular extension only. Cannot be specified together with the phoneNumber filter | [optional] 
 **phoneNumber** | **string**| Phone number of a caller/call recipient. If specified, returns all calls (both incoming and outcoming) with the mentioned phone number. Cannot be specified together with the extensionNumber filter | [optional] 
 **direction** | **string**| The direction for the result records. It is allowed to specify more than one direction. If not specified, both inbound and outbound records are returned. Multiple values are accepted | [optional] 
 **type_** | **string**| Call type of a record. It is allowed to specify more than one type. If not specified, all call types are returned. Multiple values are accepted | [optional] 
 **view** | **string**| The default value is &#39;Simple&#39; for both account and extension call log | [optional] 
 **withRecording** | **bool**| &#39;True&#39; if only recorded calls have to be returned | [optional] 
 **dateTo** | **time.Time**| The end datetime for resulting records in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is current time | [optional] 
 **dateFrom** | **time.Time**| The start datetime for resulting records in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is dateTo minus 24 hours | [optional] 
 **page** | **int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | [optional] 
 **perPage** | **int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | [optional] 

### Return type

[**InlineResponseDefault3**](inline_response_default_3.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdCallLogSyncGet**
> InlineResponseDefault16 RestapiV10AccountAccountIdExtensionExtensionIdCallLogSyncGet($accountId, $extensionId, $syncType, $syncToken, $dateFrom, $recordCount, $statusGroup)



Call Log Synchronization


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **syncType** | **string**| Type of synchronization. &#39;FSync&#39; is a default value | [optional] 
 **syncToken** | **string**| Value of syncToken property of last sync request response | [optional] 
 **dateFrom** | **time.Time**| The start datetime for resulting records in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is the current moment | [optional] 
 **recordCount** | **int32**| For FSync the parameter is mandatory, it limits the number of records to be returned in response. For ISync it specifies with how many records to extend sync Frame to the past, the maximum number of records is 250 | [optional] 
 **statusGroup** | **string**| Type of calls to be returned. The default value is &#39;All&#39; | [optional] 

### Return type

[**InlineResponseDefault16**](inline_response_default_16.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdRecordingRecordingIdContentGet**
> Binary RestapiV10AccountAccountIdRecordingRecordingIdContentGet($accountId, $recordingId)



Get Call Recording Content


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **recordingId** | **string**| Internal identifier of recording (returned in Call Log) | 

### Return type

[**Binary**](Binary.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdRecordingRecordingIdGet**
> InlineResponseDefault27 RestapiV10AccountAccountIdRecordingRecordingIdGet($accountId, $recordingId)



Get Call Recording Metadata


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **recordingId** | **string**| Internal identifier of recording (returned in Call Log) | 

### Return type

[**InlineResponseDefault27**](inline_response_default_27.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

