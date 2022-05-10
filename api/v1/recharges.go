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
	g.Meta `path:"/recharges" method:"post" tags:"文章" summary:"展示Article详情页面" `
	Id     uint `in:"path" v:"min:1#请选择查看的内容" dc:"内容id"`
}
type RechargesStoreRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
