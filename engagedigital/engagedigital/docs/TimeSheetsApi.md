# \TimeSheetsApi

All URIs are relative to *https://DOMAIN.api.engagement.dimelo.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTimeSheet**](TimeSheetsApi.md#CreateTimeSheet) | **Post** /time_sheets | Creating a time sheet
[**DeleteTimeSheet**](TimeSheetsApi.md#DeleteTimeSheet) | **Delete** /time_sheets/{timeSheetId} | Deleting a time sheet
[**GetAllTimeSheets**](TimeSheetsApi.md#GetAllTimeSheets) | **Get** /time_sheets | Getting all time sheets
[**GetTimeSheet**](TimeSheetsApi.md#GetTimeSheet) | **Get** /time_sheets/{timeSheetId} | Getting a time sheet from its id
[**UpdateTimeSheet**](TimeSheetsApi.md#UpdateTimeSheet) | **Put** /time_sheets/{timeSheetId} | Updating a time sheet



## CreateTimeSheet

> TimeSheet CreateTimeSheet(ctx, label, optional)
Creating a time sheet

This method creates a time sheet. In case of success it renders the time sheet, otherwise, it renders an error (422 HTTP code).  Authorization​: only users that can create time sheet.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**label** | **string**| The label of the time sheet. | 
 **optional** | ***CreateTimeSheetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateTimeSheetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **active** | **optional.Bool**| true or false, this field is used to enable/disable a time sheet. | 
 **sourceIds** | [**optional.Interface of []string**](string.md)| An array containing id of each source using your time sheet. | 
 **holidaysRegion** | **optional.String**| A string containing the first two letters of your country (example: \&quot;fr\&quot;/\&quot;en\&quot;/\&quot;es\&quot;), useful to bootstrap default holidays following to a country. | 
 **holidays** | **optional.String**| An array containing one or more hash of holidays, a holiday must contain a name (string) and a date (string), the date must be in a valid format, a valid format is a format corresponding to your domain’s locale). | 
 **mondayHours** | **optional.String**| this field define the time intervals of the day (in secs). An empty string means that there are no business hours on this day. For example: “a-b,c-d”: “a” is the beginning of the first interval of the day, “b” is the ending of the first interval of the day, “c” is the beginning of the second interval of the day, “d” is the ending of the second interval of the day | 
 **tuesdayHours** | **optional.String**| this field define the time intervals of the day (in secs). An empty string means that there are no business hours on this day. See &#x60;monday_hours&#x60; for the format. | 
 **wednesdayHours** | **optional.String**| this field define the time intervals of the day (in secs). An empty string means that there are no business hours on this day. See &#x60;monday_hours&#x60; for the format. | 
 **thursdayHours** | **optional.String**| this field define the time intervals of the day (in secs). An empty string means that there are no business hours on this day. See &#x60;monday_hours&#x60; for the format. | 
 **fridayHours** | **optional.String**| this field define the time intervals of the day (in secs). An empty string means that there are no business hours on this day. See &#x60;monday_hours&#x60; for the format. | 
 **saturdayHours** | **optional.String**| this field define the time intervals of the day (in secs). An empty string means that there are no business hours on this day. See &#x60;monday_hours&#x60; for the format. | 
 **sundayHours** | **optional.String**| this field define the time intervals of the day (in secs). An empty string means that there are no business hours on this day. See &#x60;monday_hours&#x60; for the format. | 

### Return type

[**TimeSheet**](TimeSheet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTimeSheet

> TimeSheet DeleteTimeSheet(ctx, timeSheetId)
Deleting a time sheet

This method destroys an existing time sheet. It renders time sheet itself. It renders a 404 if id is invalid.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timeSheetId** | **string**|  | 

### Return type

[**TimeSheet**](TimeSheet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllTimeSheets

> GetAllTimeSheetsResponse GetAllTimeSheets(ctx, optional)
Getting all time sheets

This method renders time sheets ordered by active and label.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllTimeSheetsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllTimeSheetsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| The record index to start. Default value is 0. | 
 **limit** | **optional.Int32**| The max number of records to return. Default value is 30, max value is 150. | 

### Return type

[**GetAllTimeSheetsResponse**](GetAllTimeSheetsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeSheet

> TimeSheet GetTimeSheet(ctx, timeSheetId)
Getting a time sheet from its id

This method renders a time sheet from given id.  Authorization​: only users that can see time sheets in administration section.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timeSheetId** | **string**|  | 

### Return type

[**TimeSheet**](TimeSheet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTimeSheet

> TimeSheet UpdateTimeSheet(ctx, timeSheetId, optional)
Updating a time sheet

This method updates an existing team from given attributes and renders it in case of success.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timeSheetId** | **string**|  | 
 **optional** | ***UpdateTimeSheetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateTimeSheetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **active** | **optional.Bool**| true or false, this field is used to enable/disable a time sheet. | 
 **label** | **optional.String**| The label of the time sheet. | 
 **sourceIds** | [**optional.Interface of []string**](string.md)| An array containing id of each source using your time sheet. | 
 **holidays** | **optional.String**| An array containing one or more hash of holidays, a holiday must contain a name (string) and a date (string), the date must be in a valid format, a valid format is a format corresponding to your domain’s locale). | 
 **mondayHours** | **optional.String**| this field define the time intervals of the day (in secs). An empty string means that there are no business hours on this day. For example: “a-b,c-d”: “a” is the beginning of the first interval of the day, “b” is the ending of the first interval of the day, “c” is the beginning of the second interval of the day, “d” is the ending of the second interval of the day | 
 **tuesdayHours** | **optional.String**| this field define the time intervals of the day (in secs). An empty string means that there are no business hours on this day. See &#x60;monday_hours&#x60; for the format. | 
 **wednesdayHours** | **optional.String**| this field define the time intervals of the day (in secs). An empty string means that there are no business hours on this day. See &#x60;monday_hours&#x60; for the format. | 
 **thursdayHours** | **optional.String**| this field define the time intervals of the day (in secs). An empty string means that there are no business hours on this day. See &#x60;monday_hours&#x60; for the format. | 
 **fridayHours** | **optional.String**| this field define the time intervals of the day (in secs). An empty string means that there are no business hours on this day. See &#x60;monday_hours&#x60; for the format. | 
 **saturdayHours** | **optional.String**| this field define the time intervals of the day (in secs). An empty string means that there are no business hours on this day. See &#x60;monday_hours&#x60; for the format. | 
 **sundayHours** | **optional.String**| this field define the time intervals of the day (in secs). An empty string means that there are no business hours on this day. See &#x60;monday_hours&#x60; for the format. | 

### Return type

[**TimeSheet**](TimeSheet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

