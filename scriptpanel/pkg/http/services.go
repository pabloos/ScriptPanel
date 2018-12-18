package http

import (
	"ScriptPanel/scriptpanel/pkg/objects"
	"mime/multipart"
)

// AuthService represents a authentication server, as it can be:
// - a openLDAP server
type AuthService interface {
	Login(username, department, company, password string) bool
	Signup(username, department, company, password string) bool
}

// StoreService represents a type of a bunch of data (maybe a DB or just an json array)
type StoreService interface {
	FindUserScripts(username, department, company string) (userScripts objects.ScriptCollection)
	InsertScript(script objects.Script) (mongoInsertError error)
}

// FileService represents a file service that could be:
// - a FTP server
type FileService interface {
	Upload(objects.Script, multipart.File) error
}

// ExecEnv represents an execution enviroment, as it can be:
// - the local host (via unix sockets)
// - a DinD container (via tcp)
// - a remote host (via tcp)
type ExecEnv interface {
	RunScript(script objects.RunRequest) string
}
