package order

import (
	"avito/db"
	"net/http"
	"os"
	"text/template"
)

var tmpl *template.Template

func Load() {
	var cd, _ = os.Getwd()
	cd = cd + "/cp/view/order/*"
	t, e := template.ParseGlob(cd)
	var a = template.Must(t, e)

	cd, _ = os.Getwd()
	cd = cd + "/cp/view/shared/*"
	t, e = template.ParseGlob(cd)
	a, _ = a.ParseGlob(cd)

	tmpl = a
}

func Index(w http.ResponseWriter, r *http.Request) {
	res, err := db.AllOrders()
	if err != nil {
		e := map[string]interface{}{"err": err}
		tmpl.ExecuteTemplate(w, "Error", e)
		return
	}
	tmpl.ExecuteTemplate(w, "Index", res)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	nId := r.URL.Query().Get("id")
	res, err := db.OnePage(nId)
	if err != nil {
		e := map[string]interface{}{"err": err}
		tmpl.ExecuteTemplate(w, "Error", e)
		return
	}
	res.Id = nId
	tmpl.ExecuteTemplate(w, "Edit", res)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	nId := r.URL.Query().Get("id")
	err := db.DelPage(nId)
	if err != nil {
		e := map[string]interface{}{"err": err}
		tmpl.ExecuteTemplate(w, "Error", e)
		return
	}
	http.Redirect(w, r, "/", 301)
}
