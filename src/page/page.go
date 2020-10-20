package page
import(
	"github.com/surdeus/gowik/src/path"
	"io/ioutil"
	"html/template"
	"strings"
)

type Page struct {
	User string
	Title string
	Body template.HTML
}

func
(p *Page)Save() error {
	file := path.PageFile(p.User, p.Title)
	e := ioutil.WriteFile(file, []byte(strings.ReplaceAll(string(p.Body), "\r", "")), 0600)
	return e
}

func
Get(u string, t string) (*Page, error) {
	b, e := ioutil.ReadFile(path.PageFile(u, t))
	if e !=nil {
		return nil, e
	}
	return &Page{User: u, Title: t, Body: template.HTML(b)}, nil
}
