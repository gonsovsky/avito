package cp

import (
	"avito/cp/controller/avito"
	"avito/cp/controller/order"
	. "avito/shared"
	"net/http"
)

func NewCP() {
	avito.Load()
	http.HandleFunc("/", avito.Index)
	http.HandleFunc("/original", avito.Original)
	http.HandleFunc("/fake", avito.Fake)
	http.HandleFunc("/reload", avito.Reload)
	http.HandleFunc("/new", avito.New)
	http.HandleFunc("/edit", avito.Edit)
	http.HandleFunc("/insert", avito.Insert)
	http.HandleFunc("/update", avito.Update)
	http.HandleFunc("/delete", avito.Delete)

	order.Load()
	http.HandleFunc("/orders", order.Index)
	http.HandleFunc("/order/delete", order.Delete)
	http.HandleFunc("/order/details", order.Edit)

	error := http.ListenAndServe("0.0.0.0:9001", nil)
	if error != nil {
		Log(error)
	}
}
