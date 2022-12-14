/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.54
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A problem that indicates a particular Tweet, User, etc. is not available to you.
type ResourceUnavailableProblem struct {
	Detail       string `json:"detail,omitempty"`
	Status       int32  `json:"status,omitempty"`
	Title        string `json:"title"`
	Type_        string `json:"type"`
	Parameter    string `json:"parameter"`
	ResourceId   string `json:"resource_id"`
	ResourceType string `json:"resource_type"`
}
