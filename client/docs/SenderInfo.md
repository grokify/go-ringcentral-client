# SenderInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | **string** | Phone number in E.164 (with &#39;+&#39; sign) format | [optional] 
**ExtensionNumber** | **string** | Extension number | [optional] 
**Location** | **string** | Contains party location (city, state) if one can be determined from phoneNumber. This property is filled only when phoneNumber is not empty and server can calculate location information from it (for example, this information is unavailable for US toll-free numbers) | [optional] 
**Name** | **string** | Symbolic name associated with a caller/callee. If the phone does not belong to the known extension, only the location is returned, the name is not determined then | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


