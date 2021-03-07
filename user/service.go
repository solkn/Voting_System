package user

import "github.com/solkn/Voting_System/entity"


type UserServices interface {

	Users()([]entity.User,[]error)
	User(id uint)(*entity.User,[]error)
	StoreUser(user *entity.User)(*entity.User,[]error)
	UpdateUser(user *entity.User)(*entity.User,[]error)
	DeleteUser(id uint)(*entity.User,[]error)
	UserByUserName(user entity.User) (*entity.User, []error)
	PhoneExists(phone string) bool
	EmailExists(email string) bool
	UserRoles(*entity.User) ([]entity.Role, []error)
}
type RoleServices interface {
	Roles() ([]entity.Role, []error)
	Role(id uint) (*entity.Role, []error)
	RoleByName(name string) (*entity.Role, []error)
	UpdateRole(role *entity.Role) (*entity.Role, []error)
	DeleteRole(id uint) (*entity.Role, []error)
	StoreRole(role *entity.Role) (*entity.Role, []error)
}