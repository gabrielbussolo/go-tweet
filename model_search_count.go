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

// Represent a Search Count Result.
type SearchCount struct {
	End        *time.Time `json:"end"`
	Start      *time.Time `json:"start"`
	TweetCount int32      `json:"tweet_count"`
}
