# DictionaryGreetingInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of a greeting | [optional] [default to null]
**Uri** | **string** | Link to a greeting | [optional] [default to null]
**Name** | **string** | Name of a greeting | [optional] [default to null]
**UsageType** | **string** | Usage type of a greeting, specifying if the greeting is applied for user extension or department extension &#x3D; [&#39;UserExtensionAnsweringRule&#39;, &#39;ExtensionAnsweringRule&#39;, &#39;DepartmentExtensionAnsweringRule&#39;, &#39;CompanyAnsweringRule&#39;, &#39;CompanyAfterHoursAnsweringRule&#39;] | [optional] [default to null]
**Text** | **string** | Text of a greeting, if any | [optional] [default to null]
**ContentUri** | **string** | Link to a greeting content (audio file), if any | [optional] [default to null]
**Type_** | **string** | Type of a greeting, specifying the case when the greeting is played. See Greeting Types &#x3D; [&#39;Introductory&#39;, &#39;Announcement&#39;, &#39;ConnectingMessage&#39;, &#39;ConnectingAudio&#39;, &#39;Voicemail&#39;, &#39;Unavailable&#39;, &#39;InterruptPrompt&#39;, &#39;HoldMusic&#39;, &#39;Company&#39;] | [optional] [default to null]
**Category** | **string** | Category of a greeting, specifying data form. The category value &#39;None&#39; specifies that greetings of a certain type (&#39;Introductory&#39;, &#39;ConnectingAudio&#39;, etc.) are switched off for an extension &#x3D; [&#39;Music&#39;, &#39;Message&#39;, &#39;Ring Tones&#39;, &#39;None&#39;] | [optional] [default to null]
**Navigation** | [***NavigationInfo**](NavigationInfo.md) | Information on navigation | [optional] [default to null]
**Paging** | [***PagingInfo**](PagingInfo.md) | Information on paging | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


