# MeetingInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | **string** | Canonical URI of a meeting resource | [optional] [default to null]
**Id** | **string** | Internal identifier of a meeting as retrieved from Zoom | [optional] [default to null]
**Topic** | **string** | Topic of a meeting | [optional] [default to null]
**MeetingType** | **string** | Type of a meeting | [optional] [default to null]
**Password** | **string** | Password required to join a meeting | [optional] [default to null]
**Status** | **string** | Current status of a meeting | [optional] [default to null]
**Links** | [**LinksInfo**](LinksInfo.md) | Links to start/join the meeting | [optional] [default to null]
**Schedule** | [**MeetingScheduleInfo**](MeetingScheduleInfo.md) | Schedule of a meeting | [optional] [default to null]
**AllowJoinBeforeHost** | **bool** | If &#39;True&#39; then the meeting can be joined and started by any participant (not host only). Supported for the meetings of &#39;Scheduled&#39; and &#39;Recurring&#39; type. | [optional] [default to null]
**StartHostVideo** | **bool** | Enables starting video when host joins the meeting | [optional] [default to null]
**StartParticipantsVideo** | **bool** | Enables starting video when participants join the meeting | [optional] [default to null]
**AudioOptions** | **[]string** | Meeting audio options. Possible values are &#39;Phone&#39;, &#39;ComputerAudio&#39; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


