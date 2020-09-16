package handle

import(
	"html/template"
	"net/http"
	"regexp"
	"github.com/surdeus/gowik/pass"
	"github.com/surdeus/gowik/page"
	"github.com/surdeus/gowik/render"
	"github.com/surdeus/gowik/markdown"
	"github.com/surdeus/gowik/sanitizer"
)

var ValidPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9_]+)/([a-zA-Z0-9_]+)$")

func
MakeHandler(fn func(http.ResponseWriter, *http.Request, string, string) ) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := ValidPath.FindStringSubmatch(r.URL.Path)
		if m==nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2], m[3])
	}
}

func
Edit(w http.ResponseWriter, r *http.Request, u string, t string) {
	p, e := page.Get(u, t)
	if e!=nil { /* Create new page. */
		p = &page.Page{User: u, Title: t}
	}

	render.WriteTemplate(w, "edit", p)
}

func
View(w http.ResponseWriter, r *http.Request, u string, t string) {
	p, e := page.Get(u, t)
	if e!=nil { /* Not existing page. */
		http.Redirect(w, r, "/edit/"+"/"+u+"/"+t, http.StatusFound)
		return
	}

	buf := markdown.Process([]byte(p.Body))
	p.Body = template.HTML(sanitizer.Sanitize(buf))
	render.WriteTemplate(w, "view", p)
}

func
Save(w http.ResponseWriter, r *http.Request, u string, t string) {
	var(
		e error
	)
	body := r.FormValue("body")
	pwd := r.FormValue("password")

	if !pass.IsCorrect(u, pwd) {
		http.NotFound(w, r)
		return
	}

	p := &page.Page{User: u, Title: t, Body: template.HTML(body)}

	e = p.Save()
	if e!=nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/"+"/"+u+"/"+t, http.StatusFound)
}

func
Root(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/view/admin/home", http.StatusFound)
}
