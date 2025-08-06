package presistence

import (
	"go-web-layout/internal/domain"

	"gorm.io/gorm"
)


type GormUserRepository struct{
	Db *gorm.DB
}


func NewGormUserRepository(Db *gorm.DB)*GormUserRepository{
	return &GormUserRepository{
		Db: Db,
	}
}




func (gm *GormUserRepository)GetById(Id string)(*domain.User,error){
	return nil,nil
}



func (gm *GormUserRepository)Create(user *domain.User)error{
	return nil
}


func (gm *GormUserRepository)Update(user *domain.User)error{
	return nil
}



func (gm *GormUserRepository)Delete(user *domain.User)error{
	return nil
}