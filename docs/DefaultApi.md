# \DefaultApi

All URIs are relative to *https://api.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestapiV10AccountAccountIdIvrMenusPost**](DefaultApi.md#RestapiV10AccountAccountIdIvrMenusPost) | **Post** /restapi/v1.0/account/{accountId}/ivr-menus | Create IVR Menu


# **RestapiV10AccountAccountIdIvrMenusPost**
> IvrMenuInfo RestapiV10AccountAccountIdIvrMenusPost(ctx, accountId, body)
Create IVR Menu

<p style='font-style:italic;'>Since 1.0.32 (Release 9.3)</p><p>Creates a company IVR menu</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditAccounts</td><td>Viewing and updating user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **body** | [**IvrMenuInfo**](IvrMenuInfo.md)| JSON body | 

### Return type

[**IvrMenuInfo**](IVRMenuInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

