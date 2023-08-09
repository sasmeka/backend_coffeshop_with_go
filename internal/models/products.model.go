package models

import "time"

type Products struct {
	Id_product              string                    `json:"id_product" db:"id_product" form:"id_product"`
	Name_product            string                    `json:"name_product" db:"name_product" form:"name_product"`
	Description             string                    `json:"description" db:"description" form:"description"`
	Image                   string                    `json:"favorite" db:"favorite" form:"favorite"`
	Favorite                string                    `json:"image" db:"image" form:"image"`
	Create_at               *time.Time                `json:"create_at" db:"create_at" form:"create_at"`
	Update_at               *time.Time                `json:"update_at" db:"update_at" form:"update_at"`
	Product_size            []Product_Size            `json:"product_size" db:"product_size"`
	Product_delivery_method []Product_Delivery_Method `json:"product_delivery_method" db:"product_delivery_method"`
}

type Productsset struct {
	Id_product              string         `json:"id_product" db:"id_product" form:"id_product"`
	Name_product            string         `json:"name_product" db:"name_product" form:"name_product"`
	Description             string         `json:"description" db:"description" form:"description"`
	Favorite                string         `json:"favorite" db:"favorite" form:"favorite"`
	Image                   string         `json:"image" db:"image" form:"image"`
	Create_at               *time.Time     `json:"create_at" db:"create_at" form:"create_at"`
	Update_at               *time.Time     `json:"update_at" db:"update_at" form:"update_at"`
	Product_size            []Product_Size `json:"product_size" db:"product_size"`
	Product_delivery_method []string       `json:"product_delivery_method" db:"product_delivery_method"`
}

type Product_Size struct {
	Id_product_size string `json:"id_product_size,omitempty" db:"id_product_size"`
	Id_size         string `json:"id_size" db:"id_size"`
	Name_size       string `json:"name_size" db:"name_size"`
	Abbreviation    string `json:"abbreviation" db:"abbreviation"`
	Stok            int    `json:"stok" db:"stok" form:"stok"`
	Price           int    `json:"price" db:"price" form:"price"`
}
type Product_Delivery_Method struct {
	Id_dm   string `json:"id_dm" db:"id_dm"`
	Name_dm string `json:"name_dm" db:"name_dm"`
}

type Products_Sizes struct {
	Id_product_size string `json:"id_product_size" db:"id_product_size"`
	Id_product      string `json:"id_product" db:"id_product"`
	Id_size         string `json:"id_size" db:"id_size"`
}

type Products_Delivery_Methods struct {
	Id_product_delivery_method string `json:"id_product_delivery_method" db:"id_product_delivery_method"`
	Id_product                 string `json:"id_product" db:"id_product"`
	Id_dm                      string `json:"id_dm" db:"id_dm"`
}
