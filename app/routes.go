package app

import (
	"github.com/rdelvallej32/bookstore_users-api/controllers/health"
	"github.com/rdelvallej32/bookstore_users-api/controllers/users"
)

func setRoutes() {
	router.GET("/health", health.Health)
	router.GET("/users/:user_id", users.Get)
	// router.GET("/users/search", controllers.SeachUser)
	router.POST("/users", users.Create)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
}
