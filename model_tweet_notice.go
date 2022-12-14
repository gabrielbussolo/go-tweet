/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.54
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"time"
)

type TweetNotice struct {
	// If the label is being applied or removed. Possible values are ‘apply’ or ‘remove’.
	Application string `json:"application"`
	// Information shown on the Tweet label
	Details string `json:"details,omitempty"`
	// Event time.
	EventAt time.Time `json:"event_at"`
	// The type of label on the Tweet
	EventType string `json:"event_type"`
	// Link to more information about this kind of label
	ExtendedDetailsUrl string `json:"extended_details_url,omitempty"`
	// Title/header of the Tweet label
	LabelTitle string                      `json:"label_title,omitempty"`
	Tweet      *TweetComplianceSchemaTweet `json:"tweet"`
}
