package config

import (
	"context"
)

type Operation func(ctx context.Context) error

type Metas struct {
	Next       interface{} `json:"next"`
	Prev       interface{} `json:"prev"`
	Last_page  interface{} `json:"last_page"`
	Total_data interface{} `json:"total_data"`
}

type Result struct {
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta"`
	Message interface{} `json:"message"`
}

// var CorsConfig = cors.Config{
// 	AllowOrigins:     []string{"http://www.google.com", "http://127.0.0.1:5500"},
// 	AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "HEAD", "OPTIONS"},
// 	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
// 	AllowCredentials: true,
// }
