# \SettingsApi

All URIs are relative to *https://DOMAIN.api.engagement.dimelo.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllSettings**](SettingsApi.md#GetAllSettings) | **Get** /settings | Getting all settings
[**UpdateSettings**](SettingsApi.md#UpdateSettings) | **Put** /settings | Updating settings



## GetAllSettings

> Settings GetAllSettings(ctx, )
Getting all settings

This method renders all settings of your domain.  Authorization​: only users that can update settings.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Settings**](Settings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSettings

> Settings UpdateSettings(ctx, optional)
Updating settings

This method updates the current domain settings.  Authorization​: only users that can update settings.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSettingsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activityPresenceThreshold** | **optional.Int32**| (in hours). | 
 **activityTracking** | **optional.Bool**| ​Enable activity tracking (Boolean) | 
 **beginningOfWeek** | **optional.String**| (Day of week) | 
 **categoryTagging** | **optional.Bool**| A​ ctivate the forced categorization by source.​ (Boolean) | 
 **contentLanguages** | **optional.String**| (See format) | 
 **dashboard** | **optional.Bool**| Activate the dashboard (Boolean) | 
 **denyIframeIntegration** | **optional.Bool**| Prevent the DD to be embed by other websites (Boolean) | 
 **disablePasswordAutocomplete** | **optional.Bool**| (Boolean) | 
 **expirePasswordAfter** | **optional.Int32**| password expiration delay (in seconds) | 
 **expirePasswordEnabled** | **optional.Bool**| enable password expiration (Boolean) | 
 **exportInSeconds** | **optional.Bool**| provide durations in seconds in export (Boolean) | 
 **foldUselessContents** | **optional.Bool**| fold archived contents (Boolean) | 
 **fteDuration** | **optional.Int32**| FTE data period (in hours) | 
 **identityMerge** | **optional.Bool**| enable identity merge (Boolean) | 
 **interventionDeferRates** | [**optional.Interface of []int32**](int32.md)| (Array of times in seconds) | 
 **interventionDeferThreshold** | **optional.Int32**| (in seconds) | 
 **interventionRates** | [**optional.Interface of []int32**](int32.md)| (Array of times in seconds) | 
 **locale** | **optional.String**| locale code (String) | 
 **multiLang** | **optional.Bool**| activate multi language support for messages (Boolean) | 
 **name** | **optional.String**| Name of the Dimelo Digital (String) | 
 **passwordArchivableEnabled** | **optional.Bool**| prohibit reusing old passwords (Boolean) | 
 **passwordArchivableSize** | **optional.Int32**| number of archived passwords | 
 **passwordMinLength** | **optional.Int32**| minimum character length | 
 **passwordNonWord** | **optional.Bool**| should contain at least 1 non alphanumeric char (Boolean) | 
 **passwordNumbers** | **optional.Bool**| should contain at least 1 number (Boolean) | 
 **passwordRecoveryDisabled** | **optional.Bool**| disable password recovery by email (Boolean) | 
 **pushEnabled** | **optional.Bool**| Enable push mode (Boolean) | 
 **replyAsAnyIdentity** | **optional.Bool**| Enable reply as any identity (Boolean) | 
 **rtlSupport** | **optional.Bool**| Enable right to left support (Boolean) | 
 **selfApprovalRequired** | **optional.Bool**| ​Allow authors to ask approval of their messages (Boolean) | 
 **sessionTimeout** | **optional.Int32**| Session timeout (in minutes) | 
 **spellchecking** | **optional.Bool**| Enable spellchecking (Boolean) | 
 **style** | **optional.String**| Defines the DD’s design (String) | 
 **thirdPartyServicesDisabled** | **optional.Bool**| Disable third-party services (tracking...) (Boolean) | 
 **timezone** | **optional.String**| Use the timezone endpoint to get the timezone name (String) | 
 **trackJs** | **optional.Bool**| Track JS errors (Boolean) | 
 **type_** | **optional.String**| Can be ‘demo’, ‘production’ or ‘archived’ | 
 **urgentTaskThreshold** | **optional.Int32**| Chat max response time (in seconds) | 
 **useSystemFont** | **optional.Bool**| Experimental (Boolean) | 

### Return type

[**Settings**](Settings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

