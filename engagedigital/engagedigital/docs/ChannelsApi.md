# \ChannelsApi

All URIs are relative to *https://DOMAIN.api.engagement.dimelo.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllChannels**](ChannelsApi.md#GetAllChannels) | **Get** /channels | Getting all channels
[**GetChannel**](ChannelsApi.md#GetChannel) | **Get** /channels/{channelId} | Getting a channel from its id
[**UpdateChannel**](ChannelsApi.md#UpdateChannel) | **Put** /channels/{channelId} | Updating a channel



## GetAllChannels

> GetAllChannelsResponse GetAllChannels(ctx, optional)
Getting all channels

This method renders all channels ordered by date of creation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllChannelsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllChannelsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| The record index to start. Default value is 0. | 
 **limit** | **optional.Int32**| The max number of records to return. Default value is 30, max value is 150. | 

### Return type

[**GetAllChannelsResponse**](GetAllChannelsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannel

> Channel GetChannel(ctx, channelId)
Getting a channel from its id

This method renders a channel from given id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string**|  | 

### Return type

[**Channel**](Channel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChannel

> Channel UpdateChannel(ctx, channelId, optional)
Updating a channel

This method updates an existing channel from given attributes and renders it in case of success.  Authorization​: only users that are able to update channels.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string**|  | 
 **optional** | ***UpdateChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateChannelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| The name of the channel. | 
 **sourceIds** | [**optional.Interface of []string**](string.md)| An array containing id of each source assigned to a channel (multiple). | 
 **softCapability** | **optional.Int32**| Number of tasks that can be assigned to agent by the routing before they are considered \&quot;occupied\&quot;. | 
 **hardCapability** | **optional.Int32**| M​aximum number of tasks that can be assigned to agents. | 
 **taskTimeoutSeconds** | **optional.Int32**| this field defines the time before a task expires (in seconds). | 

### Return type

[**Channel**](Channel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

