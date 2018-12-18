package http

import (
	"ScriptPanel/scriptpanel/pkg/objects"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

func Upload(store StoreService, fileServer FileService) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		defer request.Body.Close()

		request.ParseMultipartForm(32 << 20)

		file, handler, uploadErr := request.FormFile("uploadfile")

		if uploadErr != nil {
			fmt.Fprintln(os.Stderr, "Error subiendo el archivo: ", uploadErr)

			return
		}

		script := objects.Script{
			Description: request.Form.Get("description"),
			Company:     request.Form.Get("company"),
			Department:  request.Form.Get("department"),
			Username:    request.Form.Get("name"),
			Filename:    handler.Filename,
			Language:    filepath.Ext(handler.Filename)[1:], //extract the extension from the filename and remove the dot
		}

		formFlags := request.Form.Get("flags")
		formArgs := request.Form.Get("args")

		json.Unmarshal([]byte(formFlags), &script.Flags) //REMEMBER: THESE ARE ARRAYS
		json.Unmarshal([]byte(formArgs), &script.Args)

		var wg sync.WaitGroup

		wg.Add(2)

		go func() {
			defer wg.Done()

			err := store.InsertScript(script)

			if err != nil {
				log.Println(fmt.Sprintf("Store error: %v", err))
			}
		}()

		go func() {
			defer wg.Done()

			err := fileServer.Upload(script, file)

			if err != nil {
				log.Println(fmt.Sprintf("Store error: %v", err))
			}
		}()

		wg.Wait()

		response.WriteHeader(http.StatusOK)
	}
}
