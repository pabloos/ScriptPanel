package http

import (
	ldap "ScriptPanel/scriptpanel/pkg/auth/ldap"
	"ScriptPanel/scriptpanel/pkg/ftp"
	dind "ScriptPanel/scriptpanel/pkg/runner/dind"
	"ScriptPanel/scriptpanel/pkg/store/mongodb"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	publicRoute = "../public"
	protocol    = ":http" //YOU ARE USING HTTP NOT HTTPS
)

func NewHTTPServer() *http.Server {
	router := mux.NewRouter()

	// service dependencies

	// TODO: make async
	ftp := ftp.NewFTPServer()
	ldap := ldap.NewLDAPServer()
	dind := dind.NewDinDServer()
	mongo := mongodb.NewMongoStore()

	router.HandleFunc("/login", Login(ldap)).Methods("POST")
	router.HandleFunc("/signup", Signup(ldap)).Methods("POST")

	router.HandleFunc("/upload", Upload(mongo, ftp)).Methods("POST")
	router.HandleFunc("/download", Download(mongo, ftp)).Methods("POST")

	router.HandleFunc("/getUserScripts", GetUserScripts(mongo)).Methods("POST")

	router.HandleFunc("/runScript", RunScript(dind)).Methods("POST")

	/* TODO:
	- download script
	- CRUD script
	- CRUD user
	*/

	router.PathPrefix("/").Handler(http.FileServer(http.Dir(publicRoute))).Methods("GET") //serves the index.html

	return &http.Server{
		Handler: router,
		Addr:    protocol,

		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
}
