package web

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"reflect"
)

type cardAction struct {
	Text      string
	Href      string
	ExtraText string
}

type card struct {
	Title   string
	Content template.HTML
	Actions []cardAction
}

type modalContent struct {
	Card card
}

func addShort(w http.ResponseWriter, r *http.Request) {
	longUrl := r.URL.Query().Get("url")
	shortUrl, err := shortener.Shorten(longUrl)
	var content modalContent
	if reflect.TypeOf(&url.Error{}) == reflect.TypeOf(err) {
		content = modalContent{
			Card: card{
				Title:   "Could not create short link",
				Content: template.HTML(fmt.Sprintf("<p>You provided an invalid url: \"%s\"</p>", longUrl)),
			},
		}
	} else if err != nil {
		content = modalContent{
			Card: card{
				Title:   "An error occurred on our site",
				Content: template.HTML("<p>Please come back later or send an email to <a href=\"mailto:support@gur.la\">support[at]gur.la</a>.</p>"),
			},
		}
	} else {
		scheme := "http:"
		if r.TLS != nil {
			scheme = "https:"
		}
		fullShort := fmt.Sprintf("%s//%s/%s", scheme, r.Host, *shortUrl)
		content = modalContent{
			Card: card{
				Title:   "Your short link",
				Content: template.HTML(fmt.Sprintf("<p>Your short link is: <span class=\"short\">%s</span></p>", fullShort)),
				Actions: []cardAction{{Text: *shortUrl, Href: fullShort, ExtraText: "Go to:"}},
			},
		}
	}
	tmpl["modal.gohtml"].Execute(w, content)
}
