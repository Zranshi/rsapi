package util

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func DbErr(err error) error {
	return status.Errorf(codes.Internal, "db error: %v", err)
}

func IllegalArgumentErr(params string) error {
	return status.Errorf(codes.InvalidArgument, "illegal argument: %s", params)
}
