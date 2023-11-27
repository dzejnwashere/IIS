package main

import (
	"IIS/auth"
	"IIS/db"
	"fmt"
	"html/template"
	"net/http"
)

type PlanStruct struct {
	Jizdy []db.Jizda_t
}

func plan(writer http.ResponseWriter, request *http.Request) {
	files, err := template.ParseFiles("res/tmpl/plan.html")
	if err != nil {
		fmt.Fprintf(writer, err.Error())
	}
	planStruct := PlanStruct{Jizdy: db.GetMyRides(int(auth.GetUserId(request)))}
	err = files.Execute(writer, planStruct)
	if err != nil {
		return
	}
}
