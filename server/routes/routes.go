package routes

import (
	"github.com/gorilla/mux"
	"github.com/shangsuru/ctf-website/controllers"
)

func Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/challenges", controllers.GetChallengesEndpoint).Methods("GET")
	return router
}
