package riddles

import (
	"errors"
	"strings"
	"testing"
)

func Test_PasswordGenerator(t *testing.T) {
	t.Parallel()

	countFn := func(password, chars string) int {
		var count int
		for _, c := range chars {
			count += strings.Count(password, string(c))
		}
		return count
	}

	tests := []struct {
		name    string
		length  int
		minDig  int
		minSpec int
		wantErr error
	}{
		{
			name:    "should generate password with length=10, minDigits=2, minSpecialChars=1",
			length:  10,
			minDig:  2,
			minSpec: 1,
		}, {
			name:    "should generate password with length=7, minDigits=0, minSpecialChars=7",
			length:  7,
			minDig:  0,
			minSpec: 7,
		}, {
			name:    "should generate password with length=6, minDigits=6, minSpecialChars=0",
			length:  6,
			minDig:  6,
			minSpec: 0,
		}, {
			name:    "should generate password with length=5, minDigits=0, minSpecialChars=0",
			length:  5,
			minDig:  0,
			minSpec: 0,
		}, {
			name:    "should fail generate password: length cannot be zero",
			length:  0,
			minDig:  0,
			minSpec: 0,
			wantErr: ErrPasswordLength,
		}, {
			name:    "should fail generate password: length cannot less than sum of minDigits + minSpecialChars",
			length:  4,
			minDig:  3,
			minSpec: 2,
			wantErr: ErrPasswordLengthLessThanSum,
		}, {
			name:    "should fail generate password: minDigits cannot be lees than zero",
			length:  4,
			minDig:  -1,
			minSpec: 2,
			wantErr: ErrPasswordInvalidMinDigits,
		}, {
			name:    "should fail generate password: minSpecialChars cannot be lees than zero",
			length:  4,
			minDig:  2,
			minSpec: -1,
			wantErr: ErrPasswordInvalidMinSpecialChars,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := PasswordGenerator(tt.length, tt.minDig, tt.minSpec)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("PasswordGenerator() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr == nil && len(got) != tt.length {
				t.Errorf("PasswordGenerator() length expected=%d, got=%d", tt.length, len(got))
			}

			if total := countFn(got, digits); tt.wantErr == nil && total < tt.minDig {
				t.Errorf("PasswordGenerator() digit expected=%d, got=%d", tt.minDig, total)
			}

			if total := countFn(got, special); tt.wantErr == nil && total < tt.minSpec {
				t.Errorf("PasswordGenerator() chars expected=%d, got=%d", tt.minSpec, total)
			}
		})
	}
}
