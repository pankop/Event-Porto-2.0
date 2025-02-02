package utils

import "github.com/Caknoooo/go-gin-clean-starter/utils/pagination"

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Error   any    `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
	Meta    any    `json:"meta,omitempty"`
}

type EmptyObj struct{}

func BuildResponseSuccess(message string, data any, meta ...pagination.Meta) Response {
	res := Response{
		Status:  true,
		Message: message,
		Data:    data,
	}

	if len(meta) > 0 {
		res.Meta = meta
	}
	return res
}

func BuildResponseFailed(message string, err string, data any) Response {
	res := Response{
		Status:  false,
		Message: message,
		Error:   err,
		Data:    data,
	}
	return res
}
