package Users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func User(c *gin.Context){
	params:=c.Param("name")
	fmt.Println("hey hey")
	fmt.Println(params)
	c.String(http.StatusOK,params)
}