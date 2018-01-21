# DeviceResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of a device | [optional] [default to null]
**Uri** | **string** | Canonical URI of a device | [optional] [default to null]
**Sku** | **string** | Device identification number (stock keeping unit) in the format TP-ID [-AT-AC], where TP is device type (HP for RC HardPhone, DV for all other devices including softphone); ID - device model ID; AT -addon type ID; AC - addon count (if any). For example &#39;HP-56-2-2&#39; | [optional] [default to null]
**Type_** | **string** | Device type. The default value is &#39;HardPhone&#39; | [optional] [default to null]
**Status** | **string** | Status of a device &#x3D; [&#39;Online&#39;, &#39;Offline&#39;] | [optional] [default to null]
**Name** | **string** | Device name. Mandatory if ordering SoftPhone or OtherPhone . Optional for HardPhone . If not specified for HardPhone, then device model name is used as device name | [optional] [default to null]
**Serial** | **string** | Serial number for HardPhone (is returned only when the phone is shipped and provisioned); endpoint_id for softphone and mobile applications | [optional] [default to null]
**ComputerName** | **string** | PC name for softphone | [optional] [default to null]
**Model** | [***DeviceModelResource**](DeviceModelResource.md) | HardPhone model information | [optional] [default to null]
**Extension** | [***ExtensionResource**](ExtensionResource.md) | This attribute can be omitted for unassigned devices | [optional] [default to null]
**PhoneLines** | [***PhoneLineResource**](PhoneLineResource.md) | Phone lines information | [optional] [default to null]
**EmergencyServiceAddress** | [***EmergencyServiceAddressResource**](EmergencyServiceAddressResource.md) |  Address for emergency cases. The same emergency address is assigned to all numbers of a single device , | [optional] [default to null]
**Shipping** | [***ShippingResource**](ShippingResource.md) | Shipping information, according to which devices (in case of HardPhone ) or e911 stickers (in case of SoftPhone and OtherPhone ) will be delivered to the customer | [optional] [default to null]
**BoxBillingId** | **int32** | Box billing identifier of a device. Applicable only for HardPhones. It is an alternative way to identify the device to be ordered. Either model structure, or boxBillingId must be specified for HardPhone | [optional] [default to null]
**LinePooling** | **string** | Pooling type of a deviceHost - device with standalone paid phone line which can be linked to Glip/Softphone instanceGuest - device with a linked phone lineNone - device without a phone line or with specific line (free, BLA, etc.) &#x3D; [&#39;Host&#39;, &#39;Guest&#39;, &#39;None&#39;] | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


