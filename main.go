package main

import (
	"IIS/auth"
	"IIS/db"
	"IIS/typedef"
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

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", index)
	r.HandleFunc("/admin", static_site("admin", typedef.AdminPerm))
	r.HandleFunc("/demo", demo)
	r.HandleFunc("/login", login)
	r.HandleFunc("/logout", logout)
	r.HandleFunc("/usrmngmt", user_management)
	r.HandleFunc("/remove", remove)

	db.InitDB()

	http.ListenAndServe(":53714", r)
}
