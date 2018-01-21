# QueueInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransferMode** | **string** | Specifies how calls are transferred to group members | [optional] [default to null]
**FixedOrderAgents** | [**[]FixedOrderAgents**](FixedOrderAgents.md) | Information on a call forwarding rule | [optional] [default to null]
**HoldAudioInterruptionMode** | **string** | Connecting audio interruption mode | [optional] [default to null]
**HoldAudioInterruptionPeriod** | **int32** | Connecting audio interruption message period in seconds | [optional] [default to null]
**AgentTimeout** | **int32** | Maximum time in seconds to wait for a call queue member before trying the next member | [optional] [default to null]
**WrapUpTime** | **int32** | Minimum post-call wrap up time in seconds before agent status is automatically set | [optional] [default to null]
**HoldTime** | **int32** | Maximum hold time in seconds to wait for an available call queue member | [optional] [default to null]
**MaxCallers** | **int32** | Maximum count of callers on hold | [optional] [default to null]
**MaxCallersAction** | **string** | Action which should be taken if count of callers on hold exceeds the maximum | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


