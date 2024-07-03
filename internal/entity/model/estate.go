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
