package webserver

import (
	"encoding/json"
	"gta4roy/app/log"
	"gta4roy/app/model"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

const (
	BaseURL = "/api/v1/directory"

	HealthChecURL = "/health"

	AddPhoneURL = BaseURL + "/phone"

	ModifyPhoneURL = BaseURL + "/phone/{phno:[0-9]{10}$}"

	SearchPhoneURL = BaseURL + "/phone/{phno:[0-9]{10}$}"

	GetAllPhoneURL = BaseURL + "/phone"

	DeletePhoneURL = BaseURL + "/phone/{phno:[0-9]{10}$}"
)

var routes = Routes{
	Route{"HealthCheck", "GET", HealthChecURL, handlerHealthCheck},
	Route{"AddPhone", "POST", AddPhoneURL, handlerAddPhoneNumber},
	Route{"UpdatePhone", "PUT", ModifyPhoneURL, handlerUpdateNumber},
	Route{"SearchPhone", "GET", SearchPhoneURL, handlerSearchNumber},
	Route{"GetAllPhone", "GET", GetAllPhoneURL, handlerGetAllNumber},
	Route{"DeletePhone", "DELETE", DeletePhoneURL, handlerDeleteNumber},
}

func logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Trace.Println("%s %s 5s %s", r.Method, r.RequestURI, name, time.Since(start))
	})
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger(handler, route.Name)
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}
	return router
}

func handlerHealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Trace.Println("handlerHealthCheck Request")
	var response model.ResponseModel

	response.Status = model.CODE_SUCCESS
	response.Message = "ALIVE"
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)

}

func handlerAddPhoneNumber(w http.ResponseWriter, r *http.Request) {
	log.Trace.Println("handlerAddPhoneNumber Request")

	log.Trace.Println("parsing the input parameters ")
	var personDataRequestSet model.PersonModel

	var response model.ResponseModel

	response.Status = model.CODE_SUCCESS
	response.Message = "OK"

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil || reqBody == nil {
		log.Error.Println("Error Reading the body of the request %v", err)
		response.Status = model.CODE_WRONG_INPUTS
		response.Message = err.Error()
	}

	err = json.Unmarshal(reqBody, &personDataRequestSet)
	if err != nil {
		log.Error.Println("Error Reading the body of the request %v", err)
		response.Status = model.CODE_WRONG_INPUTS
		response.Message = err.Error()
	}
	log.Trace.Println("Data Passed ", personDataRequestSet.ToString())
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if response.Status == model.CODE_SUCCESS {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(response)

}

func handlerUpdateNumber(w http.ResponseWriter, r *http.Request) {
	log.Trace.Println("handlerUpdateNumber Request")
	vars := mux.Vars(r)
	log.Trace.Println("Phone no received for update", vars["phno"])

	log.Trace.Println("parsing the input parameters ")
	var personDataRequestSet model.PersonModel

	var response model.ResponseModel

	response.Status = model.CODE_SUCCESS
	response.Message = "OK"

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil || reqBody == nil {
		log.Error.Println("Error Reading the body of the request %v", err)
		response.Status = model.CODE_WRONG_INPUTS
		response.Message = err.Error()
	}

	err = json.Unmarshal(reqBody, &personDataRequestSet)
	if err != nil {
		log.Error.Println("Error Reading the body of the request %v", err)
		response.Status = model.CODE_WRONG_INPUTS
		response.Message = err.Error()
	}
	log.Trace.Println("Data Passed ", personDataRequestSet.ToString())
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if response.Status == model.CODE_SUCCESS {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(response)

}

func handlerSearchNumber(w http.ResponseWriter, r *http.Request) {
	log.Trace.Println("handlerSearchNumber Request")
	vars := mux.Vars(r)
	log.Trace.Println("Phone no received for search ", vars["phno"])
	w.WriteHeader(http.StatusOK)

}

func handlerGetAllNumber(w http.ResponseWriter, r *http.Request) {
	log.Trace.Println("handlerGetAllNumber Request")

	w.WriteHeader(http.StatusOK)

}
func handlerDeleteNumber(w http.ResponseWriter, r *http.Request) {
	log.Trace.Println("handlerDeleteNumber Request")
	vars := mux.Vars(r)
	log.Trace.Println("Phone no received for deletion", vars["phno"])
	w.WriteHeader(http.StatusOK)

}
