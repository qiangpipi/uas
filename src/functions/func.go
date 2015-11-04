package functions

import (
	"db"
	. "logs"
	"net/http"
	"strings"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	//Security check
	//Reponse the home page
}
func UserAdd(w http.ResponseWriter, r *http.Request) {
	//useradd?username=<>
	//Parse username
	username := ""
	if err := r.ParseForm(); err != nil {
	}
	for k, v := range r.Form {
		Debug(k, " = ", v)
		if k == "username" {
			username = strings.Join(v, "")
		}
	}
	if username == "" || len(r.Form) > 1 {
		//Response error paras
		Debug("username = ", username)
		Error("Bad para")
	} else {
		//Add user in database
		Info("Adding user in database")
		if err := db.UserAdd(username); err != nil {
		} else {
			//Response to client
		}
	}
}

func Passwd(w http.ResponseWriter, r *http.Request) {
	//passwd?username=<>&password=<>&newpass=<>&newpassagain=<>
	//Parse username=<>&password=<>&newpass=<>&newpassagain=<>
	username := ""
	password := ""
	newpass := ""
	newpassagain := ""
	if err := r.ParseForm(); err != nil {
	}
	for k, v := range r.Form {
		Debug(k, " = ", v)
		if k == "username" {
			username = strings.Join(v, "")
		}
		if k == "password" {
			password = strings.Join(v, "")
		}
		if k == "newpass" {
			newpass = strings.Join(v, "")
		}
		if k == "newpassagain" {
			newpassagain = strings.Join(v, "")
		}
	}
	if len(r.Form) != 4 || username == "" || password == "" || newpass == "" || newpassagain == "" {
		Debug("username = ", username)
		Debug("password = ", password)
		Debug("newpass = ", newpass)
		Debug("newpassagain = ", newpassagain)
		Error("Bad para")
	} else {
		//Change password in database
		Info("Changing password in database")
	}
}

func UserDel(w http.ResponseWriter, r *http.Request) {
	//userdel?username=<>
	//Parse username
}

func DataBackup(w http.ResponseWriter, r *http.Request) {
	//databackup?TBD
	//Parse TBD
}

func DataRestore(w http.ResponseWriter, r *http.Request) {
	//datarestore?TBD
	//Parse TBD
}
