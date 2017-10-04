# Body14

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | [**RingOutRequestFrom**](RingOut.Request.From.md) | Phone number of the caller. This number corresponds to the 1st leg of the RingOut call. This number can be one of user&#39;s configured forwarding numbers or arbitrary number | [optional] [default to null]
**To** | [**RingOutRequestTo**](RingOut.Request.To.md) | Phone number of the called party. This number corresponds to the 2nd leg of the RingOut call | [optional] [default to null]
**CallerId** | [**RingOutRequestTo**](RingOut.Request.To.md) | The number which will be displayed to the called party | [optional] [default to null]
**PlayPrompt** | **bool** | The audio prompt that the calling party hears when the call is connected | [optional] [default to null]
**Country** | [**RingOutRequestCountryInfo**](RingOut.Request.CountryInfo.md) | Optional. Dialing plan country data. If not specified, then extension home country is applied by default | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


