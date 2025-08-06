package config

import (
	"github.com/spf13/viper"
)


type Config struct{
	AppKey string
	AppPort string
	AppHost string

	DbDriver string
	DbHost string
	DbPort string
	DBName string
	DbUsername string
	DbPassword string
}




func LoadConfig()(*Config,error){
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig();
	if err!=nil{
		return nil,err
	}

	var config *Config

	merr := viper.Unmarshal(&config)

	

	if merr!=nil{
		return nil,merr
	}

	return config,nil


}