package repositories

import (
	"errors"
	"sasmeka/coffeeshop/internal/models"

	"github.com/jmoiron/sqlx"
)

type Repo_Auth struct {
	*sqlx.DB
}

func New_Auth(db *sqlx.DB) *Repo_Auth {
	return &Repo_Auth{db}
}

func (r *Repo_Auth) Get_User(data *models.Users) (*models.Users, error) {
	var result models.Users

	q := `SELECT id_user, email, "role", "pass" FROM public.users WHERE email = ?`

	if err := r.Get(&result, r.Rebind(q), data.Email); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New("email not found")
		}

		return nil, err
	}
	return &result, nil
}
func (r *Repo_Auth) Register_rep(data *models.Users) (string, error) {
	query := `INSERT INTO public.users(
		gender, 
		phone, 
		email, 
		pass,
		image
	)VALUES(
		:gender, 
		:phone, 
		:email, 
		:pass,
		:image
	);`
	if data.Phone == "" || data.Email == "" || data.Pass == "" {
		return "", errors.New("all forms must be filled")
	}
	if data.Gender == "" {
		data.Gender = "male"
	}
	if data.Image == "" {
		data.Image = "/static/img/Default_Profile.png"
	}
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "register successful.", nil
}
