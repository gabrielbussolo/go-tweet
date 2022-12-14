/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.54
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Place struct {
	ContainedWithin []string `json:"contained_within,omitempty"`
	// The full name of the county in which this place exists.
	Country     string `json:"country,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	// The full name of this place.
	FullName string `json:"full_name"`
	Geo      *Geo   `json:"geo,omitempty"`
	Id       string `json:"id"`
	// The human readable name of this place.
	Name      string     `json:"name,omitempty"`
	PlaceType *PlaceType `json:"place_type,omitempty"`
}
