# AnsweringRuleInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | **string** | Canonical URI to the answering rule resource | [optional] [default to null]
**Id** | **string** | Internal identifier of an answering rule | [optional] [default to null]
**Type_** | **string** | Type of an answering rule | [optional] [default to null]
**Name** | **string** | Name of an answering rule specified by user | [optional] [default to null]
**Enabled** | **bool** | Specifies if an answering rule is active or inactive | [optional] [default to null]
**Schedule** | [***ScheduleInfo**](ScheduleInfo.md) | Schedule when an answering rule should be applied | [optional] [default to null]
**CalledNumbers** | [**[]CalledNumberInfo**](CalledNumberInfo.md) | Answering rules are applied when calling to selected number(s) | [optional] [default to null]
**Callers** | [**[]CallersInfo**](CallersInfo.md) | Answering rules are applied when calls are received from specified caller(s) | [optional] [default to null]
**CallHandlingAction** | **string** | Specifies how incoming calls are forwarded | [optional] [default to null]
**Forwarding** | [***ForwardingInfo**](ForwardingInfo.md) | Forwarding parameters. Returned if &#39;ForwardCalls&#39; is specified in &#39;callHandlingAction&#39;. These settings determine the forwarding numbers to which the call will be forwarded | [optional] [default to null]
**UnconditionalForwarding** | [***UnconditionalForwardingInfo**](UnconditionalForwardingInfo.md) | Unconditional forwarding parameters. Returned if &#39;UnconditionalForwarding&#39; is specified in &#39;callHandlingAction&#39; | [optional] [default to null]
**Queue** | [***QueueInfo**](QueueInfo.md) | Queue settings applied for department (call queue) extension type, with the &#39;AgentQueue&#39; value specified as a call handling action | [optional] [default to null]
**Transfer** | [***TransferredExtensionInfo**](TransferredExtensionInfo.md) | Transfer settings applied for department (call queue) extension type, with &#39;TransferToExtension&#39; call handling action | [optional] [default to null]
**Voicemail** | [***VoicemailInfo**](VoicemailInfo.md) | Specifies whether to take a voicemail and who should do it | [optional] [default to null]
**Greetings** | [**[]GreetingInfo**](GreetingInfo.md) | Greetings applied for an answering rule; only predefined greetings can be applied, see Dictionary Greeting List | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


