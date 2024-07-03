package model

import "time"

type TrxTree struct {
	ID         int64      `xorm:"'id' autoincr pk" json:"-"`
	UUID       string     `xorm:"'uuid'" json:"uuid"`
	PositionX  int32      `xorm:"'position_x'" json:"position_x"`
	PositionY  int32      `xorm:"'position_y'" json:"position_y"`
	Height     int        `xorm:"'height'" json:"height"`
	EstateID   int64      `xorm:"'id_mst_estate'" json:"-"`
	CreateTime time.Time  `xorm:"created 'create_time'" json:"-"`
	UpdateTime time.Time  `xorm:"updated 'update_time'" json:"-"`
	DeleteTime *time.Time `xorm:"deleted 'delete_time'" json:"-"`
}

type InsertNewTreeRequest struct {
	PositionX  int32  `json:"x"`
	PositionY  int32  `json:"y"`
	Height     int    `json:"height"`
	EstateUUID string `json:"uuid_mst_estate"`
}

type InsertNewTreeResponse struct {
	UUID string `json:"id"`
}
