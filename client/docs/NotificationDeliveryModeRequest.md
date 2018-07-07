# NotificationDeliveryModeRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransportType** | **string** | Notifications transportation provider name. &#39;APNS&#39; (Apple Push Notifications Service) | [optional] 
**Address** | **string** | Mandatory for &#39;APNS&#39; and &#39;WebHook&#39; transport types. For &#39;APNS&#39; - internal identifier of a device &#39;device_token&#39; for &#39;WebHook&#39; - URL of a consumer service (cannot be changed during subscription update) | [optional] 
**Encryption** | **bool** | Optional parameter. Specifies if the message will be encrypted or not. If request contains any presence event filter the value by default is &#39;True&#39; (even if specified as &#39;false&#39;). If request contains only message event filters the value by default is &#39;False&#39; | [optional] 
**CertificateName** | **string** | For &#39;PubNub/APNS&#39; and &#39;PubNub/GCM&#39; transport types. Name of a certificate | [optional] 
**RegistrationId** | **string** | For &#39;PubNub/APNS&#39; and &#39;PubNub/GCM&#39; transport types. Identifier of a registration | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


