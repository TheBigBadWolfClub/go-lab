package riddles

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

const stringsTest = "test"

func Test_sumAllNumbers(t *testing.T) {
	tests := []struct {
		str  string
		want int
	}{
		{
			str:  "hello",
			want: 0,
		}, {
			str:  "12345",
			want: 15,
		}, {
			str:  "1aBc5",
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(stringsTest, func(t *testing.T) {
			if got := sumAllNumbers(tt.str); got != tt.want {
				t.Errorf("sumAllNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxNumeric(t *testing.T) {
	tests := []struct {
		str  string
		want int
	}{
		{
			str:  "hello",
			want: 0,
		}, {
			str:  "12345",
			want: 5,
		}, {
			str:  "1aBc5V",
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(stringsTest, func(t *testing.T) {
			if got := maxNumeric(tt.str); got != tt.want {
				t.Errorf("sumAllNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_duplicatedChars(t *testing.T) {
	tests := []struct {
		str  string
		want []string
	}{
		{
			str:  "edecba",
			want: []string{"e"},
		}, {
			str:  "abccdeff",
			want: []string{"c", "f"},
		},
	}
	for _, tt := range tests {
		t.Run(stringsTest, func(t *testing.T) {
			got := duplicatedChars(tt.str)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("duplicatedChars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_areAnagrams(t *testing.T) {
	tests := []struct {
		str1 string
		str2 string
		want bool
	}{
		{
			str1: "",
			str2: "",
			want: true,
		}, {
			str1: "hello",
			str2: "holle",
			want: true,
		}, {
			str1: "army",
			str2: "mary",
			want: true,
		}, {
			str1: "army",
			str2: "mari",
			want: false,
		}, {
			str1: "army",
			str2: "irma",
			want: false,
		}, {
			str1: "aaa",
			str2: "aaa",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(stringsTest, func(t *testing.T) {
			if got := areAnagrams(tt.str1, tt.str2); got != tt.want {
				t.Errorf("areAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_allPermutations(t *testing.T) {
	tests := []struct {
		str  string
		want []string
	}{
		{
			str:  "a",
			want: []string{"a"},
		}, {
			str:  "ab",
			want: []string{"ab", "ba"},
		}, {
			str:  "abc",
			want: []string{"abc", "acb", "bac", "bca", "cba", "cab"},
		},
	}
	for _, tt := range tests {
		t.Run(stringsTest, func(t *testing.T) {
			got := allPermutations(tt.str, 0)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("allPermutations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseRecursion(t *testing.T) {
	tests := []struct {
		str  string
		want string
	}{
		{
			str:  "hello",
			want: "olleh",
		}, {
			str:  "software",
			want: "erawtfos",
		},
	}
	for _, tt := range tests {
		t.Run(stringsTest, func(t *testing.T) {
			got := reverseRecursion(tt.str)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("reverseRecursion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countCharOccurrences(t *testing.T) {
	tests := []struct {
		str  string
		char uint8
		want int
	}{
		{
			str:  "hello",
			char: 'l',
			want: 2,
		}, {
			str:  "12345",
			char: '1',
			want: 1,
		}, {
			str:  "1aBBBc5V",
			char: 'B',
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(stringsTest, func(t *testing.T) {
			if got := countCharOccurrences(tt.str, tt.char); got != tt.want {
				t.Errorf("countCharOccurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstNonRepeatable(t *testing.T) {
	tests := []struct {
		str  string
		want int32
	}{
		{
			str:  "hello",
			want: 'h',
		}, {
			str:  "12345",
			want: '1',
		}, {
			str:  "BBBc5V",
			want: 'c',
		},
	}
	for _, tt := range tests {
		t.Run(stringsTest, func(t *testing.T) {
			if got := firstNonRepeatable(tt.str); got != tt.want {
				t.Errorf("countCharOccurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseWords(t *testing.T) {
	tests := []struct {
		str  string
		want string
	}{
		{
			str:  "hello world",
			want: "world hello",
		}, {
			str:  "good morning nice people",
			want: "people nice morning good",
		}, {
			str:  "good morning nice people again",
			want: "again people nice morning good",
		},
	}
	for _, tt := range tests {
		t.Run(stringsTest, func(t *testing.T) {
			got := reverseWords(tt.str)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_areRotation(t *testing.T) {
	tests := []struct {
		strA string
		strB string
		want bool
	}{
		{
			strA: "programing",
			strB: "ingprogram",
			want: true,
		}, {
			strA: "programing",
			strB: "ingprograX",
			want: false,
		}, {
			strA: "programing",
			strB: "ingprogra",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(stringsTest, func(t *testing.T) {
			got := areRotation(tt.strA, tt.strB)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("areRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		strA string
		want bool
	}{
		{
			strA: "ana",
			want: true,
		}, {
			strA: "anana",
			want: true,
		}, {
			strA: "anaa",
			want: false,
		}, {
			strA: "abana",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(stringsTest, func(t *testing.T) {
			got := isPalindrome(tt.strA)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestPalindrome(t *testing.T) {
	tests := []struct {
		strA string
		want string
	}{
		{
			strA: "ana",
			want: "ana",
		}, {
			strA: "xptoanaxpto",
			want: "ana",
		}, {
			strA: "anaa",
			want: "ana",
		}, {
			strA: "anamamma",
			want: "amma",
		},
	}
	for _, tt := range tests {
		t.Run(stringsTest, func(t *testing.T) {
			got := longestPalindrome(tt.strA)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findLongestSubstring(t *testing.T) {
	tests := []struct {
		strA string
		want string
	}{
		{
			strA: "",
			want: "",
		}, {
			strA: "abcdeefg",
			want: "abcde",
		}, {
			strA: "abb123456bba",
			want: "b123456",
		}, {
			strA: "asdfg",
			want: "asdfg",
		},
	}
	for _, tt := range tests {
		t.Run(stringsTest, func(t *testing.T) {
			got := findLongestSubstring(tt.strA)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("findLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeDuplicate(t *testing.T) {
	tests := []struct {
		strA string
		want string
	}{
		{
			strA: "",
			want: "",
		}, {
			strA: "abccd",
			want: "abcd",
		}, {
			strA: "abbcdde",
			want: "abcde",
		}, {
			strA: "aabcc",
			want: "abc",
		}, {
			strA: "abcabc",
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(stringsTest, func(t *testing.T) {
			got := removeDuplicate(tt.strA)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("findLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMaxWordOccurring(t *testing.T) {
	tests := []struct {
		str  []string
		want string
	}{
		{
			str:  []string{"hello miss", "hello mrs", "hell no", "hello world"},
			want: "hello",
		}, {
			str:  []string{"miss", "mrs", "no", "world"},
			want: "miss",
		}, {
			str:  []string{"hello", "hello", "hello"},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(stringsTest, func(t *testing.T) {
			if got := findMaxWordOccurring(tt.str); got != tt.want {
				t.Errorf("countCharOccurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}
