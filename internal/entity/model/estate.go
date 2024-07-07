package model

type EstateDB struct {
	ID     int64  `xorm:"'id' autoincr pk" json:"-"`
	UUID   string `xorm:"'uuid'" json:"id,omitempty"`
	Width  int32  `xorm:"'width'" json:"width"`
	Length int32  `xorm:"'length'" json:"length"`
}

type InsertEstateResponse struct {
	ID string `json:"id"`
}

type InsertEstateRequest struct {
	Width  int32 `xorm:"'width'" json:"width"`
	Length int32 `xorm:"'length'" json:"length"`
}

var (
	MockEstateDB = []EstateDB{}
)

type EstateJoinTrxTree struct {
	Estate EstateDB `xorm:"extends"`
	Tree   TrxTree  `xorm:"extends"`
}

type EstateStats struct {
	TreeCount     int16 `xorm:"'tree_count'" json:"count"`
	TreeMaxHeight int   `xorm:"'max_height'" json:"max"`
	TreeMinHeight int   `xorm:"'min_height'" json:"min"`
	TreeMedian    int   `xorm:"'median_height'" json:"median"`
}

type EstateDronePlanResponse struct {
	Distance int32        `json:"distance"`
	Rest     *Coordinates `json:"rest,omitempty"`
}

type Coordinates struct {
	PositionX int32 `json:"x"`
	PositionY int32 `json:"y"`
}

type TreeHeight struct {
	Height    int32 `xorm:"height"`
	PositionX int32 `xorm:"position_x"`
	PositionY int32 `xorm:"position_y"`
}

type GetDronePlanParams struct {
	UUID        string `schema:"uuid"`
	MaxDistance int32  `schema:"max-distance"`
}
