# MakeRingOutRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | [***MakeRingOutCallerInfoRequestFrom**](MakeRingOutCallerInfoRequestFrom.md) | Phone number of the caller. This number corresponds to the 1st leg of the RingOut call. This number can be one of user&#39;s configured forwarding numbers or arbitrary number | [default to null]
**To** | [***MakeRingOutCallerInfoRequestTo**](MakeRingOutCallerInfoRequestTo.md) | Phone number of the called party. This number corresponds to the 2nd leg of the RingOut call | [default to null]
**CallerId** | [***MakeRingOutCallerInfoRequestTo**](MakeRingOutCallerInfoRequestTo.md) | The number which will be displayed to the called party | [optional] [default to null]
**PlayPrompt** | **bool** | The audio prompt that the calling party hears when the call is connected | [optional] [default to null]
**Country** | [***MakeRingOutCoutryInfo**](MakeRingOutCoutryInfo.md) | Optional. Dialing plan country data. If not specified, then extension home country is applied by default | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


