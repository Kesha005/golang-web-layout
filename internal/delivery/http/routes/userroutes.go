package routes

import (
	"go-web-layout/internal/delivery/http/handler"

	"github.com/gin-gonic/gin"
)



func RegisteUserRoutes(r *gin.Engine,h *handler.UserHandler){
	r.GET("/users/:id", h.GetUser)
    r.POST("/users", h.CreateUser)
    r.PUT("/users/:id", h.UpdateUser)
    r.DELETE("/users/:id", h.DeleteUser)
}