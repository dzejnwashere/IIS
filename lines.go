package main

import (
	"IIS/auth"
	"IIS/db"
	"IIS/typedef"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func lines(writer http.ResponseWriter, request *http.Request) {
	if !auth.HasPermission(request, typedef.SpravcePerm) {
		writer.WriteHeader(403)
		fmt.Fprint(writer, "403 insufficient permissions")
		return
	}
	lines := db.GetAllLines()

	files, err := template.ParseFiles("res/tmpl/lines.html")
	if err != nil {
		fmt.Fprintf(writer, err.Error())
	}

	err = files.Execute(writer, lines)
	if err != nil {
		return
	}
}

func line_stops(writer http.ResponseWriter, request *http.Request) {
	if !auth.HasPermission(request, typedef.SpravcePerm) {
		writer.WriteHeader(403)
		fmt.Fprint(writer, "403 insufficient permissions")
		return
	}
	lineID, err := strconv.Atoi(request.URL.Query().Get("line_id"))
	if err != nil {
		log.Print("Error converting line_id to int in /line_stops", err)
		writer.WriteHeader(400)
	}
	stops := db.GetLineStops(lineID)
	for _, stop := range stops {
		fmt.Fprintf(writer, "%d;%s;%s\n", stop.Stop_id, stop.Stop_name, stop.Time)
	}
}
