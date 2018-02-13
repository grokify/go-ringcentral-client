# GetDeviceInfoResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of a device | [optional] [default to null]
**Uri** | **string** | Canonical URI of a device | [optional] [default to null]
**Sku** | **string** | Device identification number (stock keeping unit) in the format TP-ID [-AT-AC], where TP is device type (HP for RC HardPhone, DV for all other devices including softphone); ID - device model ID; AT -addon type ID; AC - addon count (if any). For example &#39;HP-56-2-2&#39; | [optional] [default to null]
**Type_** | **string** | Device type. The default value is &#39;HardPhone&#39; | [optional] [default to null]
**Name** | **string** | Device name. Mandatory if ordering  SoftPhone  or  OtherPhone . Optional for  HardPhone . If not specified for HardPhone, then device model  name is used as device  name | [optional] [default to null]
**Serial** | **string** | Serial number for HardPhone (is returned only when the phone is shipped and provisioned); endpoint_id for softphone and mobile applications | [optional] [default to null]
**ComputerName** | **string** | PC name for softphone | [optional] [default to null]
**Model** | [***ModelInfo**](ModelInfo.md) | HardPhone model information | [optional] [default to null]
**Extension** | [***ExtensionInfoIntId**](ExtensionInfoIntId.md) | This attribute can be omitted for unassigned devices | [optional] [default to null]
**EmergencyServiceAddress** | [***EmergencyAddressInfo**](EmergencyAddressInfo.md) | Address for emergency cases. The same emergency address is assigned to all the numbers of one device | [optional] [default to null]
**PhoneLines** | [**[]PhoneLinesInfo**](PhoneLinesInfo.md) | Phone lines information | [optional] [default to null]
**Shipping** | [***ShippingInfo**](ShippingInfo.md) | Shipping information, according to which devices (in case of HardPhone ) or e911 stickers (in case of  SoftPhone  and  OtherPhone ) will be delivered to the customer | [optional] [default to null]
**BoxBillingId** | **int32** | Box billing identifier of a device. Applicable only for HardPhones. It is an alternative way to identify the device to be ordered. EitherT? model  structure, or  boxBillingId  must be specified forT?HardPhone | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


