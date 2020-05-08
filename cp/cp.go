package main

import (
	"avito/db"
	. "avito/shared"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseGlob("C:\\Users\\K\\go\\src\\avito\\cp\\form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	res, err := db.AllPages()
	if err != nil {
		panic(err.Error())
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
		panic(err.Error())
	}
	tmpl.ExecuteTemplate(w, "Edit", res)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	var page = AvitoPage{}
	if r.Method == "POST" {
		page.Url = r.FormValue("url")
		page.Hint = r.FormValue("hint")
		_, err := db.NewPage(page)
		if err != nil {
			panic(err.Error())
		}
	}
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var page = AvitoPage{}
	if r.Method == "POST" {
		page.Url = r.FormValue("url")
		page.Hint = r.FormValue("hint")
		_, err := db.NewPage(page)
		if err != nil {
			panic(err.Error())
		}
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	nId := r.URL.Query().Get("id")
	err := db.DelPage(nId)
	if err != nil {
		panic(err.Error())
	}
	http.Redirect(w, r, "/", 301)
}

func Original(w http.ResponseWriter, r *http.Request) {
	nId := r.URL.Query().Get("id")
	res, err := db.OnePage(nId)
	if err != nil {
		panic(err.Error())
	}
	http.Redirect(w, r, res.Url, 301)
}

func Fake(w http.ResponseWriter, r *http.Request) {
	nId := r.URL.Query().Get("id")
	res, err := db.OnePage(nId)
	if err != nil {
		panic(err.Error())
	}
	http.Redirect(w, r, res.Url, 301)
}

func Reload(w http.ResponseWriter, r *http.Request) {
	nId := r.URL.Query().Get("id")
	res, err := db.OnePage(nId)
	if err != nil {
		panic(err.Error())
	}
	http.Redirect(w, r, res.Url, 301)
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/original", Original)
	http.HandleFunc("/fake", Fake)
	http.HandleFunc("/reload", Reload)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe("0.0.0.0:9001", nil)
}
