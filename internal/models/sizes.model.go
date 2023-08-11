package models

import "time"

type Sizes struct {
	Id_size      string     `db:"id_size" form:"id_size" valid:"-"`
	Name_size    string     `db:"name_size" form:"name_size" valid:"required"`
	Abbreviation string     `db:"abbreviation" form:"abbreviation" valid:"required"`
	Create_at    *time.Time `db:"create_at" form:"create_at" valid:"-"`
	Update_at    *time.Time `db:"update_at" form:"update_at" valid:"-"`
}
