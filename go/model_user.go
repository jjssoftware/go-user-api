/*
 * User API
 *
 * An API for managing system users
 *
 * API version: 1.0.3
 * Contact: joe@jjssoftware.co.uk
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

import (
	"time"
)

type User struct {
	Id int32 `json:"id"`

	UniqueId string `json:"uniqueId"`

	CreateTimestamp time.Time `json:"createTimestamp"`

	UserName string `json:"userName"`

	FirstName string `json:"firstName"`

	Surname string `json:"surname"`
}
