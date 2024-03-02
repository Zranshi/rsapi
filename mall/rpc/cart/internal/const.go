package internal

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrRdbInsert = status.Error(codes.Internal, "rdb insert error")
	ErrRdbUpdate = status.Error(codes.Internal, "rdb update error")
	ErrRdbDelete = status.Error(codes.Internal, "rdb delete error")
	ErrRdbFind   = status.Error(codes.Internal, "rdb find error")
)
