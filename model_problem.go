/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.54
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// An HTTP Problem Details object, as defined in IETF RFC 7807 (https://tools.ietf.org/html/rfc7807).
type Problem struct {
	Detail string `json:"detail,omitempty"`
	Status int32  `json:"status,omitempty"`
	Title  string `json:"title"`
	Type_  string `json:"type"`
}
