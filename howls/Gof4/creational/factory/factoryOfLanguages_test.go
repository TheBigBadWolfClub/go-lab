package factory

import (
	"reflect"
	"testing"
)

func Test_getLanguage(t *testing.T) {
	type args struct {
		langType string
	}
	tests := []struct {
		name     string
		langType LangType
		want     LanguageEr
		wantErr  bool
	}{
		{
			name:     "factory of go",
			langType: goLang,
			want:     newGo(),
			wantErr:  false,
		}, {
			name:     "factory of rust",
			langType: rustLang,
			want:     newRust(),
			wantErr:  false,
		}, {
			name:     "factory of unknown",
			langType: "unknown",
			want:     nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := factoryOfLanguages(tt.langType)
			if (err != nil) != tt.wantErr {
				t.Errorf("factoryOfLanguages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("factoryOfLanguages() got = %v, want %v", got, tt.want)
			}
		})
	}
}
