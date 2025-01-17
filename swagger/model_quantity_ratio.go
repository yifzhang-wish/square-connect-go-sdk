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

// A whole number or unreduced fractional ratio.
type QuantityRatio struct {
	// The whole or fractional quantity as the numerator.
	Quantity int32 `json:"quantity,omitempty"`
	// The whole or fractional quantity as the denominator.  In the case of fractional quantity this field is the denominator and quantity is the numerator. When unspecified, the value is `1`. For example, when `quantity=3` and `quantity_donominator` is unspecified, the quantity ratio is `3` or `3/1`.
	QuantityDenominator int32 `json:"quantity_denominator,omitempty"`
}
