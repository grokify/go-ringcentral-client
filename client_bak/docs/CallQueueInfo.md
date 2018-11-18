# CallQueueInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SlaGoal** | **int32** | Target percentage of calls that must be answered by agents within the service level time threshold | [optional] 
**SlaThresholdSeconds** | **int32** | Period of time in seconds that is considered to be an acceptable service level | [optional] 
**IncludeAbandonedCalls** | **bool** | If &#39;True&#39; abandoned calls (hanged up prior to being served) are included into service level calculation | [optional] 
**AbandonedThresholdSeconds** | **int32** | Period of time in seconds specifying abandoned calls duration - calls that are shorter will not be included into the calculation of service level.; zero value means that abandoned calls of any duration will be included into calculation | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


