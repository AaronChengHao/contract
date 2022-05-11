package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RechargesIndexReq struct {
	g.Meta `path:"/recharges" tags:"index" method:"get" summary:"index api"`
}
type RechargesIndexRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
type RechargesStoreReq struct {
	g.Meta    `path:"/recharges" method:"post" tags:"文章" summary:"展示Article详情页面" `
	Address   string `json:"address"    v:"required#地址不能为空" dc:"地址"`
	GameCoin  int    `json:"gameCoin"    v:"required#游戏币不能为空"      dc:"游戏币"`
	TokenCoin int    `json:"tokenCoin"   v:"required#代币数量不能为空"      dc:"代币数量"`
}
type RechargesStoreRes struct {
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`
	Msg string `json:"msg"`
}
