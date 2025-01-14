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

// A [CatalogObject](entity:CatalogObject) instance of the `ITEM` type, also referred to as an item, in the catalog.
type CatalogItem struct {
	// The item's name. This is a searchable attribute for use in applicable query filters, its value must not be empty, and the length is of Unicode code points.
	Name *string `json:"name,omitempty"`
	// The item's description. This is a searchable attribute for use in applicable query filters, and its value length is of Unicode code points.
	Description *string `json:"description,omitempty"`
	// The text of the item's display label in the Square Point of Sale app. Only up to the first five characters of the string are used. This attribute is searchable, and its value length is of Unicode code points.
	Abbreviation string `json:"abbreviation,omitempty"`
	// The color of the item's display label in the Square Point of Sale app. This must be a valid hex color code.
	LabelColor string `json:"label_color,omitempty"`
	// If `true`, the item can be added to shipping orders from the merchant's online store.
	AvailableOnline bool `json:"available_online,omitempty"`
	// If `true`, the item can be added to pickup orders from the merchant's online store.
	AvailableForPickup bool `json:"available_for_pickup,omitempty"`
	// If `true`, the item can be added to electronically fulfilled orders from the merchant's online store.
	AvailableElectronically bool `json:"available_electronically,omitempty"`
	// The ID of the item's category, if any.
	CategoryId string `json:"category_id,omitempty"`
	// A set of IDs indicating the taxes enabled for this item. When updating an item, any taxes listed here will be added to the item. Taxes may also be added to or deleted from an item using `UpdateItemTaxes`.
	TaxIds []string `json:"tax_ids,omitempty"`
	// A set of `CatalogItemModifierListInfo` objects representing the modifier lists that apply to this item, along with the overrides and min and max limits that are specific to this item. Modifier lists may also be added to or deleted from an item using `UpdateItemModifierLists`.
	ModifierListInfo []CatalogItemModifierListInfo `json:"modifier_list_info,omitempty"`
	// __Retired__. The URL of an image representing this item. Retired in favor of `image_id` in [CatalogObject](entity:CatalogObject).
	ImageUrl string `json:"image_url,omitempty"`
	// A list of [CatalogItemVariation](entity:CatalogItemVariation) objects for this item. An item must have at least one variation.
	Variations  []CatalogObject         `json:"variations,omitempty"`
	ProductType *CatalogItemProductType `json:"product_type,omitempty"`
	// If `false`, the Square Point of Sale app will present the `CatalogItem`'s details screen immediately, allowing the merchant to choose `CatalogModifier`s before adding the item to the cart.  This is the default behavior.  If `true`, the Square Point of Sale app will immediately add the item to the cart with the pre-selected modifiers, and merchants can edit modifiers by drilling down onto the item's details.  Third-party clients are encouraged to implement similar behaviors.
	SkipModifierScreen bool `json:"skip_modifier_screen,omitempty"`
	// List of item options IDs for this item. Used to manage and group item variations in a specified order.  Maximum: 6 item options.
	ItemOptions []CatalogItemOptionForItem `json:"item_options,omitempty"`
	// hidden field
	EcomUri string `json:"ecom_uri,omitempty"`
	// hidden field
	EcomImageUris []string `json:"ecom_image_uris,omitempty"`
	// hidden field
	EcomAvailable bool `json:"ecom_available,omitempty"`
	// hidden field
	EcomVisibility string `json:"ecom_visibility,omitempty"`
	// A name to sort the item by. If this name is unspecified, namely, the `sort_name` field is absent, the regular `name` field is used for sorting.  It is currently supported for sellers of the Japanese locale only.
	SortName string `json:"sort_name,omitempty"`
}
