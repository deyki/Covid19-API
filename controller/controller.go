package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Covid19MicroServiceV2/deyki/v2/service"
	"github.com/gorilla/mux"
)


func getCountryByCountryCode(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	countryCode := params["countryCode"]

	country, err := service.GetCountryByCountryCode(countryCode)
	if err != nil {
		json.NewEncoder(w).Encode(&err)
		return
	}

	json.NewEncoder(w).Encode(&country)
}


func GorillaMuxRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/{countryCode}", getCountryByCountryCode).Methods("GET")
	return router
}