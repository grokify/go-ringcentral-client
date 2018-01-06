# CallLogRecord

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of a cal log record | [optional] [default to null]
**Uri** | **string** | Canonical URI of a call log record | [optional] [default to null]
**SessionId** | **string** | Internal identifier of a call session | [optional] [default to null]
**From** | [***CallLogCallerInfo**](CallLogCallerInfo.md) | Caller information | [optional] [default to null]
**To** | [***CallLogCallerInfo**](CallLogCallerInfo.md) | Callee information | [optional] [default to null]
**Type_** | **string** | Call type | [optional] [default to null]
**Direction** | **string** | Call direction | [optional] [default to null]
**Action** | **string** | Action description of the call operation | [optional] [default to null]
**Result** | **string** | Status description of the call operation | [optional] [default to null]
**StartTime** | [**time.Time**](time.Time.md) | The call start datetime in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z | [optional] [default to null]
**Duration** | **int32** | Call duration in seconds | [optional] [default to null]
**Recording** | [***RecordingInfo**](RecordingInfo.md) | Call recording data. Returned if the call is recorded, the withRecording parameter is set to &#39;True&#39; in this case | [optional] [default to null]
**LastModifiedTime** | [**time.Time**](time.Time.md) | For &#39;Detailed&#39; view only. The datetime when the call log record was modified in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z | [optional] [default to null]
**Transport** | **string** | For &#39;Detailed&#39; view only. Call transport | [optional] [default to null]
**Legs** | [**[]CallLogRecordLegInfo**](CallLogRecordLegInfo.md) | For &#39;Detailed&#39; view only. Leg description | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


