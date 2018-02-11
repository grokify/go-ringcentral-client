# IncomingCallEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aps** | [***ApsInfo**](APSInfo.md) | Apple Push Notification Service Info | [optional] [default to null]
**Event** | **string** | Event filter URI | [optional] [default to null]
**Uuid** | **string** | Universally unique identifier of a notification | [optional] [default to null]
**SubscriptionId** | **string** | Internal identifier of a subscription | [optional] [default to null]
**Timestamp** | [**time.Time**](time.Time.md) | The datetime of a call action in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z | [optional] [default to null]
**ExtensionId** | **string** | Internal identifier of an extension | [optional] [default to null]
**Action** | **string** | Calling action, for example &#39;StartRing&#39; | [optional] [default to null]
**SessionId** | **string** | Identifier of a call session | [optional] [default to null]
**ServerId** | **string** | Identifier of a server | [optional] [default to null]
**From** | **string** | Phone number of a caller | [optional] [default to null]
**FromName** | **string** | Caller name | [optional] [default to null]
**To** | **string** | Phone number of a callee | [optional] [default to null]
**ToName** | **string** | Callee name | [optional] [default to null]
**Sid** | **string** | Unique identifier of a session | [optional] [default to null]
**ToUrl** | **string** | SIP proxy registration name | [optional] [default to null]
**SrvLvl** | **string** | User data | [optional] [default to null]
**SrvLvlExt** | **string** | User data | [optional] [default to null]
**RecUrl** | **string** | File containing recorded caller name | [optional] [default to null]
**PnTtl** | **int32** | Notification lifetime value in seconds, the default value is 30 seconds | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


