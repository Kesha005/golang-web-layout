package handler

import (
	"go-web-layout/internal/application"

	"github.com/gin-gonic/gin"
)



type UserHandler struct{
	userh *application.UserService
}


func NewUserHandler(service *application.UserService)*UserHandler{
	return &UserHandler{
		userh: service,
	}
}


func(ush *UserHandler)GetUser(ctx *gin.Context){

}

func(ush *UserHandler)CreateUser(ctx *gin.Context){
	
}


func(ush *UserHandler)UpdateUser(ctx *gin.Context){
	
}

func(ush *UserHandler)DeleteUser(ctx *gin.Context){
	
}