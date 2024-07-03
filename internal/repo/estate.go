package repo

import (
	"context"

	model "github.com/faisalhardin/sawitpro/internal/entity/model"
	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
)

const (
	MstEstateTable = "swp_mst_estate"

	WrapErrMsgPrefix                 = "EstateDBRepo."
	WrapMsgInsertEstate              = WrapErrMsgPrefix + "InsertEstate"
	WrapMsgGetEstateJoinTreeByParams = WrapErrMsgPrefix + "GetEstateJoinTreeByParams"
)

type Conn struct {
	XormEngine *xorm.Engine
}

func NewEstateDBRepo(repo *Conn) *Conn {
	return repo
}

func (c *Conn) InsertEstate(ctx context.Context, estate *model.EstateDB) (err error) {

	session := c.XormEngine.NewSession().Table(MstEstateTable)

	_, err = session.InsertOne(estate)
	if err != nil {
		err = errors.Wrap(err, WrapMsgInsertEstate)
		return
	}

	return nil
}

func (c *Conn) GetEstateByUUID(ctx context.Context, uuid string) (estate model.EstateDB, found bool, err error) {
	session := c.XormEngine.NewSession().Table(MstEstateTable)

	found, err = session.
		Where("uuid = ?", uuid).
		Get(&estate)
	if err != nil {
		err = errors.Wrap(err, WrapMsgInsertEstate)
		return
	}

	return
}

func (c *Conn) GetEstateJoinTreeByParams(ctx context.Context, params model.InsertNewTreeRequest) (resp []model.EstateJoinTrxTree, err error) {
	session := c.XormEngine.NewSession().Table(MstEstateTable)

	err = session.
		Alias("sme").
		Join("LEFT", "swp_trx_tree_estate stte", "sme.id = stte.estate_id and stte.position_x = ? and stte.position_y = ?", params.PositionX, params.PositionY).
		Where("sme.uuid = ?", params.EstateUUID).
		Find(&resp)
	if err != nil {
		err = errors.Wrap(err, WrapMsgGetEstateJoinTreeByParams)
		return
	}

	return
}
