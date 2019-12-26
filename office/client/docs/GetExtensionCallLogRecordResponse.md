# GetExtensionCallLogRecordResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of a cal log record | [optional] 
**Uri** | **string** | Canonical URI of a call log record | [optional] 
**SessionId** | **string** | Internal identifier of a call session | [optional] 
**From** | [**CallLogCallerInfo**](CallLogCallerInfo.md) |  | [optional] 
**To** | [**CallLogCallerInfo**](CallLogCallerInfo.md) |  | [optional] 
**Type** | **string** | Call type | [optional] 
**Direction** | **string** | Call direction | [optional] 
**Action** | **string** | Action description of the call operation | [optional] 
**Result** | **string** | Status description of the call operation | [optional] 
**StartTime** | [**time.Time**](time.Time.md) | The call start datetime in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z | [optional] 
**Duration** | **int32** | Call duration in seconds | [optional] 
**Recording** | [**RecordingInfo**](RecordingInfo.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


