/*
 * Object Storage
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package object_storage
// MetricsData The data points in a metrics collection
type MetricsData struct {
	Matrix DataMatrix `json:"matrix,omitempty"`
	Vector DataVector `json:"vector,omitempty"`
}