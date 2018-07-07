# GetPresenceInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | **string** | Canonical URI of a presence info resource | [optional] 
**AllowSeeMyPresence** | **bool** | If &#39;True&#39; enables other extensions to see the extension presence status | [optional] 
**DndStatus** | **string** | Extended DnD (Do not Disturb) status. Cannot be set for Department/Announcement/Voicemail (Take Messages Only)/Fax User/Shared Lines Group/Paging Only Group/IVR Menu/Application Extension/Park Location extensions. The &#39;DoNotAcceptDepartmentCalls&#39; and &#39;TakeDepartmentCallsOnly&#39; values are applicable only for extensions - members of a Department; if these values are set for department outsiders, the 400 Bad Request error code is returned. The &#39;TakeDepartmentCallsOnly&#39; status can be set through the old RingCentral user interface and is available for some migrated accounts only. | [optional] 
**Extension** | [**GetPresenceExtensionInfo**](GetPresenceExtensionInfo.md) |  | [optional] 
**Message** | **string** | Custom status message (as previously published by user) | [optional] 
**PickUpCallsOnHold** | **bool** | If &#39;True&#39; enables the extension user to pick up a monitored line on hold | [optional] 
**PresenceStatus** | **string** | Aggregated presence status, calculated from a number of sources | [optional] 
**RingOnMonitoredCall** | **bool** | If &#39;True&#39; enables to ring extension phone, if any user monitored by this extension is ringing | [optional] 
**TelephonyStatus** | **string** | Telephony presence status | [optional] 
**UserStatus** | **string** | User-defined presence status (as previously published by the user) | [optional] 
**ActiveCalls** | [**[]ActiveCallInfo**](ActiveCallInfo.md) | Information on active calls | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


