package app

import (
	"github.com/BookStore-Users-Service/Controllers/HealthCheck"
	"github.com/BookStore-Users-Service/Controllers/Users"
)

func MapUrls(){
 _router.GET("/heathCheck", HealthCheck.HealthCheck)
 //Users Controller Urls
 _router.POST("/user", Users.User)
}