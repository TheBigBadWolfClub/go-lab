package deck

import "strings"

const sepString = ":"
const STATIC = ""

type CardStr string
type SuitStr string

func (s SuitStr) metadata() (SuitMeta, bool) {
	for _, d := range SuitMeta.all(STATIC) {
		if strings.Contains(string(d), string(s)) {
			return d, true
		}
	}
	return "", false
}

func (c CardStr) metadata() (CardMeta, bool) {
	for _, d := range CardMeta.all(STATIC) {
		if strings.Contains(string(d), string(c)) {
			return d, true
		}
	}
	return "", false
}

func (s SuitStr) Valid() bool {
	for _, v := range SuitMeta.metaLake(STATIC) {
		if v == s {
			return true
		}
	}
	return false
}

func (c CardStr) Valid() bool {
	for _, v := range CardMeta.metaLake(STATIC) {
		if v == c {
			return true
		}
	}
	return false
}

type SuitMeta string
type CardMeta string
type MetaID int

const (
	SPADES   SuitMeta = "1:♠:spades:black"
	HEARTS   SuitMeta = "2:♥:heart:red"
	CLUBS    SuitMeta = "3:♣:clubs:black"
	DIAMONDS SuitMeta = "4:♦:diamonds:red"
)

const (
	A   CardMeta = "1:A:ace"
	K   CardMeta = "2:K:king"
	J   CardMeta = "3:J:knight"
	Q   CardMeta = "4:Q:queen"
	C10 CardMeta = "5:10:ten"
	C9  CardMeta = "6:9:nine"
	C8  CardMeta = "7:8:eight"
	C7  CardMeta = "8:7:seven"
	C6  CardMeta = "9:6:six"
	C5  CardMeta = "10:5:five"
	C4  CardMeta = "11:4:four"
	C3  CardMeta = "12:3:three"
	C2  CardMeta = "13:2:two"
)

const (
	RANK  MetaID = 0
	CODE  MetaID = 1
	NAME  MetaID = 2
	COLOR MetaID = 3
)

func (SuitMeta) all() []SuitMeta {
	return []SuitMeta{SPADES, HEARTS, CLUBS, DIAMONDS}
}

func (CardMeta) all() []CardMeta {
	return []CardMeta{A, K, J, Q, C10, C9, C8, C7, C6, C5, C4, C3, C2}
}

func (s SuitMeta) mapped() map[MetaID]SuitStr {
	values := strings.Split(string(s), sepString)
	m := map[MetaID]SuitStr{}
	m[RANK] = SuitStr(values[RANK])
	m[CODE] = SuitStr(values[CODE])
	m[NAME] = SuitStr(values[NAME])
	m[COLOR] = SuitStr(values[COLOR])
	return m
}

func (c CardMeta) mapped() map[MetaID]CardStr {
	values := strings.Split(string(c), sepString)
	m := map[MetaID]CardStr{}
	m[RANK] = CardStr(values[RANK])
	m[CODE] = CardStr(values[CODE])
	m[NAME] = CardStr(values[NAME])
	return m
}

func (SuitMeta) metaLake() []SuitStr {
	var strs []SuitStr
	for _, v := range SuitMeta.all(STATIC) {
		for _, str := range strings.Split(string(v), sepString) {
			strs = append(strs, SuitStr(str))
		}
	}
	return strs
}

func (CardMeta) metaLake() []CardStr {
	var strs []CardStr
	for _, v := range CardMeta.all(STATIC) {
		for _, str := range strings.Split(string(v), sepString) {
			strs = append(strs, CardStr(str))
		}
	}
	return strs
}

func (SuitMeta) get(value SuitStr) (SuitMeta, bool) {
	for _, v := range SuitMeta.all(STATIC) {
		if strings.Contains(string(v), string(value)) {
			return v, true
		}
	}
	return "", false
}

func (CardMeta) get(value string) (CardMeta, bool) {
	for _, v := range CardMeta.all(STATIC) {
		if strings.Contains(string(v), value) {
			return v, true
		}
	}
	return "", false
}
