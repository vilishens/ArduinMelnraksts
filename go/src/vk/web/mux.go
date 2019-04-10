package web

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	vscanip "vk/code/net/scanpoints"
	vomni "vk/omnibus"
	xrun "vk/run/a_runningpoints"
	vutils "vk/utils"

	"github.com/gorilla/mux"
)

/*
const STATIC_URL string = "/externals/"
const STATIC_ROOT string = "externals/"

const MANS_URL string = vomni.WebPrefix + "static/"
const MANS_ROOT string = "static/"

const BWANA string = "/static/"
*/

var rtr = mux.NewRouter()

/*
type Context struct {
	Title  string
	Static string
}

func home(w http.ResponseWriter, req *http.Request) {
	context := Context{Title: "Welcome!"}
	render(w, "index", context)
}

func about(w http.ResponseWriter, req *http.Request) {

	fmt.Println("Rogozin")

	context := Context{Title: "About"}
	render(w, "about", context)
}

func login(w http.ResponseWriter, req *http.Request) {
	context := Context{Title: "Login input"}
	render(w, "login", context)
}

func loginInp() (f func(http.ResponseWriter, *http.Request)) {
	return login
}

func aboutIt() (f func(http.ResponseWriter, *http.Request)) {
	return about
}
*/
func setMux() {
	rtr.HandleFunc("/about", pageAbout) //wrappedTmpl("about", ""))
	rtr.HandleFunc("/login", pageLogin) //     //.Methods("GET")
	rtr.HandleFunc("/", pageHome)
	rtr.HandleFunc("/home", pageHome)

	rtr.HandleFunc("/pointlist", tmplPointList)
	rtr.HandleFunc("/pointlist/data", tmplPointListData)

	rtr.HandleFunc("/point/{point}/{todo}", pointToDo)
	rtr.HandleFunc("/point/handlecfg/{point}/{todo}", handleCfg)
	rtr.HandleFunc("/station/{todo}", handleStation)

	http.HandleFunc(vomni.WebPrefix, StaticFile) // usually read files required for templates css, js, ...

	http.Handle("/", rtr)
}

func tmplPointList(w http.ResponseWriter, r *http.Request) {
	this_tmpl := "pointlist"

	fmt.Println("Kiriloff ", this_tmpl, " polina")

	err := tmpls.ExecuteTemplate(w, this_tmpl, r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func tmplPointListData(w http.ResponseWriter, r *http.Request) {
	//this_tmpl := "pointlist"

	fmt.Println("Maxim ", " polina")

	//	err := tmpls.ExecuteTemplate(w, this_tmpl, r)

	data := pointList()

	a, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(a)
}

/*
func petitBurdokovX(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("5 bocheks")

	data := pointCfg("meskals")

	a, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(a)
}
*/

func pageLogin(w http.ResponseWriter, r *http.Request) {
	pageStatic("login", w, r)
}

func pageHome(w http.ResponseWriter, r *http.Request) {
	pageStatic("home", w, r)
}

func pageAbout(w http.ResponseWriter, r *http.Request) {
	pageStatic("about", w, r)
}

func pageStatic(tmpl string, w http.ResponseWriter, r *http.Request) {

	var data interface{}

	err := tmpls.ExecuteTemplate(w, tmpl, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func pointToDo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	todo := vars["todo"]
	point := vars["point"]

	var err error
	var data interface{}

	switch todo {
	case "showcfg":
		tmplStr := "pointcfg"
		data = pointCfg(point)

		refl := reflect.ValueOf(data)

		zType := refl.FieldByName("Type")

		switch zType.Int() {
		case vomni.PointTypeRelayOnOffInterval:
			tmplStr = "cfgrelayonoffinterval"
		default:
			tmplStr = "pointcfg"
		}

		err = tmpls.ExecuteTemplate(w, tmplStr, point)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	case "getpointcfg":
		data := pointCfg(point)

		a, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(a)

	default:
		http.NotFound(w, r)
	}
}

func handleCfg(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	todo := strings.ToUpper(vars["todo"])
	point := vars["point"]
	var data interface{}

	switch todo {
	case "LOADCFG", "SAVECFG":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err.Error())
		}
		err = json.Unmarshal(body, &data)
		if err != nil {
			panic(err.Error())
		}
	case "FREEZE", "UNFREEZE", "LOADDEFAULTCFG", "LOADSAVEDCFG":
	default:
		log.Fatal(fmt.Sprintf("===> Don't know what to do with \"%s\"", todo))
	}

	responseOK(w)
	xrun.ReceivedWebMsg(point, todo, data)
}

func handleStation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todo := strings.ToUpper(vars["todo"])

	switch todo {
	case "SCANIP":

		chDone := make(chan bool)
		chErr := make(chan error)

		go vscanip.ScanPoints(chDone, chErr)

		responseOK(w)

		select {
		case <-chDone:
		case err := <-chErr:
			vomni.LogErr.Println(vutils.ErrFuncLine(err))
		}
	case "RESTART":
		vomni.RootDone <- vomni.DoneRestart
		responseOK(w)

	case "EXIT":
		vomni.RootDone <- vomni.DoneStop

	default:
		log.Fatal(fmt.Sprintf("===> Don't know what to do with \"%s\"", todo))
	}

	responseOK(w)
}

func responseOK(w http.ResponseWriter) {
	type resp struct {
		RC string
	}

	a, err := json.Marshal(resp{RC: "OK"})
	if nil != err {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(a)
}

/*
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
*/
func StaticFile(w http.ResponseWriter, req *http.Request) {
	staticFile := req.URL.Path[len(vomni.WebPrefix):]

	if len(staticFile) != 0 {
		f, err := http.Dir(vomni.WebStaticPath).Open(staticFile)
		if err == nil {
			content := io.ReadSeeker(f)
			http.ServeContent(w, req, staticFile, time.Now(), content)
			return
		}
	}
	http.NotFound(w, req)
}

/*
func Sta(w http.ResponseWriter, req *http.Request) {
	static_file := req.URL.Path[len(vomni.WebPrefix):]

	fmt.Println("URL    ", req.URL.Path)
	fmt.Println("STATIC ", static_file)

	if len(static_file) != 0 {
		f, err := http.Dir(vomni.WebStaticPath).Open(static_file)
		if err == nil {
			content := io.ReadSeeker(f)
			http.ServeContent(w, req, static_file, time.Now(), content)
			return
		}
	}
	http.NotFound(w, req)
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

func wrappedTmpl(tmplName string, str string) http.HandlerFunc {

	this_tmpl := ""

	//thisTmpl

	if "point-events" == tmplName {

		fmt.Println("******************************************************** Alexandr", tmplName, " Gims ", str)

		//	eventPointName = "bazargans"
	} else {
		//		fmt.Println("******************************************************** Alexandr", tmplName, " Gims ", str)
		fmt.Println("======================================================== Grigory ", tmplName, " Gims ", str)
	}

	fn := func(w http.ResponseWriter, r *http.Request) {
		this_tmpl = tmplName

		fmt.Println("Dokazateljstvo", tmplName, " sitro ", str)

		err := tmpls.ExecuteTemplate(w, this_tmpl, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	return fn
}

func wrappedDynamic(tmplName string, w http.ResponseWriter, r *http.Request, vars map[string]string) {

	pointName := vars["pointName"]
	files := allPageTmpls(tmplName)

	tmpls1 := template.Must(template.New("dyna").Funcs(tmplFuncs).ParseFiles(files...))

	type tmplBodyData struct {
		Events     []string
		EventCount int
	}

	lst := pointLastEvents(pointName)

	bodyData := tmplBodyData{
		Events:     lst,
		EventCount: len(lst),
	}

	data := make(map[string]interface{})

	data["body"] = bodyData

	err := tmpls1.ExecuteTemplate(w, tmplName, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
*/
