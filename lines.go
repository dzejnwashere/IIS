package main

import (
	"IIS/auth"
	"IIS/db"
	"IIS/typedef"
	"fmt"
	"github.com/gorilla/mux"
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
	if request.Method == "POST" {
		linename, err := strconv.Atoi(request.FormValue("linename"))
		if err != nil {
			log.Print("POST /lines No line namesupplied")
			writer.WriteHeader(400)
			return
		}
		err = db.CreateLine(strconv.Itoa(linename))
		if err != nil {
			log.Print("POST /lines Error inserting into db")
			writer.WriteHeader(400)
			return
		}
	}
	if request.Method == "GET" || request.Method == "POST" {
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
	if request.Method == "DELETE" {
		vars := mux.Vars(request)
		lineID, ok := vars["lineno"]
		lineIDI, err := strconv.Atoi(lineID)
		if !ok {
			log.Print("DELETE /lines No line id supplied")
			writer.WriteHeader(400)
			return
		}
		if err != nil {
			log.Print("Error converting line_id to int in /lines", err)
			writer.WriteHeader(400)
			return
		}
		err = db.DeleteLine(lineIDI)
		if err != nil {
			writer.WriteHeader(400)
			log.Println(err)
			fmt.Fprint(writer, err.Error())
			return
		}
	}
}

func line_stops(writer http.ResponseWriter, request *http.Request) {
	if !auth.HasPermission(request, typedef.SpravcePerm) {
		writer.WriteHeader(403)
		fmt.Fprint(writer, "403 insufficient permissions")
		return
	}
	if request.Method == "POST" {
		lineStopID, err := strconv.Atoi(request.FormValue("line_stop_id"))
		isNew := false
		if err != nil {
			isNew = true
		}
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
		if isNew {
			err = db.AddLineStops(db.Stop_line_t{
				Stop_id:   stopID,
				Stop_name: "IDK",
				Time:      timeVal,
				Line_id:   lineID,
			})
		} else {
			err = db.UpdateLineStops(db.Stop_line_t{
				Stop_id:      stopID,
				Stop_line_id: lineStopID,
				Stop_name:    "IDK",
				Time:         timeVal,
				Line_id:      lineID,
			})
		}

		if err != nil {
			writer.WriteHeader(400)
			_, _ = fmt.Fprintf(writer, "Error inserting into the db: %s", err.Error())
			return
		}
	}
	if request.Method == "DELETE" {
		vars := mux.Vars(request)
		lineID, ok := vars["lsno"]
		lineIDI, err := strconv.Atoi(lineID)
		if !ok {
			log.Print("DELETE /line_stop No line id supplied")
			writer.WriteHeader(400)
			return
		}
		if err != nil {
			log.Print("Error converting line_stop_id to int in /line_stop", err)
			writer.WriteHeader(400)
			return
		}
		err = db.DeleteStopLine(lineIDI)
		if err != nil {
			writer.WriteHeader(400)
			log.Println(err)
			fmt.Fprint(writer, err.Error())
			return
		}
	}
	if request.Method == "GET" || request.Method == "POST" {
		lineid, err := strconv.Atoi(request.FormValue("line_id"))
		if err != nil {
			log.Print("Error converting line_id to int in /line_stops", err)
			writer.WriteHeader(400)
			return
		}
		stops := db.GetLineStops(lineid)
		for _, stop := range stops {
			_, err := fmt.Fprintf(writer, "%d;%d;%s;%s\n", stop.Stop_line_id, stop.Stop_id, stop.Stop_name, stop.Time)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

type StructForSpoje struct {
	Linky    []db.Line_t
	Manage   bool
	Spoje    []db.Spoj_t
	DnyJizdy []db.DenJizdy_t
}

func spoje(writer http.ResponseWriter, request *http.Request) {
	var no string
	var ok bool
	if request.Method == "POST" {
		time := request.FormValue("time")
		dayIdS := request.FormValue("day_id")
		dayId, err := strconv.Atoi(dayIdS)
		if err != nil {
			log.Print("Error converting day_id to int in /spoje", err)
			writer.WriteHeader(400)
			return
		}
		spojIdS := request.FormValue("spoj_id")
		spojId, err := strconv.Atoi(spojIdS)
		isNew := false
		if err != nil {
			isNew = true
		}
		lineIdS := request.FormValue("line_id")
		lineId, err := strconv.Atoi(lineIdS)
		if err != nil {
			log.Print("Error converting lineId to int in /spoje", err)
			writer.WriteHeader(400)
			return
		}
		isPrimaryS := request.FormValue("isprimary")
		if isNew {
			db.CreateSpoj(db.Spoj_t{
				Id:           spojId,
				Linka:        lineId,
				CasOdjezdu:   time,
				PrimarniSmer: isPrimaryS == "true",
				DenJizdy:     dayId,
			})
		} else {
			db.UpdateSpoj(db.Spoj_t{
				Id:           spojId,
				Linka:        lineId,
				CasOdjezdu:   time,
				PrimarniSmer: isPrimaryS == "true",
				DenJizdy:     dayId,
			})
		}
		ok = true
		no = lineIdS
	}
	if request.Method == "GET" {
		no, ok = mux.Vars(request)["spojno"]
	}
	if request.Method == "GET" || request.Method == "POST" {
		if ok {
			lineno, err := strconv.Atoi(no)
			if err != nil {
				http.Error(writer, "invalid ine no", 400)
			}
			spojs := db.GetSpojeByLine(lineno)
			for _, spoj := range spojs {
				fmt.Fprintf(writer, "%d;%s;%t;%d\n", spoj.Id, spoj.CasOdjezdu, spoj.PrimarniSmer, spoj.DenJizdy)
			}
			return
		} else {
			lns := db.GetAllLines()
			spojs := db.GetAllSpoje()
			jizdy := db.GetAllDnyJizdy()
			str := StructForSpoje{
				Linky:    lns,
				Manage:   auth.HasPermission(request, typedef.SpravcePerm),
				Spoje:    spojs,
				DnyJizdy: jizdy,
			}

			files, err := template.ParseFiles("res/tmpl/spoje.html")
			if err != nil {
				fmt.Fprintf(writer, err.Error())
			}

			err = files.Execute(writer, str)
			if err != nil {
				return
			}
		}
	}
}
