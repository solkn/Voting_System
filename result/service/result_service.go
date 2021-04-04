package service

import (
	"github.com/solkn/Voting_System/entity"
	"github.com/solkn/Voting_System/result"
)

type ResultService struct {
	resultRepo result.ResultRepository
}

func NewResultService(resultRepo result.ResultRepository) *ResultService {
	return &ResultService{resultRepo: resultRepo}
}

func (r *ResultService) Results() ([]entity.Result, []error) {
	results,errs:= r.resultRepo.Results()
	if len(errs) > 0 {
		return nil,errs
	}
	return results,errs

	
}

func (r *ResultService) Result(id uint) (*entity.Result, []error) {
	result,errs:=r.resultRepo.Result(id)
	if len(errs) > 0{
		return nil,errs

	}
	return result,errs
}

func (r *ResultService) StoreResult(result *entity.Result) (*entity.Result, []error) {
	rslt, errs := r.resultRepo.StoreResult(result)
	if (len(errs) > 0) {
		return nil, errs
	}
	return rslt, errs
}
func (r *ResultService) UpdateResult(result *entity.Result) (*entity.Result, []error) {
	rslt,errs:=r.resultRepo.UpdateResult(result)
	if(len(errs)>0){
		return nil,errs

	}
	return rslt,errs

}

func (r *ResultService) DeleteResult(id uint) (*entity.Result, []error) {
	result,errs:=r.resultRepo.DeleteResult(id)
	if(len(errs)>0){
		return nil,errs
	}
	return result,errs
}
