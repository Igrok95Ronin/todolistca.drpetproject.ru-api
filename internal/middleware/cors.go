package middleware

import (
	"github.com/rs/cors"
	"net/http"
)

func CorsSettings() *cors.Cors {
	return cors.New(cors.Options{
		// Указываем разрешенные HTTP методы, только POST
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodPut,
			http.MethodOptions,
		},
		// Указываем разрешенные источники домен
		AllowedOrigins: []string{"https://todolist.drpetproject.ru"},

		AllowedHeaders: []string{
			"X-Api-Password",
			"Content-Type",
			"Authorization",
		},
	})
}
