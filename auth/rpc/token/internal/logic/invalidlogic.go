package logic

import (
	"context"
	"errors"

	"rsapi/auth/rpc/token/internal"
	"rsapi/auth/rpc/token/internal/svc"
	"rsapi/auth/rpc/token/token"

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
			return resp, internal.ErrRedisHdel
		}
	} else if in.Token != nil {
		if claims, err := ParseJwt(*in.Token, l.svcCtx.Config.Service.JwtSecret); err != nil {
			if errors.Is(err, internal.ErrJwtExpired) {
				_, err := kv.Hdel("token", claims.Email)
				if err != nil {
					return resp, internal.ErrRedisHdel
				}
				return resp, internal.ErrJwtExpired
			}
			return resp, internal.ErrJwtParse
		} else {
			if _, err := kv.Hdel("token", claims.Email); err != nil {
				return resp, internal.ErrRedisHdel
			}
		}
	} else {
		return resp, internal.ErrInvalidParam
	}
	return resp, nil
}
