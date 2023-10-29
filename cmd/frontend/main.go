package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/eluamous-droid/godo/pkg/web"
)

func main() {

	component := web.Hello("Mikkel")

	http.Handle("/", templ.Handler(component))

	http.ListenAndServe(":3000", nil)
}
