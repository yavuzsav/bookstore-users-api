package app

import (
	"github.com/yavuzsav/bookstore-users-api/controllers/ping"
	"github.com/yavuzsav/bookstore-users-api/controllers/users"
)

func mapUrls()  {
	router.GET("/ping", ping.Ping)

	router.GET("api/users/:user_id", users.GetUser)
	router.POST("api/users", users.CreateUser)
}
