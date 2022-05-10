package service

import (
	"context"
	"contract/internal/model"
	"contract/internal/service/internal/dao"
	"github.com/gogf/gf/v2/database/gdb"
)

type sRecharge struct{}

func Recharge() *sRecharge {
	return &sRecharge{}
}

// Create 创建回复
func (s *sRecharge) Create(ctx context.Context, in model.RechargeCreateInput) error {
	return dao.Recharges.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 覆盖用户ID
		//in.Id = 1111
		_, err := dao.Recharges.Ctx(ctx).Data(in).Insert()
		//if err == nil {
		//	err = Content().AddReplyCount(ctx, in.TargetId, 1)
		//}
		return err
	})
}
