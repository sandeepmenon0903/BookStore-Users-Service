package app

import "github.com/gin-gonic/gin"
var(
	_router = gin.Default()
)
func InitiateApp(){
	MapUrls()
	_router.Run(":8082")
}
