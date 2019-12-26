# \IdentityGroupsApi

All URIs are relative to *https://DOMAIN.api.engagement.dimelo.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllIdentityGroups**](IdentityGroupsApi.md#GetAllIdentityGroups) | **Get** /identity_groups | Getting all identity groups
[**GetIdentityGroup**](IdentityGroupsApi.md#GetIdentityGroup) | **Get** /identity_groups/{identityGroupId} | Getting an identity group from its id
[**UpdateIdentityGroup**](IdentityGroupsApi.md#UpdateIdentityGroup) | **Put** /identity_groups/{identityGroupId} | Updating an identity group



## GetAllIdentityGroups

> GetAllIdentityGroupsResponse GetAllIdentityGroups(ctx, optional)
Getting all identity groups

This method renders identity groups ordered by creation date (descending). Note that identity_group are created in a lazily only when data are manually added to an identity OR a two identity are merged altogether. That means that some identity DON’T have identity_group, and identity_group do not cover all identities.  Authorization​: no.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllIdentityGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllIdentityGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firstname** | **optional.String**| To filter groups on given firstname. | 
 **lastname** | **optional.String**| To filter groups on given lastname. | 
 **email** | **optional.String**| To filter groups that have given email. | 
 **uuid** | **optional.String**| To filter groups that have given uuid. | 
 **sort** | **optional.String**| To change the criteria chosen to sort the identities. The value can be “created_at” or “updated_at”. | 
 **offset** | **optional.Int32**| The record index to start. Default value is 0. | 
 **limit** | **optional.Int32**| The max number of records to return. Default value is 30, max value is 150. | 

### Return type

[**GetAllIdentityGroupsResponse**](GetAllIdentityGroupsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityGroup

> IdentityGroup GetIdentityGroup(ctx, identityGroupId)
Getting an identity group from its id

This method renders an identity group from given id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityGroupId** | **string**|  | 

### Return type

[**IdentityGroup**](IdentityGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdentityGroup

> IdentityGroup UpdateIdentityGroup(ctx, identityGroupId, optional)
Updating an identity group

This method updates an identity group from given attributes and renders it in case of success.  Authorization​: no.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityGroupId** | **string**|  | 
 **optional** | ***UpdateIdentityGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateIdentityGroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **company** | **optional.String**| Identity company. | 
 **customFieldValuesCustomFieldKey** | **optional.String**| Identity custom field with key « custom_field_key ». It | 
 **emails** | [**optional.Interface of []string**](string.md)| Identity emails (multiple). | 
 **firstname** | **optional.String**| Identity firstname. | 
 **gender** | **optional.String**| Identity’s gender. It can be \&quot;man\&quot;, \&quot;woman\&quot; or empty. | 
 **homePhones** | [**optional.Interface of []string**](string.md)| Identity home phones (mutiple). | 
 **lastname** | **optional.String**| Identity lastname. | 
 **mobilePhones** | [**optional.Interface of []string**](string.md)| Identity mobile phones (multiple). | 
 **notes** | **optional.String**| Identity notes. | 
 **tagIds** | [**optional.Interface of []string**](string.md)| Identity tag ids (multiple). | 

### Return type

[**IdentityGroup**](IdentityGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

