# \InterventionsApi

All URIs are relative to *https://DOMAIN.api.engagement.dimelo.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelIntervention**](InterventionsApi.md#CancelIntervention) | **Delete** /interventions/{interventionId}/cancel | Cancelling an intervention
[**CategorizeIntervention**](InterventionsApi.md#CategorizeIntervention) | **Put** /interventions/{interventionId}/update_categories | Categorizing an intervention
[**CloseIntervention**](InterventionsApi.md#CloseIntervention) | **Put** /interventions/{interventionId}/close | Closing an intervention
[**CreateIntervention**](InterventionsApi.md#CreateIntervention) | **Post** /interventions | Creating an intervention
[**GetAllInterventions**](InterventionsApi.md#GetAllInterventions) | **Get** /interventions | Getting all interventions
[**GetIntervention**](InterventionsApi.md#GetIntervention) | **Get** /interventions/{interventionId} | Getting an intervention from its id
[**ReassignIntervention**](InterventionsApi.md#ReassignIntervention) | **Put** /interventions/{interventionId}/reassign | Reassigning an intervention



## CancelIntervention

> Intervention CancelIntervention(ctx, interventionId)
Cancelling an intervention

This method cancels (destroys) an intervention. It renders intervention itself. If token’s user does not have “read” on intervention’s source a 404 HTTP response will be returned.  Caveats:  * If the intervention is already being canceled, it will return a 409 error. * To be able to close an intervention, it must meet the following criteria otherwise a 403 will be raised:   * Intervention MUST NOT already be closed  * Intervention MUST NOT have agent replies   * Access-Token agent MUST have read access on the source  Authorization​: no, but it renders an error if intervention can’t be destroyed (see caveats).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interventionId** | **string**|  | 

### Return type

[**Intervention**](Intervention.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CategorizeIntervention

> Intervention CategorizeIntervention(ctx, interventionId, categoryIds)
Categorizing an intervention

This method updates the categories of an intervention. If token’s user does not have “read” on the intervention’s source a 404 HTTP response will be returned.  Authorization​: no.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interventionId** | **string**|  | 
**categoryIds** | [**[]string**](string.md)|  | 

### Return type

[**Intervention**](Intervention.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloseIntervention

> Intervention CloseIntervention(ctx, interventionId)
Closing an intervention

This method closes an intervention. Caveats:  * If the intervention is already being closed, it will return a 409 error. * To be able to close an intervention, it must meet the following criteria otherwise a 403 will be raised:    * Intervention MUST NOT already be closed    * Intervention MUST have agent replies   * Access-Token agent MUST be the owner of the intervention or have the permission to edit permissions    * Access-Token agent MUST have read access on the source  Authorization​: no, but it renders an error if intervention can’t be closed (see caveats)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interventionId** | **string**|  | 

### Return type

[**Intervention**](Intervention.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIntervention

> GetAllInterventionsResponse CreateIntervention(ctx, contentId)
Creating an intervention

This method creates a new intervention or reopen it. In case of success it renders the intervention, otherwise, it renders an error (422 HTTP code). This method opens intervention as access token’s user.  Authorization​: no, but it renders an error if intervention can’t be created or reopened (already opened, etc.).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentId** | **string**| The content to create intervention on (mandatory). | 

### Return type

[**GetAllInterventionsResponse**](GetAllInterventionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllInterventions

> GetAllInterventionsResponse GetAllInterventions(ctx, optional)
Getting all interventions

This method renders interventions ordered by creation date (descending). Only interventions in sources where token’s user has “read” permission are returned.  Authorization​: no.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllInterventionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllInterventionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **threadId** | **optional.String**| To filter interventions on given thread id. | 
 **userId** | **optional.String**| To filter interventions on given user id. | 
 **identityGroupId** | **optional.String**| To filter interventions on given identity_group_id. This will return interventions associated to any identity in the indentity_group. | 
 **identityId** | [**optional.Interface of []string**](string.md)| To filter interventions on given identity_id(s). Can be a single value or an array. | 
 **sort** | **optional.String**| To change the criteria chosen to sort the interventions. The value can be “created_at” or “updated_at”. | 
 **offset** | **optional.Int32**| The record index to start. Default value is 0. | 
 **limit** | **optional.Int32**| The max number of records to return. Default value is 30, max value is 150. | 

### Return type

[**GetAllInterventionsResponse**](GetAllInterventionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntervention

> Intervention GetIntervention(ctx, interventionId)
Getting an intervention from its id

This method renders an intervention from given id. If token’s user does not have “read” on intervention’s source a 404 HTTP response will be returned.  Authorization​: no.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interventionId** | **string**|  | 

### Return type

[**Intervention**](Intervention.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReassignIntervention

> Intervention ReassignIntervention(ctx, interventionId, userId)
Reassigning an intervention

This method updates the user in charge of the intervention  Authorization​: Only users who can update interventions.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interventionId** | **string**|  | 
**userId** | **string**|  | 

### Return type

[**Intervention**](Intervention.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

