package models

import "time"

type Sizes struct {
	Id_size      string     `db:"id_size" form:"id_size"`
	Name_size    string     `db:"name_size" form:"name_size"`
	Abbreviation string     `db:"abbreviation" form:"abbreviation"`
	Create_at    *time.Time `db:"create_at" form:"create_at"`
	Update_at    *time.Time `db:"update_at" form:"update_at"`
}
