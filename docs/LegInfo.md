# LegInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action description of the call operation | [optional] [default to null]
**Direction** | **string** | Call direction | [optional] [default to null]
**Duration** | **int32** | Call duration in seconds | [optional] [default to null]
**Extension** | [**LegInfoExtensionInfo**](LegInfo.ExtensionInfo.md) | Information on extension | [optional] [default to null]
**LegType** | **string** | Leg type | [optional] [default to null]
**StartTime** | [**time.Time**](time.Time.md) | The call start datetime in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z | [optional] [default to null]
**Type_** | **string** | Call type | [optional] [default to null]
**Result** | **string** | Status description of the call operation | [optional] [default to null]
**From** | [**CallerInfo**](CallerInfo.md) | Caller information | [optional] [default to null]
**To** | [**CallerInfo**](CallerInfo.md) | Callee information | [optional] [default to null]
**Transport** | **string** | Call transport | [optional] [default to null]
**Recording** | [**RecordingInfo**](RecordingInfo.md) | Call recording data. Returned if the call is recorded | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


