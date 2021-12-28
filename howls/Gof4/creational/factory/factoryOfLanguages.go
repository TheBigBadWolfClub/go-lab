package factory

import "fmt"

type LangType string

const (
	goLang   LangType = "GO"
	rustLang LangType = "Rust"
)

func factoryOfLanguages(langType LangType) (LanguageEr, error) {
	if langType == goLang {
		return newGo(), nil
	}

	if langType == rustLang {
		return newRust(), nil
	}

	return nil, fmt.Errorf("no language")
}
