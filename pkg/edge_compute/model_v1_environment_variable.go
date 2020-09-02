/*
 * Edge Compute
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package edge_compute
// V1EnvironmentVariable The location to obtain a value for an environment variable
type V1EnvironmentVariable struct {
	// An environment variable's value
	Value string `json:"value,omitempty"`
	// A sensitive environment variable that should not be exposed
	SecretValue string `json:"secretValue,omitempty"`
}