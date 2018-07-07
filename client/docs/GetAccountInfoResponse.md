# GetAccountInfoResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of an account | [optional] 
**Uri** | **string** | Canonical URI of an account | [optional] 
**MainNumber** | **string** | Main phone number of the current account | [optional] 
**Operator** | [**GetExtensionInfoResponse**](GetExtensionInfoResponse.md) |  | [optional] 
**PartnerId** | **string** | Additional account identifier, developed and applied by the client | [optional] 
**ServiceInfo** | [**ServiceInfo**](ServiceInfo.md) |  | [optional] 
**SetupWizardState** | **string** | Specifies account configuration wizard state (web service setup). The default value is &#39;NotStarted&#39; | [optional] 
**Status** | **string** | Status of the current account | [optional] 
**StatusInfo** | [**AccountStatusInfo**](AccountStatusInfo.md) |  | [optional] 
**RegionalSettings** | [**RegionalSettings**](RegionalSettings.md) |  | [optional] 
**Federated** | **bool** | Specifies whether an account is included into any federation of accounts or not | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


