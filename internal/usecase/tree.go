package usecase

import (
	"context"

	model "github.com/faisalhardin/sawitpro/internal/entity/model"
	"github.com/faisalhardin/sawitpro/internal/utils"
	"github.com/pkg/errors"
)

var (
	WrapMsgInsertNewTree = WrapErrMsgPrefix + "InsertNewTree"
)

func (uc *EstateUC) InsertNewTree(ctx context.Context, req model.InsertNewTreeRequest) (resp model.InsertNewTreeResponse, err error) {

	mstEstate, err := uc.EstateDBRepo.GetEstateJoinTreeByParams(ctx, req)
	if err != nil {
		err = errors.Wrap(err, WrapMsgInsertNewTree)
		return
	}

	if len(mstEstate) == 0 {
		err = utils.SetNewNotFound("Validation", "Estate is not found")
		return
	}

	if len(mstEstate) > 0 && mstEstate[0].Tree.ID > 0 {
		err = utils.SetNewBadRequest("Validation", "Plot already has tree")
		return
	}

	if len(mstEstate) > 0 && mstEstate[0].Estate.Length < req.PositionX || req.PositionX <= 0 || mstEstate[0].Estate.Width < req.PositionY || req.PositionY <= 0 {
		err = utils.SetNewBadRequest("Validation", "Position is out of bound")
		return
	}

	newTree := model.TrxTree{
		PositionX: req.PositionX,
		PositionY: req.PositionY,
		Height:    req.Height,
		EstateID:  mstEstate[0].Estate.ID,
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
