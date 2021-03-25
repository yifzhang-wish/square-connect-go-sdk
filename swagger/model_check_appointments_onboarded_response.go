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

type CheckAppointmentsOnboardedResponse struct {
	// Indicates whether the seller has enabled the Square Appointments service (`true`) or not (`false`).
	AppointmentsOnboarded bool `json:"appointments_onboarded,omitempty"`
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}