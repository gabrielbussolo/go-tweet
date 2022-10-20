/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.54
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// ReplySettings : Shows who can reply a Tweet. Fields returned are everyone, mentioned_users, and following.
type ReplySettings string

// List of ReplySettings
const (
	EVERYONE_ReplySettings        ReplySettings = "everyone"
	MENTIONED_USERS_ReplySettings ReplySettings = "mentionedUsers"
	FOLLOWING_ReplySettings       ReplySettings = "following"
	OTHER_ReplySettings           ReplySettings = "other"
)
