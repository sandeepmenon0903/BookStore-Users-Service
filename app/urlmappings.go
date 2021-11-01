package app

import (
	"github.com/sandeepmenon0903/BookStore-Users-Service/Controllers/HealthCheck"
	"github.com/sandeepmenon0903/BookStore-Users-Service/Controllers/Users"
)

func MapUrls() {
	_router.GET("/heathCheck", HealthCheck.HealthCheck)
	//Users Controller Urls
	_router.POST("/user", Users.CreateUser)
	_router.GET("/user/:userId", Users.GetUser)
	_router.POST("/user/:userId", Users.UpdateUser)
	_router.PATCH("/user/:userId", Users.UpdateUser)
	_router.DELETE("/user/:userId", Users.DeleteUser)
	_router.GET("internal/user/search", Users.Search)
}
