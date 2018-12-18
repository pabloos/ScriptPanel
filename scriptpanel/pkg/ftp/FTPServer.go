package ftp

import (
	"ScriptPanel/scriptpanel/pkg/objects"
	"fmt"
	"log"
	"mime/multipart"
	"os"

	"github.com/jlaffaye/ftp"
)

const (
	ftpUsernameVAR = "FTP_USER"
	ftpPasswordVAR = "FTP_PASS"

	host = "ftp.scriptpanel.com"
	port = 21
	// dir = "10.10.10.6:21"
)

func getCredentials() (string, string) {
	return os.Getenv(ftpUsernameVAR), os.Getenv(ftpPasswordVAR)
}

// FTPServer implements a store service
type Server struct {
	*ftp.ServerConn
}

// NewFTPServer returns a new and configured FTP server driver
func NewFTPServer() *Server {
	conn, connectionError := ftp.Connect(fmt.Sprintf("%s:%d", host, port))

	if connectionError != nil {
		fmt.Fprintf(os.Stderr, "FTP connection error: %v", connectionError)
		return nil
	}

	err := conn.Login(getCredentials())
	if err != nil {
		log.Println(err)
	}

	return &Server{
		conn,
	}
}

// Upload sends a files tot eh ftp server
func (ftp *Server) Upload(script objects.Script, file multipart.File) error {
	err1 := ftp.MakeDir(script.Company)
	if err1 != nil {
		log.Println(err1)
	}

	err2 := ftp.MakeDir(fmt.Sprintf("%s/%s", script.Company, script.Department))
	if err2 != nil {
		log.Println(err2)
	}

	err3 := ftp.MakeDir(fmt.Sprintf("%s/%s/%s", script.Company, script.Department, script.Username))
	if err3 != nil {
		log.Println(err3)
	}

	return ftp.Stor(fmt.Sprintf("%s/%s/%s/%s", script.Company, script.Department, script.Username, script.Filename), file)
}
