package repositories

import (
	"errors"
	"sasmeka/coffeeshop/internal/models"

	"github.com/jmoiron/sqlx"
)

type Repo_Users struct {
	*sqlx.DB
}

func New_Users(db *sqlx.DB) *Repo_Users {
	return &Repo_Users{db}
}

func (r *Repo_Users) Get_Users(data *models.Users, page int, limit int) ([]models.Users, error) {
	users_data := []models.Users{}
	r.Select(&users_data, `SELECT id_user,displayname, first_name, last_name, gender, phone, email, birth_date, status_verification, "role", image, create_at, update_at FROM public.users LIMIT $1 OFFSET $2`, limit, page)
	if len(users_data) == 0 {
		return nil, errors.New("data not found.")
	}
	return users_data, nil
}

func (r *Repo_Users) Get_Count_by_Id(id string) int {
	var count_data int
	r.Get(&count_data, "SELECT count(*) FROM public.users WHERE id_user=$1", id)
	return count_data
}

func (r *Repo_Users) Get_Count_Users() int {
	var id int
	r.Get(&id, "SELECT count(*) FROM public.users")
	return id
}

func (r *Repo_Users) Insert_User(data *models.Users) (string, error) {
	query := `INSERT INTO public.users(
			displayname, 
			first_name, 
			last_name, 
			gender, 
			phone, 
			email, 
			pass, 
			birth_date,
			image
		)VALUES(
			:displayname, 
			:first_name, 
			:last_name, 
			:gender, 
			:phone, 
			:email, 
			:pass, 
			:birth_date,
			:image
		);`
	if data.Displayname == "" || data.First_name == "" || data.Last_name == "" || data.Phone == "" || data.Email == "" || data.Pass == "" || data.Birth_date == "" {
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
	return "add user data successful.", nil
}
func (r *Repo_Users) Update_User(data *models.Users) (string, error) {
	query := `UPDATE public.users SET
			displayname=:displayname, 
			first_name=:first_name, 
			last_name=:last_name, 
			gender=:gender, 
			phone=:phone, 
			email=:email, 
			pass=:pass, 
			birth_date=:birth_date,
			image=:image,
			update_at=now()
			WHERE id_user=:id_user;`
	if data.Displayname == "" || data.First_name == "" || data.Last_name == "" || data.Phone == "" || data.Email == "" || data.Pass == "" || data.Birth_date == "" {
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
	return "update user data successful", nil
}
func (r *Repo_Users) Delete_User(data *models.Users) (string, error) {
	query := `DELETE FROM public.users WHERE id_user=:id_user;`
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "delete user data successful", nil
}
