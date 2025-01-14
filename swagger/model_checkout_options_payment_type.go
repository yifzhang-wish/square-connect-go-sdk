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

type CheckoutOptionsPaymentType string

// List of CheckoutOptionsPaymentType
const (
	PAYMENT_TYPE_DO_NOT_USE_CheckoutOptionsPaymentType     CheckoutOptionsPaymentType = "PAYMENT_TYPE_DO_NOT_USE"
	CARD_PRESENT_CheckoutOptionsPaymentType                CheckoutOptionsPaymentType = "CARD_PRESENT"
	MANUAL_CARD_ENTRY_CheckoutOptionsPaymentType           CheckoutOptionsPaymentType = "MANUAL_CARD_ENTRY"
	FELICA_ID_CheckoutOptionsPaymentType                   CheckoutOptionsPaymentType = "FELICA_ID"
	FELICA_QUICPAY_CheckoutOptionsPaymentType              CheckoutOptionsPaymentType = "FELICA_QUICPAY"
	FELICA_TRANSPORTATION_GROUP_CheckoutOptionsPaymentType CheckoutOptionsPaymentType = "FELICA_TRANSPORTATION_GROUP"
	FELICA_ALL_CheckoutOptionsPaymentType                  CheckoutOptionsPaymentType = "FELICA_ALL"
)
