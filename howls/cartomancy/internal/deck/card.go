package deck

type SuitID struct {
	rank   SuitStr
	Symbol SuitStr
	Color  SuitStr
	Label  SuitStr
	SuitMeta
}

type CardID struct {
	Rank CardStr
	Name CardStr
	Code CardStr
	CardMeta
}

func (s SuitMeta) suiteID() SuitID {
	m := s.mapped()
	return SuitID{
		SuitMeta: s,
		rank:     SuitStr(m[RANK]),
		Symbol:   SuitStr(m[CODE]),
		Label:    SuitStr(m[NAME]),
		Color:    SuitStr(m[COLOR]),
	}
}

func (c CardMeta) cardID() CardID {
	m := c.mapped()
	return CardID{
		CardMeta: c,
		Rank:     CardStr(m[RANK]),
		Code:     CardStr(m[CODE]),
		Name:     CardStr(m[NAME]),
	}
}

type Card struct {
	CardID
	SuitID
}

func NewCard(c CardID, s SuitID) Card {
	return Card{
		CardID: c,
		SuitID: s,
	}
}

func (SuitID) List() []SuitID {
	var list []SuitID
	for _, s := range SuitMeta.all(STATIC) {
		list = append(list, s.suiteID())
	}
	return list
}

func (CardID) List() []CardID {
	var list []CardID
	for _, c := range CardMeta.all(STATIC) {
		list = append(list, c.cardID())
	}
	return list
}

func (c Card) Equal(o Card) bool {
	return c == o
}
