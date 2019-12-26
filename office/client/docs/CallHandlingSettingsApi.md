# \CallHandlingSettingsApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlockNumber**](CallHandlingSettingsApi.md#BlockNumber) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/blocked-number | Add Blocked Numbers
[**CreateAnsweringRuleInfo**](CallHandlingSettingsApi.md#CreateAnsweringRuleInfo) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/answering-rule | Create Custom Call Handling Rules
[**CreateCompanyAnsweringRuleInfo**](CallHandlingSettingsApi.md#CreateCompanyAnsweringRuleInfo) | **Post** /restapi/v1.0/account/{accountId}/answering-rule | Create Company Call Handling Rule
[**CreateExtensionForwardingNumber**](CallHandlingSettingsApi.md#CreateExtensionForwardingNumber) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/forwarding-number | Create Forwarding Numbers
[**CreateGreeting**](CallHandlingSettingsApi.md#CreateGreeting) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/greeting | Create User Custom Greeting
[**CreateIVRMenu**](CallHandlingSettingsApi.md#CreateIVRMenu) | **Post** /restapi/v1.0/account/{accountId}/ivr-menus | Create IVR Menu
[**CreatePrompts**](CallHandlingSettingsApi.md#CreatePrompts) | **Post** /restapi/v1.0/account/{accountId}/ivr-prompts | Create IVR Prompt
[**DeleteAnsweringRule**](CallHandlingSettingsApi.md#DeleteAnsweringRule) | **Delete** /restapi/v1.0/account/{accountId}/extension/{extensionId}/answering-rule/{ruleId} | Delete Call Handling Rule
[**DeleteExtensionForwardingNumber**](CallHandlingSettingsApi.md#DeleteExtensionForwardingNumber) | **Delete** /restapi/v1.0/account/{accountId}/extension/{extensionId}/forwarding-number/{forwardingNumberId} | Delete Forwarding Number
[**DeleteIVRPrompt**](CallHandlingSettingsApi.md#DeleteIVRPrompt) | **Delete** /restapi/v1.0/account/{accountId}/ivr-prompts/{promptId} | Delete IVR Prompt
[**GetCompanyGreeting**](CallHandlingSettingsApi.md#GetCompanyGreeting) | **Post** /restapi/v1.0/account/{accountId}/greeting | Create Custom Company Greeting
[**GetForwardingNumber**](CallHandlingSettingsApi.md#GetForwardingNumber) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/forwarding-number/{forwardingNumberId} | Get Forwarding Number
[**GetGreeting**](CallHandlingSettingsApi.md#GetGreeting) | **Get** /restapi/v1.0/dictionary/greeting/{greetingId} | Get Greeting Info
[**GetGreetingByID**](CallHandlingSettingsApi.md#GetGreetingByID) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/greeting/{greetingId} | Get Custom Greeting Info
[**GetGreetings**](CallHandlingSettingsApi.md#GetGreetings) | **Get** /restapi/v1.0/dictionary/greeting | Get Standard Greetings
[**GetIVRMenuById**](CallHandlingSettingsApi.md#GetIVRMenuById) | **Get** /restapi/v1.0/account/{accountId}/ivr-menus/{ivrMenuId} | Get IVR Menu
[**GetIVRPrompt**](CallHandlingSettingsApi.md#GetIVRPrompt) | **Get** /restapi/v1.0/account/{accountId}/ivr-prompts/{promptId} | Get IVR Prompt
[**GetIVRPromptContent**](CallHandlingSettingsApi.md#GetIVRPromptContent) | **Get** /restapi/v1.0/account/{accountId}/ivr-prompts/{promptId}/content | Get IVR Prompt Content
[**GetIVRPrompts**](CallHandlingSettingsApi.md#GetIVRPrompts) | **Get** /restapi/v1.0/account/{accountId}/ivr-prompts | Get IVR Prompts
[**ListBlockedNumbers**](CallHandlingSettingsApi.md#ListBlockedNumbers) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/blocked-number | Get Blocked Numbers
[**ListCompanyAnsweringRule**](CallHandlingSettingsApi.md#ListCompanyAnsweringRule) | **Get** /restapi/v1.0/account/{accountId}/answering-rule | Get Company Call Handling Rules
[**ListExtensionForwardingNumbers**](CallHandlingSettingsApi.md#ListExtensionForwardingNumbers) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/forwarding-number | Get Forwarding Numbers
[**LoadAnsweringRuleInfo**](CallHandlingSettingsApi.md#LoadAnsweringRuleInfo) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/answering-rule/{ruleId} | Get Call Handling Rule
[**LoadAnsweringRulesList**](CallHandlingSettingsApi.md#LoadAnsweringRulesList) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/answering-rule | Get Call Handling Rules
[**LoadBlockedNumber**](CallHandlingSettingsApi.md#LoadBlockedNumber) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/blocked-number/{blockedNumberId} | Get Blocked Number
[**LoadBusinesshoursInfo**](CallHandlingSettingsApi.md#LoadBusinesshoursInfo) | **Get** /restapi/v1.0/account/{accountId}/business-hours | Get Company Business Hours
[**LoadCompanyAnsweringRuleInfo**](CallHandlingSettingsApi.md#LoadCompanyAnsweringRuleInfo) | **Get** /restapi/v1.0/account/{accountId}/answering-rule/{ruleId} | Get Company Call Handling Rule
[**LoadUserBusinessHours**](CallHandlingSettingsApi.md#LoadUserBusinessHours) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/business-hours | Get User Business Hours
[**RestapiV10AccountAccountIdIvrMenusIvrMenuIdPut**](CallHandlingSettingsApi.md#RestapiV10AccountAccountIdIvrMenusIvrMenuIdPut) | **Put** /restapi/v1.0/account/{accountId}/ivr-menus/{ivrMenuId} | Update IVR Menu
[**UnblockNumber**](CallHandlingSettingsApi.md#UnblockNumber) | **Delete** /restapi/v1.0/account/{accountId}/extension/{extensionId}/blocked-number/{blockedNumberId} | Delete Blocked Number
[**UpdateAnsweringRuleInfo**](CallHandlingSettingsApi.md#UpdateAnsweringRuleInfo) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/answering-rule/{ruleId} | Update Custom Call Handling Rule
[**UpdateBlockedNumber**](CallHandlingSettingsApi.md#UpdateBlockedNumber) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/blocked-number/{blockedNumberId} | Update Blocked Number
[**UpdateCompanyAnsweringRuleInfo**](CallHandlingSettingsApi.md#UpdateCompanyAnsweringRuleInfo) | **Put** /restapi/v1.0/account/{accountId}/answering-rule/{ruleId} | Update Company Call Handling Rule
[**UpdateCompanyBusinessHours**](CallHandlingSettingsApi.md#UpdateCompanyBusinessHours) | **Put** /restapi/v1.0/account/{accountId}/business-hours | Update Company Business Hours
[**UpdateExtensionForwardingNumber**](CallHandlingSettingsApi.md#UpdateExtensionForwardingNumber) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/forwarding-number/{forwardingNumberId} | Update Forwarding Numbers
[**UpdateUserBusinessHours**](CallHandlingSettingsApi.md#UpdateUserBusinessHours) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/business-hours | Update User Business Hours


# **BlockNumber**
> BlockedNumberInfo BlockNumber(ctx, accountId, extensionId, optional)
Add Blocked Numbers

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditExtensions</td><td>Viewing and updating my extension info (includes extension name, number, email and phone number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***BlockNumberOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BlockNumberOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **blockedNumberInfo** | [**optional.Interface of BlockedNumberInfo**](BlockedNumberInfo.md)|  | 

### Return type

[**BlockedNumberInfo**](BlockedNumberInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAnsweringRuleInfo**
> AnsweringRuleInfo CreateAnsweringRuleInfo(ctx, accountId, extensionId, createAnsweringRuleRequest)
Create Custom Call Handling Rules

<p style='font-style:italic;'>Since 1.0.24 (Release 8.0)</p><p>Creates a custom answering rule for a particular caller ID.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditExtensions</td><td>Viewing and updating my extension info (includes extension name, number, email and phone number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
  **createAnsweringRuleRequest** | [**CreateAnsweringRuleRequest**](CreateAnsweringRuleRequest.md)| JSON body | 

### Return type

[**AnsweringRuleInfo**](AnsweringRuleInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCompanyAnsweringRuleInfo**
> CompanyAnsweringRuleInfo CreateCompanyAnsweringRuleInfo(ctx, accountId, companyAnsweringRuleRequest)
Create Company Call Handling Rule

<p style='font-style:italic;'>Since 1.0.31 (Release 9.2)</p><p>Creates a company answering rule for a particular caller ID.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr><tr><td class='code'>EditExtensions</td><td>Viewing and updating user extension info (includes extension name, number, email and phone number, assigned phone numbers, devices and other extension settings)</td></tr><tr><td class='code'>EditAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **companyAnsweringRuleRequest** | [**CompanyAnsweringRuleRequest**](CompanyAnsweringRuleRequest.md)| JSON body | 

### Return type

[**CompanyAnsweringRuleInfo**](CompanyAnsweringRuleInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateExtensionForwardingNumber**
> ForwardingNumberInfo CreateExtensionForwardingNumber(ctx, accountId, extensionId, createForwardingNumberRequest)
Create Forwarding Numbers

<p style='font-style:italic;'>Since 1.0.12 (Release 6.4)</p><p>Adds a new forwarding number to the forwarding number list.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditExtensions</td><td>Viewing and updating user extension info (includes extension name, number, email and phone number, assigned phone numbers, devices and other extension settings)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
  **createForwardingNumberRequest** | [**CreateForwardingNumberRequest**](CreateForwardingNumberRequest.md)| JSON body | 

### Return type

[**ForwardingNumberInfo**](ForwardingNumberInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateGreeting**
> CustomCompanyGreetingInfo CreateGreeting(ctx, accountId, extensionId, optional)
Create User Custom Greeting

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***CreateGreetingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateGreetingOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **customGreetingRequest** | [**optional.Interface of CustomGreetingRequest**](CustomGreetingRequest.md)|  | 

### Return type

[**CustomCompanyGreetingInfo**](CustomCompanyGreetingInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateIVRMenu**
> IvrMenuInfo CreateIVRMenu(ctx, accountId, ivrMenuInfo)
Create IVR Menu

<p style='font-style:italic;'>Since 1.0.32 (Release 9.3)</p><p>Creates a company IVR menu</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditAccounts</td><td>Viewing and updating user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 
  **ivrMenuInfo** | [**IvrMenuInfo**](IvrMenuInfo.md)| JSON body | 

### Return type

[**IvrMenuInfo**](IVRMenuInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePrompts**
> PromptInfo CreatePrompts(ctx, accountId, attachment, optional)
Create IVR Prompt

<p style='font-style:italic;'>Since 1.0.32 (Release 9.3)</p><p>Creates an IVR prompt.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 
  **attachment** | ***os.File*****os.File**| Audio file that will be used as a prompt. Attachment cannot be empty, only audio files are supported | 
 **optional** | ***CreatePromptsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreatePromptsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **optional.String**| Description of file contents. | 

### Return type

[**PromptInfo**](PromptInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAnsweringRule**
> DeleteAnsweringRule(ctx, accountId, extensionId, ruleId)
Delete Call Handling Rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
  **ruleId** | **string**| Internal identifier of an answering rule | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteExtensionForwardingNumber**
> DeleteExtensionForwardingNumber(ctx, accountId, extensionId, forwardingNumberId)
Delete Forwarding Number

<p style='font-style:italic;'>Since 1.0.24 (Release 8.0)</p><p>Deletes a forwarding number from the forwarding number list by its ID.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditExtensions</td><td>Viewing and updating user extension info (includes extension name, number, email and phone number, assigned phone numbers, devices and other extension settings)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
  **forwardingNumberId** | **string**| Internal identifier of a forwarding number | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIVRPrompt**
> DeleteIVRPrompt(ctx, accountId, promptId)
Delete IVR Prompt

<p style='font-style:italic;'>Since 1.0.32 (Release 9.3)</p><p>Deletes an IVR prompt by ID</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Heavy</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 
  **promptId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCompanyGreeting**
> CustomCompanyGreetingInfo GetCompanyGreeting(ctx, accountId, customCompanyGreetingRequest)
Create Custom Company Greeting

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **customCompanyGreetingRequest** | [**CustomCompanyGreetingRequest**](CustomCompanyGreetingRequest.md)| JSON body | 

### Return type

[**CustomCompanyGreetingInfo**](CustomCompanyGreetingInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetForwardingNumber**
> ForwardingNumberResource GetForwardingNumber(ctx, forwardingNumberId, extensionId, accountId)
Get Forwarding Number

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **forwardingNumberId** | **string**|  | 
  **extensionId** | **string**|  | 
  **accountId** | **string**|  | 

### Return type

[**ForwardingNumberResource**](ForwardingNumberResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGreeting**
> DictionaryGreetingInfo GetGreeting(ctx, greetingId)
Get Greeting Info

<p style='font-style:italic;'>Since 8.2 (Release 1.0.26)</p><p>Returns a standard greeting by ID</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>View Greetings</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **greetingId** | **string**|  | 

### Return type

[**DictionaryGreetingInfo**](DictionaryGreetingInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGreetingByID**
> CustomCompanyGreetingInfo GetGreetingByID(ctx, accountId, extensionId, greetingId)
Get Custom Greeting Info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
  **greetingId** | **int32**|  | 

### Return type

[**CustomCompanyGreetingInfo**](CustomCompanyGreetingInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGreetings**
> DictionaryGreetingList GetGreetings(ctx, optional)
Get Standard Greetings

<p style='font-style:italic;'>Since 8.2 (Release 1.0.26)</p><p>Returns a list of predefined standard greetings. Please note: Custom greetings recorded by user are not returned in response to this request. See Get Extension Custom Greetings.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>View Greetings</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetGreetingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetGreetingsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | 
 **perPage** | **optional.Int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | 
 **type_** | **optional.String**| Type of a greeting, specifying the case when the greeting is played | 
 **usageType** | **optional.String**| Usage type of a greeting, specifying if the greeting is applied for user extension or department extension | 

### Return type

[**DictionaryGreetingList**](DictionaryGreetingList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIVRMenuById**
> IvrMenuInfo GetIVRMenuById(ctx, accountId, ivrMenuId)
Get IVR Menu

<p style='font-style:italic;'>Since 1.0.32 (Release 9.3)</p><p>Returns a company IVR menu by ID</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 
  **ivrMenuId** | **string**|  | 

### Return type

[**IvrMenuInfo**](IVRMenuInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIVRPrompt**
> PromptInfo GetIVRPrompt(ctx, accountId, promptId)
Get IVR Prompt

<p style='font-style:italic;'>Since 1.0.32 (Release 9.3)</p><p>Returns an IVR prompt by ID</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 
  **promptId** | **string**|  | 

### Return type

[**PromptInfo**](PromptInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIVRPromptContent**
> GetIVRPromptContent(ctx, accountId, promptId)
Get IVR Prompt Content

<p style='font-style:italic;'>Since 1.0.32 (Release 9.3)</p><p>Returns media content of an IVR prompt by its ID</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 
  **promptId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIVRPrompts**
> IvrPrompts GetIVRPrompts(ctx, accountId)
Get IVR Prompts

<p style='font-style:italic;'>Since 1.0.32 (Release 9.3)</p><p>Returns a list of IVR prompts.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 

### Return type

[**IvrPrompts**](IVRPrompts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBlockedNumbers**
> BlockedNumbersList ListBlockedNumbers(ctx, accountId, extensionId)
Get Blocked Numbers

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]

### Return type

[**BlockedNumbersList**](BlockedNumbersList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCompanyAnsweringRule**
> CompanyAnsweringRuleList ListCompanyAnsweringRule(ctx, accountId)
Get Company Call Handling Rules

<p style='font-style:italic;'>Since 1.0.31 (Release 9.2)</p><p>Returns a list of company answering rules.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]

### Return type

[**CompanyAnsweringRuleList**](CompanyAnsweringRuleList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListExtensionForwardingNumbers**
> GetExtensionForwardingNumberListResponse ListExtensionForwardingNumbers(ctx, accountId, extensionId, optional)
Get Forwarding Numbers

<p style='font-style:italic;'>Since 1.0.7 (Release 5.16)</p><p>Returns the list of extension phone numbers used for call forwarding and call flip. The returned list contains all the extension phone numbers that are used for call forwarding and call flip.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***ListExtensionForwardingNumbersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListExtensionForwardingNumbersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | 
 **perPage** | **optional.Int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | 

### Return type

[**GetExtensionForwardingNumberListResponse**](GetExtensionForwardingNumberListResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadAnsweringRuleInfo**
> AnsweringRuleInfo LoadAnsweringRuleInfo(ctx, accountId, extensionId, ruleId)
Get Call Handling Rule

<p style='font-style:italic;'>Since 1.0.15 (Release 7.0)</p><p>Returns an answering rule by ID.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
  **ruleId** | **string**| Internal identifier of an answering rule. The value can be standard digital ID or specific ID - either business-hours-rule or after-hours-rule | 

### Return type

[**AnsweringRuleInfo**](AnsweringRuleInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadAnsweringRulesList**
> LoadAnsweringRulesList(ctx, accountId, extensionId, optional)
Get Call Handling Rules

<p style='font-style:italic;'>Since 1.0.15 (Release 7.0)</p><p>Returns the extension answering rules.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***LoadAnsweringRulesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoadAnsweringRulesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.String**|  | [default to 1]
 **perPage** | **optional.String**|  | [default to 100]

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadBlockedNumber**
> BlockedNumberInfo LoadBlockedNumber(ctx, accountId, extensionId, blockedNumberId)
Get Blocked Number

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
  **blockedNumberId** | **int32**| Internal identifiers of a blocked number list entry | 

### Return type

[**BlockedNumberInfo**](BlockedNumberInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadBusinesshoursInfo**
> CompanyBusinessHours LoadBusinesshoursInfo(ctx, accountId)
Get Company Business Hours

<p style='font-style:italic;'>Since 1.0.32 (Release 9.3)</p><p>Returns company hours when answering rules are to be applied.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]

### Return type

[**CompanyBusinessHours**](CompanyBusinessHours.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadCompanyAnsweringRuleInfo**
> CompanyAnsweringRuleInfo LoadCompanyAnsweringRuleInfo(ctx, accountId, ruleId)
Get Company Call Handling Rule

<p style='font-style:italic;'>Since 1.0.31 (Release 9.2)</p><p>Returns a company answering rule by ID.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **ruleId** | **string**| Internal identifier of an answering rule. The value can be standard digital ID or specific ID - either business-hours-rule or after-hours-rule | 

### Return type

[**CompanyAnsweringRuleInfo**](CompanyAnsweringRuleInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadUserBusinessHours**
> GetUserBusinessHoursResponse LoadUserBusinessHours(ctx, accountId, extensionId)
Get User Business Hours

<p style='font-style:italic;'>Since 1.0.15 (Release 7.0)</p><p>Returns the extension user hours when answering rules are to be applied.</p><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]

### Return type

[**GetUserBusinessHoursResponse**](GetUserBusinessHoursResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdIvrMenusIvrMenuIdPut**
> IvrMenuInfo RestapiV10AccountAccountIdIvrMenusIvrMenuIdPut(ctx, accountId, ivrMenuId, ivrMenuInfo)
Update IVR Menu

<p style='font-style:italic;'>Since 1.0.32 (Release 9.3)</p><p>Returns a company IVR menu by ID</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 
  **ivrMenuId** | **string**|  | 
  **ivrMenuInfo** | [**IvrMenuInfo**](IvrMenuInfo.md)| JSON body | 

### Return type

[**IvrMenuInfo**](IVRMenuInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnblockNumber**
> UnblockNumber(ctx, accountId, extensionId, blockedNumberId)
Delete Blocked Number

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditExtensions</td><td>Viewing and updating my extension info (includes extension name, number, email and phone number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
  **blockedNumberId** | **int32**| Internal identifiers of a blocked number list entry | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAnsweringRuleInfo**
> AnsweringRuleInfo UpdateAnsweringRuleInfo(ctx, accountId, extensionId, ruleId, updateAnsweringRuleRequest)
Update Custom Call Handling Rule

<p style='font-style:italic;'>Since 1.0.24 (Release 8.0)</p><p>Updates a custom answering rule for a particular caller ID.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditExtensions</td><td>Viewing and updating my extension info (includes extension name, number, email and phone number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
  **ruleId** | **string**| Internal identifier of an answering rule | 
  **updateAnsweringRuleRequest** | [**UpdateAnsweringRuleRequest**](UpdateAnsweringRuleRequest.md)| JSON body | 

### Return type

[**AnsweringRuleInfo**](AnsweringRuleInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBlockedNumber**
> BlockedNumberInfo UpdateBlockedNumber(ctx, accountId, extensionId, blockedNumberId, optional)
Update Blocked Number

<p style='font-style:italic;'></p><p></p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditExtensions</td><td>Viewing and updating my extension info (includes extension name, number, email and phone number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
  **blockedNumberId** | **int32**| Internal identifier of a blocked number list entry | 
 **optional** | ***UpdateBlockedNumberOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateBlockedNumberOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **blockedNumberInfo** | [**optional.Interface of BlockedNumberInfo**](BlockedNumberInfo.md)|  | 

### Return type

[**BlockedNumberInfo**](BlockedNumberInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCompanyAnsweringRuleInfo**
> CompanyAnsweringRuleInfo UpdateCompanyAnsweringRuleInfo(ctx, accountId, ruleId, companyAnsweringRuleUpdate)
Update Company Call Handling Rule

<p style='font-style:italic;'>Since 1.0.31 (Release 9.2)</p><p>Updates a company answering rule.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>ReadAccounts</td><td>Viewing user account info (including name, business name, address and phone number/account number)</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **ruleId** | **string**| Internal identifier of an answering rule. The value can be standard digital ID or specific ID - either business-hours-rule or after-hours-rule | 
  **companyAnsweringRuleUpdate** | [**CompanyAnsweringRuleUpdate**](CompanyAnsweringRuleUpdate.md)| JSON body | 

### Return type

[**CompanyAnsweringRuleInfo**](CompanyAnsweringRuleInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCompanyBusinessHours**
> CompanyBusinessHours UpdateCompanyBusinessHours(ctx, accountId, companyBusinessHoursUpdateRequest)
Update Company Business Hours

<p style='font-style:italic;'>Since 1.0.24 (Release 8.0)</p><p>Updates company hours when answering rules are to be applied.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditExtensions</td><td>Viewing and updating my extension info (includes extension name, number, email and phone number)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **companyBusinessHoursUpdateRequest** | [**CompanyBusinessHoursUpdateRequest**](CompanyBusinessHoursUpdateRequest.md)| JSON body | 

### Return type

[**CompanyBusinessHours**](CompanyBusinessHours.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateExtensionForwardingNumber**
> ForwardingNumberInfo UpdateExtensionForwardingNumber(ctx, accountId, extensionId, forwardingNumberId, updateForwardingNumberRequest)
Update Forwarding Numbers

<p style='font-style:italic;'>Since 1.0.24 (Release 8.0)</p><p>Updates an existent forwarding number from the forwarding number list.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditExtensions</td><td>Viewing and updating user extension info (includes extension name, number, email and phone number, assigned phone numbers, devices and other extension settings)</td></tr></tbody></table><h4>Usage Plan Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
  **forwardingNumberId** | **string**| Internal identifier of a forwarding number; returned in response in the id field | 
  **updateForwardingNumberRequest** | [**UpdateForwardingNumberRequest**](UpdateForwardingNumberRequest.md)| JSON body | 

### Return type

[**ForwardingNumberInfo**](ForwardingNumberInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserBusinessHours**
> UserBusinessHoursUpdate UpdateUserBusinessHours(ctx, accountId, extensionId, userBusinessHoursUpdateRequest)
Update User Business Hours

<p style='font-style:italic;'>Since 1.0.28 (Release 8.4)</p><p>Updates the extension user hours when answering rules are to be applied.</p><h4>Required Permissions</h4><table class='fullwidth'><thead><tr><th>Permission</th><th>Description</th></tr></thead><tbody><tr><td class='code'>EditExtensions</td><td>Viewing and updating user extension info (includes extension name, number, email and phone number, assigned phone numbers, devices and other extension settings)</td></tr></tbody></table><h4>API Group</h4><p>Medium</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
  **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
  **userBusinessHoursUpdateRequest** | [**UserBusinessHoursUpdateRequest**](UserBusinessHoursUpdateRequest.md)| JSON body | 

### Return type

[**UserBusinessHoursUpdate**](UserBusinessHoursUpdate.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

