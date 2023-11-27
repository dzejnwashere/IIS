package main

import (
	"IIS/auth"
	"IIS/db"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strconv"
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

type JizdyStruct struct {
	DnyJizdy []db.DenJizdy_t
}

func jizdy(writer http.ResponseWriter, request *http.Request) {
	files, err := template.ParseFiles("res/tmpl/jizdy.html")
	if err != nil {
		fmt.Println("b" + err.Error())
		fmt.Fprintf(writer, err.Error())
	}
	dnyJizdy := db.GetAllDnyJizdy()
	jizdyStruct := JizdyStruct{DnyJizdy: dnyJizdy}
	err = files.Execute(writer, jizdyStruct)
	if err != nil {
		return
	}
}

func jizdyAPI(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		s, ok := mux.Vars(request)["den"]
		if !ok {
			http.Error(writer, "missing day", 400)
		}
		rides := db.GetDayRides(s)
		for _, ride := range rides {
			fmt.Fprintf(writer, "%d;%s;%s;%s;%s;%s;%d;%s\n", ride.Id, ride.LineName, ride.StartStop.Name, ride.StartTime, ride.EndStop.Name, ride.EndTime, ride.Driver, ride.Vuz)
		}
	}
}
func jizdyRidici(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		s, ok := mux.Vars(request)["jizdaID"]
		if !ok {
			http.Error(writer, "missing day", 400)
		}
		rides := db.GetDayRides(s)
		for _, ride := range rides {
			fmt.Fprintf(writer, "%d;%s;%s;%s;%s;%s;%d;%s\n", ride.Id, ride.LineName, ride.StartStop.Name, ride.StartTime, ride.EndStop.Name, ride.EndTime, ride.Driver, ride.Vuz)
		}
	}
}

func kategorie_dne(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		s, ok := mux.Vars(request)["den"]
		if !ok {
			http.Error(writer, "missing day", 400)
		}
		fmt.Fprint(writer, db.GetDayCategory(s))
	}
	if request.Method == "POST" {
		s, ok := mux.Vars(request)["den"]
		if !ok {
			http.Error(writer, "missing day", 400)
		}
		dayIdS := request.FormValue("den_jizdy")
		dayId, err := strconv.Atoi(dayIdS)
		if err != nil {
			log.Print("Error converting day_id to int in /jizdy", err)
			writer.WriteHeader(400)
			return
		}
		fmt.Printf("dayid %d", dayId)
		err = db.SetDayCategory(s, dayId)
		if err != nil {

			log.Print("Error setting day in kategorie_dne", err)
			writer.WriteHeader(400)
			return
		}
	}

}
