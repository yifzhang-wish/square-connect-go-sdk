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

// Return value from a [CompletePayment](#endpoint-payments-completepayment) call.
type CompletePaymentResponse struct {
	// Information on errors encountered during the request
	Errors  []ModelError `json:"errors,omitempty"`
	Payment *Payment     `json:"payment,omitempty"`
}