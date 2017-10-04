# \MeetingsApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestapiV10AccountAccountIdExtensionExtensionIdMeetingGet**](MeetingsApi.md#RestapiV10AccountAccountIdExtensionExtensionIdMeetingGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/meeting | 
[**RestapiV10AccountAccountIdExtensionExtensionIdMeetingMeetingIdDelete**](MeetingsApi.md#RestapiV10AccountAccountIdExtensionExtensionIdMeetingMeetingIdDelete) | **Delete** /restapi/v1.0/account/{accountId}/extension/{extensionId}/meeting/{meetingId} | 
[**RestapiV10AccountAccountIdExtensionExtensionIdMeetingMeetingIdEndPost**](MeetingsApi.md#RestapiV10AccountAccountIdExtensionExtensionIdMeetingMeetingIdEndPost) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/meeting/{meetingId}/end | 
[**RestapiV10AccountAccountIdExtensionExtensionIdMeetingMeetingIdGet**](MeetingsApi.md#RestapiV10AccountAccountIdExtensionExtensionIdMeetingMeetingIdGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/meeting/{meetingId} | 
[**RestapiV10AccountAccountIdExtensionExtensionIdMeetingMeetingIdPut**](MeetingsApi.md#RestapiV10AccountAccountIdExtensionExtensionIdMeetingMeetingIdPut) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/meeting/{meetingId} | 
[**RestapiV10AccountAccountIdExtensionExtensionIdMeetingPost**](MeetingsApi.md#RestapiV10AccountAccountIdExtensionExtensionIdMeetingPost) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/meeting | 
[**RestapiV10AccountAccountIdExtensionExtensionIdMeetingServiceInfoGet**](MeetingsApi.md#RestapiV10AccountAccountIdExtensionExtensionIdMeetingServiceInfoGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/meeting/service-info | 


# **RestapiV10AccountAccountIdExtensionExtensionIdMeetingGet**
> InlineResponseDefault20 RestapiV10AccountAccountIdExtensionExtensionIdMeetingGet($accountId, $extensionId)



Get Meetings List


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 

### Return type

[**InlineResponseDefault20**](inline_response_default_20.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdMeetingMeetingIdDelete**
> RestapiV10AccountAccountIdExtensionExtensionIdMeetingMeetingIdDelete($accountId, $extensionId, $meetingId)



Delete Meeting


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **meetingId** | **string**| Internal identifier of a meeting | 

### Return type

void (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdMeetingMeetingIdEndPost**
> RestapiV10AccountAccountIdExtensionExtensionIdMeetingMeetingIdEndPost($accountId, $extensionId, $meetingId)



End Current Meeting


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **meetingId** | **string**| Internal identifier of a meeting | 

### Return type

void (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdMeetingMeetingIdGet**
> MeetingInfo RestapiV10AccountAccountIdExtensionExtensionIdMeetingMeetingIdGet($accountId, $extensionId, $meetingId)



Get Meeting


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **meetingId** | **string**| Internal identifier of a meeting | 

### Return type

[**MeetingInfo**](MeetingInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdMeetingMeetingIdPut**
> MeetingInfo RestapiV10AccountAccountIdExtensionExtensionIdMeetingMeetingIdPut($accountId, $extensionId, $meetingId, $body)



Update Meeting


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **meetingId** | **string**| Internal identifier of a meeting | 
 **body** | [**Body12**](Body12.md)|  | [optional] 

### Return type

[**MeetingInfo**](MeetingInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdMeetingPost**
> MeetingInfo RestapiV10AccountAccountIdExtensionExtensionIdMeetingPost($accountId, $extensionId, $body)



Create Meeting


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **body** | [**Body11**](Body11.md)|  | [optional] 

### Return type

[**MeetingInfo**](MeetingInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdMeetingServiceInfoGet**
> MeetingServiceInfo RestapiV10AccountAccountIdExtensionExtensionIdMeetingServiceInfoGet($accountId, $extensionId)



Get Meeting Service Info


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 

### Return type

[**MeetingServiceInfo**](MeetingServiceInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

