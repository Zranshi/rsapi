package logic

import (
	"context"
	"rsapi/auth/rpc/account/account"
	"rsapi/auth/rpc/account/internal"
	"rsapi/auth/rpc/account/internal/svc"
	"rsapi/auth/rpc/token/tokenclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *account.LoginReq) (*account.LoginRes, error) {
	var (
		resp = new(account.LoginRes)
	)
	// check account
	if u, err := l.svcCtx.UserM.FindOneByEmail(l.ctx, in.Email); err != nil || u == nil {
		return resp, internal.ErrAccountNotExist
	} else {
		key, err := pwd2Key(in.Pwd, l.svcCtx.Config.Service.PasswordSalt)
		if err != nil {
			return resp, internal.ErrHashPwd
		}
		if u.Key != key {
			return resp, internal.ErrPwdNotMatch
		}

		rpcResp, err := l.svcCtx.TokenRpc.Generate(l.ctx, &tokenclient.GenerateReq{Email: in.Email, Expire: 3600 * 24})
		if err != nil {
			return resp, internal.ErrRpcCall
		}
		resp.Token = rpcResp.Token
		return resp, nil
	}
}
