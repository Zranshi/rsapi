package logic

import (
	"context"
	"errors"

	"rsapi/auth/rpc/token/internal"
	"rsapi/auth/rpc/token/internal/svc"
	"rsapi/auth/rpc/token/token"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyLogic {
	return &VerifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyLogic) Verify(in *token.VerifyReq) (*token.VerifyRes, error) {
	resp := &token.VerifyRes{Token: in.Token}
	kv := l.svcCtx.KvConn
	// verify token

	// parse token
	claims, err := ParseJwt(in.Token, l.svcCtx.Config.Service.JwtSecret)
	if err != nil {
		if errors.Is(err, internal.ErrJwtExpired) {
			_, err := kv.Hdel("token", claims.Email)
			if err != nil {
				return resp, internal.ErrJwtExpired
			}
			return resp, internal.ErrJwtExpired
		}
		return nil, internal.ErrJwtParse
	}

	// Verify token
	if token, err := kv.Hget("token", claims.Email); err != nil && token != in.Token {
		return resp, internal.ErrJwtExpired
	}

	// refresh token
	if in.Refreash {
		if token, err := GenerateJwt(claims, l.svcCtx.Config.Service.JwtSecret); err != nil {
			return resp, internal.ErrJwtGen
		} else {
			resp.Token = token
		}

		// save token in redis
		if err := kv.Hset("token", claims.Email, resp.Token); err != nil {
			return resp, internal.ErrRedisHSet
		}
	}

	return resp, nil
}
