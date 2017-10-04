# Body

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseType** | **string** | Must be set to code | [optional] [default to null]
**ClientId** | **string** | Required. Enter your application key (Production or Sandbox) here | [optional] [default to null]
**RedirectUri** | **string** | Required. This is a callback URI which determines where the response will be sent to. The value of this parameter must exactly match one of the URIs you have provided for your app upon registration. This URI can be HTTP/HTTPS address for web applications or custom scheme URI for mobile or desktop applications. | [optional] [default to null]
**State** | **string** | Optional, recommended. An opaque value used by the client to maintain state between the request and callback. The authorization server includes this value when redirecting the user-agent back to the client. The parameter should be used for preventing cross-site request forgery | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


