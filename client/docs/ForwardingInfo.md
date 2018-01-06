# ForwardingInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifyMySoftPhones** | **bool** | Specifies if the user&#39;s softphone(s) are notified before forwarding the incoming call to desk phones and forwarding numbers | [optional] [default to null]
**NotifyAdminSoftPhones** | **bool** | Specifies if the administrator&#39;s softphone is notified before forwarding the incoming call to desk phones and forwarding numbers. The default value is &#39;False&#39; | [optional] [default to null]
**SoftPhonesRingCount** | **int32** | Number of rings before forwarding starts | [optional] [default to null]
**RingingMode** | **string** | Specifies the order in which forwarding numbers ring. &#39;Sequentially&#39; means that forwarding numbers are ringing one at a time, in order of priority. &#39;Simultaneously&#39; means that forwarding numbers are ring all at the same time | [optional] [default to null]
**Rules** | [**[]RuleInfo**](RuleInfo.md) | Information on a call forwarding rule | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


