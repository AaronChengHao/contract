// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Recharges is the golang structure for table recharges.
type Recharges struct {
	Id        uint        `json:"id"        ` //
	Address   string      `json:"address"   ` //
	GameCoin  int         `json:"gameCoin"  ` //
	TokenCoin int         `json:"tokenCoin" ` //
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
}