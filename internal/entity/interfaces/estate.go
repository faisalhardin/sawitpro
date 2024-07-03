package interfaces

import (
	"context"
	"net/http"

	model "github.com/faisalhardin/sawitpro/internal/entity/model"
)

type EstateRepo interface {
	InsertEstate(ctx context.Context, estate *model.EstateDB) (err error)
	GetEstateByUUID(ctx context.Context, uuid string) (estate model.EstateDB, found bool, err error)
	InsertTree(ctx context.Context, trxTree *model.TrxTree) (err error)
}

type EstateUsecase interface {
	InsertEstate(ctx context.Context, req model.InsertEstateRequest) (resp model.InsertEstateResponse, err error)
	InsertNewTree(ctx context.Context, req model.InsertNewTreeRequest) (resp model.InsertNewTreeResponse, err error)
}

type EstateHandler interface {
	InsertEstate(w http.ResponseWriter, r *http.Request)
	InsertTree(w http.ResponseWriter, r *http.Request)
}
