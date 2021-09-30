package utils

import (
	"helloadmin/pkg/error"
)

type (
	Meta struct {
		Page  int   `json:"page"`
		Size  int   `json:"size"`
		Total int64 `json:"total"`
	}

	Data struct {
		*Meta `json:"meta,omitempty"`
		Items interface{} `json:"items"`
	}

	Result struct {
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)

func Success(httpStatus, status int, data interface{}, meta *Meta) (int, Result) {
	if meta != nil {
		return httpStatus, Result{
			Status:  status,
			Message: "Success",
			Data: &Data{
				meta,
				data,
			},
		}
	}
	return httpStatus, Result{Status: status, Message: "Success", Data: data}
}

func Error(httpStatus, status int) (int, Result) {
	return httpStatus, Result{Status: status, Message: error.Message(status), Data: nil}
}
