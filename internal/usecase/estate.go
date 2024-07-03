package usecase

import (
	"context"

	estateRepo "github.com/faisalhardin/sawitpro/internal/entity/interfaces"
	model "github.com/faisalhardin/sawitpro/internal/entity/model"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

const (
	WrapErrMsgPrefix    = "EstateUsecase."
	WrapMsgInsertEstate = WrapErrMsgPrefix + "InsertEstate"
)

type EstateUC struct {
	EstateDBRepo estateRepo.EstateRepo
}

func NewEstateUC(uc *EstateUC) *EstateUC {
	return uc
}

func (uc *EstateUC) InsertEstate(ctx context.Context, req model.InsertEstateRequest) (resp model.InsertEstateResponse, err error) {

	uuidV4, err := uuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, WrapErrMsgPrefix)
		return
	}

	dbModel := model.EstateDB{
		Width:  req.Width,
		Length: req.Length,
		UUID:   uuidV4.String(),
	}

	err = uc.EstateDBRepo.InsertEstate(ctx, &dbModel)
	if err != nil {
		err = errors.Wrap(err, WrapErrMsgPrefix)
		return
	}

	return model.InsertEstateResponse{
		ID: dbModel.UUID,
	}, nil
}

func (uc *EstateUC) GetEstateStatsByUUID(ctx context.Context, uuid string) (resp model.EstateStats, err error) {

	resp, err = uc.EstateDBRepo.GetEstateStats(ctx, uuid)
	if err != nil {
		err = errors.Wrap(err, WrapErrMsgPrefix)
		return
	}

	return
}
