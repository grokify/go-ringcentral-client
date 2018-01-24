# \ReportingApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoadReportingSettings**](ReportingApi.md#LoadReportingSettings) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/reporting/settings | Get Reporting Settings
[**UpdateReportingSettings**](ReportingApi.md#UpdateReportingSettings) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/reporting/settings | Update Reporting Settings


# **LoadReportingSettings**
> ReportingSettings LoadReportingSettings(ctx, accountId, extensionId)
Get Reporting Settings

<p style='font-style:italic;'>Since 1.0.17 (Release 7.2)</p><p>Returns user-defined configuration of CFA (Customer Facing Analytics) reports.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 

### Return type

[**ReportingSettings**](ReportingSettings.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateReportingSettings**
> ReportingSettings UpdateReportingSettings(ctx, accountId, extensionId, body)
Update Reporting Settings

<p style='font-style:italic;'></p><p>Sets user-defined configuration of CFA (Customer Facing Analytics) reports.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditAccounts</td><td>Viewing and updating user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
  **body** | [**SetReportingSettingsRequest**](SetReportingSettingsRequest.md)| JSON body | 

### Return type

[**ReportingSettings**](ReportingSettings.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

