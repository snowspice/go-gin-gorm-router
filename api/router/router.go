package router
import (
	"github.com/gin-gonic/gin"
	. "github.com/snowspice/go-gin-gorm-router/api/apis"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/users", Users)

	router.POST("/user", Add)

	router.PUT("/user/:id", Update)

	router.DELETE("/user/:id", Delete)

	router.POST("/user/page",FindPage)

	return router
}