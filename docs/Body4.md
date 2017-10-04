# Body4

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Specifies if the rule is active or inactive. The default value is &#39;True&#39; | [optional] [default to null]
**Type_** | **string** | Type of an answering rule, the supported value is &#39;Custom&#39; | [optional] [default to null]
**Name** | **string** | Name of an answering rule specified by user. Max number of symbols is 30 | [optional] [default to null]
**Callers** | [**[]CallersInfo**](CallersInfo.md) | Answering rule will be applied when calls are received from the specified caller(s) | [optional] [default to null]
**CalledNumbers** | [**[]CalledNumberInfo**](CalledNumberInfo.md) | Answering rule will be applied when calling the specified number(s) | [optional] [default to null]
**Schedule** | [**AnsweringRuleScheduleInfo**](AnsweringRule.ScheduleInfo.md) | Schedule when an answering rule should be applied | [optional] [default to null]
**CallHandlingAction** | **string** | Specifies how incoming calls should be forwarded. The default value is &#39;ForwardCalls&#39; | [optional] [default to null]
**Forwarding** | [**ForwardingInfo**](ForwardingInfo.md) | Forwarding parameters. If the &#39;callHandlingAction&#39; parameter value is set to &#39;ForwardCalls&#39; - should be specified . The settings determine the forwarding numbers to which the call should be forwarded. If not specified in request, then the business-hours forwarding rules are set by default | [optional] [default to null]
**UnconditionalForwarding** | [**UnconditionalForwardingInfo**](UnconditionalForwardingInfo.md) | Unconditional forwarding parameters. If the &#39;callHandlingAction&#39; parameter value is set to &#39;UnconditionalForwarding&#39; - should be specified | [optional] [default to null]
**VoiceMail** | [**VoicemailInfo**](VoicemailInfo.md) | Specifies whether to take a voicemail and who should do it | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


