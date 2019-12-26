# \CustomFieldsApi

All URIs are relative to *https://DOMAIN.api.engagement.dimelo.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomField**](CustomFieldsApi.md#CreateCustomField) | **Post** /custom_fields | Creating a custom field
[**DeleteCustomField**](CustomFieldsApi.md#DeleteCustomField) | **Delete** /custom_fields/{customFieldId} | Deleting a custom field
[**GetAllCustomFields**](CustomFieldsApi.md#GetAllCustomFields) | **Get** /custom_fields | Getting all custom fields
[**GetCustomField**](CustomFieldsApi.md#GetCustomField) | **Get** /custom_fields/{customFieldId} | Getting a custom field from its id
[**UpdateCustomField**](CustomFieldsApi.md#UpdateCustomField) | **Put** /custom_fields/{customFieldId} | Updating a custom field



## CreateCustomField

> CustomField CreateCustomField(ctx, associatedTypeName, label, optional)
Creating a custom field

This method creates a custom field. In case of success it renders the custom field, otherwise, it renders an error (422 HTTP code).  Authorization​: only users that can create custom fields.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**associatedTypeName** | **string**| The associated type of custom field. It can be IdentityGroup or Intervention. | 
**label** | **string**| The label of the custom field. | 
 **optional** | ***CreateCustomFieldOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCustomFieldOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **key** | **optional.String**| The key of the custom field (example: customer_id). This is used to determine how it is stored on identity groups. | 
 **type_** | **optional.String**| The type of the custom field. It can be string, boolean, text, integer, float, single_choice, | 
 **choices** | [**optional.Interface of []string**](string.md)| A list of choices to be for single_choice, or multiple_choice types. This must be given | 
 **multiple** | **optional.Bool**| true or false, this as no effect on single_choice, multiple_choice or boolean types | 
 **position** | **optional.Int32**| an integer that indicates custom field’s position between others (default: -1). | [default to -1]

### Return type

[**CustomField**](CustomField.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomField

> CustomField DeleteCustomField(ctx, customFieldId)
Deleting a custom field

This method destroys an existing custom field. It renders custom field itself. It renders a 404 if id is invalid.  Authorization​: only users that are able to destroy custom fields..

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customFieldId** | **string**|  | 

### Return type

[**CustomField**](CustomField.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllCustomFields

> GetAllCustomFieldsResponse GetAllCustomFields(ctx, optional)
Getting all custom fields

This method renders custom fields ordered by position (ascending).  Authorization​: only users that can see custom fields in administration section.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllCustomFieldsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllCustomFieldsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| The record index to start. Default value is 0. | 
 **limit** | **optional.Int32**| The max number of records to return. Default value is 30, max value is 150. | 

### Return type

[**GetAllCustomFieldsResponse**](GetAllCustomFieldsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomField

> CustomField GetCustomField(ctx, customFieldId)
Getting a custom field from its id

This method renders a custom field from given id.  Authorization​: only users that can see custom fields in administration section.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customFieldId** | **string**|  | 

### Return type

[**CustomField**](CustomField.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomField

> CustomField UpdateCustomField(ctx, customFieldId, optional)
Updating a custom field

This method updates an existing custom field from given attributes and renders it in case of success.  Authorization​: only users that are able to update custom fields.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customFieldId** | **string**|  | 
 **optional** | ***UpdateCustomFieldOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCustomFieldOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **label** | **optional.String**| Custom field’s label. | 
 **choices** | [**optional.Interface of []string**](string.md)|  | 
 **position** | **optional.Int32**| Custom field’s position. | 

### Return type

[**CustomField**](CustomField.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

