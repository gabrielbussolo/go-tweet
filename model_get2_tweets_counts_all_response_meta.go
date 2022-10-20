/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.54
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Get2TweetsCountsAllResponseMeta struct {
	NewestId        string `json:"newest_id,omitempty"`
	NextToken       string `json:"next_token,omitempty"`
	OldestId        string `json:"oldest_id,omitempty"`
	TotalTweetCount int32  `json:"total_tweet_count,omitempty"`
}
