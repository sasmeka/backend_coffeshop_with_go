package repositories

import (
	"errors"
	"math"
	"sasmeka/coffeeshop/config"
	"sasmeka/coffeeshop/internal/models"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type Repo_Sizes_IF interface {
	Get_Data(data *models.Sizes, page string, limit string) (*config.Result, error)
	Get_Count_by_Id(id string) int
	Get_Count_Data() int
	Insert_Data(data *models.Sizes) (string, error)
	Update_Data(data *models.Sizes) (string, error)
	Delete_Data(data *models.Sizes) (string, error)
}

type Repo_Sizes struct {
	*sqlx.DB
}

func New_Sizes(db *sqlx.DB) *Repo_Sizes {
	return &Repo_Sizes{db}
}

func (r *Repo_Sizes) Get_Data(data *models.Sizes, page string, limit string) (*config.Result, error) {
	sizes_data := []models.Sizes{}
	var meta_size config.Metas

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
		meta_size.Next = ""
	} else {
		if float64(page_int) == math.Ceil(float64(count_data)/float64(limit_int)) {
			meta_size.Next = ""
		} else {
			meta_size.Next = strconv.Itoa(page_int + 1)
		}
	}

	if page_int == 1 {
		meta_size.Prev = ""
	} else {
		meta_size.Prev = strconv.Itoa(page_int - 1)
	}

	if int(math.Ceil(float64(count_data)/float64(limit_int))) != 0 {
		meta_size.Last_page = strconv.Itoa(int(math.Ceil(float64(count_data) / float64(limit_int))))
	} else {
		meta_size.Last_page = ""
	}

	if count_data != 0 {
		meta_size.Total_data = strconv.Itoa(count_data)
	} else {
		meta_size.Total_data = ""
	}

	r.Select(&sizes_data, `SELECT * FROM public.sizes LIMIT $1 OFFSET $2`, limit_int, offset)
	if len(sizes_data) == 0 {
		return nil, errors.New("data not found.")
	}
	return &config.Result{Data: sizes_data, Meta: meta_size}, nil
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
