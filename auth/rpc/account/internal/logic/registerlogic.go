package logic

import (
	"context"
	"rsapi/auth/rpc/account/account"
	"rsapi/auth/rpc/account/internal/svc"
	"rsapi/model"
	"rsapi/util"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *account.RegisterReq) (*account.RegisterRes, error) {
	var (
		resp = new(account.RegisterRes)
	)
	key, err := pwd2Key(in.Pwd, l.svcCtx.Config.Service.PasswordSalt)
	if err != nil {
		return resp, util.ParseErr(err)
	}
	now := time.Now()
	u := &model.AuthUser{Email: in.Email, Name: in.Name, Key: key, CreateAt: now, UpdateAt: now}
	if _, err = l.svcCtx.UserM.Insert(l.ctx, u); err != nil {
		return resp, util.DbErr(err)
	}

	resp.Email = in.Email
	return resp, err
}
