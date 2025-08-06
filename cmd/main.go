package main

import (
	"fmt"
	"go-web-layout/internal/config"
	"log"
)





func main(){
	


	config,err := config.LoadConfig()
	if err!=nil{
		log.Printf("error loading config: %s",err.Error())
		return 
	}

	fmt.Println(config)

	
}


