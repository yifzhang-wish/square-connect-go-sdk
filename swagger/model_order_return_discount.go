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

// Represents a discount being returned that applies to one or more return line items in an order.  Fixed-amount, order-scoped discounts are distributed across all non-zero return line item totals. The amount distributed to each return line item is relative to that item’s contribution to the order subtotal.
type OrderReturnDiscount struct {
	// A unique ID that identifies the returned discount only within this order.
	Uid string `json:"uid,omitempty"`
	// The discount `uid` from the order that contains the original application of this discount.
	SourceDiscountUid string `json:"source_discount_uid,omitempty"`
	// The catalog object ID referencing [CatalogDiscount](entity:CatalogDiscount).
	CatalogObjectId string `json:"catalog_object_id,omitempty"`
	// The discount's name.
	Name  string                     `json:"name,omitempty"`
	Type_ *OrderLineItemDiscountType `json:"type,omitempty"`
	// The percentage of the tax, as a string representation of a decimal number. A value of `\"7.25\"` corresponds to a percentage of 7.25%.  `percentage` is not set for amount-based discounts.
	Percentage   string                      `json:"percentage,omitempty"`
	AmountMoney  *Money                      `json:"amount_money,omitempty"`
	AppliedMoney *Money                      `json:"applied_money,omitempty"`
	Scope        *OrderLineItemDiscountScope `json:"scope,omitempty"`
}
