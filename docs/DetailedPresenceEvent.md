# DetailedPresenceEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtensionId** | **string** | Internal identifier of an extension. Optional parameter | [optional] [default to null]
**TelephonyStatus** | **string** | Telephony presence status. Returned if telephony status is changed. See Telephony Status Values | [optional] [default to null]
**TerminationType** | **string** | Type of call termination. Supported for calls in &#39;NoCall&#39; status. If the returned termination type is &#39;Intermediate&#39; it means the call is not actually ended, the connection is established on one of the devices | [optional] [default to null]
**ActiveCalls** | [**[]DetailedPresenceEventActiveCallInfo**](DetailedPresenceEvent.ActiveCallInfo.md) | Collection of Active Call Info | [optional] [default to null]
**Sequence** | **int32** | Order number of a notification to state the chronology | [optional] [default to null]
**PresenceStatus** | **string** | Aggregated presence status, calculated from a number of sources | [optional] [default to null]
**UserStatus** | **string** | User-defined presence status (as previously published by the user) | [optional] [default to null]
**DndStatus** | **string** | Extended DnD (Do not Disturb) status | [optional] [default to null]
**AllowSeeMyPresence** | **bool** | If &#39;True&#39; enables other extensions to see the extension presence status | [optional] [default to null]
**RingOnMonitoredCall** | **bool** | If &#39;True&#39; enables to ring extension phone, if any user monitored by this extension is ringing | [optional] [default to null]
**PickUpCallsOnHold** | **bool** | If &#39;True&#39; enables the extension user to pick up a monitored line on hold | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


