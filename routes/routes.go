package routes

import (
	user "CRUD/controllers/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

//StartGin function
func StartGin() {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/users", user.GetAllUser)
		api.GET("/users/:id", user.GetUser)
		api.POST("/users", user.CreateUser)
		api.PUT("/users/:id", user.UpdateUser)
		api.DELETE("/users/:id", user.DeleteUser)
	}
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":8000")
}
