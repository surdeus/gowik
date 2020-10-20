package render

import(
	"html/template"
	"net/http"
	"github.com/surdeus/gowik/src/page"
)

var(
	Templates *template.Template
)

func
WriteTemplate(w http.ResponseWriter, tmpl string, p *page.Page) {
	e := Templates.ExecuteTemplate(w, tmpl, p)
	if e!=nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
	}
}
