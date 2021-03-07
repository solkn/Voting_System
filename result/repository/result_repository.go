package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/solkn/Voting_System/entity"
)

type ResultGormRepo struct {
	conn *gorm.DB
}

func NewResultGormRepo(conn *gorm.DB) *ResultGormRepo {
	return &ResultGormRepo{conn: conn}
}

func (r *ResultGormRepo) Results() ([]entity.Result, []error) {
	var results []entity.Result

	errs := r.conn.Preload("result").Find(&results).GetErrors()

	if(len(errs)>0){
		return nil,errs
	}

	return results,errs
}

func (r *ResultGormRepo) Result(id uint) (*entity.Result, []error) {
	result:= entity.Result{}

	errs:= r.conn.Preload("result").First("result",id).GetErrors()
	if(len(errs) > 0){

		return nil,errs
	}
	return &result,errs
}

func (r *ResultGormRepo) StoreResult(result *entity.Result) (*entity.Result, []error) {
	errs:= r.conn.Create(result).GetErrors()
	if(len(errs) >0){
		return nil,errs
	}
	return result,errs
}

func (r *ResultGormRepo) UpdateResult(result *entity.Result) (*entity.Result, []error) {

	errs:= r.conn.Save(result).GetErrors()
	if(len(errs) >0){
		return nil, errs
	}

	return result,errs
}

func (r *ResultGormRepo) DeleteResult(id uint) (*entity.Result, []error) {
	result,errs:= r.Result(id)

	if(len(errs) > 0){

		return nil,errs
	}
	errs = r.conn.Delete(result,id).GetErrors()

	if(len(errs) > 0){
		return nil,errs

	}

	return result,errs
}
