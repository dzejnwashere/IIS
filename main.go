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

	users := db.GetAllUsers()
	/*files, err := template.ParseFiles("res/tmpl/usrmngmt.html")
	if err != nil {
		fmt.Fprintf(writer, err.Error())
	}*/

	tmpl, err := template.New("usrmngmt").Parse(`
<!DOCTYPE html>
<html>
<head>
    <style>
        table {
            border-collapse: collapse;
            width: 100%;
        }

        th, td {
            text-align: left;
            padding: 8px;
        }

        tr:nth-child(even){background-color: #f2f2f2}

        th {
            background-color: #D22B2B;
            color: white;
        }
    </style>
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js"></script>
    <script type="text/javascript">
        function deleteUser(userID) {
            // Add logic to handle delete action for the user with the specified ID
            if (userID === 1) {
                alert("You can't delete admin.");
            } else {
                alert("Delete user with ID: " + userID);

                $.ajax({
                    method: "GET",
                    url: "/remove",
                    data: { userID: userID },
                    success: function(response) {
                        alert("User with ID " + userID + " removed successfully!");
                        // Update the table with the new data
                        $('#table-management').html(response);
                    },
                    error: function() {
                        alert("An error occurred while removing the user.");
                    }
                });
            }
        }

        // ... (other functions remain the same)
    </script>
</head>
<body>
    <table id="table-management">
        <tr>
            <th>ID</th>
            <th>User name</th>
            <th>Permissions</th>
            <th>Delete</th>
            <th>Deactivate</th>
            <th>Edit</th>
        </tr>
        {{range .}}
            <tr>
                <td>{{.ID}}</td>
                <td>{{.Username}}</td>
                <td>{{.Permissions}}</td>
                <td><button class="delete-btn" onclick="deleteUser({{.ID}})">❌</button></td>
                <td><button class="deactivate-btn" onclick="deactivateUser({{.ID}})">🚫</button></td>
                <td><button class="edit-btn" onclick="editUser({{.ID}})">✏️</button></td>
            </tr>
        {{end}}
    </table>
</body>
</html>`)

	if err != nil {
		http.Error(writer, "Error rendering template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(writer, users)
	if err != nil {
		return
	}
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

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", index)
	r.HandleFunc("/admin", static_site("admin", typedef.AdminPerm))
	r.HandleFunc("/login", login)
	r.HandleFunc("/logout", logout)
	r.HandleFunc("/usrmngmt", user_management)
	r.HandleFunc("/remove", remove)

	db.InitDB()

	http.ListenAndServe(":53714", r)
}
