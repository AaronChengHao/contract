package service

import (
	"context"
	"contract/internal/model"
	"contract/internal/model/entity"
	"contract/internal/service/internal/dao"
	"fmt"
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
		record, err := dao.Accounts.Ctx(ctx).Where("address=?", in.Address).One()
		if len(record) == 0 {
			dao.Accounts.Ctx(ctx).Insert(entity.Accounts{Address: in.Address, Coin: in.GameCoin})
		} else {
			_, _ = dao.Accounts.Ctx(ctx).Where("address=?", in.Address).Increment("coin", in.GameCoin)
		}
		fmt.Println(record, len(record), err)
		//if err == nil {
		//	err = Content().AddReplyCount(ctx, in.TargetId, 1)
		//}
		return err
	})
}
