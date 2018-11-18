/*
 * RingCentral Connect Platform API Explorer
 *
 * <p>This is a beta interactive API explorer for the RingCentral Connect Platform. To use this service, you will need to have an account with the proper credentials to generate an OAuth2 access token.</p><p><h2>Quick Start</h2></p><ol><li>1) Go to <b>Authentication > /oauth/token</b></li><li>2) Enter <b>app_key, app_secret, username, password</b> fields and then click \"Try it out!\"</li><li>3) Upon success, your access_token is loaded and you can access any form requiring authorization.</li></ol><h2>Links</h2><ul><li><a href=\"https://github.com/ringcentral\" target=\"_blank\">RingCentral SDKs on Github</a></li><li><a href=\"mailto:devsupport@ringcentral.com\">RingCentral Developer Support Email</a></li></ul>
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech) DO NOT EDIT.

package ringcentral

type ParsePhoneNumberResponse struct {
	// Canonical URI of a resource
	Uri string `json:"uri,omitempty"`
	// Information on a user home country
	HomeCountry []GetCountryInfoNumberParser `json:"homeCountry"`
	// Parsed phone numbers data
	PhoneNumbers []PhoneNumberInfoNumberParser `json:"phoneNumbers"`
	// One of the numbers to be parsed, passed as a string in response
	OriginalString string `json:"originalString,omitempty"`
	// Area code of the location (3-digit usually), according to the NANP number format, that can be summarized as NPA-NXX-xxxx and covers Canada, the United States, parts of the Caribbean Sea, and some Atlantic and Pacific islands. See North American Numbering Plan for details
	AreaCode string `json:"areaCode,omitempty"`
	// Domestic format of a phone number
	FormattedNational string `json:"formattedNational,omitempty"`
	// International format of a phone number
	FormattedInternational string `json:"formattedInternational,omitempty"`
	// Dialing format of a phone number
	Dialable string `json:"dialable,omitempty"`
	// E.164 (11-digits) format of a phone number
	E164 string `json:"e164,omitempty"`
	//  True  if the number is in a special format (for example N11 code)
	Special bool `json:"special,omitempty"`
	// E.164 (11-digits) format of a phone number without the plus sign ('+')
	Normalized string `json:"normalized,omitempty"`
	// Information on a country the phone number belongs to
	Country []GetCountryInfoNumberParser `json:"country,omitempty"`
}