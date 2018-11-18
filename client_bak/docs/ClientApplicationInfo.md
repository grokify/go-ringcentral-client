# ClientApplicationInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detected** | **bool** | &#39;True&#39;, if the server succeeded detecting application info, sufficient to return provisioning info | 
**UserAgent** | **string** | The value of &#39;User-Agent&#39; header, as it was passed in request | [optional] 
**AppId** | **string** | Application identifier (from authorization session) | [optional] 
**AppName** | **string** | Application name (from authorization session, but must match &#39;User-Agent&#39;) | [optional] 
**AppVersion** | **string** | Application version (parsed from &#39;User-Agent&#39;) | [optional] 
**AppPlatform** | **string** | Application platform operation system (parsed from &#39;User-Agent&#39;: Windows, MacOS, Android, iOS | [optional] 
**AppPlatformVersion** | **string** | Application platform operation system version (parsed from &#39;User-Agent&#39;) | [optional] 
**Locale** | **string** | Locale, parsed from &#39;Accept-Language&#39;. Currently en-GB and en-US locales are supported. The default value is en-US | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


