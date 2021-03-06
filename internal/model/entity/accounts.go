// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Accounts is the golang structure for table accounts.
type Accounts struct {
	Id        uint        `json:"id"        ` //
	Address   string      `json:"address"   ` //
	Coin      int         `json:"coin"      ` //
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
}
