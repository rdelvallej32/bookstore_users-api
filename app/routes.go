package app

import (
	"github.com/rdelvallej32/bookstore_users-api/controllers/health"
	"github.com/rdelvallej32/bookstore_users-api/controllers/users"
)

func setRoutes() {
	router.GET("/health", health.Health)
	router.GET("/users/:user_id", users.GetUser)
	// router.GET("/users/search", controllers.SeachUser)
	router.POST("/users", users.CreateUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
}
