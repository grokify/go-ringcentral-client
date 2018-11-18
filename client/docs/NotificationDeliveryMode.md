# NotificationDeliveryMode

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransportType** | **string** | Notifications transportation provider name. &#39;APNS&#39; (Apple Push Notifications Service) | [optional] 
**Encryption** | **bool** | Optional parameter. Specifies if the message will be encrypted or not. For APNS transport type the value is always  false | [optional] 
**Address** | **string** | PubNub channel name. For APNS transport type - internal identifier of a device  device_token | [optional] 
**SubscriberKey** | **string** | PubNub subscriber credentials required to subscribe to the channel | [optional] 
**SecretKey** | **string** | PubNub subscriber credentials required to subscribe to the channel. Optional (for PubNub transport type only) | [optional] 
**EncryptionAlgorithm** | **string** | Encryption algorithm &#39;AES&#39; (for PubNub transport type only) | [optional] 
**EncryptionKey** | **string** | Key for notification message decryption (for PubNub transport type only) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


