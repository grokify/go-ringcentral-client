# AnsweringRuleResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | **string** |  | [optional] 
**Id** | **string** |  | [optional] 
**Type** | **string** |  | [optional] 
**Name** | **string** |  | [optional] 
**Enabled** | **bool** |  | [optional] [default to false]
**Schedule** | [**ScheduleResource**](ScheduleResource.md) |  | [optional] 
**Callers** | [**[]CallerResource**](CallerResource.md) |  | [optional] 
**CalledNumbers** | [**[]PhoneNumberResource**](PhoneNumberResource.md) |  | [optional] 
**CallHandlingAction** | **string** |  | [optional] 
**Forwarding** | [**AnsweringForwardingResource**](AnsweringForwardingResource.md) |  | [optional] 
**UnconditionalForwarding** | [**UnconditionalForwardingResource**](UnconditionalForwardingResource.md) |  | [optional] 
**Voicemail** | [**VoicemailSettingsResource**](VoicemailSettingsResource.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


