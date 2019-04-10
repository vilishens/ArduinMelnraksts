package templates

import (
	"fmt"
	"html/template"
	"net/http"
	"vk/cli"

	"github.com/gorilla/mux"
)

var tmplFiles []string

func init() {
	//	allTmplFiles()
	tmplFuncs["appTitle"] = title //      func(id string) string { return fmt.Sprintf("%v", flag.Lookup(id).Value) },
}

func allTmplFiles() string {
	//	base_path := cli.Param(cli.TmplPath)

	tmplX_Files := []string{
		"tmpl/about/about.tmpl",
		"tmpl/about/about_body.tmpl",
		"tmpl/base/base_footer.tmpl",
		"tmpl/base/base_header.tmpl",
		"tmpl/base/base_include.tmpl",
		"tmpl/base/base_navigation.tmpl",
		"tmpl/home/home.tmpl",
		"tmpl/home/home_body.tmpl",
	}

	tmplFiles = tmplX_Files
	/*
		for _, str := range tmplX_files {
			tmplFiles = append(tmplFiles, str)
		}
	*/
	fmt.Println(tmplFiles)

	//	return strings.Join(tmplX_files, ",")
	return ""
	/*
		tmpl_files := []string{
			"base/base_header.tmpl",
			//* "base/base_navigation.tmpl",
			//"base/base_js.tmpl",
			//		"base/login_1.tmpl",
			//		"base/login.tmpl",
			//		"home/kira1.a"}
			"home/home.tmpl",
			"home/home_body.tmpl"}
		//"home/manukjans.ashots",
		//"home/kilsei.murs"}
		//		"settings/restore/restore.tmpl",
		//		"settings/restore/restore_body.tmpl",
		//		"settings/communication/communication.tmpl",
		//		"settings/communication/communication_body.tmpl",
		//		"settings/communication/communication_select_user.tmpl",
		//		"settings/communication/communication_select_user_settings.tmpl",
		//		"settings/communication/communication_select_user_settings_new_connection.tmpl",
		//		"settings/communication/communication_user_connections.tmpl",
		//		"settings/communication/communication_user_connections_by_type.tmpl",
		//		"settings/users/users.tmpl",
		//		"settings/users/users_body.tmpl",
		//		"settings/users/users_change_pass.tmpl",
		//		"settings/users/users_new_user.tmpl",
		//		"settings/users/users_input_hidden.tmpl" }

		for _, str := range tmpl_files {
			tmplFiles = append(tmplFiles, base_path+str)
		}
	*/
}

var tmplFuncs = template.FuncMap{
	"servPrefix": func() string { return cli.Param(cli.CliFileServPrefix) },
	"appTitle":   title()}

//	"langTxt":             func(id string) string {
//		return jsonData.LangTxt[id][info.AppLang()]
//	},
//	"lang_flags":          func() map[string]string {
//		return jsonData.LangTxt["lang_flag"]
//	},
//	"tmplDataTxt":         func(id string) string {
//		return jsonData.TmplDataTxt[id]
//	}

func title() string {
	return "TITLE*TITLE*TITLE"
}

var Tmpls *template.Template
var Rtr *mux.Router

func SetTemplates() {
	setTemplates()
	setMux()
}

func setTemplatesX() {
	fmt.Println("KILOMETER")
	allTmplFiles()
	gra := "tmpl/about/about.tmpl"
	Tmpls = template.Must(template.New("app").Funcs(tmplFuncs).ParseFiles(gra))
	//	Tmpls = template.Must(template.New("app").Funcs(tmplFuncs).ParseFiles(tmplFiles[2]))
	//Tmpls = template.Must(template.New("app").Funcs(tmplFuncs).ParseFiles("tmpl/about/about.tmpl"))
	/*
		Tmpls = template.Must(template.New("app").Funcs(tmplFuncs).ParseFiles(
			"tmpl/about/about.tmpl",
			"tmpl/about/about_body.tmpl",
			"tmpl/base/base_footer.tmpl",
			"tmpl/base/base_header.tmpl",
			"tmpl/base/base_include.tmpl",
			"tmpl/base/base_navigation.tmpl",
			"tmpl/home/home.tmpl",
			"tmpl/home/home_body.tmpl",
		))
	*/
}

func setTemplates() {
	Tmpls = template.Must(template.New("app").Funcs(tmplFuncs).ParseFiles(
		"tmpl/about/about.tmpl",
		"tmpl/about/about_body.tmpl",
		"tmpl/base/base_footer.tmpl",
		"tmpl/base/base_header.tmpl",
		"tmpl/base/base_include.tmpl",
		"tmpl/base/base_navigation.tmpl",
		"tmpl/home/home.tmpl",
		"tmpl/home/home_body.tmpl",
	))
}

func setMux() {
	Rtr = mux.NewRouter()

	Rtr.HandleFunc("/", makeWrappedTmplHandler("home_tmpl", "", ""))
	Rtr.HandleFunc("/about", makeWrappedTmplHandler("about_tmpl", "", ""))
	http.Handle("/", Rtr)

}

func SetMux() {
	Rtr = mux.NewRouter()

	Rtr.HandleFunc("/", makeWrappedTmplHandler("home_tmpl", "", ""))
	Rtr.HandleFunc("/about", makeWrappedTmplHandler("about_tmpl", "", ""))
	http.Handle("/", Rtr)

}

func makeWrappedTmplHandler(tmplName string, flag_value string, match string) http.HandlerFunc {

	this_tmpl := ""

	fn := func(w http.ResponseWriter, r *http.Request) {
		this_tmpl = tmplName
		err := Tmpls.ExecuteTemplate(w, this_tmpl, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	return fn
}
