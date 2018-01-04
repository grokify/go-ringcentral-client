# UserCallLogRecord

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of a cal log record | [optional] [default to null]
**Uri** | **string** | Canonical URI of a call log record | [optional] [default to null]
**SessionId** | **string** | Internal identifier of a call session | [optional] [default to null]
**From** | [***CallLogCallerInfo**](CallLogCallerInfo.md) | Caller information | [optional] [default to null]
**To** | [***CallLogCallerInfo**](CallLogCallerInfo.md) | Callee information | [optional] [default to null]
**Type_** | **string** | Call type &#x3D; [&#39;Voice&#39;, &#39;Fax&#39;] | [optional] [default to null]
**Direction** | **string** | Call direction &#x3D; [&#39;Inbound&#39;, &#39;Outbound&#39;] | [optional] [default to null]
**StartTime** | **string** | The call start datetime in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z | [optional] [default to null]
**Duration** | **int32** | Call duration in seconds | [optional] [default to null]
**Recording** | [***RecordingInfo**](RecordingInfo.md) | Call recording data. Returned if the call is recorded. Each call recording is stored in the system for 90 days. But if the number of recordings exceeds the admissible limit (100,000 recordings per account) then the older recordings are replaced with the new ones. Thus a link to an older recording in a certain call log record becomes unavailable | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


