# \OrderingApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestapiV10AccountAccountIdOrderOrderIdGet**](OrderingApi.md#RestapiV10AccountAccountIdOrderOrderIdGet) | **Get** /restapi/v1.0/account/{accountId}/order/{orderId} | 
[**RestapiV10AccountAccountIdOrderPost**](OrderingApi.md#RestapiV10AccountAccountIdOrderPost) | **Post** /restapi/v1.0/account/{accountId}/order | 


# **RestapiV10AccountAccountIdOrderOrderIdGet**
> InlineResponseDefault25 RestapiV10AccountAccountIdOrderOrderIdGet($accountId, $orderId)



Get Order by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **orderId** | **string**| Internal identifier of an order | 

### Return type

[**InlineResponseDefault25**](inline_response_default_25.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdOrderPost**
> InlineResponseDefault24 RestapiV10AccountAccountIdOrderPost($accountId, $body)



Create New Order


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **body** | [**Body16**](Body16.md)|  | [optional] 

### Return type

[**InlineResponseDefault24**](inline_response_default_24.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

