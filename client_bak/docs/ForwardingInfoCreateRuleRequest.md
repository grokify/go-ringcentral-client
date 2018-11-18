# ForwardingInfoCreateRuleRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifyMySoftPhones** | **bool** | Specifies if the first ring on desktop/mobile apps is enabled. The default value is &#39;True&#39; | [optional] 
**NotifyAdminSoftPhones** | **bool** | Specifies if the administrator&#39;s softphone (desktop application) is notified before forwarding the incoming call to desk phones and forwarding numbers. The default value is &#39;True&#39; | [optional] 
**SoftPhonesRingCount** | **int32** | Specifies delay between ring on apps and starting of a call forwarding. The default value is 1 | [optional] 
**RingingMode** | **string** | Specifies the order in which forwarding numbers ring. &#39;Sequentially&#39; means that forwarding numbers are ringing one at a time, in order of priority. &#39;Simultaneously&#39; means that forwarding numbers are ringing all at the same time. The default value is &#39;Sequentially&#39; | [optional] 
**Rules** | [**[]RuleInfoCreateRuleRequest**](RuleInfoCreateRuleRequest.md) | Information on a call forwarding rule | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


