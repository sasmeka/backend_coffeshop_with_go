package repositories

import (
	"errors"
	"sasmeka/coffeeshop/internal/models"

	"github.com/jmoiron/sqlx"
)

type Repo_Sizes struct {
	*sqlx.DB
}

func New_Sizes(db *sqlx.DB) *Repo_Sizes {
	return &Repo_Sizes{db}
}

func (r *Repo_Sizes) Get_Data(data *models.Sizes, page int, limit int) ([]models.Sizes, error) {
	sizes_data := []models.Sizes{}
	r.Select(&sizes_data, `SELECT * FROM public.sizes LIMIT $1 OFFSET $2`, limit, page)
	if len(sizes_data) == 0 {
		return nil, errors.New("data note found.")
	}
	return sizes_data, nil
}

func (r *Repo_Sizes) Get_Count_by_Id(id string) int {
	var count_data int
	r.Get(&count_data, "SELECT count(*) FROM public.sizes WHERE id_size=$1", id)
	return count_data
}

func (r *Repo_Sizes) Get_Count_Data() int {
	var id int
	r.Get(&id, "SELECT count(*) FROM public.sizes")
	return id
}

func (r *Repo_Sizes) Insert_Data(data *models.Sizes) (string, error) {
	query := `INSERT INTO public.sizes(
			name_size, 
			abbreviation
		)VALUES(
			:name_size, 
			:abbreviation
		);`
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "add size data successful", nil
}
func (r *Repo_Sizes) Update_Data(data *models.Sizes) (string, error) {
	query := `UPDATE public.sizes SET
			name_size=:name_size, 
			abbreviation=:abbreviation,
			update_at=now()
			WHERE id_size=:id_size;`
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "update user data successful", nil
}
func (r *Repo_Sizes) Delete_Data(data *models.Sizes) (string, error) {
	query := `DELETE FROM public.sizes WHERE id_size=:id_size;`
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "delete user data successful", nil
}
