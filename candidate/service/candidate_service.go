package service

import (
	"github.com/solkn/Voting_System/candidate"
	"github.com/solkn/Voting_System/entity"
)

type CandidateService struct {
	candidateRepo candidate.CandidateRepository
}

func NewCandidateService(candidateRepo candidate.CandidateRepository)*CandidateService {
	return &CandidateService{candidateRepo: candidateRepo}
}

func (c *CandidateService) Candidates() ([]entity.Candidate, []error) {
	candidates,errs:= c.candidateRepo.Candidates()
	if(len(errs)>0){
		return nil,errs
	}
	return candidates,errs
}

func (c *CandidateService) Candidate(id uint) (*entity.Candidate, []error) {
	candidate,errs:=c.candidateRepo.Candidate(id)
	if(len(errs)>0){
		return nil,errs
	}
	return candidate,errs
}

func (c *CandidateService) StoreCandidate(candidate *entity.Candidate) (*entity.Candidate, []error) {
	candt,errs:=c.candidateRepo.StoreCandidate(candidate)
	if(len(errs)>0){
		return nil,errs
	}
	return candt,errs
}

func (c *CandidateService) UpdateCandidate(candidate *entity.Candidate) (*entity.Candidate, []error) {
	candt,errs:=c.candidateRepo.UpdateCandidate(candidate)
	if(len(errs)>0){
		return nil,errs
	}
	return candt,errs
}

func (c *CandidateService) DeleteCandidate(id uint) (*entity.Candidate, []error) {
	candidate,errs:=c.candidateRepo.DeleteCandidate(id)
	if(len(errs)>0){
		return nil,errs
	}
	return candidate,errs
}
