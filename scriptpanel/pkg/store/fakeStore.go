package store

import "ScriptPanel/scriptpanel/pkg/objects"

// FakeStore implements a store service only for unit testing purposes
type FakeStore struct {
	Scripts objects.ScriptCollection
}

//FindUserScripts returns all the scripts in the store owned by a user
func (collection *FakeStore) FindUserScripts(username, department, company string) (userScripts objects.ScriptCollection) {
	for _, script := range collection.Scripts {
		if script.Username == username && script.Department == department && script.Company == company {
			userScripts = append(userScripts, script)
		}
	}

	return userScripts
}

//InsertScript puts a new script in the store
func (collection *FakeStore) InsertScript(script objects.Script) (mongoInsertError error) {

	return nil
}
