package web

import (
	"github.com/flweber/gurl/config"
	"github.com/go-redis/redis/v9"
	"net/http"
)

type errorData struct {
	Code        int
	Text        string
	Description string
}

type errorPageData struct {
	Conf  config.PageConfig
	Error errorData
}

func redirect(w http.ResponseWriter, r *http.Request, shortUrl string) {
	longUrl, err := shortener.GetUrl(shortUrl)
	w.Header().Set("content-type", "text/html")
	if err == redis.Nil {
		data := errorPageData{Conf: conf, Error: errorData{Code: 404, Text: "Not Found", Description: "The short url you're looking for doesn't exist."}}
		tmpl["error.gohtml"].ExecuteTemplate(w, "error", data)
	} else if err != nil {
		data := errorPageData{Conf: conf, Error: errorData{Code: 500, Text: "Internal Server Error", Description: "An error occurred on our site. Please try again later."}}
		tmpl["error.gohtml"].ExecuteTemplate(w, "error", data)
	} else {
		http.Redirect(w, r, *longUrl, http.StatusTemporaryRedirect)
	}
}
