package controller

import (
	"context"
	v1 "contract/api/v1"
	"contract/internal/service"
)

var (
	Accounts = cAccounts{}
)

type cAccounts struct {
}

func (c *cAccounts) Show(ctx context.Context, req *v1.AccountsShowReq) (res *v1.AccountsShowRes, err error) {
	record, err := service.Account().Show(ctx, req.Address)
	if len(record) > 0 {
		return &v1.AccountsShowRes{Address: record["address"], Coin: record["coin"]}, nil
	}
	return &v1.AccountsShowRes{Address: req.Address, Coin: 0}, nil
}
