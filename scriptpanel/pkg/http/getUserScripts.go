package http

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetUserScripts(store StoreService) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		userInfo := User{}

		decodeError := json.NewDecoder(request.Body).Decode(&userInfo)
		if decodeError != nil {
			log.Println(decodeError)
		}

		userScripts := store.FindUserScripts(userInfo.Name, userInfo.Department, userInfo.Company)

		if len(userScripts) == 0 && userScripts != nil {
			log.Println("the user has no scripts yet")
		}

		res, errorJSON := json.Marshal(userScripts)

		if errorJSON != nil {
			log.Println("ERROR JSON: CAN'T CONVERT TO JSON")
		}

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		response.Write(res)
	}
}
