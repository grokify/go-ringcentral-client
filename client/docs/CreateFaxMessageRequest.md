# CreateFaxMessageRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | [**[]MessageStoreCallerInfoRequest**](MessageStoreCallerInfoRequest.md) | Recipient information. Phone number property is mandatory. Optional for resend fax request | 
**Resolution** | **string** | Fax resolution | [optional] 
**SendTime** | [**time.Time**](time.Time.md) | The datetime to send fax at, in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. If time is not specified, the fax will be send immediately | [optional] 
**CoverPageText** | **string** | Optional. Cover page text, entered by the fax sender and printed on the cover page. Maximum length is limited to 1024 symbols | [optional] 
**OriginalMessageId** | **string** | Internal identifier of the original fax message which needs to be resent. Mandatory for resend fax request | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


