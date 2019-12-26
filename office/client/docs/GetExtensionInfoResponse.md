# GetExtensionInfoResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Internal identifier of an extension | 
**Uri** | **string** | Canonical URI of an extension | 
**Contact** | [**ContactInfo**](ContactInfo.md) |  | [optional] 
**Departments** | [**[]DepartmentInfo**](DepartmentInfo.md) | Information on department extension(s), to which the requested extension belongs. Returned only for user extensions, members of department, requested by single extensionId | [optional] 
**ExtensionNumber** | **string** | Number of department extension | [optional] 
**Name** | **string** | Extension user name | [optional] 
**PartnerId** | **string** | For Partner Applications Internal identifier of an extension created by partner. The RingCentral supports the mapping of accounts and stores the corresponding account ID/extension ID for each partner ID of a client application. In request URIs partner IDs are accepted instead of regular RingCentral native IDs as path parameters using pid &#x3D; XXX clause. Though in response URIs contain the corresponding account IDs and extension IDs. In all request and response bodies these values are reflected via partnerId attributes of account and extension | [optional] 
**Permissions** | [**ExtensionPermissions**](ExtensionPermissions.md) |  | [optional] 
**ProfileImage** | [**ProfileImageInfo**](ProfileImageInfo.md) |  | 
**References** | [**[]ReferenceInfo**](ReferenceInfo.md) | List of non-RC internal identifiers assigned to an extension | [optional] 
**RegionalSettings** | [**RegionalSettings**](RegionalSettings.md) |  | [optional] 
**ServiceFeatures** | [**[]ExtensionServiceFeatureInfo**](ExtensionServiceFeatureInfo.md) | Extension service features returned in response only when the logged-in user requests his/her own extension info, see also Extension Service Features | [optional] 
**SetupWizardState** | **string** | Specifies extension configuration wizard state (web service setup). The default value is &#39;NotStarted&#39; | [optional] 
**Status** | **string** | Extension current state. If the status is &#39;Unassigned&#39;. Returned for all extensions | 
**StatusInfo** | [**ExtensionStatusInfo**](ExtensionStatusInfo.md) |  | [optional] 
**Type** | **string** | Extension type | 
**CallQueueInfo** | [**CallQueueInfo**](CallQueueInfo.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


