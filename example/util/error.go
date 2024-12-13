package util

import (
	"strings"

	"google.golang.org/grpc/codes"
)

func NewError() *Error {
	return &Error{}
}

func (e *Error) ParsingPrefix(err error, optMessage ...string) (codes.Code, string) {
	errMsg := err.Error()
	colonIndex := strings.Index(errMsg, ":")
	if colonIndex > 0 {
		codeStr := errMsg[:colonIndex]
		message := errMsg[colonIndex+1:]
		if len(optMessage) > 0 {
			message = optMessage[0]
		}
		switch codeStr {
		case "1":
			return 1, message
		case "2":
			return 2, message
		case "3":
			return 3, message
		case "4":
			return 4, message
		case "5":
			return 5, message
		case "6":
			return 6, message
		case "7":
			return 7, message
		case "8":
			return 8, message
		case "9":
			return 9, message
		case "10":
			return 10, message
		case "11":
			return 11, message
		case "13":
			return 13, message
		case "14":
			return 14, message
		case "15":
			return 15, message
		case "16":
			return 16, message
		}
	}

	return 13, "format error tidak sesuai"
}
