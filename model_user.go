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

// The Twitter User object.
type User struct {
	// Creation time of this User.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// The text of this User's profile description (also known as bio), if the User provided one.
	Description string        `json:"description,omitempty"`
	Entities    *UserEntities `json:"entities,omitempty"`
	Id          string        `json:"id"`
	// The location specified in the User's profile, if the User provided one. As this is a freeform value, it may not indicate a valid location, but it may be fuzzily evaluated when performing searches with location queries.
	Location string `json:"location,omitempty"`
	// The friendly name of this User, as shown on their profile.
	Name          string `json:"name"`
	PinnedTweetId string `json:"pinned_tweet_id,omitempty"`
	// The URL to the profile image for this User.
	ProfileImageUrl string `json:"profile_image_url,omitempty"`
	// Indicates if this User has chosen to protect their Tweets (in other words, if this User's Tweets are private).
	Protected     bool               `json:"protected,omitempty"`
	PublicMetrics *UserPublicMetrics `json:"public_metrics,omitempty"`
	// The URL specified in the User's profile.
	Url      string `json:"url,omitempty"`
	Username string `json:"username"`
	// Indicate if this User is a verified Twitter User.
	Verified bool          `json:"verified,omitempty"`
	Withheld *UserWithheld `json:"withheld,omitempty"`
}