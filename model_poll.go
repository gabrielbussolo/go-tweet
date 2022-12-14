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

// Represent a Poll attached to a Tweet.
type Poll struct {
	DurationMinutes int32        `json:"duration_minutes,omitempty"`
	EndDatetime     time.Time    `json:"end_datetime,omitempty"`
	Id              string       `json:"id"`
	Options         []PollOption `json:"options"`
	VotingStatus    string       `json:"voting_status,omitempty"`
}
