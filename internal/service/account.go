package service

import (
	"context"
	"contract/internal/service/internal/dao"
	"github.com/gogf/gf/v2/database/gdb"
)

type sAccount struct {
}

func Account() *sAccount {
	return &sAccount{}
}

func (s *sAccount) Show(ctx context.Context, address string) (gdb.Record, error) {
	record, err := dao.Accounts.Ctx(ctx).Where("address=?", address).One()
	return record, err
}
