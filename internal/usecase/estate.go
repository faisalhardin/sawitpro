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

var (
	newUUID   = NewUUIDString
	uuidNewV4 = uuid.NewV4
)

type EstateUC struct {
	EstateDBRepo estateRepo.EstateRepo
}

func NewEstateUC(uc *EstateUC) *EstateUC {
	return uc
}

func (uc *EstateUC) InsertEstate(ctx context.Context, req model.InsertEstateRequest) (resp model.InsertEstateResponse, err error) {

	uuidStr, err := newUUID()
	if err != nil {
		err = errors.Wrap(err, WrapErrMsgPrefix)
		return
	}

	dbModel := model.EstateDB{
		Width:  req.Width,
		Length: req.Length,
		UUID:   uuidStr,
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

func (uc *EstateUC) GetDronePlanByEstateUUID(ctx context.Context, params model.GetDronePlanParams) (resp model.EstateDronePlanResponse, err error) {

	treesHeights, err := uc.EstateDBRepo.GetEstateTreesHeightPosition(ctx, params.UUID)
	if err != nil {
		err = errors.Wrap(err, WrapMsgGetDronePlanByEstateUUID)
		return
	}

	if len(treesHeights) == 0 {
		err = utils.SetNewNotFound("Not found", "Trees not found")
		return
	}

	var (
		distanceTraversed int32 = 0
		currHeight        int32 = 0
		maxDistance       int32 = params.MaxDistance
		isCheckBattery    bool  = false
		lastCoordinates   model.TreeHeight
	)

	if maxDistance > 0 {
		isCheckBattery = true
	}

	distanceTraversed = 1 // takeoff

	for i, tree := range treesHeights {

		distanceTraversed += utils.Abs(currHeight - tree.Height)

		if isCheckBattery {
			if distanceTraversed > maxDistance {
				break
			}
			lastCoordinates = tree

		}

		hDist := GridLength
		if i == 0 || i == len(treesHeights)-1 {
			hDist = GridLength / 2
		}
		distanceTraversed += hDist
		currHeight = tree.Height

	}

	distanceTraversed += 1 + currHeight // land

	resp = model.EstateDronePlanResponse{
		Distance: distanceTraversed,
	}

	if isCheckBattery &&
		lastCoordinates.PositionX != 0 &&
		lastCoordinates != treesHeights[len(treesHeights)-1] {

		resp = model.EstateDronePlanResponse{
			Distance: maxDistance,
			Rest: &model.Coordinates{
				PositionX: lastCoordinates.PositionX,
				PositionY: lastCoordinates.PositionY,
			},
		}

	}

	return resp, nil
}

func NewUUIDString() (string, error) {
	uuidV4, err := uuidNewV4()
	if err != nil {
		return "", err
	}

	return uuidV4.String(), nil
}
