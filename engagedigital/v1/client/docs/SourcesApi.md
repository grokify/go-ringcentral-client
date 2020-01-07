# \SourcesApi

All URIs are relative to *https://DOMAIN.api.engagement.dimelo.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllSources**](SourcesApi.md#GetAllSources) | **Get** /content_sources | Getting all sources
[**GetSource**](SourcesApi.md#GetSource) | **Get** /content_sources/{sourceId} | Getting a source from its id
[**UpdateSource**](SourcesApi.md#UpdateSource) | **Put** /content_sources/{sourceId} | Updating a source



## GetAllSources

> GetAllSourcesResponse GetAllSources(ctx, optional)
Getting all sources

This method renders sources ordered by creation date (ascending).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllSourcesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllSourcesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| The record index to start. Default value is 0. | 
 **limit** | **optional.Int32**| The max number of records to return. Default value is 30, max value is 150. | 

### Return type

[**GetAllSourcesResponse**](GetAllSourcesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSource

> Source GetSource(ctx, sourceId)
Getting a source from its id

This method renders a source from given id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string**|  | 

### Return type

[**Source**](Source.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSource

> Source UpdateSource(ctx, sourceId, optional)
Updating a source

This method updates an existing source from given attributes and renders it in case of success.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string**|  | 
 **optional** | ***UpdateSourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| Source name | 
 **active** | **optional.Bool**| Activate/deactivate the source (Boolean) | 
 **channelId** | **optional.String**| Channel | 
 **color** | **optional.Int32**| Color of the icon (see S​ ource colors​) (Integer) | 
 **slaResponse** | **optional.Int32**| Response time (seconds) | 
 **slaExpiredStrategy** | **optional.String**| SLA expired strategy (\&quot;max\&quot;, \&quot;half\&quot; or \&quot;base\&quot;) | 
 **interventionMessagesBoost** | **optional.Int32**| Priority boost of messages with intervention (Integer) | 
 **transferredTasksBoost** | **optional.Int32**| Priority boost of transferred tasks (Integer) | 
 **hiddenFromStats** | **optional.Bool**| Hide from statistics (Boolean) | 
 **defaultCategoryIds** | [**optional.Interface of []string**](string.md)| Default categories | 
 **userThreadDefaultCategoryIds** | [**optional.Interface of []string**](string.md)| Default categories (agent messages) | 
 **defaultContentLanguage** | **optional.String**| Default content language | 
 **autoDetectContentLanguage** | **optional.Bool**| Auto-detect content language (Boolean) | 
 **contentArchiving** | **optional.Bool**| Automatic archiving of old contents (Boolean) | 
 **contentArchivingPeriod** | **optional.Int32**| Archive contents older than (seconds) | 

### Return type

[**Source**](Source.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

