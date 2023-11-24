package main

import (
	"IIS/auth"
	"IIS/db"
	"IIS/typedef"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
)
import _ "github.com/go-sql-driver/mysql"

func login(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil || len(request.Form) == 0 {
		static_site("login", typedef.PublicPerm)(writer, request)
		return
	}
	jwtToken, err := auth.Authenticate(request.PostFormValue("username"), request.PostFormValue("password"))
	if err != nil {
		writer.WriteHeader(401)
		fmt.Fprintf(writer, err.Error())
		return
	}
	cookie := http.Cookie{
		Name:   "iisauth",
		Value:  jwtToken,
		MaxAge: 600,
	}
	http.SetCookie(writer, &cookie)
	http.Redirect(writer, request, "/", 302)
}

func logout(writer http.ResponseWriter, request *http.Request) {
	cookie := http.Cookie{
		Name:   "iisauth",
		Value:  "",
		MaxAge: 0,
	}
	http.SetCookie(writer, &cookie)
	http.Redirect(writer, request, "/", 302)
}

type IndexPageData struct {
	LoggedIn bool
	Admin    bool
	Username string
}

func index(writer http.ResponseWriter, request *http.Request) {
	files, err := template.ParseFiles("res/tmpl/index.html")
	if err != nil {
		fmt.Fprintf(writer, err.Error())
	}
	username, _ := db.GetUsername(auth.GetUserId(request))
	data := IndexPageData{
		LoggedIn: auth.HasPermission(request, typedef.UnprotectedPerm),
		Admin:    auth.HasPermission(request, typedef.AdminPerm),
		Username: username,
	}
	err = files.Execute(writer, data)
	if err != nil {
		return
	}
}

func user_management(writer http.ResponseWriter, request *http.Request) {
	users := db.GetAllUsers()
	fmt.Println(users)

	files, err := template.ParseFiles("res/tmpl/usrmngmt.html")
	if err != nil {
		fmt.Fprintf(writer, err.Error())
	}

	err = files.Execute(writer, users)
	if err != nil {
		return
	}
}

func remove(writer http.ResponseWriter, request *http.Request) {
	userID := request.URL.Query().Get("userID")
	fmt.Println(userID)
	myUserID, _ := strconv.Atoi(userID)
	db.RemoveUser(myUserID)

	writer.WriteHeader(200)
}

// Creates a function for serving a static file. Accepts required permission for displaying the site, -1 bypasses the check
func static_site(template_name string, perm typedef.Permission) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		if !auth.HasPermission(request, perm) {
			writer.WriteHeader(403)
			fmt.Fprintf(writer, "403 insufficient permissions")
			return
		}
		files, err := template.ParseFiles("res/tmpl/" + template_name + ".html")
		if err != nil {
			fmt.Fprintf(writer, err.Error())
		}
		err = files.Execute(writer, nil)
		if err != nil {
			return
		}
	}
}

func update_perms(writer http.ResponseWriter, request *http.Request) {
	if !auth.HasPermission(request, typedef.AdminPerm) {
		writer.WriteHeader(403)
		fmt.Fprint(writer, "403 insufficient permissions")
		return
	}
	userID, _ := strconv.Atoi(request.URL.Query().Get("userID"))
	perms, _ := strconv.Atoi(request.URL.Query().Get("perms"))

	fmt.Println(userID)
	fmt.Println(perms)

	err := db.UpdatePermissions(userID, perms)
	if err != nil {
		writer.WriteHeader(403)
		fmt.Fprint(writer, "403 insufficient permissions")
		return
	}
}

func demo(writer http.ResponseWriter, request *http.Request) {
	if !auth.HasPermission(request, typedef.AdminPerm) {
		writer.WriteHeader(403)
		fmt.Fprint(writer, "403 insufficient permissions")
		return
	}
	err := db.FeedDemoData()
	if err != nil {
		writer.WriteHeader(403)
		fmt.Fprint(writer, err.Error())
		return
	}
	writer.WriteHeader(200)
	fmt.Fprint(writer, "Succesfully inserted demo data")
}

func doc(writer http.ResponseWriter, request *http.Request) {
	files, err := template.ParseFiles("res/tmpl/doc.html")
	if err != nil {
		fmt.Fprintf(writer, err.Error())
	}

	err = files.Execute(writer, nil)
	if err != nil {
		return
	}
}

func failures(writer http.ResponseWriter, request *http.Request) {
	failures := db.GetFailures()

	files, err := template.ParseFiles("res/tmpl/failures.html")
	if err != nil {
		fmt.Fprintf(writer, err.Error())
	}

	err = files.Execute(writer, failures)
	if err != nil {
		return
	}
}

func technical_records(writer http.ResponseWriter, request *http.Request) {
	technicalRecords := db.GetTechnicalRecords()

	files, err := template.ParseFiles("res/tmpl/technical-records.html")
	if err != nil {
		fmt.Fprintf(writer, err.Error())
	}

	err = files.Execute(writer, technicalRecords)
	if err != nil {
		return
	}
}

func get_stops(writer http.ResponseWriter, request *http.Request) {
	stops := db.GetStops()

	stopsJSON, err := json.Marshal(stops)
	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(stopsJSON)
}

func get_spzs(writer http.ResponseWriter, request *http.Request) {
	spzs := db.GetSPZs()

	spzsJSON, err := json.Marshal(spzs)
	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(spzsJSON)
}

func spz_exists(writer http.ResponseWriter, request *http.Request) {
	exists := db.SPZexists(request.URL.Query().Get("SPZ"))

	existsJSON, err := json.Marshal(exists)
	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(existsJSON)
}

func get_specific_failures_state(writer http.ResponseWriter, request *http.Request) {
	state, _ := strconv.Atoi(request.URL.Query().Get("state"))
	failures := db.GetFailuresForSpecificSPZWithSpecificState(request.URL.Query().Get("SPZ"), state)

	failuresJSON, err := json.Marshal(failures)
	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(failuresJSON)
}

func get_technicians(writer http.ResponseWriter, request *http.Request) {
	technicians := db.GetTechnicians()

	techniciansJSON, err := json.Marshal(technicians)
	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(techniciansJSON)

}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", index)
	r.HandleFunc("/admin", static_site("admin", typedef.AdminPerm))
	r.HandleFunc("/demo", demo)
	r.HandleFunc("/login", login)
	r.HandleFunc("/logout", logout)
	r.HandleFunc("/usrmngmt", user_management)
	r.HandleFunc("/remove", remove)
	r.HandleFunc("/update-perms", update_perms)
	r.HandleFunc("/failures", failures)
	r.HandleFunc("/doc", doc)
	r.HandleFunc("/technical-records", technical_records)
	r.HandleFunc("/get-spzs", get_spzs)
	r.HandleFunc("/lines", lines)
	r.HandleFunc("/line_stops", line_stops)
	r.HandleFunc("/get-stops", get_stops)
	r.HandleFunc("/get-specific-failures-state", get_specific_failures_state)
	r.HandleFunc("/spz-exists", spz_exists)
	r.HandleFunc("/get-technicians", get_technicians)

	db.InitDB()

	http.ListenAndServe(":53714", r)
}
