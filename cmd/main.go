package main

import (
	"fmt"
	"go-web-layout/internal/application"
	"go-web-layout/internal/config"
	"go-web-layout/internal/delivery/http/handler"
	"go-web-layout/internal/delivery/http/routes"
	"go-web-layout/internal/infrastructure/presistence"
	"go-web-layout/internal/pkg/db"
	"log"

	"github.com/gin-gonic/gin"
)





func main(){
	


	config,err := config.LoadConfig()
	if err!=nil{
		log.Printf("error loading config: %s",err.Error())
		return 
	}

	fmt.Println(config)

	dbdriver,dberr := db.NewGormDb(*config)



	if dberr!=nil{
		log.Printf("Error during connect to database: %s",dberr.Error())
		return 
	}

	db:= dbdriver.GetDb()



	ginRouter := gin.Default()

	repository:= presistence.NewGormUserRepository(db)
	service := application.NewUserService(repository)
	controller:= handler.NewUserHandler(service)
	routes.RegisterUserRoutes(ginRouter,controller)


	rerr :=ginRouter.Run(config.AppHost+":"+config.AppPort)
	if rerr!=nil{
		log.Printf("Error while start gin router: %s",rerr.Error())
		return 
	}

	log.Println("Server started")
}


