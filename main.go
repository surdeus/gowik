package main

import (
	"os"
	"fmt"
	"log"
	"flag"
	"net/http"
	"html/template"
	"github.com/surdeus/gowik/path"
	"github.com/surdeus/gowik/handle"
	"github.com/surdeus/gowik/render"
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
	flag.StringVar(&path.DataDir, "data", "data", "data dir")
	flag.StringVar(&path.TmplDir, "tmpl", "tmpl", "templates dir")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [options]\n", arg0)
		flag.PrintDefaults()
	}
	flag.Parse()
	if len(flag.Args()) != 0 {
		flag.Usage()
		os.Exit(1)
	}

	path.SaltFile = path.DataDir+"/salt"
	path.PageDir = path.DataDir+"/page"
	path.HashDir = path.DataDir+"/hash"

	/*if salt, err = ioutil.ReadFile(saltFile) ; err!=nil {
		log.Fatal(err)
	}*/

	render.Templates = template.Must(template.ParseFiles(path.TmplDir+"/edit", path.TmplDir+"/view"))

	http.HandleFunc("/view/", handle.MakeHandler(handle.View))
	http.HandleFunc("/save/", handle.MakeHandler(handle.Save))
	http.HandleFunc("/edit/", handle.MakeHandler(handle.Edit))
	http.HandleFunc("/", handle.Root)

	log.Fatal(http.ListenAndServe(hostString, nil))
}
