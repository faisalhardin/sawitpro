package repo

import (
	"context"

	model "github.com/faisalhardin/sawitpro/internal/entity/model"
	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
)

const (
	MstEstateTable = "swp_mst_estate"

	WrapErrMsgPrefix                    = "EstateDBRepo."
	WrapMsgInsertEstate                 = WrapErrMsgPrefix + "InsertEstate"
	WrapMsgGetEstateJoinTreeByParams    = WrapErrMsgPrefix + "GetEstateJoinTreeByParams"
	WrapMsgGetEstateStats               = WrapErrMsgPrefix + "GetEstateStats"
	WrapMsgGetEstateTreesHeightPosition = WrapErrMsgPrefix + "GetEstateTreesHeightPosition"
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
		Join("LEFT", "swp_trx_tree_estate stte", "sme.id = stte.id_mst_estate and stte.position_x = ? and stte.position_y = ?", params.PositionX, params.PositionY).
		Where("sme.uuid = ?", params.EstateUUID).
		Find(&resp)
	if err != nil {
		err = errors.Wrap(err, WrapMsgGetEstateJoinTreeByParams)
		return
	}

	return
}

func (c *Conn) GetEstateStats(ctx context.Context, uuid string) (resp model.EstateStats, err error) {
	session := c.XormEngine.NewSession()

	_, err = session.
		SQL(`WITH estate_summary AS (
				SELECT
					COUNT(*) AS tree_count,
					MAX(height) AS max_height,
					MIN(height) AS min_height,
					PERCENTILE_CONT(0.5) WITHIN GROUP (ORDER BY height) AS median_height
				FROM swp_trx_tree_estate stte
				INNER JOIN swp_mst_estate sme ON sme.id = stte.id_mst_estate AND stte.delete_time is null
				WHERE sme.uuid = ?
			)
			SELECT
				COALESCE(tree_count, 0) AS tree_count,
				COALESCE(max_height, 0) AS max_height,
				COALESCE(min_height, 0) AS min_height,
				COALESCE(median_height, 0) AS median_height
			FROM estate_summary`, uuid).
		Get(&resp)
	if err != nil {
		err = errors.Wrap(err, WrapMsgGetEstateStats)
		return
	}

	return
}

func (c *Conn) GetEstateTreesHeightPosition(ctx context.Context, estateUUID string) (heights []model.TreeHeight, err error) {

	sql := `
		WITH estate AS (
				SELECT width, length, id
				FROM swp_mst_estate
				WHERE uuid = ?
		), grid AS (
				SELECT x, y
				FROM generate_series(1, (SELECT length FROM estate)) AS x,
						generate_series(1, (SELECT width FROM estate)) AS y
		)
		SELECT g.x AS position_x, g.y AS position_y, COALESCE(t.height, 0) AS height, t.*
		FROM grid g
		LEFT JOIN swp_trx_tree_estate t ON t.position_x = g.x AND t.position_y = g.y and t.delete_time is null and t.id_mst_estate = (SELECT id FROM estate)
		ORDER BY g.y ASC, 
			CASE
			WHEN g.y % 2 = 0 THEN g.x END DESC,
			CASE
			WHEN g.y % 2 = 1 THEN g.x END ASC`

	err = c.XormEngine.SQL(sql, estateUUID).Find(&heights)
	if err != nil {
		err = errors.Wrap(err, WrapMsgGetEstateTreesHeightPosition)
		return
	}

	return

}
