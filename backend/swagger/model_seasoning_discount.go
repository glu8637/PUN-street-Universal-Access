/*
 * PUN street Universal Access - OpenAPI 3.0
 *
 * pua
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SeasoningDiscount struct {
	DiscountId int64 `json:"discount_id"`

	DiscountName string `json:"discount_name"`

	DiscountDescription string `json:"discount_description"`

	DiscountStartDate string `json:"discount_start_date"`

	DiscountEndDate string `json:"discount_end_date"`

	DiscountPercentage int64 `json:"discount_percentage"`

	Status int64 `json:"status"`
}
