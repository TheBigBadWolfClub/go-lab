package factory

type LanguageEr interface {
	setName(name string)
	setType(langType LangType)
	getName() string
	getType() LangType
}
