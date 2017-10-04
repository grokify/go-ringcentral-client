# \AuthorizationApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestapiOauthAuthorizePost**](AuthorizationApi.md#RestapiOauthAuthorizePost) | **Post** /restapi/oauth/authorize | 
[**RestapiOauthRevokePost**](AuthorizationApi.md#RestapiOauthRevokePost) | **Post** /restapi/oauth/revoke | 
[**RestapiOauthTokenPost**](AuthorizationApi.md#RestapiOauthTokenPost) | **Post** /restapi/oauth/token | 


# **RestapiOauthAuthorizePost**
> InlineResponseDefault RestapiOauthAuthorizePost($body)



OAuth2 Authorize


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Body**](Body.md)|  | [optional] 

### Return type

[**InlineResponseDefault**](inline_response_default.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiOauthRevokePost**
> RestapiOauthRevokePost($body)



OAuth2 Revoke Token


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Body1**](Body1.md)|  | [optional] 

### Return type

void (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiOauthTokenPost**
> TokenInfo RestapiOauthTokenPost($body)



OAuth2 Get Token


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Body2**](Body2.md)|  | [optional] 

### Return type

[**TokenInfo**](TokenInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

