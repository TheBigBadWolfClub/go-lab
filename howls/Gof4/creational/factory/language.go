package factory

type language struct {
	name     string
	langType LangType
}

func (l *language) setName(name string) {
	l.name = name
}

func (l *language) setType(langType LangType) {
	l.langType = langType
}

func (l *language) getName() string {
	return l.name
}

func (l *language) getType() LangType {
	return l.langType
}
