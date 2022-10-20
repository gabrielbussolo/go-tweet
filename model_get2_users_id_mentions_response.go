/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.54
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Get2UsersIdMentionsResponse struct {
	Data     []Tweet                          `json:"data,omitempty"`
	Errors   []Problem                        `json:"errors,omitempty"`
	Includes *Expansions                      `json:"includes,omitempty"`
	Meta     *Get2UsersIdMentionsResponseMeta `json:"meta,omitempty"`
}
