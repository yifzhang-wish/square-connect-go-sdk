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

// Represents the transfer of a quantity of product inventory at a particular time from one location to another.
type InventoryTransfer struct {
	// A unique ID generated by Square for the `InventoryTransfer`.
	Id string `json:"id,omitempty"`
	// An optional ID provided by the application to tie the `InventoryTransfer` to an external system.
	ReferenceId string          `json:"reference_id,omitempty"`
	State       *InventoryState `json:"state,omitempty"`
	// The Square-generated ID of the [Location](entity:Location) where the related quantity of items was tracked before the transfer.
	FromLocationId string `json:"from_location_id,omitempty"`
	// The Square-generated ID of the [Location](entity:Location) where the related quantity of items was tracked after the transfer.
	ToLocationId string `json:"to_location_id,omitempty"`
	// The Square-generated ID of the [CatalogObject](entity:CatalogObject) being tracked.
	CatalogObjectId string `json:"catalog_object_id,omitempty"`
	// The [type](entity:CatalogObjectType) of the [CatalogObject](entity:CatalogObject) being tracked.Tracking is only supported for the `ITEM_VARIATION` type.
	CatalogObjectType string `json:"catalog_object_type,omitempty"`
	// The number of items affected by the transfer as a decimal string. Can support up to 5 digits after the decimal point.
	Quantity string `json:"quantity,omitempty"`
	// A client-generated RFC 3339-formatted timestamp that indicates when the transfer took place. For write actions, the `occurred_at` timestamp cannot be older than 24 hours or in the future relative to the time of the request.
	OccurredAt string `json:"occurred_at,omitempty"`
	// An RFC 3339-formatted timestamp that indicates when Square received the transfer request.
	CreatedAt string             `json:"created_at,omitempty"`
	Source    *SourceApplication `json:"source,omitempty"`
	// The Square-generated ID of the [Employee](entity:Employee) responsible for the inventory transfer.
	EmployeeId string `json:"employee_id,omitempty"`
}
