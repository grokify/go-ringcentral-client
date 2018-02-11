# RcVideoNotificationsEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of a meeting | [optional] [default to null]
**Name** | **string** | Meeting Name | [optional] [default to null]
**Start** | **string** | Datetime of meeting start in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z | [optional] [default to null]
**Duration** | **int32** | Meeting duration | [optional] [default to null]
**Participants** | [***RcvParticipantsInfo**](RCVParticipantsInfo.md) | Meeting participants information | [optional] [default to null]
**Recorded** | **bool** | Specifies whether a meeting is recorded or not | [optional] [default to null]
**Recording** | [***RcvRecordingInfo**](RCVRecordingInfo.md) | Meeting recording information, if recorded value is &#39;True&#39; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


