/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.54
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AnimatedGif struct {
	Height          int32      `json:"height,omitempty"`
	MediaKey        string     `json:"media_key,omitempty"`
	Type_           string     `json:"type"`
	Width           int32      `json:"width,omitempty"`
	PreviewImageUrl string     `json:"preview_image_url,omitempty"`
	Variants        *[]Variant `json:"variants,omitempty"`
}