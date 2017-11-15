package main

import (
	"html/template"
	"net/http"
	"log"
	"github.com/julienschmidt/httprouter"
)

//templates
var glob *template.Template
var soc *template.Template
var pol *template.Template

//router
var mux *httprouter.Router

func init() {
	//parsing the templates
	glob = template.Must(template.ParseGlob("templates/*.html"))
	soc = template.Must(template.ParseGlob("templates/social/*.html"))
	pol = template.Must(template.ParseGlob("templates/political/*.html"))

	//adding includes to all templates
	soc.ParseFiles("templates/includes.html")
	pol.ParseFiles("templates/includes.html")
}

func main() {
	mux = httprouter.New()
	//routing specific paths
	mux.GET("/", index)

	mux.GET("/political/:article", political)
	mux.GET("/social/:article", social)

	//routing assets
	mux.ServeFiles("/assets/*filepath", http.Dir("assets"))

	//fire the server up
	log.Println("fired it up at port :8080")
	http.ListenAndServe(":8080", mux)
}

//handle the specific routes
func index(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
/*
	d := struct{
		Kibitagsets Settings
		Error bool
		SubKind formKind
	}{
		getSettings(),
		false,
		formKind{},
	}

*/

	err := glob.ExecuteTemplate(w, "index.html", nil)
	HandleError(w, err)
	return
}

func political(w http.ResponseWriter, _ *http.Request, parms httprouter.Params)  {
	err := pol.ExecuteTemplate(w, parms[0].Value + ".html", nil)
	HandleError(w, err)
	return
}

func social(w http.ResponseWriter, _ *http.Request, parms httprouter.Params)  {
	err := soc.ExecuteTemplate(w, parms[0].Value + ".html", nil)
	HandleError(w, err)
	return
}

/* send the client the answer HTTP 500 */
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
