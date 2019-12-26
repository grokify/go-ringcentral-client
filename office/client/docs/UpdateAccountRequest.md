# UpdateAccountRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Target account status. For account activation - &#39;Unconfirmed&#39;. For account confirmation - &#39;Confirmed&#39;. For changing account status - &#39;Confirmed&#39; or &#39;Disabled&#39; &#x3D; [&#39;Unconfirmed&#39;, &#39;Confirmed&#39;, &#39;Disabled&#39;], | [optional] 
**StatusInfo** | [**AccountStatusInfo**](AccountStatusInfo.md) |  | [optional] 
**TransitionInfo** | [**TransitionInfo**](TransitionInfo.md) |  | [optional] 
**PartnerId** | **string** | Additional account identifier, developed and applied on the client side | [optional] 
**ServiceInfo** | [**AccountServiceInfo**](AccountServiceInfo.md) |  | [optional] 
**RegionalSettings** | [**RegionalSettings**](RegionalSettings.md) |  | [optional] 
**OperatorId** | **string** | Identifier of extension to be set as operator for account | [optional] 
**SignupInfo** | [**SignupInfoResource**](SignupInfoResource.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


