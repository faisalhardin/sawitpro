package usecase

import (
	"context"

	estateRepo "github.com/faisalhardin/sawitpro/internal/entity/interfaces"
	model "github.com/faisalhardin/sawitpro/internal/entity/model"
	"github.com/faisalhardin/sawitpro/internal/utils"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

const (
	WrapErrMsgPrefix                = "EstateUsecase."
	WrapMsgInsertEstate             = WrapErrMsgPrefix + "InsertEstate"
	WrapMsgGetDronePlanByEstateUUID = WrapErrMsgPrefix + "GetDronePlanByEstateUUID"

	GridLength = int32(10)
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

func (uc *EstateUC) GetDronePlanByEstateUUID(ctx context.Context, uuid string) (resp model.EstateDronePlanResponse, err error) {
	treesHeights, err := uc.EstateDBRepo.GetEstateTreesHeightPosition(ctx, uuid)
	if err != nil {
		err = errors.Wrap(err, WrapMsgGetDronePlanByEstateUUID)
		return
	}

	if len(treesHeights) == 0 {
		err = utils.SetNewNotFound("Not found", "Trees not found")
		return
	}

	var (
		distanceVerticalTraversed   int32 = 0
		distanceHorizontalTraversed int32 = 0
		currHeight                  int32 = 0
	)

	for _, tree := range treesHeights {
		distanceHorizontalTraversed += GridLength
		distanceVerticalTraversed += utils.Abs(tree.Height - currHeight)
		currHeight = tree.Height
	}

	distanceVerticalTraversed += 2 + currHeight // take0ff & land
	distanceHorizontalTraversed -= GridLength   //last grid doesn't count

	return model.EstateDronePlanResponse{
		Distance: distanceVerticalTraversed + distanceHorizontalTraversed,
	}, nil
}
