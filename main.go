package main

import (
	"log"
	"io/ioutil"
	"net/http"
	"html/template"
	"regexp"
	"github.com/gomarkdown/markdown"
)


type HttpHandler struct {
	Prefix string
	Func func(w http.ResponseWriter, r http.Request)
}

type Page struct {
	Title string
	Body template.HTML
}

var stdHostString = ":8080"
var tmplPrefix = "tmpl"
var pagePrefix = "data/page"

var templates = template.Must(template.ParseFiles(tmplPrefix+"/edit", tmplPrefix+"/view"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func
getPageFilename(t string) string {
	return pagePrefix+"/"+t 
}

func
(p *Page)save() error {
	filename := getPageFilename(p.Title)
	return ioutil.WriteFile(filename, []byte(p.Body), 0600)
}

func
getPage(t string) (*Page, error) {
	filename := getPageFilename(t)
	b, e := ioutil.ReadFile(filename)
	if e!=nil {
		return nil, e
	}
	return &Page{Title: t, Body: template.HTML(b)}, nil
}

func
makeHandler(fn func(http.ResponseWriter, *http.Request, string) ) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m==nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func
viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, e := getPage(title)
	if e!=nil { /* Not existing page. */
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	p.Body = template.HTML(markdown.ToHTML( []byte(p.Body), nil, nil))
	renderTemplate(w, "view", p)
}

func
editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, e := getPage(title)
	if e!=nil { /* Create new page. */
		p = &Page{Title: title}
	}

	renderTemplate(w, "edit", p)
}

func
saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: template.HTML(body)}

	e := p.save()
	if e!=nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func
renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	e := templates.ExecuteTemplate(w, tmpl, p)
	if e!=nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
	}
}

func
rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/view/home", http.StatusFound)
}

func
main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/", rootHandler)

	log.Fatal(http.ListenAndServe(stdHostString, nil))
}
