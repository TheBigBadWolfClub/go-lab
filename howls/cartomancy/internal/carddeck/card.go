package carddeck

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"text/tabwriter"
)

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

func (s SuitMeta) suiteID() SuitID {
	m := s.mapped()
	return SuitID{
		SuitMeta: s,
		rank:     m[RANK],
		Symbol:   m[CODE],
		Label:    m[NAME],
		Color:    m[COLOR],
	}
}

func (c CardMeta) cardID() CardID {
	m := c.mapped()
	return CardID{
		CardMeta: c,
		Rank:     m[RANK],
		Code:     m[CODE],
		Name:     m[NAME],
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

func (Card) NumOfSuits() int {
	return len(SuitMeta.all(STATIC))
}

func (Card) NumOfCardType() int {
	return len(CardMeta.all(STATIC))
}

func (Card) NumOfCards() int {
	return Card{}.NumOfCardType() * Card{}.NumOfSuits()
}

func (c Card) Is(card CardMeta, suit SuitMeta) bool {
	return c.CardMeta == card && c.SuitMeta == suit
}

func (c Card) Equal(o Card) bool {
	return c == o
}

// Card implement fmters
func (c Card) String() string {
	return fmt.Sprintf("%s|%s", c.CardMeta, c.SuitMeta)
}

func (c Card) Format(s fmt.State, verb rune) {
	switch verb {
	case 'q':
		_, _ = fmt.Fprintf(s, "%s%s", c.Code, c.Symbol)
	case 'v':
		if s.Flag('#') {
			_, _ = fmt.Fprint(s, c.GoString())
			return
		}
		_ = json.NewEncoder(s).Encode(c)
	default:
		_, _ = fmt.Fprint(s, c.String())
	}
}

func (c Card) GoString() string {
	gs := &bytes.Buffer{}
	gs.WriteString("Card{\n")

	writer := tabwriter.NewWriter(gs, 0, 0, 1, ' ', tabwriter.TabIndent)
	_, _ = fmt.Fprintf(gs, "\tRank:\t\t%q\n", c.Rank)
	_, _ = fmt.Fprintf(gs, "\tCode:\t\t%q\n", c.Code)
	_, _ = fmt.Fprintf(gs, "\tName:\t\t%q\n", c.Name)
	_, _ = fmt.Fprintf(gs, "\trank:\t\t%q\n", c.rank)
	_, _ = fmt.Fprintf(gs, "\tSymbol:\t\t%q\n", c.Symbol)
	_, _ = fmt.Fprintf(gs, "\tLabel:\t\t%q\n", c.Label)
	_, _ = fmt.Fprintf(gs, "\tColor:\t\t%q\n", c.Color)
	_, _ = fmt.Fprintf(gs, "\tCardMeta:\t%q\n", c.SuitMeta)
	_, _ = fmt.Fprintf(gs, "\tSuitMeta:\t%q\n", c.SuitMeta)
	_ = writer.Flush()

	gs.WriteString("}")
	return gs.String()
}

func (c Card) CompareTo(c2 Card) int {
	r1, _ := strconv.Atoi(string(c.Rank))
	r2, _ := strconv.Atoi(string(c2.Rank))

	if r1 < r2 {
		return 1
	}
	if r1 > r2 {
		return -1
	}
	return 0
}
