/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.54
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The topic of a Space, as selected by its creator.
type Topic struct {
	// The description of the given topic.
	Description string `json:"description,omitempty"`
	Id          string `json:"id"`
	// The name of the given topic.
	Name string `json:"name"`
}
