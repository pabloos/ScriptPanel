package mongodb

import (
	"ScriptPanel/scriptpanel/pkg/objects"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoStore implements a Store service of MongoDB
type MongoStore struct {
	Session *mgo.Session
	Store   *mgo.Collection
}

// NewMongoStore returns a new (and configured) MongoDB Store object
func NewMongoStore() (store *MongoStore) {
	// session, conexionError := mgo.Dial(mongoDir)
	// if conexionError != nil {
	// 	fmt.Fprintln(os.Stderr, "Error getting a conecction: ", conexionError)
	// 	return nil
	// }

	caCert, err := ioutil.ReadFile("../../mongocerts/minica.pem")
	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()

	ok := caCertPool.AppendCertsFromPEM(caCert)

	if !ok {
		log.Fatal("no pilla el CA de mongo")
	}

	cert, err := tls.LoadX509KeyPair("../../mongocerts/cert.pem", "../../mongocerts/key.pem")
	if err != nil {
		log.Fatal(err)
	}

	dialInfo, err := mgo.ParseURL("mongo.scriptpanel.com") //https://medium.com/@slavabobik/how-to-create-ssl-connection-in-mgo-mongo-driver-ba35af45f870
	if err != nil {
		log.Println(err)
	}

	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		tlsConfig := &tls.Config{
			RootCAs:      caCertPool,
			Certificates: []tls.Certificate{cert},
		}

		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		if err != nil {
			log.Fatal(err)
		}

		return conn, err
	}

	session, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		log.Fatal("no pilla sesion")
	}

	return &MongoStore{
		Session: session,
		Store:   session.DB("scriptPanel").C("scripts"),
	}
}

// GetCredential is the method which returns the credentials in order to make a connection
func getCredential() *mgo.Credential {
	return &mgo.Credential{
		Username: os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
		Password: os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
	}
}

// FindUserScripts returns every script in MongoDB owned by a user
func (mongoStore *MongoStore) FindUserScripts(username, department, company string) (userScripts objects.ScriptCollection) {
	mongoConn := mongoStore.Session.Copy()

	if mongoLoginError := mongoStore.Session.Login(getCredential()); mongoLoginError != nil {
		fmt.Fprintf(os.Stderr, "ERROR HACIENDO LOGIN EN MONGO: %v", mongoLoginError)
	}
	defer mongoConn.Close()

	searchError := mongoStore.Store.Find(bson.M{"username": username, "department": department, "company": company}).All(&userScripts)

	if searchError != nil {
		fmt.Fprintf(os.Stderr, "Not able to find user scripts: %v", searchError)
	}
	return userScripts
}

// InsertScript stores a script object in the db
func (mongoStore *MongoStore) InsertScript(script objects.Script) (mongoInsertError error) {
	mongoConn := mongoStore.Session.Copy()

	if mongoLoginError := mongoConn.Login(getCredential()); mongoLoginError != nil {
		fmt.Fprintf(os.Stderr, "ERROR HACIENDO LOGIN EN MONGO: %v", mongoLoginError)
	}

	defer mongoConn.Close()

	return mongoStore.Store.Insert(&script)
}
