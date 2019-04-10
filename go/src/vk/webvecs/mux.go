/*/

package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	vutils "vk/utils"

	"io/ioutil"




	"github.com/gorilla/mux"
)

//var Rtr mux.Router

func setMux() {
	r := mux.NewRouter()

	r.HandleFunc("/", home)

	//	Rtr.HandleFunc("/", makeWrappedTmplHandler("home_tmpl", "", ""))
	//	Rtr.HandleFunc("/about", makeWrappedTmplHandler("about_tmpl", "", ""))
	//http.Handle("/", home)
	http.Handle("/", http.FileServer(http.Dir(webPrefix+"/html/ind.html")))
}

func home(w http.ResponseWriter, req *http.Request) {
	//	render(w, "ind.html")
	full := vutils.FileAbsPath("html", "index.html")
	txt, err := ioutil.ReadFile(full) // just pass the file name//
	if err != nil {
		return
	}
	t := new(template.Template)
	t, err = t.Parse(string(txt))
	if err != nil {
		panic("template parsing error: ")
	}
	err = t.Execute(w, "")
	if err != nil {
		panic("template executing error: ")
	}
}

func render(w http.ResponseWriter, tmpl string) {
	tmpl = fmt.Sprintf("templates/%s", tmpl)
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, "")
	if err != nil {
		log.Print("template executing error: ", err)
	}
}
*/
package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"
)

//const STATIC_URL string = "/externals/"
const STATIC_URL string = "/static/"
const STATIC_ROOT string = "externals/"

var rtr = mux.NewRouter()

type Context struct {
	Title  string
	Static string
}

func home(w http.ResponseWriter, req *http.Request) {
	context := Context{Title: "Welcome!"}
	render(w, "index", context)
}

func about(w http.ResponseWriter, req *http.Request) {
	context := Context{Title: "About"}
	render(w, "about", context)
}

func loginInp(w http.ResponseWriter, req *http.Request) {
	context := Context{Title: "Login input"}
	render(w, "login", context)
}

func setMux() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about/", about)

	rtr.HandleFunc("/login", loginInp).Methods("GET")

	//http.HandleFunc("/", Home)
	//http.HandleFunc("/about/", About)
	http.HandleFunc(STATIC_URL, StaticHandler)

	/*
		err := http.ListenAndServe(":8000", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	*/
}

/*
func render(w http.ResponseWriter, tmpl string) {
	tmpl = fmt.Sprintf("html/%s", tmpl)
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, "")
	if err != nil {
		log.Print("template executing error: ", err)
	}
}
*/

func render(w http.ResponseWriter, tmpl string, context Context) {
	context.Static = STATIC_URL
	//	tmpl_list := []string{"html/base.html",
	//		fmt.Sprintf("html/%s.html", tmpl)}
	tmpl_list := []string{fmt.Sprintf("html/%s.html", tmpl)}
	t, err := template.ParseFiles(tmpl_list...)
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, context)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

func StaticHandler(w http.ResponseWriter, req *http.Request) {
	static_file := req.URL.Path[len(STATIC_URL):]
	if len(static_file) != 0 {
		f, err := http.Dir(STATIC_ROOT).Open(static_file)
		if err == nil {
			content := io.ReadSeeker(f)
			http.ServeContent(w, req, static_file, time.Now(), content)
			return
		}
	}
	http.NotFound(w, req)
}
