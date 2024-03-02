package logic

import (
	"context"

	"rsapi/auth/rpc/account/account"
	"rsapi/auth/rpc/account/internal/svc"
	"rsapi/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InfoLogic) Info(in *account.InfoReq) (*account.InfoRes, error) {
	resp := new(account.InfoRes)
	if u, err := l.svcCtx.UserM.FindOneByEmail(l.ctx, in.Email); err != nil {
		return resp, util.DbErr(err)
	} else {
		resp.Name = u.Name
	}
	return resp, nil
}
