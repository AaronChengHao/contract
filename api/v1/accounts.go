package v1

import "github.com/gogf/gf/v2/frame/g"

type AccountsShowReq struct {
	g.Meta  `path:"/accounts/show" method:"get"`
	Address string `json:"address" v:"required#地址必须输入"`
}

type AccountsShowRes struct {
	g.Meta  `mime:"json" example:"string"`
	Address interface{} `json:"address"`
	Coin    interface{} `json:"coin"`
}
