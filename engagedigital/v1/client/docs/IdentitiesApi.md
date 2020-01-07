# \IdentitiesApi

All URIs are relative to *https://DOMAIN.api.engagement.dimelo.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllIdentities**](IdentitiesApi.md#GetAllIdentities) | **Get** /identities | Getting all identities
[**GetIdentity**](IdentitiesApi.md#GetIdentity) | **Get** /identities/{identityId} | Getting an identity from its id



## GetAllIdentities

> GetAllIdentitiesResponse GetAllIdentities(ctx, optional)
Getting all identities

This method renders identities ordered by creation date (descending). Only identities in sources where token’s user has “read” permission are returned.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllIdentitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllIdentitiesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **communityId** | **optional.String**| To filter identities on given community id. | 
 **identityGroupId** | **optional.String**| To filter on given group id. | 
 **userId** | **optional.String**| To filter identities on given user id. | 
 **sort** | **optional.String**| To change the criteria chosen to sort the identities. The value can be “created_at” or | 
 **foreignId** | **optional.String**| To filter identities on given user id | 
 **uuid** | **optional.String**| To filter identities on given uuid | 
 **offset** | **optional.Int32**| The record index to start. Default value is 0. | 
 **limit** | **optional.Int32**| The max number of records to return. Default value is 30, max value is 150. | 

### Return type

[**GetAllIdentitiesResponse**](GetAllIdentitiesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentity

> Identity GetIdentity(ctx, identityId)
Getting an identity from its id

This method renders an identity from given id. If token’s user does not have “read” on identity’s source community a 404 HTTP response will be returned.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityId** | **string**|  | 

### Return type

[**Identity**](Identity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

