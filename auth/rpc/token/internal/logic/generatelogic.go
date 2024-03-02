package logic

import (
	"context"

	"rsapi/auth/rpc/token/internal/svc"
	"rsapi/auth/rpc/token/token"
	"rsapi/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateLogic {
	return &GenerateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateLogic) Generate(in *token.GenerateReq) (*token.GenerateRes, error) {
	resp := new(token.GenerateRes)
	kv := l.svcCtx.KvConn
	// create claims
	claims := new(Claims)
	claims.Email = in.Email
	if in.Role != nil {
		claims.Role = *in.Role
	}
	claims.Expire = in.Expire
	// Gererate token
	if token, err := GenerateJwt(claims, l.svcCtx.Config.Service.JwtSecret); err != nil {
		return resp, util.ParseErr(err)
	} else {
		resp.Token = token
	}

	// save token in redis
	if err := kv.Hset("token", in.Email, resp.Token); err != nil {
		return resp, util.DbErr(err)
	}

	return resp, nil
}
