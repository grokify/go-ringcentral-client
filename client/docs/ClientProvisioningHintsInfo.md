# ClientProvisioningHintsInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialState** | [***ClientProvisioningHintInfo**](ClientProvisioningHintInfo.md) | Trial account expiration. Returned for trial accounts only | [optional] [default to null]
**UserCredentialState** | [***ClientProvisioningHintInfo**](ClientProvisioningHintInfo.md) | User credentials expiration | [optional] [default to null]
**AppVersionUpgrade** | [***ClientProvisioningHintInfo**](ClientProvisioningHintInfo.md) | Application version update. Returned only if the client current version is older than the latest version. &#39;actionRequired&#39;: &#39;true&#39; means the application requires force updating to the latest version | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


