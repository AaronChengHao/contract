package model

import "github.com/gogf/gf/v2/os/gtime"

type RechargeCreateInput struct {
	Id        uint        `json:"id"        ` //
	Address   string      `json:"address"   ` //
	GameCoin  int         `json:"gameCoin"  ` //
	TokenCoin int         `json:"tokenCoin" ` //
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
}
