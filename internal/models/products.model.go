package models

import "time"

type Products struct {
	Id_product   string     `db:"id_product" form:"id_product"`
	Name_product string     `db:"name_product" form:"name_product"`
	Description  string     `db:"description" form:"description"`
	Stok         int        `db:"stok" form:"stok"`
	Image        string     `db:"favorite" form:"favorite"`
	Favorite     string     `db:"image" form:"image"`
	Create_at    *time.Time `db:"create_at" form:"create_at"`
	Update_at    *time.Time `db:"update_at" form:"update_at"`
}

type Meta_Products struct {
	Next       string
	Prev       string
	Last_page  string
	Total_data string
}
