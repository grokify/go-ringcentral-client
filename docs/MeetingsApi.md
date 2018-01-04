# \MeetingsApi

All URIs are relative to *https://api.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMeeting**](MeetingsApi.md#CreateMeeting) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/meeting | Create Meetings
[**DeleteMeeting**](MeetingsApi.md#DeleteMeeting) | **Delete** /restapi/v1.0/account/{accountId}/extension/{extensionId}/meeting/{meetingId} | Delete Meeting
[**EndMeeting**](MeetingsApi.md#EndMeeting) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/meeting/{meetingId}/end | End Meeting
[**GetLiveMeetings**](MeetingsApi.md#GetLiveMeetings) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/meeting | Get Scheduled Meetings
[**GetMeetingDetails**](MeetingsApi.md#GetMeetingDetails) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/meeting/{meetingId} | Get Meeting Info
[**GetMeetingServiceInfo**](MeetingsApi.md#GetMeetingServiceInfo) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/meeting/service-info | Get Meeting Service Info
[**UpdateMeeting**](MeetingsApi.md#UpdateMeeting) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/meeting/{meetingId} | Update Meeting


# **CreateMeeting**
> CreateMeeting(ctx, extensionId, accountId, optional)
Create Meetings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **extensionId** | **string**|  | 
  **accountId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **extensionId** | **string**|  | 
 **accountId** | **string**|  | 
 **body** | [**MeetingRequestResource**](MeetingRequestResource.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMeeting**
> DeleteMeeting(ctx, meetingId, extensionId, accountId)
Delete Meeting

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **meetingId** | **string**|  | 
  **extensionId** | **string**|  | 
  **accountId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndMeeting**
> EndMeeting(ctx, meetingId, extensionId, accountId)
End Meeting

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **meetingId** | **string**|  | 
  **extensionId** | **string**|  | 
  **accountId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLiveMeetings**
> MeetingsResource GetLiveMeetings(ctx, extensionId, accountId)
Get Scheduled Meetings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **extensionId** | **string**|  | 
  **accountId** | **string**|  | 

### Return type

[**MeetingsResource**](MeetingsResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMeetingDetails**
> MeetingResponseResource GetMeetingDetails(ctx, meetingId, extensionId, accountId)
Get Meeting Info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **meetingId** | **string**|  | 
  **extensionId** | **string**|  | 
  **accountId** | **string**|  | 

### Return type

[**MeetingResponseResource**](MeetingResponseResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMeetingServiceInfo**
> MeetingServiceInfoResource GetMeetingServiceInfo(ctx, extensionId, accountId)
Get Meeting Service Info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **extensionId** | **string**|  | 
  **accountId** | **string**|  | 

### Return type

[**MeetingServiceInfoResource**](MeetingServiceInfoResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMeeting**
> MeetingResponseResource UpdateMeeting(ctx, meetingId, extensionId, accountId, optional)
Update Meeting

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **meetingId** | **string**|  | 
  **extensionId** | **string**|  | 
  **accountId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **meetingId** | **string**|  | 
 **extensionId** | **string**|  | 
 **accountId** | **string**|  | 
 **body** | [**MeetingRequestResource**](MeetingRequestResource.md)|  | 

### Return type

[**MeetingResponseResource**](MeetingResponseResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

