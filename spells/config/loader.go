package config

import (
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

var ErrUnsupportedType = errors.New("unsupported file type. Supported yml, yaml")

func Load(filename string, data interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	switch filepath.Ext(filename) {
	case ".yaml":
		fallthrough
	case ".yml":
		decoder := yaml.NewDecoder(file)
		return decoder.Decode(data)
	default:
		return ErrUnsupportedType
	}
}
