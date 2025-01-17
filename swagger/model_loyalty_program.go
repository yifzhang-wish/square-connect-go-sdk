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

// Represents a Square loyalty program. Loyalty programs define how buyers can earn points and redeem points for rewards.  Square sellers can have only one loyalty program, which is created and managed from the Seller Dashboard.  For more information, see [Loyalty Program Overview](https://developer.squareup.com/docs/loyalty/overview).
type LoyaltyProgram struct {
	// The Square-assigned ID of the loyalty program. Updates to  the loyalty program do not modify the identifier.
	Id     string                `json:"id"`
	Status *LoyaltyProgramStatus `json:"status"`
	// The list of rewards for buyers, sorted by ascending points.
	RewardTiers      []LoyaltyProgramRewardTier      `json:"reward_tiers"`
	ExpirationPolicy *LoyaltyProgramExpirationPolicy `json:"expiration_policy,omitempty"`
	Terminology      *LoyaltyProgramTerminology      `json:"terminology"`
	// The [locations](entity:Location) at which the program is active.
	LocationIds []string `json:"location_ids"`
	// The timestamp when the program was created, in RFC 3339 format.
	CreatedAt string `json:"created_at"`
	// The timestamp when the reward was last updated, in RFC 3339 format.
	UpdatedAt string `json:"updated_at"`
	// Defines how buyers can earn loyalty points.
	AccrualRules []LoyaltyProgramAccrualRule `json:"accrual_rules"`
}
