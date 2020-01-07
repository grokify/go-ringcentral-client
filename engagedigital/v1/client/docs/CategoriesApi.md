# \CategoriesApi

All URIs are relative to *https://DOMAIN.api.engagement.dimelo.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCategory**](CategoriesApi.md#CreateCategory) | **Post** /categories | Creating a category
[**DeleteCategory**](CategoriesApi.md#DeleteCategory) | **Delete** /categories/{categoryId} | Deleting a category
[**GetAllCategories**](CategoriesApi.md#GetAllCategories) | **Get** /categories | Getting all categories
[**GetCategory**](CategoriesApi.md#GetCategory) | **Get** /categories/{categoryId} | Getting a category from its id
[**UpdateCategory**](CategoriesApi.md#UpdateCategory) | **Put** /categories/{categoryId} | Updating a category



## CreateCategory

> Category CreateCategory(ctx, optional)
Creating a category

This method creates a new team. In case of success it renders the created tag, otherwise, it renders an error (422 HTTP code).  Note: The fields ​`mandatory`,​ `​multiple`,​ ​`post_qualification​`, `s​ource_ids`​ and `u​nselectable​` are accounted for only if the Category has no parent.  Authorization​: only users that can manage teams.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCategoryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Category name. | 
 **parentId** | **optional.String**| ID of parent category. | 
 **color** | **optional.Int32**| displayed color for the category, see Category colors.  | 
 **mandatory** | **optional.Bool**| mandatory categorization (Boolean). | 
 **multiple** | **optional.Bool**| allow to assign multiple child categories (Boolean). | 
 **postQualification** | **optional.Bool**| post qualification (Boolean). | 
 **unselectable** | **optional.Bool**| root category is unselectable (Boolean). | 
 **sourceIds** | [**optional.Interface of []string**](string.md)| List of source id. | 

### Return type

[**Category**](Category.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCategory

> Category DeleteCategory(ctx, categoryId, optional)
Deleting a category

This method destroys an existing category. It renders the category itself. It renders a 404 if id is invalid.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **string**|  | 
 **optional** | ***DeleteCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteCategoryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **takeOverCategoryId** | **optional.String**| ID of a category to recategorize (optional). | 

### Return type

[**Category**](Category.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllCategories

> GetAllCategoriesResponse GetAllCategories(ctx, optional)
Getting all categories

This method renders categories ordered by creation date (ascending).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllCategoriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parentId** | **optional.String**| To filter categories on given category parent id. | 
 **offset** | **optional.Int32**| The record index to start. Default value is 0. | 
 **limit** | **optional.Int32**| The max number of records to return. Default value is 30, max value is 150. | 

### Return type

[**GetAllCategoriesResponse**](GetAllCategoriesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategory

> Category GetCategory(ctx, categoryId)
Getting a category from its id

This method renders a category from given id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **string**|  | 

### Return type

[**Category**](Category.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCategory

> Category UpdateCategory(ctx, categoryId, optional)
Updating a category

This method creates a new team. In case of success it renders the created tag, otherwise, it renders an error (422 HTTP code).  Note: The fields ​`mandatory`,​ ​`multiple`,​ ​`post_qualification​`, `s​ource_ids​` and `u​nselectable​` are accounted for only if the Category has no parent.  Authorization​: only users that can manage teams.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **string**|  | 
 **optional** | ***UpdateCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCategoryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| Category name. | 
 **parentId** | **optional.String**| ID of parent category. | 
 **color** | **optional.Int32**| displayed color for the category, see Category colors.  | 
 **mandatory** | **optional.Bool**| mandatory categorization (Boolean). | 
 **multiple** | **optional.Bool**| allow to assign multiple child categories (Boolean). | 
 **postQualification** | **optional.Bool**| post qualification (Boolean). | 
 **unselectable** | **optional.Bool**| root category is unselectable (Boolean). | 
 **sourceIds** | [**optional.Interface of []string**](string.md)| List of source id. | 

### Return type

[**Category**](Category.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

