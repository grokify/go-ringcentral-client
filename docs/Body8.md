# Body8

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | [**[]CallerInfo**](CallerInfo.md) | Recipient information. Phone number property is mandatory. Optional for resend fax request | [optional] [default to null]
**FaxResolution** | **string** | Fax resolution | [optional] [default to null]
**SendTime** | [**time.Time**](time.Time.md) | The datetime to send fax at, in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. If time is not specified, the fax will be send immediately | [optional] [default to null]
**CoverIndex** | **int32** | Optional. Cover page index. If not specified, the default cover page which is configured in \&quot;Outbound Fax Settings\&quot; is attached. See also &#39;Available Cover Pages&#39; table below | [optional] [default to null]
**CoverPageText** | **string** | Optional. Cover page text, entered by the fax sender and printed on the cover page. Maximum length is limited to 1024 symbols | [optional] [default to null]
**OriginalMessageId** | **string** | Internal identifier of the original fax message which needs to be resent. Mandatory for resend fax request | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


