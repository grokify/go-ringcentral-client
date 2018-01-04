# CreateAnsweringRuleRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | Type of an answering rule. The &#39;Custom&#39; value should be specified | [default to null]
**Name** | **string** | Name of an answering rule specified by user | [default to null]
**Callers** | [**[]CallersInfoRequest**](CallersInfoRequest.md) | Answering rule will be applied when calls are received from the specified caller(s) | [optional] [default to null]
**Forwarding** | [***ForwardingInfoCreateRuleRequest**](ForwardingInfoCreateRuleRequest.md) | Forwarding parameters. Should be specified if the &#39;callHandlingAction&#39; parameter value is set to &#39;ForwardCalls&#39;. These settings determine the forwarding numbers to which the call should be forwarded | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


