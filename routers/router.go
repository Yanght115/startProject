package router

import "github.com/gin-gonic/gin"

func GetRouter() *gin.Engine {
	r := gin.Default()
	registApiRouter(r)
	return r
}

func registApiRouter(r *gin.Engine) {
	r.Group("/api")
	{
		r.GET("/test", func(c *gin.Context) {
			c.String(200, "test")
		})
	}
}
