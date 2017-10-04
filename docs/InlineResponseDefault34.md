# InlineResponseDefault34

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | **string** | Canonical URI of a resource | [optional] [default to null]
**HomeCountry** | [**[]ParsePhoneNumberCountryInfo**](ParsePhoneNumber.CountryInfo.md) | Information on a user home country | [optional] [default to null]
**PhoneNumbers** | [**[]ParsePhoneNumberPhoneNumberInfo**](ParsePhoneNumber.PhoneNumberInfo.md) | Parsed phone numbers data | [optional] [default to null]
**OriginalString** | **string** | One of the numbers to be parsed, passed as a string in response | [optional] [default to null]
**AreaCode** | **string** | Area code of the location (3-digit usually), according to the NANP number format, that can be summarized as NPA-NXX-xxxx and covers Canada, the United States, parts of the Caribbean Sea, and some Atlantic and Pacific islands. See North American Numbering Plan for details | [optional] [default to null]
**FormattedNational** | **string** | Domestic format of a phone number | [optional] [default to null]
**FormattedInternational** | **string** | International format of a phone number | [optional] [default to null]
**Dialable** | **string** | Dialing format of a phone number | [optional] [default to null]
**E164** | **string** | E.164 (11-digits) format of a phone number | [optional] [default to null]
**Special** | **bool** | \&quot;True\&quot; if the number is in a special format (for example N11 code) | [optional] [default to null]
**Normalized** | **string** | E.164 (11-digits) format of a phone number without the plus sign (&#39;+&#39;) | [optional] [default to null]
**Country** | [**[]ParsePhoneNumberCountryInfo**](ParsePhoneNumber.CountryInfo.md) | Information on a country the phone number belongs to | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


