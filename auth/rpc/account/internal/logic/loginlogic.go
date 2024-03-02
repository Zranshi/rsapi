package logic

import (
	"context"
	"rsapi/auth/rpc/account/account"
	"rsapi/auth/rpc/account/internal/svc"
	"rsapi/auth/rpc/token/tokenclient"
	"rsapi/util"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		return resp, util.DbErr(err)
	} else {
		key, err := pwd2Key(in.Pwd, l.svcCtx.Config.Service.PasswordSalt)
		if err != nil {
			return resp, util.ParseErr(err)
		}
		if u.Key != key {
			return resp, status.Error(codes.Unauthenticated, "pwd not match")
		}

		rpcResp, err := l.svcCtx.TokenRpc.Generate(l.ctx, &tokenclient.GenerateReq{Email: in.Email, Expire: 3600 * 24})
		if err != nil {
			return resp, util.RpcCallErr(err, "token.rpc")
		}
		resp.Token = rpcResp.Token
		return resp, nil
	}
}
