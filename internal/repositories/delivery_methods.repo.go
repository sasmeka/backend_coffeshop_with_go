package repositories

import (
	"errors"
	"math"
	"sasmeka/coffeeshop/config"
	"sasmeka/coffeeshop/internal/models"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type Repo_Delivery_Methods struct {
	*sqlx.DB
}

func New_Delivery_Methods(db *sqlx.DB) *Repo_Delivery_Methods {
	return &Repo_Delivery_Methods{db}
}

func (r *Repo_Delivery_Methods) Get_Data(data *models.Delivery_Methods, page string, limit string) (*config.Result, error) {
	deliver_method_data := []models.Delivery_Methods{}
	var meta_delivery_method config.Metas

	var offset int = 0
	var page_int, _ = strconv.Atoi(page)
	var limit_int, _ = strconv.Atoi(limit)
	if limit == "" {
		limit_int = 5
	}
	if page == "" {
		page_int = 1
	}
	if page_int > 0 {
		offset = (page_int - 1) * limit_int
	} else {
		offset = 0
	}

	count_data := r.Get_Count_Data()

	if count_data <= 0 {
		meta_delivery_method.Next = ""
	} else {
		if float64(page_int) == math.Ceil(float64(count_data)/float64(limit_int)) {
			meta_delivery_method.Next = ""
		} else {
			meta_delivery_method.Next = strconv.Itoa(page_int + 1)
		}
	}

	if page_int == 1 {
		meta_delivery_method.Prev = ""
	} else {
		meta_delivery_method.Prev = strconv.Itoa(page_int - 1)
	}

	if int(math.Ceil(float64(count_data)/float64(limit_int))) != 0 {
		meta_delivery_method.Last_page = strconv.Itoa(int(math.Ceil(float64(count_data) / float64(limit_int))))
	} else {
		meta_delivery_method.Last_page = ""
	}

	if count_data != 0 {
		meta_delivery_method.Total_data = strconv.Itoa(count_data)
	} else {
		meta_delivery_method.Last_page = ""
	}

	r.Select(&deliver_method_data, `SELECT * FROM public.delivery_methods LIMIT $1 OFFSET $2`, limit_int, offset)
	if len(deliver_method_data) == 0 {
		return nil, errors.New("data not found.")
	}
	return &config.Result{Data: deliver_method_data, Meta: meta_delivery_method}, nil
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
			name_dm=:name_dm,
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
