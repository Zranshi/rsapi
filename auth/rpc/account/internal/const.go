package internal

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// ErrPwdNotMatch ...
	ErrPwdNotMatch = status.Error(codes.Unauthenticated, "password not match")
	// ErrAccountNotExist ...
	ErrAccountNotExist = status.Error(codes.NotFound, "account not exist")
	// ErrAccountAlreadyExist ...
	ErrAccountAlreadyExist = status.Error(codes.AlreadyExists, "account not exist")
	// ErrHashPwd
	ErrHashPwd = status.Error(codes.Internal, "hash password error")
	// ErrUpdateFailed
	ErrUpdateFailed = status.Error(codes.Internal, "update failed")
	// ErrDeleteFailed
	ErrDeleteFailed = status.Error(codes.Internal, "delete failed")
	// ErrRpcCall
	ErrRpcCall = status.Error(codes.Internal, "rpc call error")
)
