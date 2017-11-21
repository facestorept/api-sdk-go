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

type Attribute struct {

	Id int64 `json:"id"`

	Position int64 `json:"position,omitempty"`

	IsSearchable bool `json:"is_searchable,omitempty"`

	Active bool `json:"active,omitempty"`

	Options *AttributeOptions `json:"options,omitempty"`

	I18n *I18n `json:"i18n,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
