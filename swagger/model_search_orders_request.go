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

// The request does not have any required fields. When given no query criteria, SearchOrders will return all results for all of the merchant’s locations. When fetching additional pages using a `cursor`, the `query` must be equal to the `query` used to fetch the first page of results.
type SearchOrdersRequest struct {
	// The location IDs for the orders to query. All locations must belong to the same merchant.  Min: 1 location IDs.  Max: 10 location IDs.
	LocationIds []string `json:"location_ids,omitempty"`
	// A pagination cursor returned by a previous call to this endpoint. Provide this to retrieve the next set of results for your original query. See [Pagination](https://developer.squareup.com/docs/basics/api101/pagination) for more information.
	Cursor string             `json:"cursor,omitempty"`
	Query  *SearchOrdersQuery `json:"query,omitempty"`
	// Maximum number of results to be returned in a single page. It is possible to receive fewer results than the specified limit on a given page.  Default: `500`
	Limit int32 `json:"limit,omitempty"`
	// Boolean that controls the format of the search results. If `true`, SearchOrders will return [`OrderEntry`](#type-orderentry) objects. If `false`, SearchOrders will return complete Order objects.  Default: `false`.
	ReturnEntries bool `json:"return_entries,omitempty"`
}