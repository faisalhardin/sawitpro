package utils

import (
	"net/http"

	model "github.com/faisalhardin/sawitpro/internal/entity/model"
)

func SetNewError(code int, errorName, errDesc string) *model.Response {
	return &model.Response{
		Code:    http.StatusBadRequest,
		Data:    errDesc,
		ErrName: errorName,
	}
}

func SetNewBadRequest(errorName, errDesc string) *model.Response {
	return SetNewError(http.StatusBadRequest, errorName, errDesc)
}

func SetNewNotFound(errorName, errDesc string) *model.Response {
	return SetNewError(http.StatusNotFound, errorName, errDesc)
}
