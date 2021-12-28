package factory

type rust struct {
	language
}

func newRust() LanguageEr {
	return &rust{
		language{
			name:     "Rust",
			langType: "Multi-paradigm",
		},
	}
}
