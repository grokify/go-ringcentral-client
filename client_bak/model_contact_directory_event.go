/*
 * RingCentral Connect Platform API Explorer
 *
 * <p>This is a beta interactive API explorer for the RingCentral Connect Platform. To use this service, you will need to have an account with the proper credentials to generate an OAuth2 access token.</p><p><h2>Quick Start</h2></p><ol><li>1) Go to <b>Authentication > /oauth/token</b></li><li>2) Enter <b>app_key, app_secret, username, password</b> fields and then click \"Try it out!\"</li><li>3) Upon success, your access_token is loaded and you can access any form requiring authorization.</li></ol><h2>Links</h2><ul><li><a href=\"https://github.com/ringcentral\" target=\"_blank\">RingCentral SDKs on Github</a></li><li><a href=\"mailto:devsupport@ringcentral.com\">RingCentral Developer Support Email</a></li></ul>
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech) DO NOT EDIT.

package ringcentral

type ContactDirectoryEvent struct {
	// Internal identifier of an extension
	Id string `json:"id,omitempty"`
	// Type of change
	EventType string `json:"eventType,omitempty"`
	// Extension Type
	Type string `json:"type,omitempty"`
	// Status of an extension
	Status string `json:"status,omitempty"`
	// First name of an extension user
	FirstName string `json:"firstName,omitempty"`
	// Last name of an extension user
	LastName string `json:"lastName,omitempty"`
	// Department Name
	Department string `json:"department,omitempty"`
	// Email of an extension user
	Email string `json:"email,omitempty"`
	// Extension number
	ExtensionNumber string                           `json:"extensionNumber,omitempty"`
	Account         CompanyDirectoryAccountInfo      `json:"account,omitempty"`
	PhoneNumbers    CompanyDirectoryPhoneNumberInfo  `json:"phoneNumbers,omitempty"`
	Site            SiteInfo                         `json:"site,omitempty"`
	ProfileImage    CompanyDirectoryProfileImageInfo `json:"profileImage,omitempty"`
}