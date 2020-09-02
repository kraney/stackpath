/*
 * Edge Compute Networking
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package edge_compute_networking
// PaginationPageInfo Information about a paginated response  This is modeled after the GraphQL Relay spec to support both cursor based pagination and traditional offset based pagination.
type PaginationPageInfo struct {
	// The total number of items in the dataset
	TotalCount string `json:"totalCount,omitempty"`
	// Whether or not a previous page of data exists
	HasPreviousPage bool `json:"hasPreviousPage,omitempty"`
	// Whether or not another page of data is available
	HasNextPage bool `json:"hasNextPage,omitempty"`
	// The cursor for the first item in the set of data returned
	StartCursor string `json:"startCursor,omitempty"`
	// The cursor for the last item in the set of data returned
	EndCursor string `json:"endCursor,omitempty"`
}