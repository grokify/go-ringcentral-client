# \EventsApi

All URIs are relative to *https://DOMAIN.api.engagement.dimelo.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllEvents**](EventsApi.md#GetAllEvents) | **Get** /events | Getting all events
[**GetEvent**](EventsApi.md#GetEvent) | **Get** /events/{eventId} | Getting an event from its id



## GetAllEvents

> GetAllEventsResponse GetAllEvents(ctx, optional)
Getting all events

This method renders events ordered by creation date (descending).  Authorization​: Only users whose role can search event permission.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| To filter events on given query. Query works exactly like threads query but only have those keywords: content, content_thread, name_in, created_before, created_after, user. Order can be created_at.desc (default) or created_at.asc. e.g. q&#x3D;name_in:\&quot;content.replied\&quot;%20content_thread:\&quot;7f946431b6eebffafae642cc\&quot;%20created_after:\&quot;2014-03-20\&quot;%20user:\&quot;4ee91f197aa58d01b500000f\&quot;%20order:\&quot;created_at.asc\&quot; * DateTime parameters should be ISO-8601 * you can specify multiple value for a given keyword: q&#x3D;name_in:’content.replied’&amp;name_in:’content.ignored’ Please refer to ​Search &amp; filtering parameters​ for more details. | 
 **offset** | **optional.Int32**| The record index to start. Default value is 0. | 
 **limit** | **optional.Int32**| The max number of records to return. Default value is 30, max value is 150. | 

### Return type

[**GetAllEventsResponse**](GetAllEventsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvent

> Event GetEvent(ctx, eventId)
Getting an event from its id

This method renders an event from given id. If token’s user role does not have “search event” permission a 404 HTTP response will be returned.  Authorization​: Only users who’s role can search event permission.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string**|  | 

### Return type

[**Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

