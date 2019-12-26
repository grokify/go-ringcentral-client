# \CountriesApi

All URIs are relative to *https://portal.vacd.biz/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAvailableCountries**](CountriesApi.md#GetAvailableCountries) | **Get** /admin/accounts/{accountId}/countries/available | Get a list of available countries



## GetAvailableCountries

> []Country GetAvailableCountries(ctx, accountId)
Get a list of available countries

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 

### Return type

[**[]Country**](Country.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

