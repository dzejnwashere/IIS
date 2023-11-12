package main

import (
	"IIS/auth"
	"IIS/db"
	"IIS/typedef"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)
import _ "github.com/go-sql-driver/mysql"

func login(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil || len(request.Form) == 0 {
		fmt.Fprintf(writer, "Accessing without form data.")
		return
	}
	jwtToken, err := auth.Authenticate(request.PostFormValue("username"), request.PostFormValue("password"))
	if err != nil {
		fmt.Fprintf(writer, err.Error())
		return
	}
	cookie := http.Cookie{
		Name:   "iisauth",
		Value:  jwtToken,
		MaxAge: 600,
	}
	http.SetCookie(writer, &cookie)
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
		files.Execute(writer, nil)
	}
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", static_site("login", typedef.UnprotectedPerm))
	r.HandleFunc("/admin", static_site("admin", typedef.AdminPerm))
	r.HandleFunc("/login", login)

	db.InitDB()

	http.ListenAndServe(":53714", r)
}
