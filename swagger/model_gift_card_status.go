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

// GiftCardStatus : Indicates the gift card state.
type GiftCardStatus string

// List of GiftCardStatus
const (
	STATUS_DO_NOT_USE_GiftCardStatus GiftCardStatus = "STATUS_DO_NOT_USE"
	NOT_ACTIVE_GiftCardStatus        GiftCardStatus = "NOT_ACTIVE"
	ACTIVE_GiftCardStatus            GiftCardStatus = "ACTIVE"
	DEACTIVATED_GiftCardStatus       GiftCardStatus = "DEACTIVATED"
	BLOCKED_GiftCardStatus           GiftCardStatus = "BLOCKED"
	PENDING_GiftCardStatus           GiftCardStatus = "PENDING"
)