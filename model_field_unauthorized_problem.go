/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.54
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A problem that indicates that you are not allowed to see a particular field on a Tweet, User, etc.
type FieldUnauthorizedProblem struct {
	Detail       string `json:"detail,omitempty"`
	Status       int32  `json:"status,omitempty"`
	Title        string `json:"title"`
	Type_        string `json:"type"`
	Field        string `json:"field"`
	ResourceType string `json:"resource_type"`
	Section      string `json:"section"`
}
