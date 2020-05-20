package cp

import (
	"avito/db"
	. "avito/shared"
	"net/http"
	"os"
	"strings"
	"text/template"
)

var tmpl *template.Template

func load() {
	var wqr, _ = os.Getwd()
	wqr = wqr + "/cp/form/*"
	var a = template.Must(template.ParseGlob(wqr))
	tmpl = a
}

func Index(w http.ResponseWriter, r *http.Request) {
	res, err := db.AllPages()
	if err != nil {
		e := map[string]interface{}{"err": err}
		tmpl.ExecuteTemplate(w, "Error", e)
		return
	}
	tmpl.ExecuteTemplate(w, "Index", res)
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
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

func Insert(w http.ResponseWriter, r *http.Request) {
	var page = AvitoPage{}
	if r.Method == "POST" {
		page.Url = r.FormValue("url")
		page.Hint = r.FormValue("hint")
		_, err := db.NewPage(page)
		if err != nil {
			e := map[string]interface{}{"err": err}
			tmpl.ExecuteTemplate(w, "Error", e)
			return
		}
	}
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var page = AvitoLitePage{}
	if r.Method == "POST" {
		nId := r.FormValue("uid")

		page.Url = r.FormValue("url")
		page.Hint = r.FormValue("hint")

		_, err := db.UpdatePageLight(nId, page)
		if err != nil {
			e := map[string]interface{}{"err": err}
			tmpl.ExecuteTemplate(w, "Error", e)
			return
		}
	}
	http.Redirect(w, r, "/", 301)
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

func Original(w http.ResponseWriter, r *http.Request) {
	nId := r.URL.Query().Get("id")
	res, err := db.OnePage(nId)
	if err != nil {
		e := map[string]interface{}{"err": err}
		tmpl.ExecuteTemplate(w, "Error", e)
		return
	}
	http.Redirect(w, r, res.Url, 301)
}

func Fake(w http.ResponseWriter, r *http.Request) {
	nId := r.URL.Query().Get("id")
	res, err := db.OnePage(nId)
	if err != nil {
		e := map[string]interface{}{"err": err}
		tmpl.ExecuteTemplate(w, "Error", e)
		return
	}
	host := strings.Split(r.Host, ":")[0]
	host = MainHost
	fake := "http://www.avito.ru." + host + "/" + res.Id
	http.Redirect(w, r, fake, 301)
}

func Reload(w http.ResponseWriter, r *http.Request) {
	nId := r.URL.Query().Get("id")
	_, err := db.Parse(nId)
	if err != nil {
		e := map[string]interface{}{"err": err}
		tmpl.ExecuteTemplate(w, "Error", e)
		return
	}
	http.Redirect(w, r, "/", 301)
}

func NewCP() {
	load()
	http.HandleFunc("/", Index)
	http.HandleFunc("/original", Original)
	http.HandleFunc("/fake", Fake)
	http.HandleFunc("/reload", Reload)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	error := http.ListenAndServe("0.0.0.0:9001", nil)
	if error != nil {
		Log(error)
	}
}
