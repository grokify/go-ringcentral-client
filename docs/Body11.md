# Body11

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | **string** | Topic of a meeting | [optional] [default to null]
**MeetingType** | **string** | Type of a meeting. &#39;Instant&#39; - meeting that is instantly started as soon as the host creates it; &#39;Scheduled&#39; - common scheduled meeting; &#39;Recurring&#39; - a recurring meeting. If the specified meeting type is &#39;Scheduled&#39; then schedule property is mandatory for request | [optional] [default to null]
**Password** | **string** | Password required to join a meeting. Max number of characters is 10 | [optional] [default to null]
**Schedule** | [**MeetingScheduleInfo**](MeetingScheduleInfo.md) | Schedule of a meeting | [optional] [default to null]
**AllowJoinBeforeHost** | **bool** | If &#39;True&#39; then the meeting can be joined and started by any participant (not host only). Supported for the meetings of &#39;Scheduled&#39; and &#39;Recurring&#39; type. | [optional] [default to null]
**StartHostVideo** | **bool** | Enables starting video when host joins the meeting | [optional] [default to null]
**StartParticipantsVideo** | **bool** | Enables starting video when participants join the meeting | [optional] [default to null]
**AudioOptions** | **[]string** | Meeting audio options. Possible values are &#39;Phone&#39;, &#39;ComputerAudio&#39; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


