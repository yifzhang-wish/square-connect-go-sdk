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

// An item variation (i.e., product) in the Catalog object model. Each item may have a maximum of 250 item variations.
type CatalogItemVariation struct {
	// The ID of the `CatalogItem` associated with this item variation.
	ItemId string `json:"item_id,omitempty"`
	// The item variation's name. This is a searchable attribute for use in applicable query filters, and its value length is of Unicode code points.
	Name string `json:"name,omitempty"`
	// The item variation's SKU, if any. This is a searchable attribute for use in applicable query filters.
	Sku string `json:"sku,omitempty"`
	// The universal product code (UPC) of the item variation, if any. This is a searchable attribute for use in applicable query filters.  The value of this attribute should be a number of 12-14 digits long.  This restriction is enforced on the Square Seller Dashboard, Square Point of Sale or Retail Point of Sale apps, where this attribute shows in the GTIN field. If a non-compliant UPC value is assigned to this attribute using the API, the value is not editable on the Seller Dashboard, Square Point of Sale or Retail Point of Sale apps unless it is updated to fit the expected format.
	Upc string `json:"upc,omitempty"`
	// The order in which this item variation should be displayed. This value is read-only. On writes, the ordinal for each item variation within a parent `CatalogItem` is set according to the item variations's position. On reads, the value is not guaranteed to be sequential or unique.
	Ordinal     int32               `json:"ordinal,omitempty"`
	PricingType *CatalogPricingType `json:"pricing_type,omitempty"`
	PriceMoney  *Money              `json:"price_money,omitempty"`
	// Per-location price and inventory overrides.
	LocationOverrides []ItemVariationLocationOverrides `json:"location_overrides,omitempty"`
	// If `true`, inventory tracking is active for the variation.
	TrackInventory     bool                `json:"track_inventory,omitempty"`
	InventoryAlertType *InventoryAlertType `json:"inventory_alert_type,omitempty"`
	// If the inventory quantity for the variation is less than or equal to this value and `inventory_alert_type` is `LOW_QUANTITY`, the variation displays an alert in the merchant dashboard.  This value is always an integer.
	InventoryAlertThreshold int64 `json:"inventory_alert_threshold,omitempty"`
	// Arbitrary user metadata to associate with the item variation. This attribute value length is of Unicode code points.
	UserData string `json:"user_data,omitempty"`
	// If the `CatalogItem` that owns this item variation is of type `APPOINTMENTS_SERVICE`, then this is the duration of the service in milliseconds. For example, a 30 minute appointment would have the value `1800000`, which is equal to 30 (minutes) * 60 (seconds per minute) * 1000 (milliseconds per second).
	ServiceDuration int64 `json:"service_duration,omitempty"`
	// If the `CatalogItem` that owns this item variation is of type `APPOINTMENTS_SERVICE`, a bool representing whether this service is available for booking.
	AvailableForBooking bool `json:"available_for_booking,omitempty"`
	// List of item option values associated with this item variation. Listed in the same order as the item options of the parent item.
	ItemOptionValues []CatalogItemOptionValueForItemVariation `json:"item_option_values,omitempty"`
	// ID of the ‘CatalogMeasurementUnit’ that is used to measure the quantity sold of this item variation. If left unset, the item will be sold in whole quantities.
	MeasurementUnitId string `json:"measurement_unit_id,omitempty"`
	// Whether stock is counted directly on this variation (TRUE) or only on its components (FALSE). For backward compatibility missing values will be interpreted as TRUE.
	Stockable bool `json:"stockable,omitempty"`
	// Tokens of employees that can perform the service represented by this variation. Only valid for variations of type `APPOINTMENTS_SERVICE`.
	TeamMemberIds       []string                `json:"team_member_ids,omitempty"`
	StockableConversion *CatalogStockConversion `json:"stockable_conversion,omitempty"`
}
