package main

import (
	"ScriptPanel/scriptpanel/pkg/circuitbreaker"
	"ScriptPanel/scriptpanel/pkg/http"
	"log"
)

func main() {
	circuitbreaker.CheckServers(circuitbreaker.Servers{
		{
			"mongodb",
			// "10.10.10.4",
			"mongo.scriptpanel.com",
			27017,
		},
		{
			"openLDAP",
			// "10.10.10.3",
			"ldap.scriptpanel.com",
			389,
		},
		{
			"FTP",
			// "10.10.10.6",
			"ftp.scriptpanel.com",
			21,
		},
		{
			"DinD (Docker in Docker)",
			// "10.10.10.12",
			"dind.scriptpanel.com",
			4444,
		},
	})

	log.Fatal(http.NewHTTPServer().ListenAndServe())
}
