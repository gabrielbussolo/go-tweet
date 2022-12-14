/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.54
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A [GeoJson Point](https://tools.ietf.org/html/rfc7946#section-3.1.2) geometry object.
type Point struct {
	Coordinates *[]float64 `json:"coordinates"`
	Type_       string     `json:"type"`
}
