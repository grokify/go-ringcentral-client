/*
 * RingCentral Engage Voice API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package engagevoice

import (
	"time"
)

type DateTimeComparableSearchField struct {
	// Values can be `>`, `<` and `=`
	Operator string    `json:"operator,omitempty"`
	Value    time.Time `json:"value,omitempty"`
}
