package ldap

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	ldapLibrary "gopkg.in/ldap.v2"
)

const (
	ldapHost = "ldap.scriptpanel.com"
	// ldapHost = "10.10.10.3"
	// port     = 389
	port = 636
)

type LDAP struct {
	*ldapLibrary.Conn
}

func NewLDAPServer() *LDAP {
	// connection, ldapConnectionError := ldapLibrary.Dial("tcp", fmt.Sprintf("%s:%d", ldapHost, port))

	// cert, err := tls.LoadX509KeyPair("../../certs/cert.pem", "../../certs/key.pem")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	caCert, err := ioutil.ReadFile("../../certs/minica.pem")
	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()

	ok := caCertPool.AppendCertsFromPEM(caCert)

	if !ok {
		log.Fatal("no pilla el CA")
	}

	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
		// Certificates: []tls.Certificate{cert},
		ServerName: "ldap.scriptpanel.com",
	}

	connection, ldapConnectionError := ldapLibrary.DialTLS("tcp", fmt.Sprintf("%s:%d", ldapHost, port), tlsConfig)

	if ldapConnectionError != nil {
		fmt.Fprintf(os.Stderr, "Error getting a LDAP connection: %v", ldapConnectionError)
	}
	defer log.Println("establecida la conexión TLS con ldap")

	defer connection.Bind("cn=admin,dc=scriptpanel,dc=com", "admin")

	return &LDAP{
		connection,
	}
}

func (ldap *LDAP) Login(username, department, company, password string) bool {
	rdn := fmt.Sprintf("uid=%s+userPassword=%s,ou=%s,ou=%s,dc=scriptpanel,dc=com", username, password, department, company)

	bindError := ldap.Bind(rdn, password)

	if bindError != nil {
		log.Println(bindError)
		log.Println(fmt.Sprintf("Usuario: %s, no ha podido hacer login", username))

		return false
	}

	return true
}

func (ldap *LDAP) addCompany(company string) bool {
	ldap.Bind("cn=admin,dc=scriptpanel,dc=com", "admin")

	err := ldap.Add(&ldapLibrary.AddRequest{
		DN: fmt.Sprintf("ou=%s,dc=scriptpanel,dc=com", company),
		Attributes: []ldapLibrary.Attribute{
			ldapLibrary.Attribute{
				Type: "objectClass",
				Vals: []string{"organizationalUnit"},
			},
			ldapLibrary.Attribute{
				Type: "ou",
				Vals: []string{company},
			},
		},
	})

	if err != nil {
		log.Println("addCompany")
		log.Println(err)

		return false
	}

	return true
}

func (ldap *LDAP) addDepartment(company, department string) bool {
	ldap.Bind("cn=admin,dc=scriptpanel,dc=com", "admin")

	err := ldap.Add(&ldapLibrary.AddRequest{
		DN: fmt.Sprintf("ou=%s,ou=%s,dc=scriptpanel,dc=com", department, company),

		Attributes: []ldapLibrary.Attribute{
			ldapLibrary.Attribute{
				Type: "objectClass",
				Vals: []string{"organizationalUnit"},
			},
			ldapLibrary.Attribute{
				Type: "ou",
				Vals: []string{department},
			},
		},
	})

	if err != nil {
		log.Println("addDepartment")
		log.Println(err)

		return false
	}

	return true
}

func (ldap *LDAP) addUser(company, department, user, pass string) bool {
	ldap.Bind("cn=admin,dc=scriptpanel,dc=com", "admin")

	err := ldap.Add(&ldapLibrary.AddRequest{
		DN: fmt.Sprintf("uid=%s+userPassword=%s,ou=%s,ou=%s,dc=scriptpanel,dc=com", user, pass, department, company),

		Attributes: []ldapLibrary.Attribute{
			ldapLibrary.Attribute{
				Type: "objectClass",
				Vals: []string{"inetOrgPerson"},
			},
			ldapLibrary.Attribute{
				Type: "sn",
				Vals: []string{user},
			},
			ldapLibrary.Attribute{
				Type: "cn",
				Vals: []string{user},
			},
		},
	})

	if err != nil {
		log.Println("addUser")
		log.Println(err)

		return false
	}

	return true
}

func (ldap *LDAP) Signup(username, department, company, password string) bool {
	checkCompany, err := ldap.Compare(fmt.Sprintf("ou=%s,dc=scriptpanel,dc=com", company), "ou", company)

	if err != nil {
		log.Println(err)
	}

	checkDepartment, err := ldap.Compare(fmt.Sprintf("ou=%s,ou=%s,dc=scriptpanel,dc=com", department, company), "ou", department)

	if err != nil {
		log.Println(err)
	}

	if !checkCompany { //if the company does not exists in the ldap db, create it ...
		ok := ldap.addCompany(company)

		if !ok {
			return false
		}
	}

	if !checkDepartment { //... the same with the department
		ok := ldap.addDepartment(company, department)

		if !ok {
			return false
		}
	}

	ok := ldap.addUser(company, department, username, password)

	if !ok {
		return false
	}

	return true
}

func (ldap *LDAP) SearchCompany(company string) {
	result, _ := ldap.Search(ldapLibrary.NewSearchRequest(
		"dc=scriptpanel,dc=com",
		ldapLibrary.ScopeWholeSubtree, ldapLibrary.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=*)(dc=%s))", company),
		[]string{"dn"},
		nil,
	))

	for _, entry := range result.Entries {
		fmt.Printf("%s: %v\n", entry.DN, entry.GetAttributeValue("cn"))
	}
}

func (ldap *LDAP) SearchDepartment(department, company string) {
	result, _ := ldap.Search(ldapLibrary.NewSearchRequest(
		fmt.Sprintf("dc=%s,dc=scriptpanel,dc=com", company),
		ldapLibrary.ScopeWholeSubtree, ldapLibrary.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=*)(dc=%s))", department),
		[]string{"dn"},
		nil,
	))

	for _, entry := range result.Entries {
		fmt.Printf("%s: %v\n", entry.DN, entry.GetAttributeValue("cn"))
	}
}

func (ldap *LDAP) SearchUser(user, department, company string) {
	result, _ := ldap.Search(ldapLibrary.NewSearchRequest(
		fmt.Sprintf("dc=%s,dc=%s,dc=scriptpanel,dc=com", department, company),
		ldapLibrary.ScopeWholeSubtree, ldapLibrary.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=*)(dc=%s))", department),
		[]string{"dn"},
		nil,
	))

	for _, entry := range result.Entries {
		fmt.Printf("%s: %v\n", entry.DN, entry.GetAttributeValue("cn"))
	}
}
