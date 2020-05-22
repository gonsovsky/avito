package api

import (
	"avito/db"
	. "avito/shared"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func allPages(w http.ResponseWriter, r *http.Request) {
	var Pages []AvitoPage
	Pages, err := db.AllPages()
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(Pages)
	}
}

func onePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	page, err := db.OnePage(id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(page)
	}
}

func newPage(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var page AvitoPage
	json.Unmarshal(reqBody, &page)
	res, err := db.NewPage(page)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(&res)
	}
}

func newOrder(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var order AvitoOrder
	json.Unmarshal(reqBody, &order)
	res, err := db.NewOrder(order)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(&res)
	}
}

func updage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reqBody, _ := ioutil.ReadAll(r.Body)
	var page AvitoPage
	json.Unmarshal(reqBody, &page)
	res, err := db.UpdatePage(vars["id"], page)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(&res)
	}
}

func delPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := db.DelPage(id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(id)
	}
}

func NewApi() {
	defer func() {
		if r := recover(); r != nil {
			Log("(NewApi) Recovered in f", r)
		}
	}()
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/pages", allPages)
	myRouter.HandleFunc("/page/{id}", updage).Methods("POST")
	myRouter.HandleFunc("/page", newPage).Methods("POST")
	myRouter.HandleFunc("/page/{id}", delPage).Methods("DELETE")
	myRouter.HandleFunc("/page/{id}", onePage)
	Log(http.ListenAndServe("0.0.0.0:9000", myRouter))
}
