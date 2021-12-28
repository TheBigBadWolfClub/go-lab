package factory

type golang struct {
	language
}

func newGo() LanguageEr {
	return &golang{
		language{
			name:     "GoLang",
			langType: "Multi-paradigm",
		},
	}
}
