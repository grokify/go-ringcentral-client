/*
 * Engage Digital API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package engagedigital

// The \"custom_status\" attribute represent the custom \"away\" status selected, it can either be: * 1. null — ​The agent is available or in the generic \"Away\" status 2. { \"id\":\"5582b1f4776562af9b000008\" } — ​The id of the custom status
type AgentCustomStatus struct {
	Id string `json:"id,omitempty"`
}
