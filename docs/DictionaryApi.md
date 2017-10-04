# \DictionaryApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestapiV10DictionaryCountryCountryIdGet**](DictionaryApi.md#RestapiV10DictionaryCountryCountryIdGet) | **Get** /restapi/v1.0/dictionary/country/{countryId} | 
[**RestapiV10DictionaryCountryGet**](DictionaryApi.md#RestapiV10DictionaryCountryGet) | **Get** /restapi/v1.0/dictionary/country | 
[**RestapiV10DictionaryLanguageGet**](DictionaryApi.md#RestapiV10DictionaryLanguageGet) | **Get** /restapi/v1.0/dictionary/language | 
[**RestapiV10DictionaryLanguageLanguageIdGet**](DictionaryApi.md#RestapiV10DictionaryLanguageLanguageIdGet) | **Get** /restapi/v1.0/dictionary/language/{languageId} | 
[**RestapiV10DictionaryLocationGet**](DictionaryApi.md#RestapiV10DictionaryLocationGet) | **Get** /restapi/v1.0/dictionary/location | 
[**RestapiV10DictionaryStateGet**](DictionaryApi.md#RestapiV10DictionaryStateGet) | **Get** /restapi/v1.0/dictionary/state | 
[**RestapiV10DictionaryStateStateIdGet**](DictionaryApi.md#RestapiV10DictionaryStateStateIdGet) | **Get** /restapi/v1.0/dictionary/state/{stateId} | 
[**RestapiV10DictionaryTimezoneGet**](DictionaryApi.md#RestapiV10DictionaryTimezoneGet) | **Get** /restapi/v1.0/dictionary/timezone | 
[**RestapiV10DictionaryTimezoneTimezoneIdGet**](DictionaryApi.md#RestapiV10DictionaryTimezoneTimezoneIdGet) | **Get** /restapi/v1.0/dictionary/timezone/{timezoneId} | 


# **RestapiV10DictionaryCountryCountryIdGet**
> FullCountryInfo RestapiV10DictionaryCountryCountryIdGet($countryId)



Get Country by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryId** | **string**| Internal identifier of a country | 

### Return type

[**FullCountryInfo**](FullCountryInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10DictionaryCountryGet**
> InlineResponseDefault29 RestapiV10DictionaryCountryGet($loginAllowed, $numberSelling, $page, $perPage)



Get Country List


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginAllowed** | **bool**| Specifies whether login with the phone numbers of this country is enabled or not | [optional] 
 **numberSelling** | **bool**| Specifies if RingCentral sells phone numbers of this country | [optional] 
 **page** | **int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | [optional] 
 **perPage** | **int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | [optional] 

### Return type

[**InlineResponseDefault29**](inline_response_default_29.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10DictionaryLanguageGet**
> InlineResponseDefault30 RestapiV10DictionaryLanguageGet()



Get Supported Language List


### Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponseDefault30**](inline_response_default_30.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10DictionaryLanguageLanguageIdGet**
> LanguageInfo RestapiV10DictionaryLanguageLanguageIdGet($languageId)



Get Language by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **languageId** | **string**| Internal identifier of a language | 

### Return type

[**LanguageInfo**](LanguageInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10DictionaryLocationGet**
> InlineResponseDefault31 RestapiV10DictionaryLocationGet($orderBy, $page, $perPage, $stateId, $withNxx)



Get Location List


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderBy** | **string**| Sorts results by the specified property. The default value is &#39;City&#39; | [optional] 
 **page** | **int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39;. | [optional] 
 **perPage** | **int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default. | [optional] 
 **stateId** | **string**| Internal identifier of a state | [optional] 
 **withNxx** | **bool**| Specifies if nxx codes are returned | [optional] 

### Return type

[**InlineResponseDefault31**](inline_response_default_31.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10DictionaryStateGet**
> InlineResponseDefault32 RestapiV10DictionaryStateGet($countryId, $page, $perPage, $withPhoneNumbers)



Get State/Province List


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryId** | **int32**| Internal identifier of a country | [optional] 
 **page** | **int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39;. | [optional] 
 **perPage** | **int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | [optional] 
 **withPhoneNumbers** | **bool**| If &#39;True&#39;, the list of states with phone numbers available for buying is returned. The default value is &#39;False&#39; | [optional] 

### Return type

[**InlineResponseDefault32**](inline_response_default_32.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10DictionaryStateStateIdGet**
> StateInfo RestapiV10DictionaryStateStateIdGet($stateId)



Get State/Province by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stateId** | **string**| Internal identifier of a state | 

### Return type

[**StateInfo**](StateInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10DictionaryTimezoneGet**
> InlineResponseDefault33 RestapiV10DictionaryTimezoneGet($page, $perPage)



Get Time Zone List


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | [optional] 
 **perPage** | **string**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | [optional] 

### Return type

[**InlineResponseDefault33**](inline_response_default_33.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10DictionaryTimezoneTimezoneIdGet**
> TimezoneInfo RestapiV10DictionaryTimezoneTimezoneIdGet($timezoneId)



Get Time Zone by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timezoneId** | **string**| Internal identifier of a timezone | 

### Return type

[**TimezoneInfo**](TimezoneInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

