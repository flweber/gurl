package web

import (
	"fmt"
	"github.com/Masterminds/sprig/v3"
	"github.com/flweber/gurl/config"
	shortenerpkg "github.com/flweber/gurl/shortener"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var shortener shortenerpkg.Shortener
var conf config.PageConfig
var tmpl map[string]*template.Template

func StartServer(shortenerParam shortenerpkg.Shortener, port int, c config.PageConfig) {
	shortener = shortenerParam
	conf = c
	prepareTemplates()
	http.HandleFunc("/", handleRoute)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func prepareTemplates() {
	tmpl = make(map[string]*template.Template)
	tmpl["page.gohtml"] = template.Must(template.New("page.gohtml").Funcs(sprig.FuncMap()).ParseFiles("web/templates/head.gohtml", "web/templates/header.gohtml", "web/templates/footer.gohtml", "web/templates/page.gohtml"))
	tmpl["error.gohtml"] = template.Must(template.New("error.gohtml").Funcs(sprig.FuncMap()).ParseFiles("web/templates/error.gohtml", "web/templates/head.gohtml", "web/templates/header.gohtml", "web/templates/footer.gohtml"))
	tmpl["modal.gohtml"] = template.Must(template.New("modal.gohtml").Funcs(sprig.FuncMap()).ParseFiles("web/templates/modal.gohtml"))
	for k := range tmpl {
		tmpl[k].Funcs(sprig.FuncMap())
	}
}

func handleRoute(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pathArgs := strings.Split(path, "/")
	if strings.HasSuffix(pathArgs[1], ".js") || strings.HasSuffix(pathArgs[1], ".css") {
		http.ServeFile(w, r, fmt.Sprintf("web/static/%s", pathArgs[1]))
	} else if len(pathArgs[1]) > 0 {
		redirect(w, r, pathArgs[1])
	} else {
		vRequest := r.Header.Get("V-Request")
		if len(vRequest) > 0 {
			addShort(w, r)
		} else {
			homePage(w, r)
		}
	}
}
