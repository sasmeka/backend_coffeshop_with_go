package repositories

import (
	"errors"
	"fmt"
	"log"
	"sasmeka/coffeeshop/internal/models"

	"github.com/jmoiron/sqlx"
)

type Repo_Products struct {
	*sqlx.DB
}

func New_Products(db *sqlx.DB) *Repo_Products {
	return &Repo_Products{db}
}

func (r *Repo_Products) Get_Data(data *models.Products, page int, limit int, search string, orderby string) ([]models.Products, error) {
	var list_products_data []models.Products
	Products_data := models.Products{}
	if search == "" {
		search = ""
	} else {
		search = fmt.Sprintf(` AND LOWER(name_product) like LOWER('%s')`, "%"+search+"%")
	}
	if orderby == "" {
		orderby = ""
	} else {
		orderby = fmt.Sprintf(` ORDER BY %s`, orderby)
	}
	rows, err := r.Queryx(`select * from products WHERE TRUE `+search+orderby+` LIMIT $1 OFFSET $2`, limit, page)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var list_product_size []models.Product_Size
		var list_product_delivery_method []models.Product_Delivery_Method
		err := rows.StructScan(&Products_data)
		if err != nil {
			log.Fatalln(err)
		}
		rows, _ := r.Queryx("select s.id_size,s.name_size,s.abbreviation,ps.price,ps.stok from products_sizes ps left join sizes s on ps.id_size=s.id_size where ps.id_product=$1", Products_data.Id_product)
		for rows.Next() {
			var product_size models.Product_Size
			err := rows.Scan(&product_size.Id_size, &product_size.Name_size, &product_size.Abbreviation, &product_size.Price, &product_size.Stok)
			if err != nil {
				log.Fatalln(err)
			}
			// fmt.Println(product_size)
			list_product_size = append(list_product_size, product_size)
		}
		rows1, _ := r.Queryx("select s.id_dm,s.name_dm from products_delivery_methods ps left join delivery_methods s on ps.id_dm=s.id_dm where ps.id_product=$1", Products_data.Id_product)
		for rows1.Next() {
			var product_delivery_method models.Product_Delivery_Method
			err1 := rows1.Scan(&product_delivery_method.Id_dm, &product_delivery_method.Name_dm)
			if err1 != nil {
				log.Fatalln(err1)
			}
			// fmt.Println(product_delivery_method)
			list_product_delivery_method = append(list_product_delivery_method, product_delivery_method)
		}

		Products_data.Product_size = list_product_size
		Products_data.Product_delivery_method = list_product_delivery_method

		rows.Close()

		// fmt.Println(Products_data)
		list_products_data = append(list_products_data, Products_data)
	}
	rows.Close()
	return list_products_data, nil
}

func (r *Repo_Products) Get_Count_by_Id(id string) int {
	var count_data int
	r.Get(&count_data, "SELECT count(*) FROM public.products WHERE id_product=$1", id)
	return count_data
}

func (r *Repo_Products) Get_Count_Data(search string) int {
	if search == "" {
		search = ""
	} else {
		search = fmt.Sprintf(` AND LOWER(name_product) like LOWER('%s')`, "%"+search+"%")
	}
	var id int
	r.Get(&id, `SELECT count(*) FROM public.products WHERE TRUE `+search)
	return id
}

func (r *Repo_Products) Insert_Data(data *models.Productsset) (string, error) {
	if data.Name_product == "" || data.Description == "" {
		return "", errors.New("field product name & description must be filled.")
	}
	if data.Image == "" {
		data.Image = "/static/product/Default_Product.png"
	}
	if data.Favorite == "" {
		data.Favorite = "0"
	}
	tx := r.MustBegin()
	var new_id string
	tx.Get(&new_id, "select uuid_generate_v4()")
	data.Id_product = new_id
	tx.NamedExec(`INSERT INTO public.products (id_product,name_product, description, image, favorite) VALUES(:id_product,:name_product, :description, :image, :favorite);`, data)
	for i := range data.Product_size {
		tx.MustExec("INSERT INTO public.products_sizes (id_product, id_size,stok,price) VALUES ($1, $2, $3, $4)", &new_id, &data.Product_size[i].Id_size, &data.Product_size[i].Stok, &data.Product_size[i].Price)
	}
	for i := range data.Product_delivery_method {
		tx.MustExec("INSERT INTO public.products_delivery_methods (id_product, id_dm) VALUES ($1, $2)", &new_id, &data.Product_delivery_method[i])
	}
	tx.Commit()

	return "add product data successful", nil
}
func (r *Repo_Products) Update_Data(data *models.Productsset) (string, error) {
	if data.Name_product == "" || data.Description == "" {
		return "", errors.New("field product name & description must be filled.")
	}
	if data.Image == "" {
		data.Image = "/static/product/Default_Product.png"
	}
	if data.Favorite == "" {
		data.Favorite = "0"
	}

	fmt.Println(data)
	var id string
	id = data.Id_product

	tx := r.MustBegin()
	tx.NamedExec(`UPDATE public.products SET name_product=:name_product, description=:description, image=:image, favorite=:favorite WHERE id_product=:id_product;`, data)
	tx.MustExec(`DELETE FROM public.products_sizes WHERE id_product=$1;`, &id)
	tx.MustExec(`DELETE FROM public.products_delivery_methods WHERE id_product=$1`, &id)
	for i := range data.Product_size {
		tx.MustExec("INSERT INTO public.products_sizes (id_product, id_size,stok,price) VALUES ($1, $2, $3, $4)", &id, &data.Product_size[i].Id_size, &data.Product_size[i].Stok, &data.Product_size[i].Price)
	}
	for i := range data.Product_delivery_method {
		tx.MustExec("INSERT INTO public.products_delivery_methods (id_product, id_dm) VALUES ($1, $2)", &id, &data.Product_delivery_method[i])
	}
	tx.Commit()

	return "update product data successful", nil
}
func (r *Repo_Products) Delete_Data(data *models.Products, data2 *models.Products_Sizes, data3 *models.Products_Delivery_Methods) (string, error) {
	tx := r.MustBegin()
	_, err1 := tx.NamedExec(`DELETE FROM public.products_sizes WHERE id_product=:id_product;`, data2)
	if err1 != nil {
		return "", err1
	}
	_, err3 := tx.NamedExec(`DELETE FROM public.products_delivery_methods WHERE id_product=:id_product;`, data3)
	if err3 != nil {
		return "", err3
	}
	_, err2 := tx.NamedExec(`DELETE FROM public.products WHERE id_product=:id_product;`, data)
	if err2 != nil {
		return "", err2
	}
	tx.Commit()
	return "delete product data successful", nil
}