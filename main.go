package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strings"

	"google.golang.org/appengine"
)

//templates
var glob *template.Template

//var soc *template.Template
//var pol *template.Template

func main() {
	glob = template.Must(template.ParseGlob("templates/*.html"))

	//test routing
	http.HandleFunc("/bruh", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `<html><body style="background: rgb(255, 255, 255);display: flex;justify-content: center;align-items: center;height: 100vh;width: 100vwmargin: 0;"><h1 style="font-size: 10rvw;font-weight: 200;text-transform: lowercase;letter-spacing: 1vw;">Siesta y fiesta</h1></body></html>`)
	})

	http.HandleFunc("/", splash)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "files/assets/img/favicon.png")
	})
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("files/assets"))))
	http.HandleFunc("/timeline/", timeline)
	http.HandleFunc("/games/", games)
	http.HandleFunc("/resources/", resources)
	http.HandleFunc("/about", about)
	http.HandleFunc("/policy", policy)
	http.HandleFunc("/welcome", welcome)
	http.Handle("/article/", http.StripPrefix("/article", http.FileServer(http.Dir("hugo/public"))))

	fmt.Println("listening at port :8080")

	//http.ListenAndServe(":8080", nil)
	appengine.Main()
}

func splash(w http.ResponseWriter, r *http.Request) {
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
	p := strings.Split(r.URL.Path, "/")

	if p[2] == "" {
		if err := glob.ExecuteTemplate(w, "games.html", nil); err != nil {
			handleError(w, err)
		}
	} else if p[2] == "index.json" {
		http.ServeFile(w, r, "files/games/index.json")
	} else {
		http.ServeFile(w, r, "files/games/"+strings.Join(p[2:], "/"))
	}
}

func resources(w http.ResponseWriter, r *http.Request) {
	p := strings.Split(r.URL.Path, "/")

	if p[2] == "" {
		//get 3 random vids
		vids := getRandom(3, len(Interviews)-1)
		arts := getRandom(3, len(Artefacts)-1)

		//write response
		if err := glob.ExecuteTemplate(w, "resources.index.html", Resource{[]Interview{Interviews[vids[0]], Interviews[vids[1]], Interviews[vids[2]]}, []Artefact{Artefacts[arts[0]], Artefacts[arts[1]], Artefacts[arts[2]]}}); err != nil {
			handleError(w, err)
		}
	} else if p[2] == "interviews" {
		if len(p) < 4 || p[3] == "" {
			if err := glob.ExecuteTemplate(w, "resources.interviews.html", Interviews); err != nil {
				handleError(w, err)
			}
		} else {
			//p[3] = Interviews[x].Key
			for _, interv := range Interviews {
				if interv.Key == p[3] {
					if err := glob.ExecuteTemplate(w, "resources.interview-single.html", interv); err != nil {
						handleError(w, err)
					}
				}
			}
		}
	} else if p[2] == "artefacts" {
		if len(p) < 4 || p[3] == "" {
			if err := glob.ExecuteTemplate(w, "resources.artefacts.html", Artefacts); err != nil {
				handleError(w, err)
			}
		} else {
			//p[3] = Artefacts[x].Key
			for _, artf := range Artefacts {
				if artf.Key == p[3] {
					if err := glob.ExecuteTemplate(w, "resources.artefact-single.html", artf); err != nil {
						handleError(w, err)
					}
				}
			}
		}
	} else if p[2] == "introduction" {
		if err := glob.ExecuteTemplate(w, "resources.introduction.html", nil); err != nil {
			handleError(w, err)
		}
	} else if p[2] == "introduction" {
		//if err := glob.ExecuteTemplate(w, "resources.introduction.html", nil); err != nil {
		//	handleError(w, err)
		//}
		fmt.Fprint(w, "hello at the teaching guide")
	} else if p[2] == "raw" {
		http.ServeFile(w, r, "files/interviews/"+strings.Join(p[3:], "/"))
	} else {
		fmt.Fprint(w, p)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	if err := glob.ExecuteTemplate(w, "about.html", nil); err != nil {
		handleError(w, err)
	}
}

func policy(w http.ResponseWriter, r *http.Request) {
	if err := glob.ExecuteTemplate(w, "policy.html", nil); err != nil {
		handleError(w, err)
	}

}

func welcome(w http.ResponseWriter, r *http.Request) {
	if err := glob.ExecuteTemplate(w, "welcome.html", nil); err != nil {
		handleError(w, err)
	}
}

func handleError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func getRandom(n int, max int) []int {
	var ret []int
	i := 0

	for i < n {
		x := rand.Intn(max)
		yep := false
		for j := 0; j < i; j++ {
			if ret[j] == x {
				yep = true
				break
			}
		}
		if !yep {
			ret = append(ret, x)
			i++
		}
	}
	return ret
}

/*

//main START

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

//main END


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
