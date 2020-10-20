package main

import (
	"os"
	"fmt"
	"log"
	"flag"
	"net/http"
	"html/template"
	"github.com/surdeus/gowik/src/path"
	"github.com/surdeus/gowik/src/handle"
	"github.com/surdeus/gowik/src/render"
)

var(
	arg0 string
	hostString string
	salt int
)

func
main() {
	arg0 = os.Args[0]
	flag.StringVar(&hostString, "host", ":8080", "host string")
	flag.StringVar(&path.WebDir, "web", "web", "web dir")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [options]\n", arg0)
		flag.PrintDefaults()
	}
	flag.Parse()
	if len(flag.Args()) != 0 {
		flag.Usage()
		os.Exit(1)
	}

	path.DataDir = path.WebDir+"/data"
	path.StaticDir = path.WebDir+"/static"
	path.TmplDir = path.WebDir+"/tmpl"
	path.SaltFile = path.DataDir+"/salt"
	path.PageDir = path.DataDir+"/page"
	path.HashDir = path.DataDir+"/hash"

	render.Templates = template.Must( template.ParseFiles(
		path.TmplDir+"/page-edit",
		path.TmplDir+"/page-view" ) )

	http.HandleFunc("/", handle.Root)

	http.HandleFunc("/user/edit/", handle.MakeUserHandler(handle.UserEdit))
	http.HandleFunc("/user/save/", handle.MakeUserHandler(handle.UserSave))

	http.HandleFunc("/page/view/", handle.MakePageHandler(handle.PageView))
	http.HandleFunc("/page/edit/", handle.MakePageHandler(handle.PageEdit))
	http.HandleFunc("/page/save/", handle.MakePageHandler(handle.PageSave))

	log.Fatal(http.ListenAndServe(hostString, nil))
}
