# CreatePagerMessageRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | [**MessageStoreCallerInfoRequest**](MessageStoreCallerInfoRequest.md) |  | 
**ReplyOn** | **int32** | Internal identifier of a message this message replies to | [optional] 
**Text** | **string** | Text of a pager message. Max length is 1024 symbols (2-byte UTF-16 encoded). If a character is encoded in 4 bytes in UTF-16 it is treated as 2 characters, thus restricting the maximum message length to 512 symbols | 
**To** | [**[]MessageStoreCallerInfoRequest**](MessageStoreCallerInfoRequest.md) | Optional if replyOn parameter is specified. Receiver of a pager message. The extensionNumber property must be filled | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


