package project

type Project struct {
	Config     string
	Design     Design
	Generators []Generator
	Pipelines  []Pipeline
}
