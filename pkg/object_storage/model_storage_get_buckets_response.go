/*
 * Object Storage
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package object_storage
// StorageGetBucketsResponse The buckets for the given stack
type StorageGetBucketsResponse struct {
	PageInfo PaginationPageInfo `json:"pageInfo,omitempty"`
	// The requested buckets
	Results []StorageBucket `json:"results,omitempty"`
}