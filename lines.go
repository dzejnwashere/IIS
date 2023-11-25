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

type StructForLines struct {
	Linky []db.Line_t
	Stops []db.Stop_t
}

func lines(writer http.ResponseWriter, request *http.Request) {
	if !auth.HasPermission(request, typedef.SpravcePerm) {
		writer.WriteHeader(403)
		fmt.Fprint(writer, "403 insufficient permissions")
		return
	}
	lines := db.GetAllLines()
	stops := db.GetStops2()

	files, err := template.ParseFiles("res/tmpl/lines.html")
	if err != nil {
		fmt.Fprintf(writer, err.Error())
	}

	err = files.Execute(writer, StructForLines{
		Linky: lines,
		Stops: stops,
	})
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
	if request.Method == "POST" {
		lineID, err := strconv.Atoi(request.FormValue("line_id"))
		if err != nil {
			log.Print("Error converting line_id to int in /line_stops", err)
			writer.WriteHeader(400)
			return
		}
		stopID, err := strconv.Atoi(request.FormValue("stop_id"))
		if err != nil {
			log.Print("Error converting stop_id to int in /line_stops", err)
			writer.WriteHeader(400)
			return
		}
		timeVal := request.FormValue("time")
		err = db.AddLineStops(db.Stop_line_t{
			Stop_id:   stopID,
			Stop_name: "IDK",
			Time:      timeVal,
			Line_id:   lineID,
		})
		if err != nil {
			writer.WriteHeader(400)
			_, _ = fmt.Fprintf(writer, "Error inserting into the db: %s", err.Error())
			return
		}
	}
	if request.Method == "GET" || request.Method == "POST" {
		lineID, err := strconv.Atoi(request.FormValue("line_id"))
		if err != nil {
			log.Print("Error converting line_id to int in /line_stops", err)
			writer.WriteHeader(400)
		}
		stops := db.GetLineStops(lineID)
		for _, stop := range stops {
			fmt.Fprintf(writer, "%d;%s;%s\n", stop.Stop_id, stop.Stop_name, stop.Time)
		}
	}
}
