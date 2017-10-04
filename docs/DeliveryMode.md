# DeliveryMode

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransportType** | **string** | Notifications transportation provider name. &#39;APNS&#39; (Apple Push Notifications Service) | [optional] [default to null]
**Encryption** | **bool** | Optional parameter. Specifies if the message will be encrypted or not. For APNS transport type the value is always \&quot;false\&quot; | [optional] [default to null]
**Address** | **string** | PubNub channel name. For APNS transport type - internal identifier of a device \&quot;device_token\&quot; | [optional] [default to null]
**SubscriberKey** | **string** | PubNub subscriber credentials required to subscribe to the channel | [optional] [default to null]
**SecretKey** | **string** | PubNub subscriber credentials required to subscribe to the channel. Optional (for PubNub transport type only) | [optional] [default to null]
**EncryptionAlgorithm** | **string** | Encryption algorithm &#39;AES&#39; (for PubNub transport type only) | [optional] [default to null]
**EncryptionKey** | **string** | Key for notification message decryption (for PubNub transport type only) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


