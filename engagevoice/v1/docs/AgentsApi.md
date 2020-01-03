# \AgentsApi

All URIs are relative to *https://portal.vacd.biz/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAgentGroups**](AgentsApi.md#GetAgentGroups) | **Get** /admin/accounts/{accountId}/agentGroups | Get Agent Groups
[**GetAgents**](AgentsApi.md#GetAgents) | **Get** /admin/accounts/{accountId}/agentGroups/{agentGroupId}/agents | Get Group Agents



## GetAgentGroups

> []AgentGroup GetAgentGroups(ctx, accountId)
Get Agent Groups

Returns a listing of agent groups for an account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 

### Return type

[**[]AgentGroup**](AgentGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgents

> []Agent GetAgents(ctx, accountId, agentGroupId)
Get Group Agents

Returns a listing of agents for an agent group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**agentGroupId** | **string**|  | 

### Return type

[**[]Agent**](Agent.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

