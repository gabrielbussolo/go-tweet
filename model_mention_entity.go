/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.54
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type MentionEntity struct {
	Id       string `json:"id,omitempty"`
	Username string `json:"username"`
	// Index (zero-based) at which position this entity ends.  The index is exclusive.
	End int32 `json:"end"`
	// Index (zero-based) at which position this entity starts.  The index is inclusive.
	Start int32 `json:"start"`
}
