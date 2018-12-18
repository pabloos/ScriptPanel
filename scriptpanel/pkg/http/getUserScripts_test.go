package http

import (
	"ScriptPanel/scriptpanel/pkg/objects"
	"ScriptPanel/scriptpanel/pkg/store"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	testUsername   = "Pedro"
	testDepartment = "Operaciones"
	testCompany    = "Arqueosur"
)

func TestGetUserScriptsHandler(t *testing.T) {
	jsonRequest := User{
		Name:       testUsername,
		Department: testDepartment,
		Company:    testCompany,
	}

	jsonBody, jsonMarshallError := json.Marshal(jsonRequest)

	if jsonMarshallError != nil {
		t.Error()
	}

	request, createRequestError := http.NewRequest("POST", "/getUserScripts", bytes.NewReader(jsonBody))

	if createRequestError != nil {
		t.Error()
	}

	response := httptest.NewRecorder()

	store := &store.FakeStore{}

	store.Scripts = objects.ScriptCollection{
		{
			Username:   "Pedro",
			Department: "Operaciones",
			Company:    "Arqueosur",
			Filename:   "script.py",
			Language:   "Python",
		},
		{
			Username:   "Pedro",
			Department: "Operaciones",
			Company:    "Arqueosur",
			Filename:   "sc.pl",
			Language:   "Perl",
		},
		{
			Username:   "Alberto",
			Department: "Operaciones",
			Company:    "Arqueosur",
			Filename:   "cotas.sh",
			Language:   "Bash",
		},
		{
			Username:   "Pedro",
			Department: "Operaciones",
			Company:    "Arqueosur",
			Filename:   "sp.sh",
			Language:   "Bash",
		},

		{
			Username:   "Pedro",
			Department: "Ventas",
			Company:    "Arqueosur",
			Filename:   "sp.sh",
			Language:   "Bash",
		},
	}

	handler := http.HandlerFunc(GetUserScripts(store))

	handler.ServeHTTP(response, request)

	var resultScripts objects.ScriptCollection

	unmarshalError := json.Unmarshal(response.Body.Bytes(), &resultScripts)

	if unmarshalError != nil {
		t.Error()
	}

	for _, script := range resultScripts {
		if script.Username != testUsername || script.Department != testDepartment || script.Company != testCompany {
			t.Fail()
		}
	}
}
