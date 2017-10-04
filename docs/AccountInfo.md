# AccountInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of an account | [optional] [default to null]
**Uri** | **string** | Canonical URI of an account | [optional] [default to null]
**MainNumber** | **string** | Main phone number of the current account | [optional] [default to null]
**Operator** | [**ExtensionInfo**](ExtensionInfo.md) | Operator&#39;s extension information. This extension will receive all calls and messages intended for the operator | [optional] [default to null]
**PartnerId** | **string** | Additional account identifier, developed and applied by the client | [optional] [default to null]
**ServiceInfo** | [**AccountServiceInfo**](Account.ServiceInfo.md) | Account service information, including brand, service plan and billing plan | [optional] [default to null]
**SetupWizardState** | **string** | Specifies account configuration wizard state (web service setup). The default value is &#39;NotStarted&#39; | [optional] [default to null]
**Status** | **string** | Status of the current account | [optional] [default to null]
**StatusInfo** | [**StatusInfo**](StatusInfo.md) | Status information (reason, comment, lifetime). Returned for &#39;Disabled&#39; status only | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


