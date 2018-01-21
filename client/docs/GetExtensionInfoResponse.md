# GetExtensionInfoResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Internal identifier of an extension | [default to null]
**Uri** | **string** | Canonical URI of an extension | [default to null]
**Contact** | [***ContactInfo**](ContactInfo.md) | Contact detailed information | [optional] [default to null]
**Departments** | [**[]DepartmentInfo**](DepartmentInfo.md) | Information on department extension(s), to which the requested extension belongs. Returned only for user extensions, members of department, requested by single extensionId | [optional] [default to null]
**ExtensionNumber** | **string** | Number of department extension | [optional] [default to null]
**Name** | **string** | Extension user name | [optional] [default to null]
**PartnerId** | **string** | For Partner Applications Internal identifier of an extension created by partner. The RingCentral supports the mapping of accounts and stores the corresponding account ID/extension ID for each partner ID of a client application. In request URIs partner IDs are accepted instead of regular RingCentral native IDs as path parameters using pid &#x3D; XXX clause. Though in response URIs contain the corresponding account IDs and extension IDs. In all request and response bodies these values are reflected via partnerId attributes of account and extension | [optional] [default to null]
**Permissions** | [***ExtensionPermissions**](ExtensionPermissions.md) | Extension permissions, corresponding to the Service Web permissions &#39;Admin&#39; and &#39;InternationalCalling&#39; | [optional] [default to null]
**ProfileImage** | [***ProfileImageInfo**](ProfileImageInfo.md) | Information on profile image | [default to null]
**References** | [**[]ReferenceInfo**](ReferenceInfo.md) | List of non-RC internal identifiers assigned to an extension | [optional] [default to null]
**RegionalSettings** | [***RegionalSettings**](RegionalSettings.md) | Extension region data (timezone, home country, language) | [optional] [default to null]
**ServiceFeatures** | [**[]ExtensionServiceFeatureInfo**](ExtensionServiceFeatureInfo.md) | Extension service features returned in response only when the logged-in user requests his/her own extension info, see also Extension Service Features | [optional] [default to null]
**SetupWizardState** | **string** | Specifies extension configuration wizard state (web service setup). The default value is &#39;NotStarted&#39; | [optional] [default to null]
**Status** | **string** | Extension current state. If the status is &#39;Unassigned&#39;. Returned for all extensions | [default to null]
**StatusInfo** | [***ExtensionStatusInfo**](ExtensionStatusInfo.md) | Status information (reason, comment). Returned for &#39;Disabled&#39; status only | [optional] [default to null]
**Type_** | **string** | Extension type | [default to null]
**CallQueueInfo** | [***CallQueueInfo**](CallQueueInfo.md) | For Department extension type only. Call queue settings | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


