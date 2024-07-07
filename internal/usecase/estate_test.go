package usecase

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/faisalhardin/sawitpro/internal/entity/mocks"
	model "github.com/faisalhardin/sawitpro/internal/entity/model"
	"github.com/gofrs/uuid"
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
		name    string
		args    args
		mock    func()
		want    model.EstateStats
		wantErr bool
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
			want:    model.EstateStats{},
			wantErr: true,
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
		ctx    context.Context
		params model.GetDronePlanParams
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
			name: "Success with max distance lower than should be traversed",
			args: args{
				ctx: context.Background(),
				params: model.GetDronePlanParams{
					UUID:        "UUID-1",
					MaxDistance: 30,
				},
			},
			mock: func() {
				mockEstateDB.EXPECT().GetEstateTreesHeightPosition(gomock.Any(), "UUID-1").
					Return([]model.TreeHeight{
						{
							Height:    0,
							PositionX: 1,
							PositionY: 1,
						},
						{
							Height:    5,
							PositionX: 2,
							PositionY: 1},
						{
							Height:    3,
							PositionX: 3,
							PositionY: 1},
						{
							Height:    4,
							PositionX: 4,
							PositionY: 1},
						{
							Height:    0,
							PositionX: 5,
							PositionY: 1},
					}, nil).Times(1)
			},
			want: model.EstateDronePlanResponse{
				Distance: 30,
				Rest: &model.Coordinates{
					PositionX: 3,
					PositionY: 1,
				},
			},
			wantErr: false,
		},
		{
			name: "Success with max distance higher than should be traversed",
			args: args{
				ctx: context.Background(),
				params: model.GetDronePlanParams{
					UUID:        "UUID-1",
					MaxDistance: 100,
				},
			},
			mock: func() {
				mockEstateDB.EXPECT().GetEstateTreesHeightPosition(gomock.Any(), "UUID-1").
					Return([]model.TreeHeight{
						{
							Height:    0,
							PositionX: 1,
							PositionY: 1,
						},
						{
							Height:    5,
							PositionX: 2,
							PositionY: 1},
						{
							Height:    3,
							PositionX: 3,
							PositionY: 1},
						{
							Height:    4,
							PositionX: 4,
							PositionY: 1},
						{
							Height:    0,
							PositionX: 5,
							PositionY: 1},
					}, nil).Times(1)
			},
			want: model.EstateDronePlanResponse{
				Distance: 54,
			},
			wantErr: false,
		},
		{
			name: "Success without max distance",
			args: args{
				ctx: context.Background(),
				params: model.GetDronePlanParams{
					UUID: "UUID-1",
				},
			},
			mock: func() {
				mockEstateDB.EXPECT().GetEstateTreesHeightPosition(gomock.Any(), "UUID-1").
					Return([]model.TreeHeight{
						{
							Height:    0,
							PositionX: 1,
							PositionY: 1,
						},
						{
							Height:    5,
							PositionX: 2,
							PositionY: 1},
						{
							Height:    3,
							PositionX: 3,
							PositionY: 1},
						{
							Height:    4,
							PositionX: 4,
							PositionY: 1},
						{
							Height:    0,
							PositionX: 5,
							PositionY: 1},
					}, nil).Times(1)
			},
			want: model.EstateDronePlanResponse{
				Distance: 54,
			},
			wantErr: false,
		},
		{
			name: "Error from Repository",
			args: args{
				ctx: context.Background(),
				params: model.GetDronePlanParams{
					UUID: "UUID-2",
				},
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
				ctx: context.Background(),
				params: model.GetDronePlanParams{
					UUID: "UUID-3",
				},
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
			got, err := uc.GetDronePlanByEstateUUID(tt.args.ctx, tt.args.params)
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
				uuidNewV4 = func() (uuid.UUID, error) {
					return uuid.FromStringOrNil("00000000-0000-0000-0000-000000000001"), nil
				}
			},
			want:    "00000000-0000-0000-0000-000000000001",
			wantErr: false,
		},
		{
			name: "Error Generating UUID",
			mock: func() {
				uuidNewV4 = func() (uuid.UUID, error) {
					return uuid.UUID{}, errors.New("error generating uuid")
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
