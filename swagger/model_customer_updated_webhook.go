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

// Published when a [customer](entity:Customer) is updated. For more information, see [Use Customer Webhooks](https://developer.squareup.com/docs/customers-api/use-the-api/customer-webhooks).  Updates to the following customer fields do not invoke a `customer.updated` event: `cards` and `segment_ids`. In addition, the `customer` object in the event notification does not include these fields.
type CustomerUpdatedWebhook struct {
	// The ID of the seller associated with the event.
	MerchantId string `json:"merchant_id,omitempty"`
	// The type of event. For this object, the value is `customer.updated`.
	Type_ string `json:"type,omitempty"`
	// The unique ID of the event, which is used for [idempotency support](https://developer.squareup.com/docs/webhooks-api/what-it-does#idempotency-support).
	EventId string `json:"event_id,omitempty"`
	// The timestamp of when the event was created, in RFC 3339 format.
	CreatedAt string                      `json:"created_at,omitempty"`
	Data      *CustomerUpdatedWebhookData `json:"data,omitempty"`
}
