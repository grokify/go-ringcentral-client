# \APIInfoApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAPIVersion**](APIInfoApi.md#GetAPIVersion) | **Get** /restapi | Get API Versions
[**GetVersionInfo**](APIInfoApi.md#GetVersionInfo) | **Get** /restapi/{apiVersion} | Get Version Info
[**LoadAPIStatus**](APIInfoApi.md#LoadAPIStatus) | **Get** /restapi/v1.0/status | Get Status


# **GetAPIVersion**
> GetVersionsResponse GetAPIVersion(ctx, )
Get API Versions

<p style='font-style:italic;'>Since 1.0.0</p><p>Returns current API version(s) and server info.</p><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetVersionsResponse**](GetVersionsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVersionInfo**
> GetVersionResponse GetVersionInfo(ctx, apiVersion)
Get Version Info

<p style='font-style:italic;'>Since 1.0.0</p><p>Returns current API version info by apiVersion.</p><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiVersion** | **string**| API version to be requested, for example &#39;v1.0&#39; | 

### Return type

[**GetVersionResponse**](GetVersionResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadAPIStatus**
> LoadAPIStatus(ctx, )
Get Status

<p style='font-style:italic;'>Since 1.0.27 (Release 8.3)</p><p>Returns the API status; status '200' means the API is working fine, and '503' means it is temporary unavailable.</p><h4>API Group</h4><p>Light</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

