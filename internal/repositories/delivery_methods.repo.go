package repositories

import (
	"errors"
	"sasmeka/coffeeshop/internal/models"

	"github.com/jmoiron/sqlx"
)

type Repo_Delivery_Methods struct {
	*sqlx.DB
}

func New_Delivery_Methods(db *sqlx.DB) *Repo_Delivery_Methods {
	return &Repo_Delivery_Methods{db}
}

func (r *Repo_Delivery_Methods) Get_Data(data *models.Delivery_Methods, page int, limit int) ([]models.Delivery_Methods, error) {
	deliver_method_data := []models.Delivery_Methods{}
	r.Select(&deliver_method_data, `SELECT * FROM public.delivery_methods LIMIT $1 OFFSET $2`, limit, page)
	if len(deliver_method_data) == 0 {
		return nil, errors.New("data not found.")
	}
	return deliver_method_data, nil
}

func (r *Repo_Delivery_Methods) Get_Count_by_Id(id string) int {
	var count_data int
	r.Get(&count_data, "SELECT count(*) FROM public.delivery_methods WHERE id_dm=$1", id)
	return count_data
}

func (r *Repo_Delivery_Methods) Get_Count_Data() int {
	var id int
	r.Get(&id, "SELECT count(*) FROM public.delivery_methods")
	return id
}

func (r *Repo_Delivery_Methods) Insert_Data(data *models.Delivery_Methods) (string, error) {
	query := `INSERT INTO public.delivery_methods(
			name_dm
		)VALUES(
			:name_dm
		);`
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "add delivery method data successful", nil
}
func (r *Repo_Delivery_Methods) Update_Data(data *models.Delivery_Methods) (string, error) {
	query := `UPDATE public.delivery_methods SET
			name_dm=:name_dm
			update_at=now()
			WHERE id_dm=:id_dm;`
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "update delivery method data successful", nil
}
func (r *Repo_Delivery_Methods) Delete_Data(data *models.Delivery_Methods) (string, error) {
	query := `DELETE FROM public.delivery_methods WHERE id_dm=:id_dm;`
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "delete delivery method data successful", nil
}
