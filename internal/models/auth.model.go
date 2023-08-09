package models

import "time"

type Auth struct {
	Id_user string `db:"id_user" form:"id_user" valid:"-"`
	Email   string `db:"email" json:"email" form:"email" valid:"required~e-mail is required"`
	Pass    string `db:"pass" json:"pass" form:"pass" valid:"required~password is required,stringlength(6|1024)~password of at least 6 characters"`
	Role    string `db:"role" form:"role" valid:"-"`
	Phone   string `db:"phone" form:"phone" valid:"-"`
}

type Users_auth struct {
	Id_user             string     `db:"id_user" form:"id_user" valid:"-"`
	Displayname         string     `db:"displayname" form:"displayname" valid:"-"`
	First_name          string     `db:"first_name" form:"first_name" valid:"-"`
	Last_name           string     `db:"last_name" form:"last_name" valid:"-"`
	Gender              string     `db:"gender" form:"gender" valid:"-"`
	Phone               string     `db:"phone" form:"phone" valid:"-"`
	Email               string     `db:"email" json:"email" form:"email" valid:"required~e-mail is required"`
	Pass                string     `db:"pass" json:"pass" form:"pass" valid:"required~password is required,stringlength(6|1024)~password of at least 6 characters"`
	Birth_date          string     `db:"birth_date" form:"birth_date" valid:"-"`
	Status_verification string     `db:"status_verification" form:"status_verification" valid:"-"`
	Role                string     `db:"role" form:"role" valid:"-"`
	Image               string     `db:"image" form:"image" valid:"-"`
	Create_at           *time.Time `db:"create_at" form:"create_at" valid:"-"`
	Update_at           *time.Time `db:"update_at" form:"update_at" valid:"-"`
}
