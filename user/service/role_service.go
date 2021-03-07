package service

import (
	"github.com/solkn/Voting_System/entity"
	"github.com/solkn/Voting_System/user"
)

type RoleService struct{
	roleRepo user.RoleRepository
}

func NewRoleService(roleRepo user.RoleRepository) *RoleService {
	return &RoleService{roleRepo: roleRepo}
}

func (r *RoleService) Roles() ([]entity.Role, []error) {
	roles,errs:=r.roleRepo.Roles()
	if(len(errs)>0){
		return nil,errs
	}
	return roles,errs
}

func (r *RoleService) Role(id uint) (*entity.Role, []error) {
    role,errs:=r.roleRepo.Role(id)
    if(len(errs)>0){
    	return nil,errs
	}
	return role,errs
}

func (r *RoleService) RoleByName(name string) (*entity.Role, []error) {
	role,errs:=r.roleRepo.RoleByName(name)
	if(len(errs)>0){
		return nil,errs
	}
	return role,errs
}

func (r *RoleService) UpdateRole(role *entity.Role) (*entity.Role, []error) {
	rol,errs:=r.roleRepo.UpdateRole(role)
	if(len(errs)>0){
		return nil,errs
	}
	return rol,errs
}

func (r *RoleService) DeleteRole(id uint) (*entity.Role, []error) {
	role,errs:=r.roleRepo.DeleteRole(id)
	if(len(errs)>0){
		return nil,errs
	}
	return role,errs
}

func (r *RoleService) StoreRole(role *entity.Role) (*entity.Role, []error) {
	rol,errs:=r.roleRepo.StoreRole(role)
	if(len(errs)>0){
		return nil,errs
	}
	return rol,errs
}
