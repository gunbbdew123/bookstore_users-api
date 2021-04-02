package app

import (
	// controllers/package name
	"github.com/gunbbdew123/bookstore_users-api/controllers/ping"
	"github.com/gunbbdew123/bookstore_users-api/controllers/users"
)

func mapUrls() {
	// we not excecute the function
	// just define that what function to be excecute
	// agaist this path /ping
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
