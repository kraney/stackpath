/*
 * Accounts and Users
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package accounts_and_users
import (
	"time"
)
// IdentityUserCredential User API credentials  API credentials are used to generate authorization tokens for use with the StackPath API. Generating API credentials creates a client ID and secret. Client secrets are exposed to the user only once on creation and are not stored at StackPath. Please record your client secret after generating API credentials.  StackPath recommends using one set of API credentials per application and operating environment.
type IdentityUserCredential struct {
	// An API credential's unique identifier
	Id string `json:"id,omitempty"`
	// An API credential's name  API credential names are typically associated with the user's application and operating environment
	Name string `json:"name,omitempty"`
	// An API credential's OAuth2 client ID
	ClientId string `json:"clientId,omitempty"`
	// The date an API credential was created
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// The date an API credential was last updated
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}