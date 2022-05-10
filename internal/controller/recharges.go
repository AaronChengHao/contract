package controller

import (
	"context"
	v1 "contract/api/v1"
	"contract/internal/model"
	"contract/internal/service"
	"fmt"
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
		Address:   "xxx",
		GameCoin:  122222,
		TokenCoin: 222222222,
		CreatedAt: nil,
		UpdatedAt: nil,
	})
	fmt.Println(err)
	//dao.Recharges.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
	//	dao.Recharges.Ctx(ctx).Data().Insert()
	//})
	//return dao.Reply.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
	//	// 覆盖用户ID
	//	in.UserId = Context().Get(ctx).User.Id
	//	_, err := dao.Reply.Ctx(ctx).Data(in).Insert()
	//	if err == nil {
	//		err = Content().AddReplyCount(ctx, in.TargetId, 1)
	//	}
	//	return err
	//})
	g.RequestFromCtx(ctx).Response.Writeln("recharges store")
	return
}
