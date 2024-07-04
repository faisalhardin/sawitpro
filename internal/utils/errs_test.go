package utils

import (
	"net/http"
	"testing"

	"github.com/faisalhardin/sawitpro/internal/entity/model"
	"github.com/stretchr/testify/assert"
)

func Test_SetNewError(t *testing.T) {
	type args struct {
		code      int
		errorName string
		errDesc   string
	}
	tests := []struct {
		name string
		args args
		want *model.Response
	}{
		{
			name: "BadRequest Error",
			args: args{
				code:      http.StatusBadRequest,
				errorName: "BadRequest",
				errDesc:   "Bad request error occurred",
			},
			want: &model.Response{
				Code:    http.StatusBadRequest,
				Data:    "Bad request error occurred",
				ErrName: "BadRequest",
			},
		},
		{
			name: "NotFound Error",
			args: args{
				code:      http.StatusNotFound,
				errorName: "NotFound",
				errDesc:   "Resource not found",
			},
			want: &model.Response{
				Code:    http.StatusNotFound,
				Data:    "Resource not found",
				ErrName: "NotFound",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SetNewError(tt.args.code, tt.args.errorName, tt.args.errDesc)
			assert.Equal(t, tt.want, got, tt.name)
		})
	}
}

func Test_SetNewBadRequest(t *testing.T) {
	type args struct {
		errorName string
		errDesc   string
	}
	tests := []struct {
		name string
		args args
		want *model.Response
	}{
		{
			name: "BadRequest Error",
			args: args{
				errorName: "BadRequest",
				errDesc:   "Bad request error occurred",
			},
			want: &model.Response{
				Code:    http.StatusBadRequest,
				Data:    "Bad request error occurred",
				ErrName: "BadRequest",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SetNewBadRequest(tt.args.errorName, tt.args.errDesc)
			assert.Equal(t, tt.want, got, tt.name)
		})
	}
}

func Test_SetNewNotFound(t *testing.T) {
	type args struct {
		errorName string
		errDesc   string
	}
	tests := []struct {
		name string
		args args
		want *model.Response
	}{
		{
			name: "NotFound Error",
			args: args{
				errorName: "NotFound",
				errDesc:   "Resource not found",
			},
			want: &model.Response{
				Code:    http.StatusNotFound,
				Data:    "Resource not found",
				ErrName: "NotFound",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SetNewNotFound(tt.args.errorName, tt.args.errDesc)
			assert.Equal(t, tt.want, got, tt.name)
		})
	}
}
