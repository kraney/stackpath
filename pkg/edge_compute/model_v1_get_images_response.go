/*
 * Edge Compute
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package edge_compute
// V1GetImagesResponse A response from a request to retrieve images for a stack
type V1GetImagesResponse struct {
	PageInfo PaginationPageInfo `json:"pageInfo,omitempty"`
	// The requested images
	Results []V1Image `json:"results,omitempty"`
}