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

// V1InventoryEntry
type V1InventoryEntry struct {
	// The variation that the entry corresponds to.
	VariationId string `json:"variation_id,omitempty"`
	// The current available quantity of the item variation.
	QuantityOnHand float64 `json:"quantity_on_hand,omitempty"`
}