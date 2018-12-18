package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type fakeAuthServer struct {
	Users []User
}

func (auth *fakeAuthServer) Signup(username, department, company, password string) bool {
	user := User{
		Name:       username,
		Company:    company,
		Department: department,
		Password:   password,
	}

	auth.Users = append(auth.Users, user)

	return true
}

func (auth *fakeAuthServer) Login(username, department, company, password string) bool {
	for _, user := range auth.Users {
		if user.Company == company && user.Department == department && user.Name == username && user.Password == password {
			return true
		}
	}
	return false
}

const (
	testUsername1   = "John"
	testDepartment1 = "Software"
	testCompany1    = "Apple"
	testPassword1   = "pass"
)

func TestLogin(t *testing.T) {
	user := User{
		Name:       testUsername1,
		Department: testDepartment1,
		Company:    testCompany1,
		Password:   testPassword1,
	}

	jsonBody, jsonMarshallError := json.Marshal(user)

	if jsonMarshallError != nil {
		t.Error("1")
	}

	request, createRequestError := http.NewRequest("POST", "/upload", bytes.NewReader(jsonBody))

	if createRequestError != nil {
		t.Error("2")
	}

	response := httptest.NewRecorder()

	auth := &fakeAuthServer{}

	auth.Users = []User{
		User{
			Name:       "John",
			Department: "Software",
			Company:    "Apple",
			Password:   "pass",
		},
	}

	handler := http.HandlerFunc(Login(auth))

	handler.ServeHTTP(response, request)

	var result User

	unmarshalError := json.Unmarshal(response.Body.Bytes(), &result)

	if unmarshalError != nil {
		t.Error("3")
	}

}
