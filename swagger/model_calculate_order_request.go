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

type CalculateOrderRequest struct {
	Order *Order `json:"order"`
	// Identifies one or more loyalty reward tiers to apply during order calculation. The discounts defined by the reward tiers are added to the order only to preview the effect of applying the specified reward(s). The reward(s) do not correspond to actual redemptions, that is, no `reward`s are created. Therefore, the reward `id`s are random strings used only to reference the reward tier.
	ProposedRewards []OrderReward `json:"proposed_rewards,omitempty"`
}