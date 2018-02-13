# \AccountProvisioningApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkAssign**](AccountProvisioningApi.md#BulkAssign) | **Post** /restapi/v1.0/account/{accountId}/department/bulk-assign | Edit Call Queue Members
[**CreateAccount**](AccountProvisioningApi.md#CreateAccount) | **Post** /restapi/v1.0/account | Create Account
[**CreateDeviceOrder**](AccountProvisioningApi.md#CreateDeviceOrder) | **Post** /restapi/v1.0/account/{accountId}/order | Create Device Order
[**CreateLicensesOrder**](AccountProvisioningApi.md#CreateLicensesOrder) | **Post** /restapi/v1.0/account/{accountId}/licenses/bulk-purchase | Order Licenses
[**DeleteLicense**](AccountProvisioningApi.md#DeleteLicense) | **Delete** /restapi/v1.0/account/{accountId}/licenses/{licenseId} | Delete License
[**DeletePhoneNumberById**](AccountProvisioningApi.md#DeletePhoneNumberById) | **Delete** /restapi/v1.0/account/{accountId}/phone-number/{phoneNumberId} | Delete Phone Number
[**GetBrandInfo**](AccountProvisioningApi.md#GetBrandInfo) | **Get** /restapi/v1.0/dictionary/brand/{brandId} | Get Brand Info
[**GetDeviceModels**](AccountProvisioningApi.md#GetDeviceModels) | **Get** /restapi/v1.0/dictionary/device | Get Device Catalog
[**GetExtensionFreeNumbers**](AccountProvisioningApi.md#GetExtensionFreeNumbers) | **Get** /restapi/v1.0/account/{accountId}/extension/free-numbers | Get Free Extension Numbers
[**GetPagingOnlyGroupDevices**](AccountProvisioningApi.md#GetPagingOnlyGroupDevices) | **Get** /restapi/v1.0/account/{accountId}/paging-only-groups/{pagingOnlyGroupId}/devices | Get Paging Only Group Devices
[**GetPagingOnlyGroupUsers**](AccountProvisioningApi.md#GetPagingOnlyGroupUsers) | **Get** /restapi/v1.0/account/{accountId}/paging-only-groups/{pagingOnlyGroupId}/users | Get Paging Only Group Users
[**GetParkLocationUsers**](AccountProvisioningApi.md#GetParkLocationUsers) | **Get** /restapi/v1.0/account/{accountId}/park-locations/{parkLocationId}/users | Get Park Location Users
[**GetServicePlanInfo**](AccountProvisioningApi.md#GetServicePlanInfo) | **Get** /restapi/v1.0/dictionary/service-plan/{servicePlanId} | Get Service Plan
[**ListAccountDevices**](AccountProvisioningApi.md#ListAccountDevices) | **Get** /restapi/v1.0/account/{accountId}/device | Get Account Devices
[**ListAccountPhoneNumbers**](AccountProvisioningApi.md#ListAccountPhoneNumbers) | **Get** /restapi/v1.0/account/{accountId}/phone-number | Get All Company Phone Numbers
[**ListDepartmentMembers**](AccountProvisioningApi.md#ListDepartmentMembers) | **Get** /restapi/v1.0/account/{accountId}/department/{departmentId}/members | Get Department Member List
[**ListExtensionDevices**](AccountProvisioningApi.md#ListExtensionDevices) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/device | Get User Device Info
[**ListLicenseTypes**](AccountProvisioningApi.md#ListLicenseTypes) | **Get** /restapi/v1.0/dictionary/license-types | Get License Types
[**ListLicenses**](AccountProvisioningApi.md#ListLicenses) | **Get** /restapi/v1.0/account/{accountId}/licenses | Get License List
[**ListShippingOptions**](AccountProvisioningApi.md#ListShippingOptions) | **Get** /restapi/v1.0/dictionary/shipping-options | Get Shipping Options
[**LoadAccountDevice**](AccountProvisioningApi.md#LoadAccountDevice) | **Get** /restapi/v1.0/account/{accountId}/device/{deviceId} | Get Device Info
[**LoadAccountPhoneNumber**](AccountProvisioningApi.md#LoadAccountPhoneNumber) | **Get** /restapi/v1.0/account/{accountId}/phone-number/{phoneNumberId} | Get Phone Number
[**LoadDeviceOrder**](AccountProvisioningApi.md#LoadDeviceOrder) | **Get** /restapi/v1.0/account/{accountId}/order/{orderId} | Get Device Order by ID
[**LoadLicense**](AccountProvisioningApi.md#LoadLicense) | **Get** /restapi/v1.0/account/{accountId}/licenses/{licenseId} | Get License
[**LookupPhoneNumbers**](AccountProvisioningApi.md#LookupPhoneNumbers) | **Post** /restapi/v1.0/number-pool/lookup | Get Numbers
[**ParsePhoneNumber**](AccountProvisioningApi.md#ParsePhoneNumber) | **Post** /restapi/v1.0/number-parser/parse | Parse Phone Number
[**ProvisionPhoneNumber**](AccountProvisioningApi.md#ProvisionPhoneNumber) | **Post** /restapi/v1.0/account/{accountId}/phone-number | Provision Phone Numbers
[**ReassignPhoneNumberById**](AccountProvisioningApi.md#ReassignPhoneNumberById) | **Put** /restapi/v1.0/account/{accountId}/phone-number/{phoneNumberId} | Reassign Phone Number
[**ReservePhoneNumbers**](AccountProvisioningApi.md#ReservePhoneNumbers) | **Post** /restapi/v1.0/number-pool/reserve | Reserve/ Un-reserve Numbers
[**UpdateDevice**](AccountProvisioningApi.md#UpdateDevice) | **Put** /restapi/v1.0/account/{accountId}/device/{deviceId} | Update Device
[**UpdatePagingOnlyGroupUsersAndDevices**](AccountProvisioningApi.md#UpdatePagingOnlyGroupUsersAndDevices) | **Post** /restapi/v1.0/account/{accountId}/paging-only-groups/{pagingOnlyGroupId}/bulk-assign | Edit Paging Group Users &amp; Devices
[**UpdateParkLocationUsers**](AccountProvisioningApi.md#UpdateParkLocationUsers) | **Post** /restapi/v1.0/account/{accountId}/park-locations/{parkLocationId}/bulk-assign | Adds and/or removes park location users


# **BulkAssign**
> BulkAssign(ctx, accountId, optional)
Edit Call Queue Members

<p style='font-style:italic;'>Since 1.0.32 (Release 9.3)</p><p>Adds and/or removes multiple call queue members.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditExtensions</td><td>Viewing and updating my extension info (includes extension name, number, email and phone number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**|  | 
 **body** | [**DepartmentBulkAssignResource**](DepartmentBulkAssignResource.md)|  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAccount**
> GetAccountInfoResponse CreateAccount(ctx, body)
Create Account

<p style='font-style:italic;'>Since 1.0.10 (Release 6.2)</p><p>Creates the account in Initial state.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>Accounts</td><td>Managing accounts: creating new accounts, viewing and updating account information, deleting existing accounts</td></tr><tr><td class='code'>EditAccounts</td><td>Viewing and updating user account info (including name, business name, address and phone number/account number)</td></tr><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr><tr><td class='code'>EditExtensions</td><td>Viewing and updating user extension info (includes extension name, number, email and phone number, assigned phone numbers, devices and other extension settings)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**CreateAccountRequest**](CreateAccountRequest.md)| JSON body | 

### Return type

[**GetAccountInfoResponse**](GetAccountInfoResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDeviceOrder**
> DeviceOrderCreation CreateDeviceOrder(ctx, accountId, optional)
Create Device Order

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditAccounts</td><td>Viewing and updating user account info (including name, business name, address and phone number/account number)</td></tr><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

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
 **async** | **bool**| Specifies whether a client expects order details to be returned or not. If &#39;True&#39; the server returns a link for tracking order status without waiting for it to be actually created. If &#39;False&#39; the server returns order details once it is created - it may take some time depending on device count. The default value is &#39;False&#39; | [default to false]
 **body** | [**DeviceResource**](DeviceResource.md)|  | 

### Return type

[**DeviceOrderCreation**](DeviceOrderCreation.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLicensesOrder**
> CreateLicensesOrder(ctx, accountId, body)
Order Licenses

<p style='font-style:italic;'>Since 1.0.31 (Release 9.2)</p><p>Purchases licenses for add-on features: Rooms, Room Connector, Webinar, Live Reports, etc.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr><tr><td class='code'>EditExtensions</td><td>Viewing and updating user extension info (includes extension name, number, email and phone number, assigned phone numbers, devices and other extension settings)</td></tr><tr><td class='code'>EditAccounts</td><td>Viewing and updating user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **body** | [**OrderLicensesRequest**](OrderLicensesRequest.md)| JSON body | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLicense**
> DeleteLicense(ctx, accountId, licenseId)
Delete License

<p style='font-style:italic;'>Since 1.0.31 (Release 9.2)</p><p>Removes a license for a specific user. Please note: It is not allowed to remove assigned licenses (only Webinars and Large Meetings can be assigned).</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr><tr><td class='code'>EditExtensions</td><td>Viewing and updating user extension info (includes extension name, number, email and phone number, assigned phone numbers, devices and other extension settings)</td></tr><tr><td class='code'>EditAccounts</td><td>Viewing and updating user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **licenseId** | **string**| Internal identifier of a license | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePhoneNumberById**
> DeletePhoneNumberById(ctx, phoneNumberId, accountId)
Delete Phone Number

<p style='font-style:italic;'>Since 1.0.11 (Release 6.3) </p><p>Deletes a phone number belonging to a certain account/extension.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditAccounts</td><td>Viewing and updating user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **phoneNumberId** | **int64**|  | 
  **accountId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBrandInfo**
> BrandResource GetBrandInfo(ctx, brandId)
Get Brand Info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **brandId** | **string**|  | 

### Return type

[**BrandResource**](BrandResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceModels**
> DeviceModelExtsResource GetDeviceModels(ctx, )
Get Device Catalog

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DeviceModelExtsResource**](DeviceModelExtsResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExtensionFreeNumbers**
> BulkProvisionUnassignedExtensionsResponseResource GetExtensionFreeNumbers(ctx, accountId)
Get Free Extension Numbers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 

### Return type

[**BulkProvisionUnassignedExtensionsResponseResource**](BulkProvisionUnassignedExtensionsResponseResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagingOnlyGroupDevices**
> PagingOnlyGroupDevices GetPagingOnlyGroupDevices(ctx, accountId, pagingOnlyGroupId, optional)
Get Paging Only Group Devices

<p style='font-style:italic;'>Since 1.0.32 (Release 9.3)</p><p>Returns the list of paging devices assigned to this group.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **pagingOnlyGroupId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**|  | 
 **pagingOnlyGroupId** | **string**|  | 
 **page** | **int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | 
 **perPage** | **int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | 

### Return type

[**PagingOnlyGroupDevices**](PagingOnlyGroupDevices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagingOnlyGroupUsers**
> PagingOnlyGroupUsers GetPagingOnlyGroupUsers(ctx, accountId, pagingOnlyGroupId, optional)
Get Paging Only Group Users

<p style='font-style:italic;'>Since 1.0.32 (Release 9.3)</p><p>Returns the list of users allowed to page this group.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **pagingOnlyGroupId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**|  | 
 **pagingOnlyGroupId** | **string**|  | 
 **page** | **int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | 
 **perPage** | **int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | 

### Return type

[**PagingOnlyGroupUsers**](PagingOnlyGroupUsers.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetParkLocationUsers**
> ParkLocationResponse GetParkLocationUsers(ctx, accountId, parkLocationId, optional)
Get Park Location Users

<p style='font-style:italic;'>Since 1.0.32 (Release 9.3)</p><p>Returns the list of users allowed to park and unpark calls to/from the park location extension specified</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **parkLocationId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**|  | 
 **parkLocationId** | **string**|  | 
 **page** | **int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | 
 **perPage** | **int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | 

### Return type

[**ParkLocationResponse**](ParkLocationResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServicePlanInfo**
> ServicePlanResource GetServicePlanInfo(ctx, servicePlanId)
Get Service Plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **servicePlanId** | **string**|  | 

### Return type

[**ServicePlanResource**](ServicePlanResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountDevices**
> GetAccountDevicesResponse ListAccountDevices(ctx, accountId)
Get Account Devices

<p style='font-style:italic;'>Since 1.0.12 (Release 6.4)</p><p>Returns all the devices for a particular extension.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 

### Return type

[**GetAccountDevicesResponse**](GetAccountDevicesResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountPhoneNumbers**
> AccountPhoneNumbers ListAccountPhoneNumbers(ctx, accountId, optional)
Get All Company Phone Numbers

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

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
 **usageType** | [**[]string**](string.md)| Usage type of the phone number | 

### Return type

[**AccountPhoneNumbers**](AccountPhoneNumbers.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDepartmentMembers**
> DepartmentMemberList ListDepartmentMembers(ctx, accountId, departmentId, optional)
Get Department Member List

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **departmentId** | **int32**| Internal identifier of a Department extension (same as extensionId but only the ID of a department extension is valid) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **departmentId** | **int32**| Internal identifier of a Department extension (same as extensionId but only the ID of a department extension is valid) | 
 **page** | **int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | 
 **perPage** | **int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | 

### Return type

[**DepartmentMemberList**](DepartmentMemberList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListExtensionDevices**
> GetExtensionDevicesResponse ListExtensionDevices(ctx, accountId, extensionId, optional)
Get User Device Info

<p style='font-style:italic;'>Since 1.0.12 (Release 6.4)</p><p>Returns all the devices for extension(s) by extension ID(s). Batch request is supported, see Batch Requests for details.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **linePooling** | **string**| Pooling type of a device | 

### Return type

[**GetExtensionDevicesResponse**](GetExtensionDevicesResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLicenseTypes**
> LicenseTypes ListLicenseTypes(ctx, )
Get License Types

<p style='font-style:italic;'>Since 1.0.31 (Release 9.2)</p><p>Returns supported license types.</p><h4>API Group</h4><p>Light</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**LicenseTypes**](LicenseTypes.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLicenses**
> LicenseList ListLicenses(ctx, accountId, optional)
Get License List

<p style='font-style:italic;'>Since 1.0.31 (Release 9.2)</p><p>Returns list of licenses for a specific user.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Light</p>

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
 **page** | **int64**| Indicates the page number to retrieve. Only positive number values are allowed. The default value is &#39;1&#39; | 
 **perPage** | **int64**| Indicates the page size (number of items). If not specified, the value is &#39;25&#39; by default | 
 **typeId** | **int64**| Internal identifier of a license type. If not specified account licenses of all types are returned | 

### Return type

[**LicenseList**](LicenseList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListShippingOptions**
> DictionaryShippingOptions ListShippingOptions(ctx, optional)
Get Shipping Options

<p style='font-style:italic;'>Since 1.0.16 (Release 7.1)</p><p>Returns the list of device shipping options with their prices, according to brand, tier, number of ordered devices.</p><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **servicePlanId** | **int32**|  | 
 **brandId** | **int32**|  | 
 **quantity** | **int32**|  | 

### Return type

[**DictionaryShippingOptions**](DictionaryShippingOptions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadAccountDevice**
> GetDeviceInfoResponse LoadAccountDevice(ctx, accountId, deviceId)
Get Device Info

<p style='font-style:italic;'>Since 1.0.9 (Release 6.1)</p><p>Returns account device(s) by their ID(s).</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **deviceId** | **string**| Internal identifier of a device | 

### Return type

[**GetDeviceInfoResponse**](GetDeviceInfoResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadAccountPhoneNumber**
> PhoneNumberInfo LoadAccountPhoneNumber(ctx, accountId, phoneNumberId)
Get Phone Number

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **phoneNumberId** | **int32**| Internal identifier of a phone number | 

### Return type

[**PhoneNumberInfo**](PhoneNumberInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadDeviceOrder**
> DeviceOrder LoadDeviceOrder(ctx, accountId, orderId)
Get Device Order by ID

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **orderId** | **int32**| Internal identifier of an order | 

### Return type

[**DeviceOrder**](DeviceOrder.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadLicense**
> LicenseInfo LoadLicense(ctx, accountId, licenseId)
Get License

<p style='font-style:italic;'>Since 1.0.31 (Release 9.2)</p><p>Returns license information by its ID.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **licenseId** | **string**| Internal identifier of a license | 

### Return type

[**LicenseInfo**](LicenseInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LookupPhoneNumbers**
> PhoneNumbers LookupPhoneNumbers(ctx, optional)
Get Numbers

<p style='font-style:italic;'>Since 1.0.10 (Release 6.2)</p><p>Returns the required numbers filtered by criteria.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>NumberLookup</td><td>Looking-up and reserving available phone number</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **areaCode** | **int32**| Area code of the location | 
 **countryCode** | **string**| Two-letter country code, complying with the ISO standard | 
 **countryId** | **string**| Internal identifier of a country; &#39;1&#39;- the US; &#39;39&#39; - Canada; &#39;224&#39; - the UK. The default value is &#39;1&#39; | 
 **exclude** | **string**| A string of digits (one and more) that should not appear among the last four digits (line part) of the phone numbers that will be returned. It is possible to specify severalT?exclude parameters. If specified, it is taken into account in all returned phone numbers both in the phone numbers satisfying to parameters of lookup and in alternative phone numbers (in case when extendedSearch is specified) | 
 **extendedSearch** | **bool**| If the value is &#39;False&#39;, then the returned numbers exactly correspond to the specified NXX, NPA and LINE or countryCode, areaCode and numberPattern parameters. If the value is &#39;True&#39;, then the resulting numbers are ranked and returned with the rank attribute values (1-10). The default value is &#39;False&#39; | 
 **line** | **string**| LINE pattern for vanity or wildcard search. Digits, Latin characters and asterisks are allowed (usually 4 characters) | 
 **numberPattern** | **string**| Phone number pattern (for wildcard or vanity search). For NANP countries (US, Canada) is concatenation of nxx (the first three digits) and line. If the first three characters are specified as not digits (e.g. 5** or CAT) then parameter extendedSearch will be ignored. | 
 **nxx** | **string**| NXX pattern for vanity or wildcard search. Digits, Latin characters and asterisks are allowed (usually 3 characters) | 
 **npa** | **string**| Area code (mandatory). For example, 800, 844, 855, 866, 877, 888 for North America; and 647 for Canada | 
 **paymentType** | [**[]string**](string.md)| Payment type. Default is &#39;Local&#39; (it should correlate with the npa provided) | 
 **perPage** | **int32**| Indicates the page size (number of items). If not specified, the value is &#39;10&#39; by default | 
 **providerId** | **int32**| Internal identifier of a phone number provider. Supported if brand is specified. If specified, area code and nxx are optional | 
 **smsEnabled** | **bool**| Specifies if SMS activation is available for the number. If specified, it is taken into account in all returned phone numbers both in the phone numbers satisfying to parameters of lookup and in alternative phone numbers (in case when extendedSearch is specified). If not specified, the value is null. | 

### Return type

[**PhoneNumbers**](PhoneNumbers.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ParsePhoneNumber**
> ParsePhoneNumberResponse ParsePhoneNumber(ctx, body, optional)
Parse Phone Number

<p style='font-style:italic;'>Since 1.0.13 (Release 6.5)</p><p>Returns one or more parsed and/or formatted phone numbers that are passed as a string.</p><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**ParsePhoneNumberRequest**](ParsePhoneNumberRequest.md)| JSON body | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ParsePhoneNumberRequest**](ParsePhoneNumberRequest.md)| JSON body | 
 **homeCountry** | **string**| Internal identifier of a home country. The default value is ISO code (ISO 3166) of the user&#39;s home country or brand country, if the user is undefined | 
 **nationalAsPriority** | **bool**| The default value is &#39;False&#39;. If &#39;True&#39;, the numbers that are closer to the home country are given higher priority | 

### Return type

[**ParsePhoneNumberResponse**](ParsePhoneNumberResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProvisionPhoneNumber**
> PhoneNumberInfo ProvisionPhoneNumber(ctx, accountId, body)
Provision Phone Numbers

<p style='font-style:italic;'>Since 1.0.11 (Release 6.3)</p><p>Provisions a phone number.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditAccounts</td><td>Viewing and updating user account info (including name, business name, address and phone number/account number)</td></tr><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
  **body** | [**ProvisionPhoneNumberRequest**](ProvisionPhoneNumberRequest.md)| JSON body | 

### Return type

[**PhoneNumberInfo**](PhoneNumberInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReassignPhoneNumberById**
> PhoneNumberResource ReassignPhoneNumberById(ctx, phoneNumberId, accountId, optional)
Reassign Phone Number

<p style='font-style:italic;'>Since 1.0.11 (Release 6.3) </p><p>Reassigns the phone number belonging to a certain account/extension. This call reassigns a phone number, currently belonging to some other extension or company (Auto-Receptionist). Please note: numbers with certain usage types 'MainCompanyNumber', 'AdditionalCompanyNumber' and 'CompanyFaxNumber' cannot be reassigned</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditAccounts</td><td>Viewing and updating user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **phoneNumberId** | **int64**|  | 
  **accountId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phoneNumberId** | **int64**|  | 
 **accountId** | **string**|  | 
 **body** | [**ReassignPhoneNumberResource**](ReassignPhoneNumberResource.md)|  | 

### Return type

[**PhoneNumberResource**](PhoneNumberResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReservePhoneNumbers**
> ReservePhoneNumbersRecordsResource ReservePhoneNumbers(ctx, body)
Reserve/ Un-reserve Numbers

<p style='font-style:italic;'>Since 1.0.10 (Release 6.2)</p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>NumberLookup</td><td>Looking-up and reserving available phone number</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**NumberPoolPhoneNumberRequest**](NumberPoolPhoneNumberRequest.md)| JSON body | 

### Return type

[**ReservePhoneNumbersRecordsResource**](ReservePhoneNumbersRecordsResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDevice**
> DeviceResource UpdateDevice(ctx, accountId, deviceId, body)
Update Device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **deviceId** | **string**|  | 
  **body** | [**AccountDeviceUpdate**](AccountDeviceUpdate.md)|  | 

### Return type

[**DeviceResource**](DeviceResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePagingOnlyGroupUsersAndDevices**
> UpdatePagingOnlyGroupUsersAndDevices(ctx, accountId, pagingOnlyGroupId, optional)
Edit Paging Group Users & Devices

<p style='font-style:italic;'>Since 1.0.32 (Release 9.3)</p><p>Adds and/or removes paging group users and devices.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditExtensions</td><td>Viewing and updating my extension info (includes extension name, number, email and phone number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **pagingOnlyGroupId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**|  | 
 **pagingOnlyGroupId** | **string**|  | 
 **body** | [**EditPagingGroupRequest**](EditPagingGroupRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateParkLocationUsers**
> UpdateParkLocationUsers(ctx, accountId, parkLocationId, optional)
Adds and/or removes park location users

<p style='font-style:italic;'>Since 1.0.32 (Release 9.3)</p><p>Adds and/or removes park location users.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditExtensions</td><td>Viewing and updating my extension info (includes extension name, number, email and phone number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **parkLocationId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**|  | 
 **parkLocationId** | **string**|  | 
 **body** | [**EditParkLocationRequest**](EditParkLocationRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

