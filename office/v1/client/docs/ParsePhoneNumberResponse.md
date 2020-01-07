# ParsePhoneNumberResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | **string** | Canonical URI of a resource | [optional] 
**HomeCountry** | [**[]GetCountryInfoNumberParser**](GetCountryInfoNumberParser.md) | Information on a user home country | 
**PhoneNumbers** | [**[]PhoneNumberInfoNumberParser**](PhoneNumberInfoNumberParser.md) | Parsed phone numbers data | 
**OriginalString** | **string** | One of the numbers to be parsed, passed as a string in response | [optional] 
**AreaCode** | **string** | Area code of the location (3-digit usually), according to the NANP number format, that can be summarized as NPA-NXX-xxxx and covers Canada, the United States, parts of the Caribbean Sea, and some Atlantic and Pacific islands. See North American Numbering Plan for details | [optional] 
**FormattedNational** | **string** | Domestic format of a phone number | [optional] 
**FormattedInternational** | **string** | International format of a phone number | [optional] 
**Dialable** | **string** | Dialing format of a phone number | [optional] 
**E164** | **string** | E.164 (11-digits) format of a phone number | [optional] 
**Special** | **bool** |  True  if the number is in a special format (for example N11 code) | [optional] 
**Normalized** | **string** | E.164 (11-digits) format of a phone number without the plus sign (&#39;+&#39;) | [optional] 
**Country** | [**[]GetCountryInfoNumberParser**](GetCountryInfoNumberParser.md) | Information on a country the phone number belongs to | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


