# \RingOutApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelRingOutCall**](RingOutApi.md#CancelRingOutCall) | **Delete** /restapi/v1.0/account/{accountId}/extension/{extensionId}/ringout/{ringoutId} | Cancel RingOut Call
[**CancelRingOutCallNew**](RingOutApi.md#CancelRingOutCallNew) | **Delete** /restapi/v1.0/account/{accountId}/extension/{extensionId}/ring-out/{ringoutId} | Cancel RingOut Call
[**CreateDirectRingOut**](RingOutApi.md#CreateDirectRingOut) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/ringout/direct | 
[**CreateDirectRingOutNew**](RingOutApi.md#CreateDirectRingOutNew) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/direct-ring-out | Make Direct RingOut Call
[**GetRingOutCallStatus**](RingOutApi.md#GetRingOutCallStatus) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/ringout/{ringoutId} | Get Status of RingOut Call
[**GetRingOutCallStatusNew**](RingOutApi.md#GetRingOutCallStatusNew) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/ring-out/{ringoutId} | Get Status of RingOut Call
[**MakeRingOutCall**](RingOutApi.md#MakeRingOutCall) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/ringout | Make RingOut Call
[**MakeRingOutCallNew**](RingOutApi.md#MakeRingOutCallNew) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/ring-out | Make RingOut Call


# **CancelRingOutCall**
> CancelRingOutCall(ctx, accountId, extensionId, ringoutId)
Cancel RingOut Call

<p style='font-style:italic;'></p><p>Cancels the 2-leg RingOut call.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>RingOut</td><td>Performing two-legged ring-out phone calls</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
  **ringoutId** | **int32**| Internal identifier of a RingOut call | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelRingOutCallNew**
> CancelRingOutCallNew(ctx, accountId, extensionId, ringoutId)
Cancel RingOut Call

<p style='font-style:italic;'></p><p>Cancels the 2-leg RingOut call.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>RingOut</td><td>Performing two-legged ring-out phone calls</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
  **ringoutId** | **int32**| Internal identifier of a RingOut call | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDirectRingOut**
> CreateDirectRingOut(ctx, extensionId, accountId, optional)


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
 **body** | [**RingOutResource**](RingOutResource.md)|  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDirectRingOutNew**
> CreateDirectRingOutNew(ctx, extensionId, accountId, optional)
Make Direct RingOut Call

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
 **body** | [**RingOutResource**](RingOutResource.md)|  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRingOutCallStatus**
> GetRingOutStatusResponse GetRingOutCallStatus(ctx, accountId, extensionId, ringoutId)
Get Status of RingOut Call

<p style='font-style:italic;'>Since 1.0.7 (Release 5.16)</p><p>Returns the status of a 2-leg RingOut call.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>RingOut</td><td>Performing two-legged ring-out phone calls</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
  **ringoutId** | **int32**| Internal identifier of a RingOut call | 

### Return type

[**GetRingOutStatusResponse**](GetRingOutStatusResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRingOutCallStatusNew**
> GetRingOutStatusResponse GetRingOutCallStatusNew(ctx, accountId, extensionId, ringoutId)
Get Status of RingOut Call

<p style='font-style:italic;'>Since 1.0.7 (Release 5.16)</p><p>Returns the status of a 2-leg RingOut call.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>RingOut</td><td>Performing two-legged ring-out phone calls</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
  **ringoutId** | **int32**| Internal identifier of a RingOut call | 

### Return type

[**GetRingOutStatusResponse**](GetRingOutStatusResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MakeRingOutCall**
> GetRingOutStatusResponse MakeRingOutCall(ctx, accountId, extensionId, body)
Make RingOut Call

<p style='font-style:italic;'>Since 1.0.7 (Release 5.16)</p><p>Makes a 2-leg RingOut call.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>RingOut</td><td>Performing two-legged ring-out phone calls</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
  **body** | [**MakeRingOutRequest**](MakeRingOutRequest.md)| JSON body | 

### Return type

[**GetRingOutStatusResponse**](GetRingOutStatusResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MakeRingOutCallNew**
> GetRingOutStatusResponse MakeRingOutCallNew(ctx, accountId, extensionId, body)
Make RingOut Call

<p style='font-style:italic;'>Since 1.0.7 (Release 5.16)</p><p>Makes a 2-leg RingOut call.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>RingOut</td><td>Performing two-legged ring-out phone calls</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
  **body** | [**MakeRingOutRequest**](MakeRingOutRequest.md)| JSON body | 

### Return type

[**GetRingOutStatusResponse**](GetRingOutStatusResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

