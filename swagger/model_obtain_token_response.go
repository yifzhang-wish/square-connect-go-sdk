/*
 * Square
 *
 * Use Square APIs to manage and run business including payment, customer, product, inventory, and employee management.
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ObtainTokenResponse struct {
	// A valid OAuth access token. OAuth access tokens are 64 bytes long. Provide the access token in a header with every request to Connect API endpoints. See [OAuth API: Walkthrough](https://developer.squareup.com/docs/oauth-api/walkthrough) for more information.
	AccessToken string `json:"access_token,omitempty"`
	// This value is always _bearer_.
	TokenType string `json:"token_type,omitempty"`
	// The date when access_token expires, in [ISO 8601](http://www.iso.org/iso/home/standards/iso8601.htm) format.
	ExpiresAt string `json:"expires_at,omitempty"`
	// The ID of the authorizing merchant's business.
	MerchantId string `json:"merchant_id,omitempty"`
	// __LEGACY FIELD__. The ID of a subscription plan the merchant signed up for. Only present if the merchant signed up for a subscription during authorization.
	SubscriptionId string `json:"subscription_id,omitempty"`
	// __LEGACY FIELD__. The ID of the subscription plan the merchant signed up for. Only present if the merchant signed up for a subscription during authorization.
	PlanId string `json:"plan_id,omitempty"`
	// Then OpenID token belonging to this this person. Only present if the OPENID scope is included in the authorize request.
	IdToken string `json:"id_token,omitempty"`
	// A refresh token. OAuth refresh tokens are 64 bytes long. For more information, see [OAuth access token management](https://developer.squareup.com/docs/oauth-api/how-it-works#oauth-access-token-management).
	RefreshToken string `json:"refresh_token,omitempty"`
	// A boolean indicating the access token is a short-lived access token. The short-lived access token returned in the response will expire in 24 hours.
	ShortLived bool `json:"short_lived,omitempty"`
}
