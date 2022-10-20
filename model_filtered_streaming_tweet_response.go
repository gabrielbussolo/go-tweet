/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.54
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A Tweet or error that can be returned by the streaming Tweet API. The values returned with a successful streamed Tweet includes the user provided rules that the Tweet matched.
type FilteredStreamingTweetResponse struct {
	Data     *Tweet      `json:"data,omitempty"`
	Errors   []Problem   `json:"errors,omitempty"`
	Includes *Expansions `json:"includes,omitempty"`
	// The list of rules which matched the Tweet
	MatchingRules []FilteredStreamingTweetResponseMatchingRules `json:"matching_rules,omitempty"`
}