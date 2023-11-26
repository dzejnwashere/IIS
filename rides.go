package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func plan(writer http.ResponseWriter, request *http.Request) {
	files, err := template.ParseFiles("res/tmpl/plan.html")
	if err != nil {
		fmt.Fprintf(writer, err.Error())
	}

	err = files.Execute(writer, nil)
	if err != nil {
		return
	}
}
