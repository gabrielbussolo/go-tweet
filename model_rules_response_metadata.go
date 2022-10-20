/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.54
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RulesResponseMetadata struct {
	NextToken string `json:"next_token,omitempty"`
	// Number of Rules in result set.
	ResultCount int32                `json:"result_count,omitempty"`
	Sent        string               `json:"sent"`
	Summary     *RulesRequestSummary `json:"summary,omitempty"`
}
