# QueueInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransferMode** | **string** | Specifies how calls are transferred to group members | [optional] 
**FixedOrderAgents** | [**[]FixedOrderAgents**](FixedOrderAgents.md) | Information on a call forwarding rule | [optional] 
**HoldAudioInterruptionMode** | **string** | Connecting audio interruption mode | [optional] 
**HoldAudioInterruptionPeriod** | **int32** | Connecting audio interruption message period in seconds | [optional] 
**AgentTimeout** | **int32** | Maximum time in seconds to wait for a call queue member before trying the next member | [optional] 
**WrapUpTime** | **int32** | Minimum post-call wrap up time in seconds before agent status is automatically set | [optional] 
**HoldTime** | **int32** | Maximum hold time in seconds to wait for an available call queue member | [optional] 
**MaxCallers** | **int32** | Maximum count of callers on hold | [optional] 
**MaxCallersAction** | **string** | Action which should be taken if count of callers on hold exceeds the maximum | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


