package repo

import (
	"context"

	model "github.com/faisalhardin/sawitpro/internal/entity/model"
	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
)

const (
	MstEstateTable = "swp_mst_estate"

	WrapErrMsgPrefix = "EstateDBRepo."
	WrapMsgInsertEstate = WrapErrMsgPrefix+"InsertEstate"
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
		err = errors.Wrap(err, "InsertEstate")
	}
	
	return nil
}