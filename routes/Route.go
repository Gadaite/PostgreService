package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)
import "PostgreService/controller"

func InitRoutes() *gin.Engine {
	route := gin.Default()
	trajactory := route.Group("/trajectory")
	{
		trajactory.GET("/objecttrajactory/:lastappeared_id", controller.ObjectTrajactoryApi)
		trajactory.GET("/machineTypeSearch", controller.MachinetypeSearch)
		trajactory.POST("/machineTypeUpdate", controller.MachinetypeUpdate)
		trajactory.POST("/machineTypeAdd", controller.MachinetypeAdd)
	}

	//	配置跨域访问
	route.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Max-Age", "3600")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	//	swagger整合
	url := ginSwagger.URL("http://localhost:2345/swagger/doc.json")
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return route
}
