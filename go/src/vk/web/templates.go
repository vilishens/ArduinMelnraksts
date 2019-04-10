package web

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"path/filepath"

	vparam "vk/params"
	vutils "vk/utils"
)

var tmpls = new(template.Template)
var tmplFiles []string
var tmplFuncs template.FuncMap

var tmplPath = "tmpl" // the base path of templates

func allPageTmpls(page string) (files []string) {

	files = append(files, allDirTmpls("base")...)
	files = append(files, allDirTmpls(page)...)

	return
}

func allDirTmpls(dir string) (files []string) {

	path := filepath.Join(vparam.Params.TemplatePath, dir)
	path = vutils.FileAbsPath(path, "")

	foundF, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, f := range foundF {
		if f.IsDir() {
			continue
		}
		if filepath.Ext(f.Name()) == vparam.Params.TemplateExt {

			fmt.Println("MIKLOVANS ", path, f.Name())

			files = append(files, filepath.Join(path, f.Name()))
		}
	}

	return
}

func init() {
	initTmpls()
	initFuncs()

	fmt.Println("ALEXEY Skaits", len(tmplFiles))

	for i := range tmplFiles {
		fmt.Println(i, "Meliora", tmplFiles[i])
	}

	tmpls = template.Must(template.New("app").Funcs(tmplFuncs).ParseFiles(tmplFiles...))

	ts := tmpls.DefinedTemplates()

	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
	fmt.Printf("Temlastes %v\n", ts)
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")

}

func initTmpls() {
	tmplFiles = []string{
		"tmpl/base/base-footer.tmpl",
		"tmpl/base/base-header.tmpl",
		"tmpl/base/base-js.tmpl",
		"tmpl/base/base-navigation.tmpl",
	}

	addPage(&tmplFiles, tmplPath, "home")
	addPage(&tmplFiles, tmplPath, "about")
	addPage(&tmplFiles, tmplPath, "login")
	addPage(&tmplFiles, tmplPath, "points/pointlist")
	addPage(&tmplFiles, tmplPath, "points/pointcfg/relayonoffinterval")
}

func addPage(files *([]string), path string, page string) {

	dir := filepath.Dir(page)
	base := filepath.Base(page)

	newF := filepath.Join(path, dir, base, base)
	*files = append(*files, newF+".tmpl")
	*files = append(*files, newF+"-body.tmpl")
}

func initFuncs() {
	tmplFuncs = make(map[string]interface{})
	tmplFuncs["raspName"] = raspName
	tmplFuncs["pointList"] = pointList
	tmplFuncs["pointCfg"] = pointCfg
	tmplFuncs["increment1"] = increment1
	tmplFuncs["pointCfgJsFile"] = pointCfgJsFile
	tmplFuncs["webPrefix"] = webPrefix
	tmplFuncs["webPref"] = webPref
}
