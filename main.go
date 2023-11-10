package main

import (
	"IIS/auth"
	"IIS/db"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)
import _ "github.com/go-sql-driver/mysql"

func login(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Attempting login from user %s.", request.Form.Get("username"))

}

// Creates a function for serving a static file. Accepts required permission for displaying the site, -1 bypasses the check
func static_site(template_name string, perm auth.Permission) func(writer http.ResponseWriter, request *http.Request) {
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

	r.HandleFunc("/", static_site("login", auth.UnprotectedPerm))
	r.HandleFunc("/admin", static_site("admin", auth.AdminPerm))
	r.HandleFunc("/login", login)

	db.InitDB()

	http.ListenAndServe(":53714", r)
}
