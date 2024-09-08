package util

import (
	"google.golang.org/grpc/codes"
)

type Error struct {
	Errors     error
	StatusCode codes.Code
}
