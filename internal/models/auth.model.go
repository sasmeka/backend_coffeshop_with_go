package models

type Auth struct {
	Id_user string `db:"id_user" form:"id_user" valid:"-"`
	Email   string `db:"email" json:"email" form:"email" valid:"required~e-mail is required"`
	Pass    string `db:"pass" json:"pass" form:"pass" valid:"required~password is required,stringlength(6|1024)~password of at least 6 characters"`
	Role    string `db:"role" form:"role" valid:"-"`
	Phone   string `db:"phone" form:"phone" valid:"required~phone is required"`
}
