package controller

import (
	"context"
	v1 "contract/api/v1"
	"contract/internal/model"
	"contract/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Recharges = cRecharges{}
)

type cRecharges struct{}

func (c *cRecharges) Index(ctx context.Context, req *v1.RechargesIndexReq) (res *v1.RechargesIndexRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("recharges index")
	return
}

func (c *cRecharges) Store(ctx context.Context, req *v1.RechargesStoreReq) (res *v1.RechargesStoreRes, err error) {
	err = service.Recharge().Create(ctx, model.RechargeCreateInput{
		//Id:        nil,
		Address:   req.Address,
		GameCoin:  req.GameCoin,
		TokenCoin: req.TokenCoin,
		CreatedAt: nil,
		UpdatedAt: nil,
	})
	return
}
