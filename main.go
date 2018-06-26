package main

import (
	"html/template"
	"net/http"
	"strings"

	"google.golang.org/appengine"
)

//templates
var glob *template.Template
var soc *template.Template
var pol *template.Template

func main() {
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

	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	//routing assets
	http.HandleFunc("/assets/", assets)

	//test routing
	http.HandleFunc("/ping", ping)

	//next version
	http.Handle("/next/", http.StripPrefix("/next", http.FileServer(http.Dir("hugo/public"))))

	//fmt.Println("listening at port :8080")
	//http.ListenAndServe(":8080", nil)
	appengine.Main()
}

//handle the ping
func ping(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This is a teapot", http.StatusTeapot)
}

//handle the specific routes
func index(w http.ResponseWriter, r *http.Request) {
	if HandleSession(w, r) == false {
		return
	}

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
	if HandleSession(w, r) == false {
		return
	}

	err := pol.ExecuteTemplate(w, getArticle(r)+".html", nil)
	HandleError(w, err)
}

func social(w http.ResponseWriter, r *http.Request) {
	if HandleSession(w, r) == false {
		return
	}

	err := soc.ExecuteTemplate(w, getArticle(r)+".html", nil)
	HandleError(w, err)
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// check if already logged in
		if CheckSession(r) == true {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

		// check passwd
		p := r.FormValue("password")

		if p != "erasmus2019" {
			err := glob.ExecuteTemplate(w, "login.html", true)
			HandleError(w, err)
			return
		}

		// set cookie "logged-in" to 1 and return
		cookie := &http.Cookie{
			Name:     "logged-in",
			Value:    "1",
			HttpOnly: true,
		}

		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if CheckSession(r) == true {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	err := glob.ExecuteTemplate(w, "login.html", nil)
	HandleError(w, err)
	return
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   "logged-in",
		Value:  "0",
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

/* send the client the answer HTTP 500 */
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)

		/*send the client the "in-work"-page*/
		err = glob.ExecuteTemplate(w, "in-work.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

/* map the requested path and return a parameter set by getCode int */
func getArticle(r *http.Request) string {
	p := strings.Split(r.URL.Path, "/")

	if p[2] == "" {
		return "index"
	} else {
		return p[2]
	}
}

/* return true when the client is loggedd-in */
func CheckSession(r *http.Request) bool {
	cookie, err := r.Cookie("logged-in")

	// no cookie
	if err == http.ErrNoCookie {
		return false
	}

	// wrong cookie
	if cookie.Value != "1" {
		return false
	}

	return true
}

/* check on login and handle when not logged in, otherwise return true */
func HandleSession(w http.ResponseWriter, r *http.Request) bool {
	if CheckSession(r) == false {
		err := glob.ExecuteTemplate(w, "login.html", nil)
		HandleError(w, err)
		return false
	}

	return true
}
