package usecase

import (
	"context"

	model "github.com/faisalhardin/sawitpro/internal/entity/model"
	"github.com/pkg/errors"
)

var (
	WrapMsgInsertNewTree = WrapErrMsgPrefix + "InsertNewTree"
)

func (uc *EstateUC) InsertNewTree(ctx context.Context, req model.InsertNewTreeRequest) (resp model.InsertNewTreeResponse, err error) {

	mstEstate, found, err := uc.EstateDBRepo.GetEstateByUUID(ctx, req.EstateUUID)
	if err != nil {
		err = errors.Wrap(err, WrapMsgInsertNewTree)
		return
	}

	if !found {
		err = errors.New("estate is not found")
		return
	}

	if mstEstate.Length < req.PositionX || req.PositionX <= 0 || mstEstate.Width < req.PositionY || req.PositionY <= 0 {
		err = errors.New("position is out of bound")
		return
	}

	newTree := model.TrxTree{
		PositionX: req.PositionX,
		PositionY: req.PositionY,
		Height:    req.Height,
		EstateID:  mstEstate.ID,
	}

	err = uc.EstateDBRepo.InsertTree(ctx, &newTree)
	if err != nil {
		err = errors.Wrap(err, WrapErrMsgPrefix)
		return
	}

	return model.InsertNewTreeResponse{
		UUID: newTree.UUID,
	}, nil
}
