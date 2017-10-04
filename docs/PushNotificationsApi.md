# \PushNotificationsApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestapiV10SubscriptionGet**](PushNotificationsApi.md#RestapiV10SubscriptionGet) | **Get** /restapi/v1.0/subscription | 
[**RestapiV10SubscriptionPost**](PushNotificationsApi.md#RestapiV10SubscriptionPost) | **Post** /restapi/v1.0/subscription | 
[**RestapiV10SubscriptionSubscriptionIdDelete**](PushNotificationsApi.md#RestapiV10SubscriptionSubscriptionIdDelete) | **Delete** /restapi/v1.0/subscription/{subscriptionId} | 
[**RestapiV10SubscriptionSubscriptionIdGet**](PushNotificationsApi.md#RestapiV10SubscriptionSubscriptionIdGet) | **Get** /restapi/v1.0/subscription/{subscriptionId} | 
[**RestapiV10SubscriptionSubscriptionIdPut**](PushNotificationsApi.md#RestapiV10SubscriptionSubscriptionIdPut) | **Put** /restapi/v1.0/subscription/{subscriptionId} | 


# **RestapiV10SubscriptionGet**
> InlineResponseDefault37 RestapiV10SubscriptionGet()



Get Subscription List


### Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponseDefault37**](inline_response_default_37.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10SubscriptionPost**
> SubscriptionInfo RestapiV10SubscriptionPost($body)



Create New Subscription


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Body20**](Body20.md)|  | [optional] 

### Return type

[**SubscriptionInfo**](SubscriptionInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10SubscriptionSubscriptionIdDelete**
> RestapiV10SubscriptionSubscriptionIdDelete($subscriptionId)



Cancel Subscription by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionId** | **string**| Internal identifier of a subscription | 

### Return type

void (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10SubscriptionSubscriptionIdGet**
> SubscriptionInfo RestapiV10SubscriptionSubscriptionIdGet($subscriptionId)



Get Subscription by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionId** | **string**| Internal identifier of a subscription | 

### Return type

[**SubscriptionInfo**](SubscriptionInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10SubscriptionSubscriptionIdPut**
> SubscriptionInfo RestapiV10SubscriptionSubscriptionIdPut($subscriptionId, $body)



Update/Renew Subscription by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionId** | **string**| Internal identifier of a subscription | 
 **body** | [**Body21**](Body21.md)|  | [optional] 

### Return type

[**SubscriptionInfo**](SubscriptionInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

