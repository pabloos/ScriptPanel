package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name           string `json:"name"`
	Company        string `json:"company"`
	Department     string `json:"department"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Authenticathed bool
}

type SignupRequest struct {
	Company    string `json:"company"`
	Department string `json:"department"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

type signupResponse struct {
	OK bool `json:"OK"`
}

type LoginRequest struct {
	Userdata User `json:"user"`
}

type LoginResponse struct {
	Authorized bool `json:"authorized"`
}

// Login is the http route interface to do a user login
func Login(auth AuthService) http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		var request LoginRequest

		err := json.NewDecoder(req.Body).Decode(&request.Userdata)

		if err != nil {
			http.Error(response, err.Error(), 400)
			log.Println(err)
		}

		res := LoginResponse{
			Authorized: auth.Login(request.Userdata.Name, request.Userdata.Department, request.Userdata.Company, request.Userdata.Password),
		}

		json, err := json.Marshal(res)

		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			fmt.Println(err)

			return
		}

		response.Write(json)
	}
}

func Signup(auth AuthService) http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()

		var signRequest SignupRequest

		err := json.NewDecoder(req.Body).Decode(&signRequest)

		if err != nil {
			http.Error(response, err.Error(), 400)
			log.Println(err)
		}

		ok := auth.Signup(signRequest.Username, signRequest.Department, signRequest.Company, signRequest.Password)

		if !ok {
			return
		}

		json, err := json.Marshal(&signupResponse{
			true,
		})

		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		response.WriteHeader(http.StatusOK)
		response.Header().Set("Content-Type", "application/json")
		response.Write(json)
	}
}
