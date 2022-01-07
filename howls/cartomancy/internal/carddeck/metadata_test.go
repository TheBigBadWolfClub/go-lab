package carddeck

import (
	"reflect"
	"strings"
	"testing"
)

func TestCardMeta_all(t *testing.T) {
	want := 13
	got := len(CardMeta.all(STATIC))
	if got != want {
		t.Errorf("CardMeta.all got Len() = %v, want len() = %v", got, want)
	}
}

func TestSuitMeta_all(t *testing.T) {
	want := 4
	got := len(SuitMeta.all(STATIC))
	if got != want {
		t.Errorf("SuitMeta.all got Len() = %v, want len() = %v", got, want)
	}
}

func TestCardMeta_mapped(t *testing.T) {
	values := strings.Split(string(A), sepString)
	tests := []struct {
		name string
		c    CardMeta
		want map[MetaID]CardStr
	}{
		{
			name: "Test for Ace",
			c:    A,
			want: map[MetaID]CardStr{
				RANK: CardStr(values[RANK]),
				CODE: CardStr(values[CODE]),
				NAME: CardStr(values[NAME]),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.mapped(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapped() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardMeta_metaLake(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "Test for the presence of some values",
			want: append(strings.Split(string(A), sepString),
				append(strings.Split(string(K), sepString),
					append(strings.Split(string(J), sepString), strings.Split(string(Q), sepString)...)...)...),
		},

		{
			name: "Test for the presence of some values",
			want: append(strings.Split(string(C2), sepString),
				append(strings.Split(string(C3), sepString),
					append(strings.Split(string(C4), sepString), strings.Split(string(C5), sepString)...)...)...),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			lake := CardMeta.metaLake(STATIC)
			for _, has := range tt.want {
				for _, got := range lake {
					if has == string(got) {
						return
					}
				}
				t.Errorf("metaLake() not found, want %v", has)
			}
		})
	}
}

func TestCardStr_Valid(t *testing.T) {
	tests := []struct {
		name string
		c    CardStr
		want bool
	}{
		{
			name: "Test for the validity of string ACE",
			c:    A.mapped()[NAME],
			want: true,
		},
		{
			name: "Test for the NOT found of garbage string ",
			c:    "carbage",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Valid(); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardStr_metadata(t *testing.T) {
	tests := []struct {
		name string
		c    CardStr
		want CardMeta
		ok   bool
	}{
		{
			name: "given CardStr as part of metadata, return full metadata ",
			c:    A.mapped()[NAME],
			want: A,
			ok:   true,
		},
		{
			name: "given FAKE CardStr as part of metadata, DO NOT return full metadata ",
			c:    "FAKE",
			want: "",
			ok:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.c.metadata()
			if got != tt.want {
				t.Errorf("metadata() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.ok {
				t.Errorf("metadata() got1 = %v, want %v", got1, tt.ok)
			}
		})
	}
}

func TestSuitMeta_mapped(t *testing.T) {
	values := strings.Split(string(SPADES), sepString)
	tests := []struct {
		name string
		s    SuitMeta
		want map[MetaID]SuitStr
	}{
		{
			name: "Test for Spades",
			s:    SPADES,
			want: map[MetaID]SuitStr{
				RANK:  SuitStr(values[RANK]),
				CODE:  SuitStr(values[CODE]),
				NAME:  SuitStr(values[NAME]),
				COLOR: SuitStr(values[COLOR]),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.mapped(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapped() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSuitMeta_metaLake(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "Test for the presence of some values",
			want: append(strings.Split(string(SPADES), sepString),
				append(strings.Split(string(DIAMONDS), sepString),
					append(strings.Split(string(HEARTS), sepString), strings.Split(string(CLUBS), sepString)...)...)...),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			lake := SuitMeta.metaLake(STATIC)
			for _, has := range tt.want {
				for _, got := range lake {
					if has == string(got) {
						return
					}
				}
				t.Errorf("metaLake() not found, want %v", has)
			}
		})
	}
}

func TestSuitStr_Valid(t *testing.T) {
	tests := []struct {
		name string
		s    SuitStr
		want bool
	}{
		{
			name: "Test for the validity of string ACE",
			s:    SPADES.mapped()[NAME],
			want: true,
		},
		{
			name: "Test for the NOT found of garbage string ",
			s:    "carbage",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Valid(); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSuitStr_metadata(t *testing.T) {
	tests := []struct {
		name string
		s    SuitStr
		want SuitMeta
		ok   bool
	}{
		{
			name: "given CardStr as part of metadata, return full metadata ",
			s:    SPADES.mapped()[NAME],
			want: SPADES,
			ok:   true,
		},
		{
			name: "given FAKE CardStr as part of metadata, DO NOT return full metadata ",
			s:    "FAKE",
			want: "",
			ok:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.metadata()
			if got != tt.want {
				t.Errorf("metadata() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.ok {
				t.Errorf("metadata() got1 = %v, want %v", got1, tt.ok)
			}
		})
	}
}
