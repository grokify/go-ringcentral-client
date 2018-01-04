# \ApplicationSettingsApi

All URIs are relative to *https://api.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSipRegistration**](ApplicationSettingsApi.md#CreateSipRegistration) | **Post** /restapi/v1.0/client-info/sip-provision | Register SIP Device
[**DeleteCustomData**](ApplicationSettingsApi.md#DeleteCustomData) | **Delete** /restapi/v1.0/client-info/custom-data/{key} | Delete Custom Data
[**GetAttachment**](ApplicationSettingsApi.md#GetAttachment) | **Get** /restapi/v1.0/client-info/custom-data/{key}/content | Get Custom Data Attachment
[**GetCustomData**](ApplicationSettingsApi.md#GetCustomData) | **Get** /restapi/v1.0/client-info/custom-data/{key} | Get Custom Data
[**GetInProductMessage**](ApplicationSettingsApi.md#GetInProductMessage) | **Get** /restapi/v1.0/client-info/banners | Get In Product Message List
[**LoadClientInfo**](ApplicationSettingsApi.md#LoadClientInfo) | **Get** /restapi/v1.0/client-info | Get App Settings
[**LoadPhoneData**](ApplicationSettingsApi.md#LoadPhoneData) | **Get** /restapi/v1.0/number-parser/phonedata.xml | Get Phonedata
[**RestapiV10ClientInfoSpecialNumberRuleGet**](ApplicationSettingsApi.md#RestapiV10ClientInfoSpecialNumberRuleGet) | **Get** /restapi/v1.0/client-info/special-number-rule | Get Special Numbers Calling Limitation
[**UpdateCustomData**](ApplicationSettingsApi.md#UpdateCustomData) | **Put** /restapi/v1.0/client-info/custom-data/{key} | Create/Update Custom Data


# **CreateSipRegistration**
> CreateSipRegistrationResponse CreateSipRegistration(ctx, body)
Register SIP Device

<p style='font-style:italic;'>Since 1.0.16 (Release 7.1)</p><p>Creates SIP registration of a device/application (WebPhone, Mobile, softphone)</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>VoipCalling</td><td>Registering as VoIP device and making VoIP calls</td></tr></tbody></table><h4>API Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**CreateSipRegistrationRequest**](CreateSipRegistrationRequest.md)| JSON body | 

### Return type

[**CreateSipRegistrationResponse**](CreateSipRegistrationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCustomData**
> DeleteCustomData(ctx, key)
Delete Custom Data

<p style='font-style:italic;'>Since 1.0.14 (Release 6.6)</p>Deletes custom data by its ID.<p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditCustomData</td><td>Viewing and updating client custom data (key-value)</td></tr></tbody></table><h4>API Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **key** | **int32**| Custom data access key. The number of unique custom data keys is limited to 100 keys per extension, summarized for all the applications. For example, if you have created 50 custom data keys under the Android mobile client application for the particular extension, then logged in the iOS application and created another 50 keys, the web client application won&#39;t be allowed to create any custom data key for that extension | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAttachment**
> GetAttachment(ctx, key, optional)
Get Custom Data Attachment

<p style='font-style:italic;'>Since 1.0.14 (Release 6.6)</p>Returns custom data attachment by ID.<p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditCustomData</td><td>Viewing and updating client custom data (key-value)</td></tr></tbody></table><h4>API Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **key** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string**|  | 
 **range_** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomData**
> CustomDataResource GetCustomData(ctx, key)
Get Custom Data

<p style='font-style:italic;'>Since 1.0.14 (Release 6.6)</p>Returns custom data of a logged-in extension.<p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditCustomData</td><td>Viewing and updating client custom data (key-value)</td></tr></tbody></table><h4>API Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **key** | **int32**| Custom data access key. The number of unique custom data keys is limited to 100 keys per extension, summarized for all the applications. For example, if you have created 50 custom data keys under the Android mobile client application for the particular extension, then logged in the iOS application and created another 50 keys, the web client application won&#39;t be allowed to create any custom data key for that extension | 

### Return type

[**CustomDataResource**](CustomDataResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInProductMessage**
> InProductMessages GetInProductMessage(ctx, bannerType)
Get In Product Message List

<p style='font-style:italic;'>Since 1.0.28 (Release 8.4)</p><p>Returns list of in-product messages: learderboards, announcements, etc. Please note: Banners are set on Marketing/Messages tab in Admin Web.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadClientInfo</td><td>Viewing of client application registered attributes and additional helper information (external URLs, hints, etc.)</td></tr></tbody></table><h4>API Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bannerType** | **string**|  | [default to LeaderBoard]

### Return type

[**InProductMessages**](InProductMessages.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadClientInfo**
> ClientApiResponse LoadClientInfo(ctx, )
Get App Settings

<p style='font-style:italic;'>Since 1.0.11 (Release 6.3)</p><p>Returns client application data: general info, specific provisioning parameters, hints, etc.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadClientInfo</td><td>Viewing of client application registered attributes and additional helper information (external URLs, hints, etc.)</td></tr></tbody></table>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClientApiResponse**](ClientApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadPhoneData**
> LoadPhoneData(ctx, )
Get Phonedata

<p style='font-style:italic;'>Since 1.0.12 (Release 6.4)</p><p>Returns configuration settings for phone number parser in a phonedata.xml file.</p><h4>Usage Plan Group</h4><p>Medium</p>

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

# **RestapiV10ClientInfoSpecialNumberRuleGet**
> NavigationInfo RestapiV10ClientInfoSpecialNumberRuleGet(ctx, optional)
Get Special Numbers Calling Limitation

<p style='font-style:italic;'>Since 1.0.12 (Release 6.4)</p><p>Returns available special numbers and rules of processing, depending on account brand and application type (mobile/softphone/web/other). A special number rule for each number includes limitations on the options: VoIP, RingOut, outgoing SMS and Fax. If the option is disallowed, the server returns the reason code and text description.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadClientInfo</td><td>Viewing of client application registered attributes and additional helper information (external URLs, hints, etc.)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryId** | **int32**| Internal identifier of a country. If not specified, the response is returned for the brand country | 

### Return type

[**NavigationInfo**](NavigationInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomData**
> CustomDataApiResponse UpdateCustomData(ctx, key, body)
Create/Update Custom Data

<p style='font-style:italic;'>Since 1.0.14 (Release 6.6)</p><p>Creates or updates custom data for the extension that is currently logged in.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditCustomData</td><td>Viewing and updating client custom data (key-value)</td></tr></tbody></table>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **key** | **int32**| Custom data access key. The number of unique custom data keys is limited to 100 keys per extension, summarized for all the applications. For example, if you have created 50 custom data keys under the Android mobile client application for the particular extension, then logged in the iOS application and created another 50 keys, the web client application won&#39;t be allowed to create any custom data key for that extension | 
  **body** | [**CustomDataRequest**](CustomDataRequest.md)| JSON body | 

### Return type

[**CustomDataApiResponse**](CustomDataApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

