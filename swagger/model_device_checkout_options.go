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

type DeviceCheckoutOptions struct {
	// The unique ID of the device intended for this `TerminalCheckout`. A list of `DeviceCode` objects can be retrieved from the /v2/devices/codes endpoint. Match a `DeviceCode.device_id` value with `device_id` to get the associated device code.
	DeviceId string `json:"device_id"`
	// Instructs the device to skip the receipt screen. Defaults to false.
	SkipReceiptScreen bool         `json:"skip_receipt_screen,omitempty"`
	TipSettings       *TipSettings `json:"tip_settings,omitempty"`
}
