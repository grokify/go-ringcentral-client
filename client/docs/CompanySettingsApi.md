# \CompanySettingsApi

All URIs are relative to *https://api.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAccount**](CompanySettingsApi.md#DeleteAccount) | **Delete** /restapi/v1.0/account/{accountId} | Delete Account
[**ListDialingPlans**](CompanySettingsApi.md#ListDialingPlans) | **Get** /restapi/v1.0/account/{accountId}/dialing-plan | Get Available Dialing Plans
[**LoadAccount**](CompanySettingsApi.md#LoadAccount) | **Get** /restapi/v1.0/account/{accountId} | Get Company Info
[**LoadAccountBusinessAddress**](CompanySettingsApi.md#LoadAccountBusinessAddress) | **Get** /restapi/v1.0/account/{accountId}/business-address | Get Company Business Address
[**LoadServiceInfo**](CompanySettingsApi.md#LoadServiceInfo) | **Get** /restapi/v1.0/account/{accountId}/service-info | Get Account Service Info
[**ModifyAccount**](CompanySettingsApi.md#ModifyAccount) | **Put** /restapi/v1.0/account/{accountId} | Update Company
[**UpdateAccountBusinessAddress**](CompanySettingsApi.md#UpdateAccountBusinessAddress) | **Put** /restapi/v1.0/account/{accountId}/business-address | Update Company Business Address
[**UpdateServiceInfo**](CompanySettingsApi.md#UpdateServiceInfo) | **Put** /restapi/v1.0/account/{accountId}/service-info | Update Account Service Info [Beta]


# **DeleteAccount**
> DeleteAccount(ctx, accountId)
Delete Account

<p style='font-style:italic;'>Since 1.0.0</p><p>Updates account information.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditAccounts</td><td>Viewing and updating user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDialingPlans**
> GetDialingPlanInfo ListDialingPlans(ctx, accountId, optional)
Get Available Dialing Plans

<p style='font-style:italic;'></p><p>Returns list of countries which can be selected for a dialing plan (to call short numbers and special services).</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **page** | **int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | 
 **perPage** | **int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | 

### Return type

[**GetDialingPlanInfo**](GetDialingPlanInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadAccount**
> GetAccountInfoResponse LoadAccount(ctx, accountId)
Get Company Info

<p style='font-style:italic;'>Since 1.0.0</p><p>Returns basic information about a particular RingCentral customer account.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 

### Return type

[**GetAccountInfoResponse**](GetAccountInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadAccountBusinessAddress**
> GetAccountInfoResponse LoadAccountBusinessAddress(ctx, accountId)
Get Company Business Address

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 

### Return type

[**GetAccountInfoResponse**](GetAccountInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadServiceInfo**
> GetServiceInfoResponse LoadServiceInfo(ctx, accountId)
Get Account Service Info

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 

### Return type

[**GetServiceInfoResponse**](GetServiceInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyAccount**
> GetAccountInfoResponse ModifyAccount(ctx, accountId, body)
Update Company

<p style='font-style:italic;'>Since 1.0.0</p><p>Updates account information.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditAccounts</td><td>Viewing and updating user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **body** | [**UpdateAccountRequest**](UpdateAccountRequest.md)| JSON body | 

### Return type

[**GetAccountInfoResponse**](GetAccountInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccountBusinessAddress**
> AccountBusinessAddressResource UpdateAccountBusinessAddress(ctx, accountId, body)
Update Company Business Address

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditAccounts</td><td>Viewing and updating user account info (including name, business name, address and phone number/account number)</td></tr><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **body** | [**ModifyAccountBusinessAddressRequest**](ModifyAccountBusinessAddressRequest.md)| JSON body | 

### Return type

[**AccountBusinessAddressResource**](AccountBusinessAddressResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceInfo**
> AccountServiceInfo UpdateServiceInfo(ctx, accountId, body)
Update Account Service Info [Beta]

<p style='font-style:italic;'>Since 1.0.29 (Release 9.0)</p><p>Updates limitations for a particular RingCentral customer account.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>Accounts</td><td>Managing accounts: creating new accounts, viewing and updating account information, deleting existing accounts</td></tr><tr><td class='code'>EditAccounts</td><td>Viewing and updating user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **body** | [**ServiceInfoRequest**](ServiceInfoRequest.md)| JSON body | 

### Return type

[**AccountServiceInfo**](AccountServiceInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

