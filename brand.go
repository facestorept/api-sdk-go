/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type Brand struct {

	Id int64 `json:"id"`

	Position int64 `json:"position,omitempty"`

	ImageSmall string `json:"image_small,omitempty"`

	ImageLarger string `json:"image_larger,omitempty"`

	Active bool `json:"active,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// channels that resource are showed
	Visibility []string `json:"visibility,omitempty"`

	// I18n fields to brands
	I18n []I18n `json:"i18n,omitempty"`
}
