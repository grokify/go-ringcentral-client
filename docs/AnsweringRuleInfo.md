# AnsweringRuleInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | **string** | Canonical URI to the answering rule resource | [optional] [default to null]
**Id** | **string** | Internal identifier of an asnwering rule | [optional] [default to null]
**Type_** | **string** | Type of an answering rule | [optional] [default to null]
**Name** | **string** | Name of an answering rule specified by user | [optional] [default to null]
**Enabled** | **bool** | Specifies if an answering rule is active or inactive | [optional] [default to null]
**Schedule** | [**ScheduleInfo**](ScheduleInfo.md) | Schedule when an answering rule should be applied | [optional] [default to null]
**CalledNumbers** | [**[]AnsweringRuleInfoCalleeInfo**](AnsweringRuleInfo.CalleeInfo.md) | Answering rules are applied when calling to selected number(s) | [optional] [default to null]
**Callers** | [**[]AnsweringRuleInfoCallerInfo**](AnsweringRuleInfo.CallerInfo.md) | Answering rules are applied when calls are received from specified caller(s) | [optional] [default to null]
**CallHandlingAction** | **string** | Specifies how incoming calls are forwarded | [optional] [default to null]
**Forwarding** | [**ForwardingInfo**](ForwardingInfo.md) | Forwarding parameters. Returned if &#39;ForwardCalls&#39; is specified in &#39;callHandlingAction&#39;. These settings determine the forwarding numbers to which the call will be forwarded | [optional] [default to null]
**UnconditionalForwarding** | [**UnconditionalForwardingInfo**](UnconditionalForwardingInfo.md) | Unconditional forwarding parameters. Returned if &#39;UnconditionalForwarding&#39; is specified in &#39;callHandlingAction&#39; | [optional] [default to null]
**Voicemail** | [**VoicemailInfo**](VoicemailInfo.md) | Specifies whether to take a voicemail and who should do it | [optional] [default to null]
**Greetings** | [**[]GreetingInfo**](GreetingInfo.md) | Predefined greetings applied for an answering rule | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


