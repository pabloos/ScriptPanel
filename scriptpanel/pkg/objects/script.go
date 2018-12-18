package objects

// Script represents the info of a scriptfile
type Script struct {
	Username    string   `json:"username" bson:"username"`
	Department  string   `json:"department" bson:"department"`
	Company     string   `json:"company" bson:"company"`
	Filename    string   `json:"filename" bson:"filename"`
	Language    string   `json:"language" bson:"language"`
	Description string   `json:"description" bson:"description"`
	Flags       []string `json:"flags" bson:"flags,omitempty"`
	Args        []string `json:"args" bson:"args,omitempty"`
}

// ScriptCollection represents a slice of Scripts objects
type ScriptCollection []Script
