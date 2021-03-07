package service

import (
	"github.com/solkn/Voting_System/entity"
	"github.com/solkn/Voting_System/user"
)

type UserService struct {
	userRepo user.UserRepository
}

func NewUserService(userRepo user.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (u *UserService) Users() ([]entity.User, []error) {
	users,errs:=u.userRepo.Users()
	if(len(errs)>0){
		return nil,errs
	}
	return users,errs
}

func (u *UserService) User(id uint) (*entity.User, []error) {
	user,errs:=u.userRepo.User(id)
	if(len(errs)>0){
		return nil,errs
	}
	return user,errs
}

func (u *UserService) StoreUser(user *entity.User) (*entity.User, []error) {
	usr,errs:=u.userRepo.StoreUser(user)
	if(len(errs)>0){
		return nil,errs
	}
	return usr,errs
}

func (u *UserService) UpdateUser(user *entity.User) (*entity.User, []error) {
	usr,errs:=u.userRepo.UpdateUser(user)
	if(len(errs)>0){
		return nil,errs
	}
	return usr,errs
}

func (u *UserService) DeleteUser(id uint) (*entity.User, []error) {
	user,errs:=u.userRepo.DeleteUser(id)
	if(len(errs)>0){
		return nil,errs
	}
	return user,errs
}

func (u *UserService) UserByUserName(user entity.User) (*entity.User, []error) {
	usr,errs:=u.userRepo.UserByUserName(user)
	if(errs != nil){
		return nil,nil
	}
	return usr,nil
}

func (u *UserService) PhoneExists(phone string) bool {

	  exist:= u.userRepo.PhoneExists(phone)
	  return exist

}

func (u *UserService) EmailExists(email string) bool {
	exist:=u.userRepo.EmailExists(email)
	return exist
}

func (u *UserService) UserRoles(user *entity.User) ([]entity.Role, []error) {
	usr,errs:=u.userRepo.UserRoles(user)
	if(errs != nil){
		return nil,errs
	}
	return usr,errs
}

