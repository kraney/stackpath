/*
 * Accounts and Users
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package accounts_and_users
// PolicyBinding An assignment of members to roles
type PolicyBinding struct {
	// The name of the role for this binding.  This is either roles/<name> or accounts/<id>/roles/<name>.
	Role string `json:"role,omitempty"`
	// The members assigned to the roles in this binding.  This is user:<name>.
	Members []string `json:"members,omitempty"`
}