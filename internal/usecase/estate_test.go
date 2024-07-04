package usecase

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/faisalhardin/sawitpro/internal/entity/mocks"
	model "github.com/faisalhardin/sawitpro/internal/entity/model"
	"github.com/stretchr/testify/assert"
)

var mockEstateDB *mocks.MockEstateRepo

func initMocks(t *testing.T) *gomock.Controller {
	ctrl := gomock.NewController(t)

	mockEstateDB = mocks.NewMockEstateRepo(ctrl)

	return ctrl
}

func Test_InsertEstate(t *testing.T) {
	ctrl := initMocks(t)
	defer ctrl.Finish()

	type args struct {
		ctx context.Context
		req model.InsertEstateRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    model.InsertEstateResponse
		patch   func()
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				req: model.InsertEstateRequest{},
			},
			wantErr: false,
			want: model.InsertEstateResponse{
				ID: "UUID-1",
			},
			patch: func() {
				newUUID = func() (string, error) {
					return "UUID-1", nil
				}

				mockEstateDB.EXPECT().InsertEstate(gomock.Any(), gomock.Any()).
					Return(nil).Times(1)
			},
		},
		{
			name: "Error Inserting into DB",
			args: args{
				ctx: context.Background(),
				req: model.InsertEstateRequest{
					Width:  10,
					Length: 20,
				},
			},
			wantErr: true,
			patch: func() {
				newUUID = func() (string, error) {
					return "UUID-1", nil
				}

				mockEstateDB.EXPECT().InsertEstate(gomock.Any(), gomock.Any()).
					Return(errors.New("database error")).Times(1)
			},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			tt.patch()
			uc := EstateUC{
				EstateDBRepo: mockEstateDB,
			}
			got, err := uc.InsertEstate(tt.args.ctx, tt.args.req)
			if tt.wantErr {
				assert.Error(t, err, tt.name)
			} else {
				assert.NoError(t, err, fmt.Sprintf("got error = %v, wantErr = %v", err, tt.wantErr))
			}

			assert.Equal(t, got, tt.want)
		})
	}
}

func Test_GetEstateStatsByUUID(t *testing.T) {
	ctrl := initMocks(t)
	defer ctrl.Finish()

	type args struct {
		ctx  context.Context
		uuid string
	}
	tests := []struct {
		name      string
		args      args
		mock      func()
		want      model.EstateStats
		wantErr   bool
		wantError error
	}{
		{
			name: "Success",
			args: args{
				ctx:  context.Background(),
				uuid: "UUID-1",
			},
			mock: func() {
				mockEstateDB.EXPECT().GetEstateStats(gomock.Any(), "UUID-1").
					Return(model.EstateStats{}, nil).Times(1)
			},
			want:    model.EstateStats{},
			wantErr: false,
		},
		{
			name: "Error from Repository",
			args: args{
				ctx:  context.Background(),
				uuid: "UUID-2",
			},
			mock: func() {
				mockEstateDB.EXPECT().GetEstateStats(gomock.Any(), "UUID-2").
					Return(model.EstateStats{}, errors.New("database error")).Times(1)
			},
			want:      model.EstateStats{},
			wantErr:   true,
			wantError: errors.New("database error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			uc := EstateUC{
				EstateDBRepo: mockEstateDB,
			}
			got, err := uc.GetEstateStatsByUUID(tt.args.ctx, tt.args.uuid)
			if tt.wantErr {
				assert.Error(t, err, tt.name)
				assert.EqualError(t, err, tt.wantError.Error(), tt.name)
			} else {
				assert.NoError(t, err, tt.name)
				assert.Equal(t, tt.want, got, tt.name)
			}
		})
	}
}

func Test_GetDronePlanByEstateUUID(t *testing.T) {
	ctrl := initMocks(t)
	defer ctrl.Finish()

	type args struct {
		ctx  context.Context
		uuid string
	}
	tests := []struct {
		name      string
		args      args
		mock      func()
		want      model.EstateDronePlanResponse
		wantErr   bool
		wantError string
	}{
		{
			name: "Success",
			args: args{
				ctx:  context.Background(),
				uuid: "UUID-1",
			},
			mock: func() {
				mockEstateDB.EXPECT().GetEstateTreesHeightPosition(gomock.Any(), "UUID-1").
					Return([]model.TreeHeight{
						{Height: 10},
						{Height: 15},
						{Height: 20},
					}, nil).Times(1)
			},
			want: model.EstateDronePlanResponse{
				Distance: 62,
			},
			wantErr: false,
		},
		{
			name: "Error from Repository",
			args: args{
				ctx:  context.Background(),
				uuid: "UUID-2",
			},
			mock: func() {
				mockEstateDB.EXPECT().GetEstateTreesHeightPosition(gomock.Any(), "UUID-2").
					Return(nil, errors.New("database error")).Times(1)
			},
			want:      model.EstateDronePlanResponse{},
			wantErr:   true,
			wantError: "database error",
		},
		{
			name: "No Trees Found",
			args: args{
				ctx:  context.Background(),
				uuid: "UUID-3",
			},
			mock: func() {
				mockEstateDB.EXPECT().GetEstateTreesHeightPosition(gomock.Any(), "UUID-3").
					Return([]model.TreeHeight{}, nil).Times(1)
			},
			want:      model.EstateDronePlanResponse{},
			wantErr:   true,
			wantError: "Not found: Trees not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			uc := EstateUC{
				EstateDBRepo: mockEstateDB,
			}
			got, err := uc.GetDronePlanByEstateUUID(tt.args.ctx, tt.args.uuid)
			if tt.wantErr {
				assert.Error(t, err, tt.name)
				assert.Contains(t, err.Error(), tt.wantError, tt.name)
			} else {
				assert.NoError(t, err, tt.name)
				assert.Equal(t, tt.want, got, tt.name)
			}
		})
	}
}

func Test_NewUUIDString(t *testing.T) {
	tests := []struct {
		name    string
		mock    func()
		want    string
		wantErr bool
	}{
		{
			name: "Success",
			mock: func() {
				newUUID = func() (string, error) {
					return "generated-uuid", nil
				}
			},
			want:    "generated-uuid",
			wantErr: false,
		},
		{
			name: "Error Generating UUID",
			mock: func() {
				newUUID = func() (string, error) {
					return "", errors.New("uuid generation error")
				}
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := NewUUIDString()
			if tt.wantErr {
				assert.Error(t, err, tt.name)
			} else {
				assert.NoError(t, err, tt.name)
				assert.Equal(t, tt.want, got, tt.name)
			}
		})
	}
}
