# \PushNotificationsApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubscription**](PushNotificationsApi.md#CreateSubscription) | **Post** /restapi/v1.0/subscription | Create Subscription
[**DeleteSubscription**](PushNotificationsApi.md#DeleteSubscription) | **Delete** /restapi/v1.0/subscription/{subscriptionId} | Cancel Subscription
[**GetSubscriptions**](PushNotificationsApi.md#GetSubscriptions) | **Get** /restapi/v1.0/subscription | Get Subscriptions
[**LoadSubscription**](PushNotificationsApi.md#LoadSubscription) | **Get** /restapi/v1.0/subscription/{subscriptionId} | Get Subscription
[**RenewSubscription**](PushNotificationsApi.md#RenewSubscription) | **Post** /restapi/v1.0/subscription/{subscriptionId}/renew | Renew Subscription
[**UpdateSubscription**](PushNotificationsApi.md#UpdateSubscription) | **Put** /restapi/v1.0/subscription/{subscriptionId} | Renew Subscription / Update Event Filters


# **CreateSubscription**
> SubscriptionInfo CreateSubscription(ctx, createSubscriptionRequest)
Create Subscription

<p style='font-style:italic;'>Since 1.0.6 (Release 5.15)</p><p>Creates a new subscription.</p><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createSubscriptionRequest** | [**CreateSubscriptionRequest**](CreateSubscriptionRequest.md)| JSON body | 

### Return type

[**SubscriptionInfo**](SubscriptionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSubscription**
> DeleteSubscription(ctx, subscriptionId)
Cancel Subscription

<p style='font-style:italic;'></p><p>Cancels the existent subscription.</p><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **string**| Internal identifier of a subscription | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubscriptions**
> RecordsCollectionResourceSubscriptionResponse GetSubscriptions(ctx, )
Get Subscriptions

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RecordsCollectionResourceSubscriptionResponse**](RecordsCollectionResourceSubscriptionResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadSubscription**
> SubscriptionInfo LoadSubscription(ctx, subscriptionId)
Get Subscription

<p style='font-style:italic;'>Since 1.0.6 (Release 5.15)</p><p>Returns the requested subscription.</p><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **int32**| Internal identifier of a subscription | 

### Return type

[**SubscriptionInfo**](SubscriptionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenewSubscription**
> SubscriptionInfo RenewSubscription(ctx, subscriptionId)
Renew Subscription

<p style='font-style:italic;'>Since 1.0.26 (Release 8.12)</p><p>Renews an existent subscription by ID by posting request with an empty body..</p><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **string**|  | 

### Return type

[**SubscriptionInfo**](SubscriptionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSubscription**
> SubscriptionInfo UpdateSubscription(ctx, subscriptionId, modifySubscriptionRequest, optional)
Renew Subscription / Update Event Filters

<p style='font-style:italic;'>Since 1.0.6 (Release 5.15)</p><p>Renews the existent subscription if the request body is empty. If event filters are specified, calling this method modifies the event filters for the existing subscription. The client application can extend or narrow the events for which it receives notifications in the frame of one subscription.</p><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **string**| Internal identifier of a subscription | 
  **modifySubscriptionRequest** | [**ModifySubscriptionRequest**](ModifySubscriptionRequest.md)| JSON body | 
 **optional** | ***UpdateSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateSubscriptionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aggregated** | **optional.Bool**| If &#39;True&#39; then aggregated presence status is returned in a notification payload | 

### Return type

[**SubscriptionInfo**](SubscriptionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

