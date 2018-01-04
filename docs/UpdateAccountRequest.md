# UpdateAccountRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Target account status. For account activation - &#39;Unconfirmed&#39;. For account confirmation - &#39;Confirmed&#39;. For changing account status - &#39;Confirmed&#39; or &#39;Disabled&#39; &#x3D; [&#39;Unconfirmed&#39;, &#39;Confirmed&#39;, &#39;Disabled&#39;], | [optional] [default to null]
**StatusInfo** | [***AccountStatusInfo**](AccountStatusInfo.md) | Status information (reason, comment, lifetime). Returned for &#39;Disabled&#39; status only | [optional] [default to null]
**TransitionInfo** | [***TransitionInfo**](TransitionInfo.md) | Email notifications setting | [optional] [default to null]
**PartnerId** | **string** | Additional account identifier, developed and applied on the client side  | [optional] [default to null]
**ServiceInfo** | [***AccountServiceInfo**](AccountServiceInfo.md) | Account service information, including brand, service plan and billing plan | [optional] [default to null]
**RegionalSettings** | [***RegionalSettings**](RegionalSettings.md) | Account level region data (web service Auto-Receptionist settings) | [optional] [default to null]
**OperatorId** | **string** | Identifier of extension to be set as operator for account | [optional] [default to null]
**SignupInfo** | [***SignupInfoResource**](SignupInfoResource.md) | Account sign up data | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


