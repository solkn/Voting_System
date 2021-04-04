package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/solkn/Voting_System/entity"
	"github.com/solkn/Voting_System/party"
	"github.com/solkn/Voting_System/utils"
	"net/http"
	"strconv"
)

type PartyApiHandler struct {
	partyServices party.PartyServices
}
func NewPartyApiHandler(partyServ party.PartyServices)*PartyApiHandler{
	return &PartyApiHandler{partyServices: partyServ}
}

func (pah *PartyApiHandler)GetParties(w http.ResponseWriter,r *http.Request)  {

	parties,errs:= pah.partyServices.Parties()
	if  len(errs)>0{
		w.Header().Set("Content-Type","application/json")
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}
	output,err:=json.MarshalIndent(parties,"","\t\t")
	if(err!=nil){
		w.Header().Set("Content-Type","application/json")
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(output)

	return
}
func (pah *PartyApiHandler)GetParty(w http.ResponseWriter,r *http.Request){
	params := mux.Vars(r)
	id,err:= strconv.Atoi(params["id"])
	if(err!=nil){
		w.Header().Set("Content-Type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}

	party, errs:= pah.partyServices.Party(uint(id))

	if (errs != nil) {
		w.Header().Set("Content-Type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return

	}

	output,err:= json.MarshalIndent(party,"","\t\t")
	if (err != nil) {
		w.Header().Set("Content-Type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return

	}
	w.Header().Set("Content-Type","application/json")
	w.Write(output)

 return
}

func (pah *PartyApiHandler)PostParty(w http.ResponseWriter,r *http.Request){
	body:= utils.BodyParser(r)
	var party entity.Party
	err:= json.Unmarshal(body,&party)
	if (err!=nil) {
		utils.ToJson(w,err.Error(),http.StatusUnprocessableEntity)
		return
	}

	storedParty, errs:=pah.partyServices.StoreParty(&party)

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
	return
}

func (pah *PartyApiHandler) PutParty(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	party, errs := pah.partyServices.Party((uint(id)))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	// l := r.ContentLength

	// body := make([]byte, l)

	// _, err = r.Body.Read(body)
	// if err != nil {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	// 	return
	// }
	body:= utils.BodyParser(r)
	err = json.Unmarshal(body, &party)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	party, errs = pah.partyServices.UpdateParty(party)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(party, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}
func (pah *PartyApiHandler) DeleteParty(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := pah.partyServices.DeleteParty(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}

