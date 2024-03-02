package logic

import (
	"context"

	"rsapi/auth/rpc/account/account"
	"rsapi/auth/rpc/account/internal/svc"
	"rsapi/util"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SignoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignoutLogic {
	return &SignoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SignoutLogic) Signout(in *account.SignOutReq) (*account.SignOutRes, error) {
	resp := new(account.SignOutRes)
	// check account
	if u, err := l.svcCtx.UserM.FindOneByEmail(l.ctx, in.Email); err != nil || u == nil {
		return resp, util.DbErr(err)
	} else {
		// check pwd
		if key, err := pwd2Key(in.Pwd, l.svcCtx.Config.Service.PasswordSalt); err != nil {
			return resp, util.ParseErr(err)
		} else if key != u.Key {
			return resp, status.Error(codes.Unauthenticated, "pwd not match")
		}
		// delete user instance
		if err := l.svcCtx.UserM.Delete(l.ctx, u.Id); err != nil {
			return resp, util.DbErr(err)
		}
		return resp, nil
	}
}
