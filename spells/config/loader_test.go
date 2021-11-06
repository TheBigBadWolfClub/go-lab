package config

import (
	"reflect"
	"testing"
)

func TestLoad_FileOpen(t *testing.T) {
	type args struct {
		filename string
		data     interface{}
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "success open yaml",
			args: args{
				filename: "config.yaml",
				data:     &struct{}{},
			},
			wantErr: false,
		},
		{name: "fail to open yaml",
			args: args{
				filename: "not_exist.yaml",
				data:     &struct{}{},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Load(tt.args.filename, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoad(t *testing.T) {
	type FakeData struct {
		Name string `yaml:"name"`
		ID   int64  `yaml:"id" validate:"required"`
	}
	type AppConfig struct {
		FakeData `yaml:"appconfig"`
	}

	type args struct {
		filename string
		data     interface{}
	}

	tests := []struct {
		name         string
		args         args
		expectedData *AppConfig
	}{
		{
			name: "success load config",
			args: args{
				filename: "config.yaml",
				data:     &AppConfig{},
			},
			expectedData: &AppConfig{
				FakeData{
					Name: "name",
					ID:   1,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := Load(tt.args.filename, tt.args.data); err != nil {
				t.Errorf("Load() unexpected error = %v", err)
			}

			if !reflect.DeepEqual(tt.expectedData, tt.args.data) {
				t.Errorf("Load() expected = %v, got = %v", tt.expectedData, tt.args.data)
			}
		})
	}
}
