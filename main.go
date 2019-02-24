package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//templates
var glob *template.Template

//var soc *template.Template
//var pol *template.Template

func main() {
	glob = template.Must(template.ParseGlob("templates/*.html"))
	/*
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

		//next version - not working!!!!!!!!!!!!!!!!!!!!!!11!!!!elf
		http.HandleFunc("/next/", next)
		http.HandleFunc("/next/timeline/", timeline)
		http.HandleFunc("/next/about", about)
		http.Handle("/next/article/", http.StripPrefix("/next/article", http.FileServer(http.Dir("hugo/public"))))

	*/
	//test routing
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/bruh", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `<html><body style="background: rgb(255, 255, 255);display: flex;justify-content: center;align-items: center;height: 100vh;width: 100%;"><h1 style="font-size: 11rem;font-weight: 200;text-transform: lowercase;letter-spacing: 1rem;">Siesta y fiesta</h1></body></html>`)
	})

	http.HandleFunc("/", next)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "hugo/static/assets/img/favicon.png")
	})
	http.Handle("/assets/", http.FileServer(http.Dir("hugo/public")))
    http.Handle("/gameAssets/", http.FileServer(http.Dir("hugo/public/games")))
	http.HandleFunc("/timeline/", timeline)
    http.HandleFunc("/games/", games)
	http.HandleFunc("/about", about)
	http.Handle("/article/", http.StripPrefix("/article", http.FileServer(http.Dir("hugo/public"))))

	fmt.Println("listening at port :8080")
	http.ListenAndServe(":8080", nil)
	//appengine.Main()
}

func next(w http.ResponseWriter, r *http.Request) {
	if err := glob.ExecuteTemplate(w, "splash.html", nil); err != nil {
		handleError(w, err)
	}

}

func timeline(w http.ResponseWriter, r *http.Request) {
	if err := glob.ExecuteTemplate(w, "timeline.html", nil); err != nil {
		handleError(w, err)
	}

}

func games(w http.ResponseWriter, r *http.Request) {
	if err := glob.ExecuteTemplate(w, "games.html", nil); err != nil {
		handleError(w, err)
	}

}

func about(w http.ResponseWriter, r *http.Request) {
	if err := glob.ExecuteTemplate(w, "about.html", nil); err != nil {
		handleError(w, err)
	}

}

//handle the ping
func ping(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This is a teapot", http.StatusTeapot)
}

func handleError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

/*
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
	assetsPath := "https://storage.googleapis.com/cold-war-cojobo.appspot.com"
	http.Redirect(w, r, assetsPath+r.URL.RequestURI(), http.StatusSeeOther)
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

// HandleError send the client the answer HTTP 500
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)

		/*send the client the "in-work"-page
		err = glob.ExecuteTemplate(w, "in-work.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

/* map the requested path and return a parameter set by getCode int
func getArticle(r *http.Request) string {
	p := strings.Split(r.URL.Path, "/")

	if p[2] == "" {
		return "index"
	}

	return p[2]
}

// CheckSession return true when the client is loggedd-in
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

// HandleSession check on login and handle when not logged in, otherwise return true
func HandleSession(w http.ResponseWriter, r *http.Request) bool {
	if CheckSession(r) == false {
		err := glob.ExecuteTemplate(w, "login.html", nil)
		HandleError(w, err)
		return false
	}

	return true
}
*/
