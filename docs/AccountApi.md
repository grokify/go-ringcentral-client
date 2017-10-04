# \AccountApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestapiV10AccountAccountIdBusinessAddressGet**](AccountApi.md#RestapiV10AccountAccountIdBusinessAddressGet) | **Get** /restapi/v1.0/account/{accountId}/business-address | 
[**RestapiV10AccountAccountIdBusinessAddressPut**](AccountApi.md#RestapiV10AccountAccountIdBusinessAddressPut) | **Put** /restapi/v1.0/account/{accountId}/business-address | 
[**RestapiV10AccountAccountIdDialingPlanGet**](AccountApi.md#RestapiV10AccountAccountIdDialingPlanGet) | **Get** /restapi/v1.0/account/{accountId}/dialing-plan | 
[**RestapiV10AccountAccountIdGet**](AccountApi.md#RestapiV10AccountAccountIdGet) | **Get** /restapi/v1.0/account/{accountId} | 
[**RestapiV10AccountAccountIdPhoneNumberGet**](AccountApi.md#RestapiV10AccountAccountIdPhoneNumberGet) | **Get** /restapi/v1.0/account/{accountId}/phone-number | 
[**RestapiV10AccountAccountIdPhoneNumberPhoneNumberIdGet**](AccountApi.md#RestapiV10AccountAccountIdPhoneNumberPhoneNumberIdGet) | **Get** /restapi/v1.0/account/{accountId}/phone-number/{phoneNumberId} | 
[**RestapiV10AccountAccountIdServiceInfoGet**](AccountApi.md#RestapiV10AccountAccountIdServiceInfoGet) | **Get** /restapi/v1.0/account/{accountId}/service-info | 


# **RestapiV10AccountAccountIdBusinessAddressGet**
> InlineResponseDefault2 RestapiV10AccountAccountIdBusinessAddressGet($accountId)



Get Company Business Address


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 

### Return type

[**InlineResponseDefault2**](inline_response_default_2.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdBusinessAddressPut**
> InlineResponseDefault2 RestapiV10AccountAccountIdBusinessAddressPut($accountId, $body)



Update Company Business Address


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **body** | [**Body3**](Body3.md)|  | [optional] 

### Return type

[**InlineResponseDefault2**](inline_response_default_2.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdDialingPlanGet**
> InlineResponseDefault6 RestapiV10AccountAccountIdDialingPlanGet($accountId, $page, $perPage)



Get IBO Dialing Plans


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **page** | **int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | [optional] 
 **perPage** | **int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | [optional] 

### Return type

[**InlineResponseDefault6**](inline_response_default_6.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdGet**
> AccountInfo RestapiV10AccountAccountIdGet($accountId)



Get Account Info by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 

### Return type

[**AccountInfo**](AccountInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdPhoneNumberGet**
> InlineResponseDefault26 RestapiV10AccountAccountIdPhoneNumberGet($accountId, $page, $perPage, $usageType)



Get Account Phone Numbers


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **page** | **int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | [optional] 
 **perPage** | **int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | [optional] 
 **usageType** | **string**| Usage type of the phone number | [optional] 

### Return type

[**InlineResponseDefault26**](inline_response_default_26.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdPhoneNumberPhoneNumberIdGet**
> PhoneNumberInfo RestapiV10AccountAccountIdPhoneNumberPhoneNumberIdGet($accountId, $phoneNumberId)



Get Phone Number by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **phoneNumberId** | **string**| Internal identifier of a phone number | 

### Return type

[**PhoneNumberInfo**](PhoneNumberInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10AccountAccountIdServiceInfoGet**
> AccountServiceInfo RestapiV10AccountAccountIdServiceInfoGet($accountId)



Get Account Service Info


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 

### Return type

[**AccountServiceInfo**](AccountServiceInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

