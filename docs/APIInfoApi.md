# \APIInfoApi

All URIs are relative to *https://api.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoadAPIStatus**](APIInfoApi.md#LoadAPIStatus) | **Get** /restapi/v1.0/status | Get Status
[**RestapiApiVersionGet**](APIInfoApi.md#RestapiApiVersionGet) | **Get** /restapi/{apiVersion} | Get Version Info
[**RestapiGet**](APIInfoApi.md#RestapiGet) | **Get** /restapi | Get API Versions


# **LoadAPIStatus**
> LoadAPIStatus(ctx, )
Get Status

<p style='font-style:italic;'>Since 1.0.27 (Release 8.3)</p><p>Returns the API status; status '200' means the API is working fine, and '503' means it is temporary unavailable.</p><h4>API Group</h4><p>Light</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiApiVersionGet**
> GetVersionResponse RestapiApiVersionGet(ctx, apiVersion)
Get Version Info

<p style='font-style:italic;'>Since 1.0.0</p><p>Returns current API version info by apiVersion.</p><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **apiVersion** | **string**| API version to be requested, for example &#39;v1.0&#39; | 

### Return type

[**GetVersionResponse**](GetVersionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiGet**
> GetVersionsResponse RestapiGet(ctx, )
Get API Versions

<p style='font-style:italic;'>Since 1.0.0</p><p>Returns current API version(s) and server info.</p><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetVersionsResponse**](GetVersionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

