package internal

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrIllegalArgument = status.Error(codes.InvalidArgument, "illegal argument")
	ErrRdbInsert       = status.Error(codes.Internal, "rdb insert error")
	ErrRdbFind         = status.Error(codes.Internal, "rdb find error")
	ErrRdbUpdate       = status.Error(codes.Internal, "rdb update error")
	ErrRdbList         = status.Error(codes.Internal, "rdb list error")
	ErrRdbDelete       = status.Error(codes.Internal, "rdb delete error")
)
