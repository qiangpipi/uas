package functions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"uas/db"
	. "uas/logs"
)

var Ip string
var Port string

type Conf struct {
	ServiceIp   string
	ServicePort string
	BindDn      string
	BindPwd     string
	BaseDn      string
	GroupSuffix string
	UserSuffix  string
	Debug       bool
}

func readFromFile(filename string, foldername string) (buf []byte) {
	//	uasroot := os.Getenv("PWD")
	uasroot, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fullpath := uasroot + "/" + foldername + "/" + filename
	buf, err := ioutil.ReadFile(fullpath)
	if err != nil {
		Error("File read failed:", err)
	}
	return buf
}

func readHomepage(fn string) []byte {
	var fd []byte
	file, err := os.Open(fn)
	defer file.Close()
	if err != nil {
		Error("function.Homepage: fail to open main page")
		fd = []byte("Fail to open main page")
	} else {
		fd, _ = ioutil.ReadAll(file)
	}
	return fd
}

func ReadConf() (res string) {
	var c Conf
	res = "OK"
	cbuf := readFromFile("uas.conf", "conf")
	if err := json.Unmarshal(cbuf, &c); err != nil {
		Error("function.ReadConf:", err)
	}
	//Ip address validation
	Ip = c.ServiceIp
	//Port validation
	Port = c.ServicePort
	db.BindDn = c.BindDn
	db.BindPwd = c.BindPwd
	db.BaseDn = c.BaseDn
	db.GroupSuffix = c.GroupSuffix
	db.UserSuffix = c.UserSuffix
	D = c.Debug
	return res
}

func CheckDb() (res string) {
	res = "OK"
	return res
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	//Security check
	//Reponse the home page
	fd := readHomepage("data/main.html")
	w.Write(fd)
}
func UserAdd(w http.ResponseWriter, r *http.Request) {
	//useradd?username=<>
	//Parse username
	username := ""
	if err := r.ParseForm(); err != nil {
	}
	for k, v := range r.Form {
		Debug("functions.UserAdd: ", k, " = ", v)
		if k == "username" {
			username = strings.Join(v, "")
		}
	}
	if username == "" || len(r.Form) > 1 {
		//Response error paras
		Debug("functions.UserAdd: username = ", username)
		Error("Bad para")
	} else {
		//Add user in database
		Info("Adding user in database")
		if err := db.UserAdd(username); err != "OK" {
		} else {
			//Response to client
			w.Write([]byte("OK"))
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
	rule := regexp.MustCompile(`[\x00-\x2F\x3A-\x40\x5B-\x60\x7B-\x7F]`)
	fd := readHomepage("data/main.html")
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
	w.Write(fd)
	if len(r.Form) != 4 {
		Debug(len(r.Form))
		w.Write([]byte("<font color=\"red\">Bad para and please try again</font>"))
	} else if username == "" || rule.MatchString(username) {
		w.Write([]byte("<font color=\"red\">Username invalid and please try againi. Rule:[0-9A-Za-z]</font>"))
	} else if password == "" || rule.MatchString(password) {
		w.Write([]byte("<font color=\"red\">Password invalid and please try again. Rule:[0-9A-Za-z]</font>"))
	} else if newpass == "" || rule.MatchString(newpass) {
		w.Write([]byte("<font color=\"red\">Newpass invalid and please try again. Rule:[0-9A-Za-z]</font>"))
	} else if newpass != newpassagain || rule.MatchString(newpassagain) {
		w.Write([]byte("<font color=\"red\">Newpass not equal newpassagain and please try again</font>"))
	} else {
		//Change password in database
		Debug("Changing password in database")
		passcheck := db.PasswordCheck(username, password)
		if passcheck == "OK" {
			changecheck := db.Passwd(username, newpass)
			if changecheck == "OK" {
				w.Write([]byte("<font color=\"green\">Password changed for <strong>" + username + "</strong></font>"))
			} else {
				w.Write([]byte("<font color=\"red\">" + changecheck + "</strong></font>"))
			}
		} else if passcheck == "Current password check failed" {
			w.Write([]byte("<font color=\"red\">" + passcheck + " <strong>" + username + "</strong></font>"))
		} else {
			w.Write([]byte("<font color=\"red\">Unknown error!</strong></font>"))
		}
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
