package http

import (
	"ScriptPanel/scriptpanel/pkg/objects"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func RunScript(env ExecEnv) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var runRequest objects.RunRequest

		decodeError := json.NewDecoder(r.Body).Decode(&runRequest)

		if decodeError != nil {
			fmt.Fprintf(os.Stderr, "Error decoding the request: %v", decodeError)
		}

		response := env.RunScript(runRequest)

		res, errorJSON := json.Marshal(response)

		if errorJSON != nil {
			log.Println(errorJSON)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
}
