/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.54
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Specifies the type of attachments (if any) present in this Tweet.
type TweetAttachments struct {
	// A list of Media Keys for each one of the media attachments (if media are attached).
	MediaKeys []string `json:"media_keys,omitempty"`
	// A list of poll IDs (if polls are attached).
	PollIds []string `json:"poll_ids,omitempty"`
}
