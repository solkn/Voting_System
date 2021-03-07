package result

import "github.com/solkn/Voting_System/entity"

type ResultServices interface {

	Results()([]entity.Result,[]error)
	Result(id uint)(*entity.Result,[]error)
	StoreResult(result *entity.Result)(*entity.Result,[]error)
	UpdateResult(result *entity.Result)(*entity.Result,[]error)
	DeleteResult(id uint)(*entity.Result,[]error)
}

