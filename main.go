package erasmus

import (
	"html/template"
	"net/http"
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

	/*from main*/
	mux = httprouter.New()
	//routing specific paths
	mux.GET("/", index)

	mux.GET("/political/:article", political)
	mux.GET("/social/:article", social)

	//routing assets
	mux.GET("/assets/*filepath", assets)
}

//handle the specific routes
func index(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	err := glob.ExecuteTemplate(w, "index.html", nil)
	HandleError(w, err)
	return
}

func assets(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	assets_path := "https://storage.googleapis.com/cold-war-cojobo.appspot.com/assets"
	http.Redirect(w, r, assets_path+p.ByName("filepath"), http.StatusSeeOther)
	return
}

func political(w http.ResponseWriter, _ *http.Request, parms httprouter.Params) {
	err := pol.ExecuteTemplate(w, parms[0].Value+".html", nil)
	HandleError(w, err)
	return
}

func social(w http.ResponseWriter, _ *http.Request, parms httprouter.Params) {
	err := soc.ExecuteTemplate(w, parms[0].Value+".html", nil)
	HandleError(w, err)
	return
}

/* send the client the answer HTTP 500 */
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
