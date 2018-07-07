# CreateSmsMessage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | [**MessageStoreCallerInfoRequest**](MessageStoreCallerInfoRequest.md) |  | 
**To** | [**[]MessageStoreCallerInfoRequest**](MessageStoreCallerInfoRequest.md) | Receiver of an SMS message. The phoneNumber property must be filled | 
**Text** | **string** | Text of a message. Max length is 1000 symbols (2-byte UTF-16 encoded). If a character is encoded in 4 bytes in UTF-16 it is treated as 2 characters, thus restricting the maximum message length to 500 symbols | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


