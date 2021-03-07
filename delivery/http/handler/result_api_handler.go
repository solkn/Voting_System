package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/solkn/Voting_System/entity"
	"github.com/solkn/Voting_System/utils"
	"github.com/solkn/Voting_System/result"
	"net/http"
	"strconv"
)

type  ResultApiHandler struct {
	resultServices result.ResultServices
}

func NewResultApiHandler(resultServ result.ResultServices) *ResultApiHandler {
	return &ResultApiHandler{resultServices: resultServ}
}

func (rah *ResultApiHandler)GetResults(w http.ResponseWriter,r *http.Request)  {

	results,errs:= rah.resultServices.Results()
	if(errs!=nil){
		w.Header().Set("Content-Type","application/json")
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}
	output,err:=json.MarshalIndent(results,"","\t\t")
	if(err!=nil){
		w.Header().Set("Content-Type","application/json")
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(output)
}
func (rah *ResultApiHandler)GetResult(w http.ResponseWriter,r *http.Request){
	params := mux.Vars(r)
	id,err:= strconv.Atoi(params["id"])
	if(err!=nil){
		w.Header().Set("Content-Type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}

	result, errs:= rah.resultServices.Result(uint(id))

	if (errs != nil) {
		w.Header().Set("Content-Type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return

	}

	output,err:= json.MarshalIndent(result,"","\t\t")
	if (err != nil) {
		w.Header().Set("Content-Type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return

	}
	w.Header().Set("Content-Type","application/json")
	w.Write(output)


}

func (rah *ResultApiHandler)PostResult(w http.ResponseWriter,r *http.Request){
	body:= utils.BodyParser(r)
	var result entity.Result
	err:= json.Unmarshal(body,&result)
	if (err!=nil) {
		utils.ToJson(w,err.Error(),http.StatusUnprocessableEntity)
		return
	}

	storedParty, errs:=rah.resultServices.StoreResult(&result)

	if(errs!=nil){
		w.Header().Set("Content-Type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}

	output,err:= json.MarshalIndent(storedParty,"","\t\t")
	if (err != nil) {
		w.Header().Set("Content-Type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return

	}
	w.Header().Set("Content-Type","application/json")
	w.Write(output)
}

func (rah *ResultApiHandler) PutResult(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	result, errs := rah.resultServices.Result(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	l := r.ContentLength

	body := make([]byte, l)

	_, err = r.Body.Read(body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	result, errs = rah.resultServices.UpdateResult(result)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(result, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}
func (rah *ResultApiHandler) DeleteResult(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := rah.resultServices.DeleteResult(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}

