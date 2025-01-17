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

// A refund of an [AdditionalRecipientReceivable](entity:AdditionalRecipientReceivable). This includes the ID of the additional recipient receivable associated to this object, as well as a reference to the [Refund](entity:Refund) that created this receivable refund.
type AdditionalRecipientReceivableRefund struct {
	// The receivable refund's unique ID, issued by Square payments servers.
	Id string `json:"id"`
	// The ID of the receivable that the refund was applied to.
	ReceivableId string `json:"receivable_id"`
	// The ID of the refund that is associated to this receivable refund.
	RefundId string `json:"refund_id"`
	// The ID of the location that created the receivable. This is the location ID on the associated transaction.
	TransactionLocationId string `json:"transaction_location_id"`
	AmountMoney           *Money `json:"amount_money"`
	// The time when the refund was created, in RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`
}
