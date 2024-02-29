package internal

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// ErrJwtGen
	ErrJwtGen = status.Error(codes.Internal, "jwt generate error")
	// ErrJwtParse
	ErrJwtParse = status.Error(codes.Internal, "jwt parse error")
	// ErrJwtExpired
	ErrJwtExpired = status.Error(codes.Unauthenticated, "jwt expired")
	// ErrRedisHset
	ErrRedisHSet = status.Error(codes.Internal, "redis Hset error")
	// ErrRedisHdel
	ErrRedisHdel = status.Error(codes.Internal, "redis Hdel error")
	// ErrRedisHGet
	ErrRedisHGet = status.Error(codes.Internal, "redis Hget error")
	// ErrInvalidParam
	ErrInvalidParam = status.Error(codes.InvalidArgument, "invalid param")
)
