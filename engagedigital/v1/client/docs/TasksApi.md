# \TasksApi

All URIs are relative to *https://DOMAIN.api.engagement.dimelo.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllTasks**](TasksApi.md#GetAllTasks) | **Get** /tasks | Getting all tasks
[**GetTask**](TasksApi.md#GetTask) | **Get** /tasks/{taskId} | Getting a task from its id
[**MoveTask**](TasksApi.md#MoveTask) | **Delete** /tasks/{taskId}/move | Move a task to another queue
[**TransferTask**](TasksApi.md#TransferTask) | **Put** /tasks/{taskId}/transfer | Transferring a task



## GetAllTasks

> GetAllTasksResponse GetAllTasks(ctx, optional)
Getting all tasks

This method renders tasks ordered by priority (highest first) and then by creation date (latest first).  Authorization​: only users that can read tasks.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllTasksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queue** | **optional.String**| To filter tasks on given queue name (filters on the “global” queue by default).The most commonly used queues are: “global” (contains every task pending assignation), “workbin_{agent_id}” (contains every tasks assigned to the {agent_id} agent, “history” (contains every processed tasks), and “undelivered” (contains every undelivered tasks). If queue is set to “workbins” all the tasks currently in a workbin will be returned. | 
 **channelId** | **optional.String**| To filter tasks on given channel id. | 
 **step** | **optional.String**| To filter tasks on the step they’re currently in. | 
 **offset** | **optional.Int32**| The record index to start. Default value is 0. | 
 **limit** | **optional.Int32**| The max number of records to return. Default value is 30, max value is 150. | 

### Return type

[**GetAllTasksResponse**](GetAllTasksResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTask

> Task GetTask(ctx, taskId)
Getting a task from its id

This method renders a task from given id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**|  | 

### Return type

[**Task**](Task.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveTask

> Task MoveTask(ctx, taskId, queue)
Move a task to another queue

This method changes a task queue and renders it in case of success. Only accepts “undelivered” and special queue defined in topology (e.g. triage).  Authorization​: only users that have the right to monitor the task view.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**|  | 
**queue** | **string**| Name of the queue task has to be moved in. | 

### Return type

[**Task**](Task.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferTask

> Task TransferTask(ctx, taskId, optional)
Transferring a task

This method transfers an existing task and renders it in case of success.  Authorization​: only users that have the right to monitor the task view.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**|  | 
 **optional** | ***TransferTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TransferTaskOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentIds** | [**optional.Interface of []string**](string.md)| List of agents to transfer the task to (multiple). | 
 **bypass** | **optional.String**| Force the transfer to the first agent in agent_ids if set. When bypass is used, | 
 **categoryIds** | [**optional.Interface of []string**](string.md)| Filter agents receiving the task depending on their categories. | 
 **language** | **optional.String**| Filter agents receiving the task depending on their spoken languages. | 
 **teamIds** | [**optional.Interface of []string**](string.md)| Filter agents receiving the task depending on their teams. | 
 **comment** | **optional.String**| Add a comment to the task. | 

### Return type

[**Task**](Task.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

