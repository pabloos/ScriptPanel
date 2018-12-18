package objects

// RunRequest represents a request to run a Script object based on his configuration (also an object)
type RunRequest struct {
	ScriptObject Script `json:"script"`
	ConfigObject Config `json:"config"`
}

// Config represents the flag and args for that could be present in a script run
type Config struct {
	Flag string   `json:"flag"`
	Args []string `json:"args"`
}
