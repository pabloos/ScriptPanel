package script

type Script struct {
	Name        string
	Path        string
	Shell       string //python, perl, bash, *sh...
	Description string
	FlagsB      bool
	Flags       []string
	InputB      bool
	Input       []string //maybe an interface
	OutputTextB bool
	OutputText  string
	OutputFileB bool
	OutputFile  []byte //maybe an interface
}
