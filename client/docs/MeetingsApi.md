# \MeetingsApi

All URIs are relative to *https://platform.devtest.ringcentral.com/*

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
> MeetingResponseResource CreateMeeting(ctx, accountId, extensionId, meetingRequestResource)
Create Meetings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 
  **extensionId** | **string**|  | 
  **meetingRequestResource** | [**MeetingRequestResource**](MeetingRequestResource.md)|  | 

### Return type

[**MeetingResponseResource**](MeetingResponseResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

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
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **meetingId** | **string**|  | 
  **extensionId** | **string**|  | 
  **accountId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndMeeting**
> EndMeeting(ctx, meetingId, extensionId, accountId)
End Meeting

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **meetingId** | **string**|  | 
  **extensionId** | **string**|  | 
  **accountId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLiveMeetings**
> MeetingsResource GetLiveMeetings(ctx, accountId, extensionId)
Get Scheduled Meetings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 
  **extensionId** | **string**|  | 

### Return type

[**MeetingsResource**](MeetingsResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMeetingDetails**
> MeetingResponseResource GetMeetingDetails(ctx, accountId, extensionId, meetingId)
Get Meeting Info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 
  **extensionId** | **string**|  | 
  **meetingId** | **string**|  | 

### Return type

[**MeetingResponseResource**](MeetingResponseResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMeetingServiceInfo**
> MeetingServiceInfoResource GetMeetingServiceInfo(ctx, extensionId, accountId)
Get Meeting Service Info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **extensionId** | **string**|  | 
  **accountId** | **string**|  | 

### Return type

[**MeetingServiceInfoResource**](MeetingServiceInfoResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMeeting**
> MeetingResponseResource UpdateMeeting(ctx, accountId, extensionId, meetingId, meetingRequestResource)
Update Meeting

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 
  **extensionId** | **string**|  | 
  **meetingId** | **string**|  | 
  **meetingRequestResource** | [**MeetingRequestResource**](MeetingRequestResource.md)|  | 

### Return type

[**MeetingResponseResource**](MeetingResponseResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

