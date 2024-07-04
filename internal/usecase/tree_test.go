package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	model "github.com/faisalhardin/sawitpro/internal/entity/model"
)

func Test_InsertNewTree(t *testing.T) {
	ctrl := initMocks(t)
	defer ctrl.Finish()

	type args struct {
		ctx context.Context
		req model.InsertNewTreeRequest
	}
	tests := []struct {
		name    string
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				req: model.InsertNewTreeRequest{
					Height:     10,
					PositionX:  2,
					PositionY:  3,
					EstateUUID: "UUID-1",
				},
			},
			mock: func() {
				mockEstateDB.EXPECT().GetEstateJoinTreeByParams(gomock.Any(), gomock.Any()).
					Return([]model.EstateJoinTrxTree{
						{
							Estate: model.EstateDB{
								UUID:   "UUID-1",
								Width:  3,
								Length: 3,
							},
						},
					}, nil).Times(1)

				mockEstateDB.EXPECT().InsertTree(gomock.Any(), gomock.Any()).
					Return(nil).Times(1)
			},
			wantErr: false,
		},
		{
			name: "Validation Error: Incorrect Tree Height",
			args: args{
				ctx: context.Background(),
				req: model.InsertNewTreeRequest{
					Height:     35,
					PositionX:  2,
					PositionY:  3,
					EstateUUID: "UUID-1",
				},
			},
			mock:    func() {},
			wantErr: true,
		},
		{
			name: "Database Error: GetEstateJoinTreeByParams",
			args: args{
				ctx: context.Background(),
				req: model.InsertNewTreeRequest{
					Height:     10,
					PositionX:  2,
					PositionY:  3,
					EstateUUID: "UUID-2",
				},
			},
			mock: func() {
				mockEstateDB.EXPECT().GetEstateJoinTreeByParams(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("database error")).Times(1)
			},
			wantErr: true,
		},
		{
			name: "Validation Error: Plot already has tree",
			args: args{
				ctx: context.Background(),
				req: model.InsertNewTreeRequest{
					Height:     10,
					PositionX:  2,
					PositionY:  3,
					EstateUUID: "UUID-3",
				},
			},
			mock: func() {
				mockEstateDB.EXPECT().GetEstateJoinTreeByParams(gomock.Any(), gomock.Any()).
					Return([]model.EstateJoinTrxTree{
						{
							Tree: model.TrxTree{ID: 1},
							Estate: model.EstateDB{
								ID:     1,
								Length: 5,
								Width:  5,
							},
						},
					}, nil).Times(1)
			},
			wantErr: true,
		},
		{
			name: "Validation Error: Position is out of bound",
			args: args{
				ctx: context.Background(),
				req: model.InsertNewTreeRequest{
					Height:     10,
					PositionX:  10,
					PositionY:  2,
					EstateUUID: "UUID-4",
				},
			},
			mock: func() {
				mockEstateDB.EXPECT().GetEstateJoinTreeByParams(gomock.Any(), gomock.Any()).
					Return([]model.EstateJoinTrxTree{
						{
							Estate: model.EstateDB{
								ID:     1,
								Length: 5,
								Width:  5,
							},
						},
					}, nil).Times(1)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			uc := EstateUC{
				EstateDBRepo: mockEstateDB,
			}
			_, err := uc.InsertNewTree(tt.args.ctx, tt.args.req)
			if tt.wantErr {
				assert.Error(t, err, tt.name)
			} else {
				assert.NoError(t, err, tt.name)
			}
		})
	}
}
