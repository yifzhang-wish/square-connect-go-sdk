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

// Response object returned by the [CreateLocation](api-endpoint:Locations-CreateLocation) endpoint.
type CreateLocationResponse struct {
	// Information on errors encountered during the request.
	Errors   []ModelError `json:"errors,omitempty"`
	Location *Location    `json:"location,omitempty"`
}
