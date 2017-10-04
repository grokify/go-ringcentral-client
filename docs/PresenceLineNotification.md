# PresenceLineNotification

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | Universally unique identifier of a notification | [optional] [default to null]
**Event** | **string** | Event filter URI | [optional] [default to null]
**SubscriptionId** | **string** | Internal identifier of a subscription | [optional] [default to null]
**Timestamp** | [**time.Time**](time.Time.md) | The datetime of sending a notification in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z | [optional] [default to null]
**Body** | [**PresenceLineEvent**](PresenceLineEvent.md) | Notification payload body | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


