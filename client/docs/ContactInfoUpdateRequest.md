# ContactInfoUpdateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | For User extension type only. Extension user first name, | [optional] 
**LastName** | **string** | For User extension type only. Extension user last name, | [optional] 
**Company** | **string** | Extension user company name | [optional] 
**Email** | **string** | Email of extension user | [optional] 
**BusinessPhone** | **string** | Extension user contact phone number in E.164 format | [optional] 
**BusinessAddress** | [**ContactAddressInfo**](ContactAddressInfo.md) |  | [optional] 
**EmailAsLoginName** | **bool** |  If &#39;True&#39; then contact email is enabled as login name for this user. Please note that email should be unique in this case. The default value is &#39;False&#39; | [optional] 
**PronouncedName** | [**PronouncedNameInfo**](PronouncedNameInfo.md) |  | [optional] 
**Department** | **string** | Extension user department, if any | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


