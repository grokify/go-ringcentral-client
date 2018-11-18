# UserCallLogRecord

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of a cal log record | [optional] 
**Uri** | **string** | Canonical URI of a call log record | [optional] 
**SessionId** | **string** | Internal identifier of a call session | [optional] 
**From** | [**CallLogCallerInfo**](CallLogCallerInfo.md) |  | [optional] 
**To** | [**CallLogCallerInfo**](CallLogCallerInfo.md) |  | [optional] 
**Type** | **string** | Call type &#x3D; [&#39;Voice&#39;, &#39;Fax&#39;] | [optional] 
**Direction** | **string** | Call direction &#x3D; [&#39;Inbound&#39;, &#39;Outbound&#39;] | [optional] 
**StartTime** | **string** | The call start datetime in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z | [optional] 
**Duration** | **int32** | Call duration in seconds | [optional] 
**Recording** | [**RecordingInfo**](RecordingInfo.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


