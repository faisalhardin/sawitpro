package interfaces

import (
	"context"
	"net/http"

	model "github.com/faisalhardin/sawitpro/internal/entity/model"
)

type EstateRepo interface {
	InsertEstate(ctx context.Context, estate *model.EstateDB) (err error)
}

type EstateUsecase interface {
	InsertEstate(ctx context.Context, req model.InsertEstateRequest) (resp model.InsertEstateResponse, err error)
}

type EstateHandler interface {
	InsertEstate(w http.ResponseWriter, r *http.Request)
}