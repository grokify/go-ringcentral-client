# ConferencingInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | **string** | Canonical URI of a conferencing | [optional] [default to null]
**AllowJoinBeforeHost** | **bool** | Determines if host user allows conference participants to join before the host | [optional] [default to null]
**HostCode** | **string** | Access code for a host user | [optional] [default to null]
**Mode** | **string** | Internal parameter specifying conferencing engine | [optional] [default to null]
**ParticipantCode** | **string** | Access code for any participant | [optional] [default to null]
**PhoneNumber** | **string** | Primary conference phone number for user&#39;s home country returned in E.164 (11-digits) format | [optional] [default to null]
**TapToJoinUri** | **string** | Short URL leading to the service web page Tap to Join for audio conference bridge | [optional] [default to null]
**PhoneNumbers** | [**[]ConferencingInfoPhoneNumberInfo**](ConferencingInfo.PhoneNumberInfo.md) | List of multiple dial-in phone numbers to connect to audio conference service, relevant for user&#39;s brand. Each number is given with the country and location information, in order to let the user choose the less expensive way to connect to a conference. The first number in the list is the primary conference number, that is default and domestic | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


