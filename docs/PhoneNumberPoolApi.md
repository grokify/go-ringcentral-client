# \PhoneNumberPoolApi

All URIs are relative to *https://platform.devtest.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestapiV10NumberPoolLookupPost**](PhoneNumberPoolApi.md#RestapiV10NumberPoolLookupPost) | **Post** /restapi/v1.0/number-pool/lookup | 
[**RestapiV10NumberPoolReservePost**](PhoneNumberPoolApi.md#RestapiV10NumberPoolReservePost) | **Post** /restapi/v1.0/number-pool/reserve | 


# **RestapiV10NumberPoolLookupPost**
> InlineResponseDefault35 RestapiV10NumberPoolLookupPost($areaCode, $countryCode, $countryId, $exclude, $extendedSearch, $line, $numberPattern, $nxx, $npa, $paymentType, $perPage, $smsEnabled)



Look up Phone Number


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **areaCode** | **int32**| Area code of the location | [optional] 
 **countryCode** | **string**| Two-letter country code, complying with the ISO standard | [optional] 
 **countryId** | **string**| Internal identifier of a country; &#39;1&#39;- the US; &#39;39&#39; - Canada; &#39;224&#39; - the UK. The default value is &#39;1&#39; | [optional] 
 **exclude** | **string**| A string of digits (one and more) that should not appear among the last four digits (line part) of the phone numbers that will be returned. It is possible to specify several exclude parameters. If specified, it is taken into account in all returned phone numbers both in the phone numbers satisfying to parameters of lookup and in alternative phone numbers (in case when extendedSearch is specified) | [optional] 
 **extendedSearch** | **bool**| If the value is &#39;False&#39;, then the returned numbers exactly correspond to the specified NXX, NPA and LINE or countryCode, areaCode and numberPattern parameters. If the value is &#39;True&#39;, then the resulting numbers are ranked and returned with the rank attribute values (1-10). The default value is &#39;False&#39; | [optional] 
 **line** | **string**| LINE pattern for vanity or wildcard search. Digits, Latin characters and asterisks are allowed (usually 4 characters) | [optional] 
 **numberPattern** | **string**| Phone number pattern (for wildcard or vanity search). For NANP countries (US, Canada) is concatenation of nxx (the first three digits) and line. If the first three characters are specified as not digits (e.g. 5** or CAT) then parameter extendedSearch will be ignored. | [optional] 
 **nxx** | **string**| NXX pattern for vanity or wildcard search. Digits, Latin characters and asterisks are allowed (usually 3 characters) | [optional] 
 **npa** | **string**| Area code (mandatory). For example, 800, 844, 855, 866, 877, 888 for North America; and 647 for Canada | [optional] 
 **paymentType** | **string**| Payment type. Default is &#39;Local&#39; (it should correlate with the npa provided) | [optional] 
 **perPage** | **int32**| Indicates the page size (number of items). If not specified, the value is &#39;10&#39; by default | [optional] 
 **smsEnabled** | **bool**| Specifies if SMS activation is available for the number. If specified, it is taken into account in all returned phone numbers both in the phone numbers satisfying to parameters of lookup and in alternative phone numbers (in case when extendedSearch is specified). If not specified, the value is null. | [optional] 

### Return type

[**InlineResponseDefault35**](inline_response_default_35.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestapiV10NumberPoolReservePost**
> InlineResponseDefault36 RestapiV10NumberPoolReservePost($body)



Reserve Phone Number


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Body19**](Body19.md)|  | [optional] 

### Return type

[**InlineResponseDefault36**](inline_response_default_36.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

