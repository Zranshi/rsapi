package logic

import (
	"context"
	"rsapi/auth/rpc/account/account"
	"rsapi/auth/rpc/account/internal"
	"rsapi/auth/rpc/account/internal/svc"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditLogic {
	return &EditLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditLogic) Edit(in *account.EditReq) (*account.EditRes, error) {
	resp := new(account.EditRes)

	u, err := l.svcCtx.UserM.FindOneByEmail(l.ctx, in.Email)
	if err != nil || u == nil {
		return resp, internal.ErrAccountNotExist
	}

	if in.Pwd != nil {
		key, err := pwd2Key(*in.Pwd, l.svcCtx.Config.Service.PasswordSalt)
		if err != nil {
			return resp, internal.ErrHashPwd
		}
		u.Key = key
		u.UpdateAt = time.Now()
	}
	if in.Name != nil {
		u.Name = *in.Name
		u.UpdateAt = time.Now()
	}

	if err := l.svcCtx.UserM.Update(l.ctx, u); err != nil {
		return resp, internal.ErrUpdateFailed
	}
	return resp, nil
}
