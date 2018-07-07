# UpdateAnsweringRuleRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of an answering rule specified by user | [optional] 
**Enabled** | **bool** | Specifies if an answering rule is active or inactive | [optional] 
**Callers** | [**[]CallersInfo**](CallersInfo.md) | Answering rules are applied when calls are received from specified caller(s) | [optional] 
**CalledNumbers** | [**[]CalledNumberInfo**](CalledNumberInfo.md) | Answering rules are applied when calling to selected number(s) | [optional] 
**Schedule** | [**ScheduleInfo**](ScheduleInfo.md) |  | [optional] 
**CallHandlingAction** | **string** | Specifies how incoming calls are forwarded | [optional] 
**Forwarding** | [**ForwardingInfo**](ForwardingInfo.md) |  | [optional] 
**UnconditionalForwarding** | [**UnconditionalForwardingInfo**](UnconditionalForwardingInfo.md) |  | [optional] 
**Queue** | [**QueueInfo**](QueueInfo.md) |  | [optional] 
**Voicemail** | [**VoicemailInfo**](VoicemailInfo.md) |  | [optional] 
**Greetings** | [**[]GreetingInfo**](GreetingInfo.md) | Greetings applied for an answering rule; only predefined greetings can be applied, see Dictionary Greeting List | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


