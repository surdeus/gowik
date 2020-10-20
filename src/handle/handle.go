package handle

import(
	"html/template"
	"net/http"
	"regexp"
	"github.com/surdeus/gowik/src/path"
	"github.com/surdeus/gowik/src/user"
	"github.com/surdeus/gowik/src/page"
	"github.com/surdeus/gowik/src/render"
	"github.com/surdeus/gowik/src/markdown"
	"github.com/surdeus/gowik/src/sanitizer"
)

var ValidPagePath = regexp.MustCompile("^/page/(edit|save|view)/([a-zA-Z0-9_-]+)/([a-zA-Z0-9_]+)$")
var ValidUserPath = regexp.MustCompile("^/user/(edit|save|make)/([a-zA-Z0-9_-]+)$")
var RootPath = "/view/admin/home"

func
MakePageHandler(fn func(http.ResponseWriter, *http.Request, string, string) ) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := ValidPagePath.FindStringSubmatch(r.URL.Path)
		if m==nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2], m[3])
	}
}

func
MakeUserHandler(fn func(http.ResponseWriter, *http.Request, string) ) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := ValidUserPath.FindStringSubmatch(r.URL.Path)
		if m==nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func
PageEdit(w http.ResponseWriter, r *http.Request, u string, t string) {
	p, e := page.Get(u, t)
	if e!=nil { /* Create new page. */
		p = &page.Page{User: u, Title: t}
	}

	render.WriteTemplate(w, "page-edit", p)
}

func
PageView(w http.ResponseWriter, r *http.Request, u string, t string) {
	p, e := page.Get(u, t)
	if e!=nil { /* Not existing page. */
		http.NotFound(w, r)
		return
	}

	buf := markdown.Process([]byte(p.Body))
	p.Body = template.HTML(sanitizer.Sanitize(buf))
	render.WriteTemplate(w, "page-view", p)
}

func
PageSave(w http.ResponseWriter, r *http.Request, u string, t string) {
	var(
		e error
	)
	body := r.FormValue("body")
	pwd := r.FormValue("password")

	if !user.IsPasswordCorrect(u, pwd) {
		http.NotFound(w, r)
		return
	}

	p := &page.Page{User: u, Title: t, Body: template.HTML(body)}

	e = p.Save()
	if e!=nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/page/view/"+"/"+u+"/"+t, http.StatusFound)
}

func
UserEdit(w http.ResponseWriter, r *http.Request, u string) {
	if !user.Exist(u) {
		http.Redirect(w, r, "/user/make/"+u, http.StatusFound)
		return
	}
	http.ServeFile(w, r, path.StaticDir+"/user-edit.htm")
}

func
UserSave(w http.ResponseWriter, r *http.Request, u string) {
}

func
Root(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, RootPath, http.StatusFound)
}
