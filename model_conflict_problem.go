/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.54
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// You cannot create a new job if one is already in progress.
type ConflictProblem struct {
	Detail string `json:"detail,omitempty"`
	Status int32  `json:"status,omitempty"`
	Title  string `json:"title"`
	Type_  string `json:"type"`
}
