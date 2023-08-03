package models

import "time"

type Delivery_Methods struct {
	Id_dm     string     `db:"id_dm" form:"id_dm"`
	Name_dm   string     `db:"name_dm" form:"name_dm"`
	Create_at *time.Time `db:"create_at" form:"create_at"`
	Update_at *time.Time `db:"update_at" form:"update_at"`
}
