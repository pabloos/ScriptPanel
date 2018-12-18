package http

import (
	"ScriptPanel/scriptpanel/pkg/objects"
	"mime/multipart"
)

type FakeFileServer struct {
	ScriptCollection []ScriptFile
}

/* func NewFakeFileServer() *FakeFileServer {
	fakeFileServer := &FakeFileServer{
		ScriptCollection: make([]ScriptFile, 0),
	}

	return fakeFileServer
} */

func (ffs *FakeFileServer) Upload(script objects.Script, file multipart.File) error {
	scriptFile := ScriptFile{
		script:  script,
		handler: file,
	}

	ffs.ScriptCollection = append(ffs.ScriptCollection, scriptFile)

	return nil
}

type ScriptFile struct {
	script  objects.Script
	handler multipart.File
}

/* func TestUploadHandler(t *testing.T) {

	request, createRequestError := http.NewRequest("POST", "/upload", bytes.NewReader(jsonBody))
	if createRequestError != nil {
		t.Error()
	}

	response := httptest.NewRecorder()

	store := &FakeStore{}
	fs := &FakeFileServer{}

	handler := http.HandlerFunc(Upload(store, fs))

	handler.ServeHTTP(response, request)

} */
