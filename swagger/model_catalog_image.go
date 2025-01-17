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

// An image file to use in Square catalogs. It can be associated with catalog items, item variations, and categories.
type CatalogImage struct {
	// The internal name to identify this image in calls to the Square API. This is a searchable attribute for use in applicable query filters using the [SearchCatalogObjects](api-endpoint:Catalog-SearchCatalogObjects). It is not unique and should not be shown in a buyer facing context.
	Name string `json:"name,omitempty"`
	// The URL of this image, generated by Square after an image is uploaded using the [CreateCatalogImage](api-endpoint:Catalog-CreateCatalogImage) endpoint.
	Url string `json:"url,omitempty"`
	// A caption that describes what is shown in the image. Displayed in the Square Online Store. This is a searchable attribute for use in applicable query filters using the [SearchCatalogObjects](api-endpoint:Catalog-SearchCatalogObjects).
	Caption string `json:"caption,omitempty"`
}
