package models

import "time"

type Users struct {
	Id_user             string     `db:"id_user" form:"id_user"`
	Displayname         string     `db:"displayname" form:"displayname"`
	First_name          string     `db:"first_name" form:"first_name"`
	Last_name           string     `db:"last_name" form:"last_name"`
	Gender              string     `db:"gender" form:"gender"`
	Phone               string     `db:"phone" form:"phone"`
	Email               string     `db:"email" form:"email"`
	Pass                string     `db:"pass" form:"pass"`
	Birth_date          string     `db:"birth_date" form:"birth_date"`
	Status_verification string     `db:"status_verification" form:"status_verification"`
	Role                string     `db:"role" form:"role"`
	Image               string     `db:"image" form:"image"`
	Create_at           *time.Time `db:"create_at" form:"create_at"`
	Update_at           *time.Time `db:"update_at" form:"update_at"`
}

type Meta_Users struct {
	Next       string
	Prev       string
	Last_page  string
	Total_data string
}
