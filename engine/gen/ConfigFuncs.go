package gen

// HOFSTADTER_START import
// HOFSTADTER_END   import

/*
Name:      Config
About:
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

func NewConfig() *Config {
	return &Config{
		NewConfigs:      map[string]TemplateConfig{},
		StaticFiles:     []StaticFilesConfig{},
		TemplateConfigs: []TemplateConfig{},
	}
}

// HOFSTADTER_BELOW

func (C *Config) Merge(M *Config) {

	// Merge Template Configs
	for i, ElemC := range C.TemplateConfigs {
		for _, ElemM := range M.TemplateConfigs {
			if ElemC.Name == ElemM.Name {
				ElemC.Templates = append(ElemC.Templates, ElemM.Templates...)
				ElemC.StaticFiles = append(ElemC.StaticFiles, ElemM.StaticFiles...)
			}
		}
		C.TemplateConfigs[i] = ElemC
	}
}
