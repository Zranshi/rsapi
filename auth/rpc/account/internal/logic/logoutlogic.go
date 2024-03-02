package logic

import (
	"context"

	"rsapi/auth/rpc/account/account"
	"rsapi/auth/rpc/account/internal/svc"
	"rsapi/auth/rpc/token/tokenclient"
	"rsapi/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogoutLogic) Logout(in *account.LogoutReq) (*account.LogoutRes, error) {
	resp := new(account.LogoutRes)
	_, err := l.svcCtx.TokenRpc.Invalid(l.ctx, &tokenclient.InvalidReq{Email: &in.Email})
	if err != nil {
		return resp, util.RpcCallErr(err, "token.rpc")
	}
	return resp, nil
}
