# Body5

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Specifies if the answering rule is active or not | [optional] [default to null]
**Name** | **string** | Custom name of an answering rule. The maximum number of characters is 64 | [optional] [default to null]
**Forwarding** | [**ForwardingInfo**](ForwardingInfo.md) | Forwarding parameters. Returned if &#39;ForwardCalls&#39; is specified in &#39;callHandlingAction&#39;. These settings determine the forwarding numbers to which the call will be forwarded | [optional] [default to null]
**Greetings** | [**[]GreetingInfo**](GreetingInfo.md) | Predefined greetings applied for an answering rule | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


