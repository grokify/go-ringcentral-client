/*
 * Engage Digital API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package engagedigital

// always contains text a​nd​ html versions.
type ContentBodyFormatted struct {
	Html string `json:"html,omitempty"`
	Text string `json:"text,omitempty"`
}