# \FoldersApi

All URIs are relative to *https://DOMAIN.api.engagement.dimelo.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFolder**](FoldersApi.md#CreateFolder) | **Post** /folders | Creating a folder
[**DeleteFolder**](FoldersApi.md#DeleteFolder) | **Delete** /folders/{folderId} | Deleting a folder
[**GetAllFolders**](FoldersApi.md#GetAllFolders) | **Get** /folders | Getting all folders
[**GetFolder**](FoldersApi.md#GetFolder) | **Get** /folders/{folderId} | Getting a folder from its id
[**UpdateFolder**](FoldersApi.md#UpdateFolder) | **Put** /folders/{folderId} | Updating a folder



## CreateFolder

> Folder CreateFolder(ctx, optional)
Creating a folder

This method creates a new folder. In case of success it renders the created folder, otherwise an error (422 HTTP code).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateFolderOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **label** | **optional.String**| Folder’s label (mandatory). | 
 **parentId** | **optional.String**| ID of the parent folder. | 
 **position** | **optional.Int32**| position of the folder.  | 
 **query** | **optional.String**| query of the folder as described in ​Search API documentation.​\\n\\nExample: “​active_and_assigned_to_me:true” | 
 **renderThreadsCount** | **optional.Bool**| boolean describing display of the number of threads.  | 
 **roleRestrictionOnly** | [**optional.Interface of []string**](string.md)| list of roles allowed to see this folder. This parameter has to be a hash otherwise it will raise a 400 error. The key should be \&quot;only\&quot;. For example: &#x60;&amp;role_restriction[only][]&#x3D;4e5596cdae70f677b5000002&#x60; | 
 **teamRestrictionOnly** | [**optional.Interface of []string**](string.md)| list of teams allowed to see this folder. Same thing as role_restriction: team_restriction parameter has to be a hash with the key \&quot;only\&quot;. | 

### Return type

[**Folder**](Folder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFolder

> Folder DeleteFolder(ctx, folderId)
Deleting a folder

This method destroys an existing folder. It renders the folder itself. It renders a 404 if id is invalid.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string**|  | 

### Return type

[**Folder**](Folder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllFolders

> GetAllFoldersResponse GetAllFolders(ctx, optional)
Getting all folders

This method renders folders.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllFoldersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| The record index to start. Default value is 0. | 
 **limit** | **optional.Int32**| The max number of records to return. Default value is 30, max value is 150. | 

### Return type

[**GetAllFoldersResponse**](GetAllFoldersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFolder

> Folder GetFolder(ctx, folderId)
Getting a folder from its id

This method renders a folder from given id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string**|  | 

### Return type

[**Folder**](Folder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFolder

> Folder UpdateFolder(ctx, folderId, optional)
Updating a folder

This method updates an existing folder from given attributes and renders it in case of success.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string**|  | 
 **optional** | ***UpdateFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateFolderOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **label** | **optional.String**| Folder’s label (mandatory). | 
 **parentId** | **optional.String**| ID of the parent folder. | 
 **position** | **optional.Int32**| position of the folder.  | 
 **query** | **optional.String**| query of the folder as described in ​Search API documentation.​\\n\\nExample: “​active_and_assigned_to_me:true” | 
 **renderThreadsCount** | **optional.Bool**| boolean describing display of the number of threads.  | 
 **roleRestrictionOnly** | [**optional.Interface of []string**](string.md)| list of roles allowed to see this folder. This parameter has to be a hash otherwise it will raise a 400 error. The key should be \&quot;only\&quot;. For example: &#x60;&amp;role_restriction[only][]&#x3D;4e5596cdae70f677b5000002&#x60; | 
 **teamRestrictionOnly** | [**optional.Interface of []string**](string.md)| list of teams allowed to see this folder. Same thing as role_restriction: team_restriction parameter has to be a hash with the key \&quot;only\&quot;. | 

### Return type

[**Folder**](Folder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

