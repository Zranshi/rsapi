package logic

import (
	"context"

	"rsapi/auth/rpc/token/internal/svc"
	"rsapi/auth/rpc/token/token"
	"rsapi/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type InvalidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInvalidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InvalidLogic {
	return &InvalidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InvalidLogic) Invalid(in *token.InvalidReq) (*token.InvalidRes, error) {
	resp := &token.InvalidRes{}
	kv := l.svcCtx.KvConn
	if in.Email != nil {
		if _, err := kv.Hdel("token", *in.Email); err != nil {
			return resp, util.DbErr(err)
		}
	} else if in.Token != nil {
		if claims, err := ParseJwt(*in.Token, l.svcCtx.Config.Service.JwtSecret); err != nil {
			_, err := kv.Hdel("token", claims.Email)
			if err != nil {
				return resp, util.DbErr(err)
			}
			return resp, util.ParseErr(err)
		} else {
			if _, err := kv.Hdel("token", claims.Email); err != nil {
				return resp, util.DbErr(err)
			}
		}
	} else {
		return resp, util.IllegalArgumentErr("need Email or token")
	}
	return resp, nil
}
