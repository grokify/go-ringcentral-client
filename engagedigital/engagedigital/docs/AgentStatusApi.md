# \AgentStatusApi

All URIs are relative to *https://DOMAIN.api.engagement.dimelo.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeAgentStatus**](AgentStatusApi.md#ChangeAgentStatus) | **Put** /status/{agentId} | Changing an agent&#39;s status
[**GetAgentStatus**](AgentStatusApi.md#GetAgentStatus) | **Get** /status/{agentId} | Get a connected agent status
[**GetAllAgentStatus**](AgentStatusApi.md#GetAllAgentStatus) | **Get** /status | Get all connected agents status



## ChangeAgentStatus

> AgentStatus ChangeAgentStatus(ctx, agentId, optional)
Changing an agent's status

This method updates an agent's availability. Can be used to set either channels statuses OR custom  status. If both parameters are provided, ignores custom status.The status parameter​ **MUST** b​e either “away” or “available”.  Authorization​: only users that have the right to monitor the task view.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**|  | 
 **optional** | ***ChangeAgentStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ChangeAgentStatusOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **optional.String**| A hash of channel_id &#x3D;&gt; availability (must contain all channels). | 
 **customStatusId** | **optional.String**| id of presence status (optional) | 

### Return type

[**AgentStatus**](AgentStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentStatus

> AgentStatus GetAgentStatus(ctx, agentId)
Get a connected agent status

This method get the status of a connected agent. Returns a 404 if the user does not exist (not_found) or if he’s not connected (disconnected).  Authorization​: only users that have the right to monitor the task view.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**|  | 

### Return type

[**AgentStatus**](AgentStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllAgentStatus

> []AgentStatus GetAllAgentStatus(ctx, )
Get all connected agents status

This method get all currently connected agents & their status.  Authorization​: only users that have the right to monitor the task view.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]AgentStatus**](AgentStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

