# \AnsweringRulesApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestapiV10AccountAccountIdExtensionExtensionIdAnsweringRuleAnsweringRuleIdDelete**](AnsweringRulesApi.md#RestapiV10AccountAccountIdExtensionExtensionIdAnsweringRuleAnsweringRuleIdDelete) | **Delete** /restapi/v1.0/account/{accountId}/extension/{extensionId}/answering-rule/{answeringRuleId} | 
[**RestapiV10AccountAccountIdExtensionExtensionIdAnsweringRuleAnsweringRuleIdGet**](AnsweringRulesApi.md#RestapiV10AccountAccountIdExtensionExtensionIdAnsweringRuleAnsweringRuleIdGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/answering-rule/{answeringRuleId} | 
[**RestapiV10AccountAccountIdExtensionExtensionIdAnsweringRuleAnsweringRuleIdPut**](AnsweringRulesApi.md#RestapiV10AccountAccountIdExtensionExtensionIdAnsweringRuleAnsweringRuleIdPut) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/answering-rule/{answeringRuleId} | 
[**RestapiV10AccountAccountIdExtensionExtensionIdAnsweringRuleGet**](AnsweringRulesApi.md#RestapiV10AccountAccountIdExtensionExtensionIdAnsweringRuleGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/answering-rule | 
[**RestapiV10AccountAccountIdExtensionExtensionIdAnsweringRulePost**](AnsweringRulesApi.md#RestapiV10AccountAccountIdExtensionExtensionIdAnsweringRulePost) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/answering-rule | 
[**RestapiV10AccountAccountIdExtensionExtensionIdBusinessHoursGet**](AnsweringRulesApi.md#RestapiV10AccountAccountIdExtensionExtensionIdBusinessHoursGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/business-hours | 
[**RestapiV10AccountAccountIdExtensionExtensionIdGreetingGreetingIdGet**](AnsweringRulesApi.md#RestapiV10AccountAccountIdExtensionExtensionIdGreetingGreetingIdGet) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/greeting/{greetingId} | 
[**RestapiV10AccountAccountIdExtensionExtensionIdGreetingPost**](AnsweringRulesApi.md#RestapiV10AccountAccountIdExtensionExtensionIdGreetingPost) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/greeting | 


# **RestapiV10AccountAccountIdExtensionExtensionIdAnsweringRuleAnsweringRuleIdDelete**
> RestapiV10AccountAccountIdExtensionExtensionIdAnsweringRuleAnsweringRuleIdDelete($accountId, $extensionId, $answeringRuleId)



Delete Answering Rule by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **answeringRuleId** | **string**| Internal identifier of an answering rule. The value can be standard digital ID or specific ID - either business-hours-rule or after-hours-rule | 

### Return type

void (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdAnsweringRuleAnsweringRuleIdGet**
> AnsweringRuleInfo RestapiV10AccountAccountIdExtensionExtensionIdAnsweringRuleAnsweringRuleIdGet($accountId, $extensionId, $answeringRuleId)



Get Custom Answering Rule by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **answeringRuleId** | **string**| Internal identifier of an answering rule. The value can be standard digital ID or specific ID - either business-hours-rule or after-hours-rule | 

### Return type

[**AnsweringRuleInfo**](AnsweringRuleInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdAnsweringRuleAnsweringRuleIdPut**
> AnsweringRuleInfo RestapiV10AccountAccountIdExtensionExtensionIdAnsweringRuleAnsweringRuleIdPut($accountId, $extensionId, $answeringRuleId, $body)



Update Answering Rule by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **answeringRuleId** | **string**| Internal identifier of an answering rule. The value can be standard digital ID or specific ID - either business-hours-rule or after-hours-rule | 
 **body** | [**Body5**](Body5.md)|  | [optional] 

### Return type

[**AnsweringRuleInfo**](AnsweringRuleInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdAnsweringRuleGet**
> InlineResponseDefault11 RestapiV10AccountAccountIdExtensionExtensionIdAnsweringRuleGet($accountId, $extensionId)



Get Answering Rules List


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 

### Return type

[**InlineResponseDefault11**](inline_response_default_11.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdAnsweringRulePost**
> AnsweringRuleInfo RestapiV10AccountAccountIdExtensionExtensionIdAnsweringRulePost($accountId, $extensionId, $body)



Create Custom Answering Rule


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **body** | [**Body4**](Body4.md)|  | [optional] 

### Return type

[**AnsweringRuleInfo**](AnsweringRuleInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdBusinessHoursGet**
> InlineResponseDefault15 RestapiV10AccountAccountIdExtensionExtensionIdBusinessHoursGet($accountId, $extensionId)



Get User Hours Setting


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 

### Return type

[**InlineResponseDefault15**](inline_response_default_15.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdGreetingGreetingIdGet**
> CustomGreetingInfo RestapiV10AccountAccountIdExtensionExtensionIdGreetingGreetingIdGet($accountId, $extensionId, $greetingId)



Get Custom Greeting by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **greetingId** | **string**| Internal identifier of a greeting | 

### Return type

[**CustomGreetingInfo**](CustomGreetingInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdExtensionExtensionIdGreetingPost**
> CustomGreetingInfo RestapiV10AccountAccountIdExtensionExtensionIdGreetingPost($accountId, $extensionId, $body)



Create Custom Greeting


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
 **body** | [**Body10**](Body10.md)|  | [optional] 

### Return type

[**CustomGreetingInfo**](CustomGreetingInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

