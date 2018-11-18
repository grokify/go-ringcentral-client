# CallLogRecordLegInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action description of the call operation | [optional] 
**Direction** | **string** | Call direction | [optional] 
**Duration** | **int32** | Call duration in seconds | [optional] 
**Extension** | [**ExtensionInfoCallLog**](ExtensionInfoCallLog.md) |  | [optional] 
**LegType** | **string** | Leg type | [optional] 
**StartTime** | [**time.Time**](time.Time.md) | The call start datetime in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z | [optional] 
**Type** | **string** | Call type | [optional] 
**Result** | **string** | Status description of the call operation | [optional] 
**From** | [**CallLogCallerInfo**](CallLogCallerInfo.md) |  | [optional] 
**To** | [**CallLogCallerInfo**](CallLogCallerInfo.md) |  | [optional] 
**Transport** | **string** | Call transport | [optional] 
**Recording** | [**RecordingInfo**](RecordingInfo.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


