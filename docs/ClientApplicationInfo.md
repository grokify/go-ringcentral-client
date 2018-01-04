# ClientApplicationInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detected** | **bool** | &#39;True&#39;, if the server succeeded detecting application info, sufficient to return provisioning info | [default to null]
**UserAgent** | **string** | The value of &#39;User-Agent&#39; header, as it was passed in request | [optional] [default to null]
**AppId** | **string** | Application identifier (from authorization session) | [optional] [default to null]
**AppName** | **string** | Application name (from authorization session, but must match &#39;User-Agent&#39;) | [optional] [default to null]
**AppVersion** | **string** | Application version (parsed from &#39;User-Agent&#39;) | [optional] [default to null]
**AppPlatform** | **string** | Application platform operation system (parsed from &#39;User-Agent&#39;: Windows, MacOS, Android, iOS | [optional] [default to null]
**AppPlatformVersion** | **string** | Application platform operation system version (parsed from &#39;User-Agent&#39;) | [optional] [default to null]
**Locale** | **string** | Locale, parsed from &#39;Accept-Language&#39;. Currently en-GB and en-US locales are supported. The default value is en-US | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


