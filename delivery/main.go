package main

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/solkn/Voting_System/delivery/http/handler"
	"github.com/solkn/Voting_System/entity"

	ur "github.com/solkn/Voting_System/user/repository"
	us "github.com/solkn/Voting_System/user/service"

	pr "github.com/solkn/Voting_System/party/repository"
	ps "github.com/solkn/Voting_System/party/service"

	rr "github.com/solkn/Voting_System/result/repository"
	rs "github.com/solkn/Voting_System/result/service"

	cr "github.com/solkn/Voting_System/candidate/repository"
	cs "github.com/solkn/Voting_System/candidate/service"



)

func createTables(dbConn *gorm.DB) []error {
	dbConn.DropTableIfExists(&entity.Role{}, &entity.User{}, &entity.Party{}, &entity.Result{},
		&entity.Candidate{}).GetErrors()
	errs := dbConn.CreateTable(&entity.Role{}, &entity.User{}, &entity.Party{},&entity.Result{},
		&entity.Candidate{}).GetErrors()
	dbConn.Debug().Model(&entity.User{}).AddForeignKey("role_id", "roles(Id)", "cascade", "cascade")
	dbConn.Debug().Model(&entity.Result{}).AddForeignKey("party_id", "parties(Id)", "cascade", "cascade")
	dbConn.Debug().Model(&entity.Candidate{}).AddForeignKey("partyID", "parties(Id)", "cascade", "cascade")

	if len(errs) > 0 {
		return errs
	}
	return nil
}

func main() {

	dbconn, err := gorm.Open("postgres",
		"postgres://postgres:solomon@localhost/vote?sslmode=disable")
	if dbconn != nil {
		defer dbconn.Close()
	}
	if err != nil {
		panic(err)
	}
	//createTables(dbconn)

     router:= mux.NewRouter()

     roleRepo:= ur.NewRoleGormRepo(dbconn)
     roleService:= us.NewRoleService(roleRepo)
     roleHandler:=handler.NewRoleApiHandler(roleService)


     userRepo:= ur.NewUserGormRepo(dbconn)
     userService:= us.NewUserService(userRepo)
     usersHandler:= handler.NewUserApiHandler(userService)

	router.HandleFunc("/v1/role", usersHandler.Authenticated(usersHandler.Authorized(roleHandler.GetRoles))).Methods("GET")
	router.HandleFunc("/v1/roles/{name}", usersHandler.Authenticated(usersHandler.Authorized(roleHandler.GetRoleByName))).Methods("GET")
	router.HandleFunc("/v1/role/{id}", usersHandler.Authenticated(usersHandler.Authorized(roleHandler.GetRoleByID))).Methods("GET")
	router.HandleFunc("/v1/role", usersHandler.Authenticated(usersHandler.Authorized(roleHandler.PostRole))).Methods("POST")
	router.HandleFunc("/v1/role/{id}", usersHandler.Authenticated(usersHandler.Authorized(roleHandler.PutRole))).Methods("PUT")
	router.HandleFunc("/v1/role/{id}", usersHandler.Authenticated(usersHandler.Authorized(roleHandler.DeleteRole))).Methods("DELETE")

	router.HandleFunc("/v1/admin/users/{id}", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.GetUser))).Methods("GET")
	router.HandleFunc("/v1/admin/users", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.GetUsers))).Methods("GET")
	router.HandleFunc("/v1/admin/users/{id}", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.PutUser))).Methods("PUT")
	router.HandleFunc("/v1/admin/users", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.PostUser))).Methods("POST")
	router.HandleFunc("/v1/admin/users/{id}", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.DeleteUser))).Methods("DELETE")
	router.HandleFunc("/v1/admin/email/{email}", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.IsEmailExists))).Methods("GET")
	router.HandleFunc("/v1/admin/phone/{phone}", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.IsPhoneExists))).Methods("GET")
	router.HandleFunc("/v1/admin/check", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.GetUserByUsernameAndPassword))).Methods("POST")

	router.HandleFunc("/v1/user/users", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.PostUser))).Methods("POST")
	router.HandleFunc("/v1/user/users", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.GetUsers))).Methods("GET")
	router.HandleFunc("/v1/user/users/{id}", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.GetUser))).Methods("GET")
	router.HandleFunc("/v1/user/users/{id}", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.PutUser))).Methods("PUT")
	router.HandleFunc("/v1/user/email/{email}", usersHandler.IsEmailExists).Methods("GET")
	router.HandleFunc("/v1/user/phone/{phone}", usersHandler.IsPhoneExists).Methods("GET")
	router.HandleFunc("/v1/user/check", usersHandler.GetUserByUsernameAndPassword).Methods("POST")
	router.HandleFunc("/v1/user/login", usersHandler.Login).Methods("POST")
	router.HandleFunc("/v1/user/signup", usersHandler.SignUp).Methods("POST")

     partyRepo:= pr.NewPartyGormRepo(dbconn)
     partyService:= ps.NewPartyService(partyRepo)
     partyHandler:= handler.NewPartyApiHandler(partyService)

	router.HandleFunc("/v1/party", partyHandler.GetParties).Methods("GET")
	router.HandleFunc("/v1/party/{id}", partyHandler.GetParty).Methods("GET")
	router.HandleFunc("/v1/party", usersHandler.Authenticated(usersHandler.Authorized(partyHandler.PostParty))).Methods("POST")
	router.HandleFunc("/v1/party/{id}", usersHandler.Authenticated(usersHandler.Authorized(partyHandler.PutParty))).Methods("PUT")
	router.HandleFunc("/v1/party/{id}", usersHandler.Authenticated(usersHandler.Authorized(partyHandler.DeleteParty))).Methods("DELETE")

    resulRepo:= rr.NewResultGormRepo(dbconn)
    resultService:= rs.NewResultService(resulRepo)
    resultsHandler:= handler.NewResultApiHandler(resultService)

	router.HandleFunc("/v1/result", resultsHandler.GetResults).Methods("GET")
	router.HandleFunc("/v1/result/{id}", resultsHandler.GetResult).Methods("GET")
	router.HandleFunc("/v1/result", usersHandler.Authenticated(usersHandler.Authorized(resultsHandler.PostResult))).Methods("POST")
	router.HandleFunc("/v1/result/{id}", usersHandler.Authenticated(usersHandler.Authorized(resultsHandler.PutResult))).Methods("PUT")
	router.HandleFunc("/v1/result/{id}", usersHandler.Authenticated(usersHandler.Authorized(resultsHandler.DeleteResult))).Methods("DELETE")

    candidateRepo:= cr.NewPartyGormRepo(dbconn)
    candidateService:= cs.NewCandidateService(candidateRepo)
    candidateHandler:= handler.NewCandidateApiHandler(candidateService)

	router.HandleFunc("/v1/candidate", candidateHandler.GetCandidates).Methods("GET")
	router.HandleFunc("/v1/candidate/{id}", candidateHandler.GetCandidate).Methods("GET")
	router.HandleFunc("/v1/candidate", usersHandler.Authenticated(usersHandler.Authorized(candidateHandler.PostCandidate))).Methods("POST")
	router.HandleFunc("/v1/candidate/{id}", usersHandler.Authenticated(usersHandler.Authorized(candidateHandler.PutCandidate))).Methods("PUT")
	router.HandleFunc("/v1/candidate/{id}", usersHandler.Authenticated(usersHandler.Authorized(candidateHandler.DeleteCandidate))).Methods("DELETE")
	err = http.ListenAndServe("192.168.56.1:8080", router)

	if err != nil {
		panic(err)
	}
}