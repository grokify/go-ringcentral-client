# CreateSubscriptionRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventFilters** | **[]string** | Mandatory. Collection of URIs to API resources (see Event Types for details). For APNS transport type only message event filter is available | [optional] [default to null]
**DeliveryMode** | [***NotificationDeliveryModeRequest**](NotificationDeliveryModeRequest.md) | Notification delivery settings | [default to null]
**ExpiresIn** | **int32** | Subscription lifetime in seconds. Max value is 7 days (604800 sec) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


