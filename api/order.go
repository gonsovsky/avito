package api

import (
	"avito/db"
	. "avito/shared"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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
