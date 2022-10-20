/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.54
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Engagement metrics for the Tweet at the time of the request.
type TweetPublicMetrics struct {
	// Number of times this Tweet has been liked.
	LikeCount int32 `json:"like_count"`
	// Number of times this Tweet has been quoted.
	QuoteCount int32 `json:"quote_count,omitempty"`
	// Number of times this Tweet has been replied to.
	ReplyCount int32 `json:"reply_count"`
	// Number of times this Tweet has been Retweeted.
	RetweetCount int32 `json:"retweet_count"`
}