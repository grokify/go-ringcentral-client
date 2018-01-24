# ExtensionCreationRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | [***ContactInfoUpdateRequest**](ContactInfoUpdateRequest.md) | Contact Information | [optional] [default to null]
**ExtensionNumber** | **string** | Number of extension | [optional] [default to null]
**Password** | **string** | Password for extension. If not specified, the password is auto-generated | [optional] [default to null]
**References** | [**[]ReferenceInfo**](ReferenceInfo.md) | List of non-RC internal identifiers assigned to an extension | [optional] [default to null]
**RegionalSettings** | [***RegionalSettings**](RegionalSettings.md) | Extension region data (timezone, home country, language) | [optional] [default to null]
**SetupWizardState** | **string** | Specifies extension configuration wizard state (web service setup). The default value is &#39;NotStarted&#39; &#x3D; [&#39;NotStarted&#39;, &#39;Incomplete&#39;, &#39;Completed&#39;] | [optional] [default to null]
**Status** | **string** | Extension current state &#x3D; [&#39;Enabled&#39;, &#39;Disabled&#39;, &#39;NotActivated&#39;, &#39;Unassigned&#39;] | [optional] [default to null]
**StatusInfo** | [***ExtensionStatusInfo**](ExtensionStatusInfo.md) | Status information (reason, comment). For &#39;Disabled&#39; status only | [optional] [default to null]
**Type_** | **string** | Extension type &#x3D; [&#39;User&#39;, &#39;VirtualUser&#39;, &#39;DigitalUser&#39;, &#39;Department&#39;] | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


