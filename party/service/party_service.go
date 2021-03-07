package service

import (
	"github.com/solkn/Voting_System/entity"
	"github.com/solkn/Voting_System/party"
)

type PartyService struct {
	partyRepo party.PartyRepository
}

func NewPartyService(partyRepos party.PartyRepository)*PartyService{
	return &PartyService{partyRepo: partyRepos}
}

func (ps *PartyService) Candidates()([]entity.Party,[]error){

	parties,errs:= ps.partyRepo.Candidates()
	if(len(errs)>0){
		return nil,errs
	}
	return parties,errs
}
func (ps *PartyService)Party(id uint)(*entity.Party,[]error)  {

	party,errs:= ps.partyRepo.Party(id)
	if(len(errs)>0){
		return nil,errs

	}
	return party,errs
}
func (ps *PartyService)UpdateParty(party *entity.Party)(*entity.Party,[]error)  {

	pt,errs:=ps.partyRepo.UpdateParty(party)
	if(len(errs)>0){
		return nil,errs

	}
	return pt,errs
}

func (ps *PartyService)StoreParty(party *entity.Party)(*entity.Party,[]error)  {

	pt,errs:= ps.partyRepo.StoreParty(party)
	if(len(errs)>0){
		return nil,errs
	}
	return pt,errs

}
func (ps *PartyService) DeleteParty(id uint)(*entity.Party,[]error)  {
	party,errs:=ps.partyRepo.DeleteParty(id)
	if(len(errs)>0){
		return nil,errs
	}
	return party,errs
}