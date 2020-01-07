# GetDeviceInfoResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of a device | [optional] 
**Uri** | **string** | Canonical URI of a device | [optional] 
**Sku** | **string** | Device identification number (stock keeping unit) in the format TP-ID [-AT-AC], where TP is device type (HP for RC HardPhone, DV for all other devices including softphone); ID - device model ID; AT -addon type ID; AC - addon count (if any). For example &#39;HP-56-2-2&#39; | [optional] 
**Type** | **string** | Device type. The default value is &#39;HardPhone&#39; | [optional] 
**Name** | **string** | Device name. Mandatory if ordering  SoftPhone  or  OtherPhone . Optional for  HardPhone . If not specified for HardPhone, then device model  name is used as device  name | [optional] 
**Serial** | **string** | Serial number for HardPhone (is returned only when the phone is shipped and provisioned); endpoint_id for softphone and mobile applications | [optional] 
**ComputerName** | **string** | PC name for softphone | [optional] 
**Model** | [**ModelInfo**](ModelInfo.md) |  | [optional] 
**Extension** | [**ExtensionInfoIntId**](ExtensionInfoIntId.md) |  | [optional] 
**EmergencyServiceAddress** | [**EmergencyAddressInfo**](EmergencyAddressInfo.md) |  | [optional] 
**PhoneLines** | [**[]PhoneLinesInfo**](PhoneLinesInfo.md) | Phone lines information | [optional] 
**Shipping** | [**ShippingInfo**](ShippingInfo.md) |  | [optional] 
**BoxBillingId** | **int32** | Box billing identifier of a device. Applicable only for HardPhones. It is an alternative way to identify the device to be ordered. EitherT? model  structure, or  boxBillingId  must be specified forT?HardPhone | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


