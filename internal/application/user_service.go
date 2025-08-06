package application

import "go-web-layout/internal/domain"


type UserService struct{
	Repo *domain.UserRepository
}



func NewUserService(userrepo *domain.UserRepository)*UserService{
	return &UserService{
		Repo: userrepo,
	}
}




func (service *UserService)GetUser(id string){

}

func (service *UserService)UpdateUser(user *domain.User){

}


func (service *UserService)CreateUser(user *domain.User){

}



func (service *UserService)DeleteUser(user *domain.User){
	
}


