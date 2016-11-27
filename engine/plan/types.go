type FileGenData struct {
	Name     string
	Template *raymond.Template
	Data     interface{}
}



type Stage struct {
	Name  string
	Steps []Step
}

type Pipeline struct {
	Name   string
	Stages []Stage
}

