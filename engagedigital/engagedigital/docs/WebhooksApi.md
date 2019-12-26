# \WebhooksApi

All URIs are relative to *https://DOMAIN.api.engagement.dimelo.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhook**](WebhooksApi.md#CreateWebhook) | **Post** /webhooks | Creating a webhook
[**DeleteWebhook**](WebhooksApi.md#DeleteWebhook) | **Delete** /webhooks/{webhookId} | Deleting a webhook
[**GetAllWebhooks**](WebhooksApi.md#GetAllWebhooks) | **Get** /webhooks | Getting all webhooks
[**GetWebhook**](WebhooksApi.md#GetWebhook) | **Get** /webhooks/{webhookId} | Getting a webhook from its id
[**UpdateWebhook**](WebhooksApi.md#UpdateWebhook) | **Put** /webhooks/{webhookId} | Updating a webhook



## CreateWebhook

> Webhook CreateWebhook(ctx, accessToken, label, url, registeredEvents, optional)
Creating a webhook

This method creates a webhook. In case of success it renders the webhook, otherwise, it renders an error (422 HTTP code).  Authorization​: All users having the manage_api_access_tokens permission or all users having an api access token.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| Access token. | 
**label** | **string**| The label of the webhook. | 
**url** | **string**| The url of a webhook. This is used to determine the endpoint of your webhook, where | 
**registeredEvents** | [**[]string**](string.md)| An array containing all the events that your webhook wants to subscribe. | 
 **optional** | ***CreateWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateWebhookOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **active** | **optional.Bool**| true or false, this field is used to enable/disable a webhook. | 
 **stagingUse** | **optional.Bool**| true or false, this field is used to determine if a webhook will be run in staging | 
 **verifyToken** | **optional.String**| The token used in your webhook. | 
 **secret** | **optional.String**| The secret key that will be served as a ​&#x60;X-Dimelo-Secret​&#x60; header in every request. | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhook

> Webhook DeleteWebhook(ctx, webhookId, accessToken)
Deleting a webhook

This method destroys an existing webhook. It renders webhook itself. It renders a 404 if id is invalid.  Authorization​: All users having the manage_api_access_tokens permission or all users having an api access token belonging to the webhook you’re deleting.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string**|  | 
**accessToken** | **string**| Access token. | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllWebhooks

> GetAllWebhooksResponse GetAllWebhooks(ctx, accessToken, optional)
Getting all webhooks

This method renders webhooks ordered by active and staging_use (descending).  Authorization​: users having manage_api_access_tokens permission can see all webhooks / users don’t having the manage_api_access_tokens permission can see only the webhooks belonging to access token created by them.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| Access token. | 
 **optional** | ***GetAllWebhooksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllWebhooksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| The record index to start. Default value is 0. | 
 **limit** | **optional.Int32**| The max number of records to return. Default value is 30, max value is 150. | 

### Return type

[**GetAllWebhooksResponse**](GetAllWebhooksResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhook

> Webhook GetWebhook(ctx, webhookId, accessToken)
Getting a webhook from its id

This method renders a webhook from given id.  Authorization​: users having manage_api_access_tokens permission can see any webhook / users don’t having the `manage_api_access_tokens` permission can see only the webhook in case the webhook is associated to an access token created by them.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string**|  | 
**accessToken** | **string**| Access token. | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhook

> Webhook UpdateWebhook(ctx, webhookId, accessToken, optional)
Updating a webhook

This method updates an existing webhook from given attributes and renders it in case of success.  Authorization​: All users having the manage_api_access_tokens permission or all users having an api access token belonging to the webhook you’re updating.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string**|  | 
**accessToken** | **string**| Access token. | 
 **optional** | ***UpdateWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateWebhookOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **active** | **optional.Bool**| true or false, this field is used to enable/disable a webhook. | 
 **label** | **optional.String**| The label of the webhook. | 
 **stagingUse** | **optional.Bool**| true or false, this field is used to determine if a webhook will be run in staging | 
 **url** | **optional.String**| The url of a webhook. This is used to determine the endpoint of your webhook, where | 
 **verifyToken** | **optional.String**| The token used in your webhook. | 
 **secret** | **optional.String**| The secret key that will be served as a ​&#x60;X-Dimelo-Secret​&#x60; header in every request. | 
 **registeredEvents** | [**optional.Interface of []string**](string.md)| An array containing all the events that your webhook wants to subscribe. | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

