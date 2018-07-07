# DeviceResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of a device | [optional] 
**Uri** | **string** | Canonical URI of a device | [optional] 
**Sku** | **string** | Device identification number (stock keeping unit) in the format TP-ID [-AT-AC], where TP is device type (HP for RC HardPhone, DV for all other devices including softphone); ID - device model ID; AT -addon type ID; AC - addon count (if any). For example &#39;HP-56-2-2&#39; | [optional] 
**Type** | **string** | Device type. The default value is &#39;HardPhone&#39; | [optional] 
**Status** | **string** | Status of a device &#x3D; [&#39;Online&#39;, &#39;Offline&#39;] | [optional] 
**Name** | **string** | Device name. Mandatory if ordering SoftPhone or OtherPhone . Optional for HardPhone . If not specified for HardPhone, then device model name is used as device name | [optional] 
**Serial** | **string** | Serial number for HardPhone (is returned only when the phone is shipped and provisioned); endpoint_id for softphone and mobile applications | [optional] 
**ComputerName** | **string** | PC name for softphone | [optional] 
**Model** | [**DeviceModelResource**](DeviceModelResource.md) |  | [optional] 
**Extension** | [**ExtensionResourceIntId**](ExtensionResourceIntId.md) |  | [optional] 
**PhoneLines** | [**[]PhoneLineResource**](PhoneLineResource.md) | Phone lines information | [optional] 
**EmergencyServiceAddress** | [**EmergencyServiceAddressResource**](EmergencyServiceAddressResource.md) |  | [optional] 
**Shipping** | [**ShippingResource**](ShippingResource.md) |  | [optional] 
**BoxBillingId** | **int32** | Box billing identifier of a device. Applicable only for HardPhones. It is an alternative way to identify the device to be ordered. Either model structure, or boxBillingId must be specified for HardPhone | [optional] 
**LinePooling** | **string** | Pooling type of a deviceHost - device with standalone paid phone line which can be linked to Glip/Softphone instanceGuest - device with a linked phone lineNone - device without a phone line or with specific line (free, BLA, etc.) &#x3D; [&#39;Host&#39;, &#39;Guest&#39;, &#39;None&#39;] | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


