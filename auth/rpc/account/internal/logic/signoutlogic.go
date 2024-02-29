package logic

import (
	"context"

	"rsapi/auth/rpc/account/account"
	"rsapi/auth/rpc/account/internal"
	"rsapi/auth/rpc/account/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
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
		return resp, internal.ErrAccountNotExist
	} else {
		// check pwd
		if key, err := pwd2Key(in.Pwd, l.svcCtx.Config.Service.PasswordSalt); err != nil {
			return resp, internal.ErrHashPwd
		} else if key != u.Key {
			return resp, internal.ErrPwdNotMatch
		}
		// delete user instance
		if err := l.svcCtx.UserM.Delete(l.ctx, u.Id); err != nil {
			return resp, internal.ErrDeleteFailed
		}
		return resp, nil
	}
}
