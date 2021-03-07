package repository

import (
"github.com/jinzhu/gorm"
"github.com/solkn/Voting_System/entity"
)

type CandidateGormRepo struct {

	conn *gorm.DB
}
func NewPartyGormRepo(db *gorm.DB)*CandidateGormRepo {
	return &CandidateGormRepo{conn: db}
}

func (candidateRepo *CandidateGormRepo) Candidates()([]entity.Candidate,[]error){

	var candidates []entity.Candidate

	errs := candidateRepo.conn.Preloads("party").Find(&candidates).GetErrors()

	if(len(errs)>0){
		return nil,errs
	}

	return candidates,errs
}

func (candidateRepo *CandidateGormRepo)Candidate(id uint)(*entity.Candidate,[]error){

	candidate:= entity.Candidate{}

	errs:= candidateRepo.conn.Preloads("party").First("candidate",id).GetErrors()
	if(len(errs) > 0){

		return nil,errs
	}
	return &candidate,errs
}

func (candidateRepo *CandidateGormRepo)UpdateCandidate(party *entity.Candidate)(*entity.Candidate,[]error){

	errs:= candidateRepo.conn.Save(party).GetErrors()
	if(len(errs) >0){
		return nil, errs
	}

	return party,errs
}
func (candidateRepo *CandidateGormRepo)StoreCandidate(candidate *entity.Candidate)(*entity.Candidate,[]error){
	errs:= candidateRepo.conn.Create(candidate).GetErrors()
	if(len(errs) >0){
		return nil,errs
	}
	return candidate,errs
}
func (candidateRepo *CandidateGormRepo) DeleteCandidate(id uint)(*entity.Candidate,[]error){

	candidate,errs:= candidateRepo.Candidate(id)

	if(len(errs) > 0){

		return nil,errs
	}
	errs = candidateRepo.conn.Delete(candidate,id).GetErrors()

	if(len(errs) > 0){
		return nil,errs

	}

	return candidate,errs
}
