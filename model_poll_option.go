/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.54
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Describes a choice in a Poll object.
type PollOption struct {
	Label string `json:"label"`
	// Position of this choice in the poll.
	Position int32 `json:"position"`
	// Number of users who voted for this choice.
	Votes int32 `json:"votes"`
}
