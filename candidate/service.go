package candidate

import "github.com/solkn/Voting_System/entity"

type CandidateServices interface {

	Candidates()([]entity.Candidate,[]error)
	Candidate(id uint)(*entity.Candidate,[]error)
	StoreCandidate(candidate *entity.Candidate)(*entity.Candidate,[]error)
	UpdateCandidate(candidate *entity.Candidate)(*entity.Candidate,[]error)
	DeleteCandidate(id uint)(*entity.Candidate,[]error)
}

