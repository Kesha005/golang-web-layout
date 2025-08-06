package db

import (
	"errors"
	"fmt"
	"go-web-layout/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)




type Database interface{
	GetDb()*gorm.DB
}



type GormDb struct{
	conn *gorm.DB
}

func NewGormDb(cfg config.Config)(Database,error){
	var dialector gorm.Dialector
	switch cfg.DbDriver{


	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.DbUsername,
			cfg.DbPassword,
			cfg.DbHost,
			cfg.DbPort,
			cfg.DBName,
		)
		dialector = mysql.Open(dsn)
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			cfg.DbHost,
			cfg.DbUsername,
			cfg.DbPassword,
			cfg.DBName,
			cfg.DbPort,
		)
		dialector = postgres.Open(dsn)
	default :
		return nil,errors.New("invalid database driver")
	}

	db,err:= gorm.Open(dialector,&gorm.Config{})
	if err!=nil{
		return nil,err
	}


	return GormDb{conn: db},nil
}



func (db GormDb)GetDb()*gorm.DB{
	return db.conn
}