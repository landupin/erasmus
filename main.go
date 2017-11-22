package erasmus

import (
	"html/template"
	"net/http"
	"strings"
)

//templates
var glob *template.Template
var soc *template.Template
var pol *template.Template

func init() {
	//parsing the templates
	glob = template.Must(template.ParseGlob("templates/*.html"))
	soc = template.Must(template.ParseGlob("templates/social/*.html"))
	pol = template.Must(template.ParseGlob("templates/political/*.html"))

	//adding includes to all templates
	soc.ParseFiles("templates/includes.html")
	pol.ParseFiles("templates/includes.html")

	//routing specific paths
	http.HandleFunc("/", index)

	http.HandleFunc("/political/", political)
	http.HandleFunc("/social/", social)

	//routing assets
	http.HandleFunc("/assets/", assets)
}

//handle the specific routes
func index(w http.ResponseWriter, _ *http.Request) {
	err := glob.ExecuteTemplate(w, "index.html", nil)
	HandleError(w, err)
	return
}

func assets(w http.ResponseWriter, r *http.Request) {
	assets_path := "https://storage.googleapis.com/cold-war-cojobo.appspot.com"
	http.Redirect(w, r, assets_path+r.URL.RequestURI(), http.StatusSeeOther)
	return
}

func political(w http.ResponseWriter, r *http.Request) {
	err := pol.ExecuteTemplate(w, getArticle(r)+".html", nil)
	HandleError(w, err)
}

func social(w http.ResponseWriter, r *http.Request) {
	err := pol.ExecuteTemplate(w, getArticle(r)+".html", nil)
	HandleError(w, err)
}

/* send the client the answer HTTP 500 */
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/* map the requested path and return a parameter set by getCode int */
func getArticle(r *http.Request) (string) {
	p := strings.Split(r.URL.Path, "/")

	if p[2]=="" {
		return "index"
	} else {
		 return p[2]
	}
}
