type FileGenData struct {
	Name     string
	Template *raymond.Template
	Data     interface{}
}

type Plan struct {
	template  string
	design    interface{}
	outfile   string
	templates []string
}



type Stage struct {
	Name  string
	Steps []Step
}

type Pipeline struct {
	Name   string
	Stages []Stage
}

