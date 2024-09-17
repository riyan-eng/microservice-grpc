package util

import (
	"fmt"
	"server/infrastructure"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PaginationMeta struct {
	Page       *int32 `json:"page"`
	Limit      *int32 `json:"per_page"`
	CountRows  *int32 `json:"total"`
	CountPages *int32 `json:"total_pages"`
}

type SuccessResponse struct {
	Data     interface{}          `json:"data"`
	Message  string               `json:"message" example:"Successfully display data."`
	Meta     interface{}          `json:"meta,omitempty"`
	Response SuccessResponseChild `json:"response"`
}

type ErrorResponse struct {
	Error    any                `json:"errors"`
	Message  string             `json:"message" example:"Failed to display data."`
	Response ErrorResponseChild `json:"response"`
}

type ImportResponse struct {
	Errors     []ImportError        `json:"errors"`
	TotalInput int                  `json:"total_input"`
	Success    int                  `json:"success"`
	Failed     int                  `json:"failed"`
	Response   SuccessResponseChild `json:"response"`
}

type ImportError struct {
	Row    int `json:"nomor"`
	Errors any `json:"error"`
}

type SuccessResponseChild struct {
	RequestId string `json:"request_id" example:"8d0dc325-9fa4-430a-a46f-2042140c81ff"`
	Code      int    `json:"code" example:"200"`
	Message   string `json:"message" example:"OK"`
}

type ErrorResponseChild struct {
	RequestId string `json:"request_id" example:"8d0dc325-9fa4-430a-a46f-2042140c81ff"`
	Code      int    `json:"code" example:"400"`
	Message   string `json:"message" example:"Bad Request"`
}

type repsonseInterface interface {
	Success(data any, meta any, message string, statusCode ...int)
	Error(errors any, message string, statusCode ...int)
	Import(errors []ImportError, totalInput int, failed int)
	GrpcError(errors error, message string, statusCode ...int)
}

type responseStruct struct {
	c *gin.Context
}

func NewResponse(c *gin.Context) repsonseInterface {
	return &responseStruct{
		c: c,
	}
}

func (m *responseStruct) Success(data any, meta any, message string, statusCode ...int) {
	code := 200
	if len(statusCode) > 0 {
		code = statusCode[0]
	}

	requestId := requestid.Get(m.c)
	m.c.JSON(code, SuccessResponse{
		Data:    data,
		Meta:    meta,
		Message: message,
		Response: SuccessResponseChild{
			RequestId: requestId,
			Code:      code,
			Message:   statusMessages[code],
		},
	})
}

func (m *responseStruct) Error(errors any, message string, statusCode ...int) {
	code := 500
	if len(statusCode) > 0 && statusCode[0] != 0 {
		code = statusCode[0]
	}

	if message == "" {
		message = infrastructure.Localize(localizeResponseCode[code])
	}

	requestId := requestid.Get(m.c)

	m.c.Set("error", errors)
	m.c.AbortWithStatusJSON(code, ErrorResponse{
		Error:   errors,
		Message: message,
		Response: ErrorResponseChild{
			RequestId: requestId,
			Code:      code,
			Message:   statusMessages[code],
		},
	})
}

func (m *responseStruct) Import(errors []ImportError, totalInput int, failed int) {
	requestId := requestid.Get(m.c)
	code := 200
	m.c.JSON(200, ImportResponse{
		Errors:     errors,
		TotalInput: totalInput,
		Success:    totalInput - failed,
		Failed:     failed,
		Response: SuccessResponseChild{
			RequestId: requestId,
			Code:      code,
			Message:   statusMessages[code],
		},
	})
}

func (m *responseStruct) GrpcError(errors error, message string, statusCode ...int) {
	code := 500
	var err error = fmt.Errorf("")
	grpcErr, ok := status.FromError(errors)
	if ok {
		grpcCode := grpcErr.Code()
		grpcMessage := grpcErr.Message()
		if grpcMessage != "" {
			err = fmt.Errorf(grpcMessage)
		}

		switch grpcCode {
		case codes.InvalidArgument:
			code = 400
		case codes.NotFound:
			code = 404
		case codes.Unauthenticated:
			code = 401
		case codes.PermissionDenied:
			code = 403
		case codes.AlreadyExists:
			code = 409
		case codes.DeadlineExceeded:
			code = 504
		case codes.Unavailable:
			code = 500
		case codes.Internal:
			code = 500
		default:
			code = 500
		}
	}
	if len(statusCode) > 0 && statusCode[0] != 0 {
		code = statusCode[0]
	}

	if message == "" {
		fmt.Println(code)
		message = infrastructure.Localize(localizeResponseCode[code])
	}

	requestId := requestid.Get(m.c)

	m.c.Set("error", errors.Error())
	m.c.AbortWithStatusJSON(code, ErrorResponse{
		Error:   err.Error(),
		Message: message,
		Response: ErrorResponseChild{
			RequestId: requestId,
			Code:      code,
			Message:   statusMessages[code],
		},
	})
}

var localizeResponseCode = map[int]string{
	400: "BAD_REQUEST",
	401: "UNAUTHORIZED",
	403: "PERMISSION_DENIED",
	409: "CONFLICT",
	404: "NOT_FOUND",
	500: "BAD_SYSTEM",
}

var statusMessages = map[int]string{
	200: "OK",
	201: "Created",
	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	404: "Not Found",
	405: "Method Not Allowed",
	406: "Not Acceptable",
	407: "Proxy Authentication Required",
	408: "Request Timeout",
	409: "Conflict",
	410: "Gone",
	411: "Length Required",
	412: "Precondition Failed",
	413: "Payload Too Large",
	414: "URI Too Long",
	415: "Unsupported Media Type",
	416: "Range Not Satisfiable",
	417: "Expectation Failed",
	418: "I'm a teapot",
	421: "Misdirected Request",
	422: "Unprocessable Entity",
	423: "Locked",
	424: "Failed Dependency",
	425: "Too Early",
	426: "Upgrade Required",
	428: "Precondition Required",
	429: "Too Many Requests",
	431: "Request Header Fields Too Large",
	451: "Unavailable For Legal Reasons",
	500: "Internal Server Error",
	501: "Not Implemented",
	502: "Bad Gateway",
	503: "Service Unavailable",
	504: "Gateway Timeout",
	505: "HTTP Version Not Supported",
	506: "Variant Also Negotiates",
	507: "Insufficient Storage",
	508: "Loop Detected",
	510: "Not Extended",
	511: "Network Authentication Required",
}
