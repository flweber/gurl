package web

import (
	"github.com/flweber/gurl/config"
	"net/http"
)

type homePageData struct {
	Conf config.PageConfig
}

func homePage(w http.ResponseWriter, r *http.Request) {
	data := homePageData{Conf: conf}
	tmpl["page.gohtml"].ExecuteTemplate(w, "index", data)
}
