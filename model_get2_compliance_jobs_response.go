/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.54
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Get2ComplianceJobsResponse struct {
	Data   []ComplianceJob                 `json:"data,omitempty"`
	Errors []Problem                       `json:"errors,omitempty"`
	Meta   *Get2ComplianceJobsResponseMeta `json:"meta,omitempty"`
}
