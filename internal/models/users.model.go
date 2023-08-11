package models

import (
	"time"
)

type Users struct {
	Id_user             string  `db:"id_user" json:"id_user" form:"id_user" valid:"-"`
	Displayname         *string `db:"displayname" json:"displayname" form:"displayname" valid:"-"`
	First_name          *string `db:"first_name" json:"first_name" form:"first_name" valid:"-"`
	Last_name           *string `db:"last_name" json:"last_name" form:"last_name" valid:"-"`
	Gender              string  `db:"gender" json:"gender" form:"gender" valid:"-"`
	Phone               string  `db:"phone" json:"phone" form:"phone" valid:"-"`
	Email               string  `db:"email" json:"email" form:"email" valid:"required~e-mail is required"`
	Pass                string  `db:"pass" json:"pass,omitempty" form:"pass" valid:"required~password is required,stringlength(6|1024)~password of at least 6 characters"`
	Birth_date          *string `db:"birth_date" json:"birth_date" form:"birth_date" valid:"-"`
	Status_verification string  `db:"status_verification" json:"status_verification" form:"status_verification" valid:"-"`
	Role                string  `db:"role" json:"role" form:"role" valid:"-"`
	Image               string  `db:"image" json:"image,omitempty" valid:"-"`
	// File                multipart.FileHeader `json:"file" form:"file"`
	Create_at *time.Time `db:"create_at" json:"create_at" valid:"-"`
	Update_at *time.Time `db:"update_at" json:"update_at" valid:"-"`
}
