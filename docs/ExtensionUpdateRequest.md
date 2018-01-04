# ExtensionUpdateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | [optional] [default to null]
**StatusInfo** | [***ExtensionStatusInfo**](ExtensionStatusInfo.md) |  | [optional] [default to null]
**Reason** | **string** | Type of suspension | [optional] [default to null]
**Comment** | **string** | Free Form user comment | [optional] [default to null]
**ExtensionNumber** | **string** | Extension number available | [optional] [default to null]
**Contact** | [***ContactInfoUpdateRequest**](ContactInfoUpdateRequest.md) |  | [optional] [default to null]
**RegionalSettings** | [***ExtensionRegionalSettingRequest**](ExtensionRegionalSettingRequest.md) |  | [optional] [default to null]
**SetupWizardState** | **string** |  | [optional] [default to null]
**PartnerId** | **string** |  Extension partner identifier | [optional] [default to null]
**IvrPin** | **string** | IVR PIN | [optional] [default to null]
**Password** | **string** | Password for extension | [optional] [default to null]
**CallQueueInfo** | [***CallQueueInfoRequest**](CallQueueInfoRequest.md) | For Department extension type only. Call queue settings | [optional] [default to null]
**Transition** | **string** | For NotActivated extensions only. Welcome email setting | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


