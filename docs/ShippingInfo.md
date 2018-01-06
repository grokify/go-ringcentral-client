# ShippingInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Shipping status of the order item. It is set to &#39;Initial&#39; when the order is submitted. Then it is changed to &#39;Accepted&#39; when a distributor starts processing the order. Finally it is changed to Shipped which means that distributor has shipped the device. | [optional] [default to null]
**Carrier** | **string** | Shipping carrier name. Appears only if the device status is  Shipped  | [optional] [default to null]
**TrackingNumber** | **string** | Carrier-specific tracking number. Appears only if the device status is  Shipped | [optional] [default to null]
**Method** | [***MethodInfo**](MethodInfo.md) | Shipping method information | [default to null]
**Address** | [***ShippingAddressInfo**](ShippingAddressInfo.md) | Shipping address for the order. If it coincides with the Emergency Service Address, then can be omitted. By default the same value as the emergencyServiceAddress. Multiple addresses can be specified; in case an order contains several devices, they can be delivered to different addresses | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


