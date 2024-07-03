package repo

import (
	"context"

	model "github.com/faisalhardin/sawitpro/internal/entity/model"
	"github.com/pkg/errors"
)

const (
	MstTrxTreeTable = "swp_trx_tree_estate"

	WrapMsgInsertTree = WrapErrMsgPrefix + "InsertTree"
)

func (c *Conn) InsertTree(ctx context.Context, trxTree *model.TrxTree) (err error) {

	session := c.XormEngine.NewSession().Table(MstTrxTreeTable)

	sql := `INSERT INTO public.swp_trx_tree_estate (position_x, position_y, height, estate_id, create_time, update_time) VALUES (?, ?, ?, ?, ?, ?) RETURNING uuid`

	_, err = session.SQL(
		sql,
		trxTree.PositionX,
		trxTree.PositionY,
		trxTree.Height,
		trxTree.EstateID,
		trxTree.CreateTime,
		trxTree.UpdateTime).
		Get(trxTree)
	if err != nil {
		err = errors.Wrap(err, WrapMsgInsertTree)
		return
	}

	return nil
}
