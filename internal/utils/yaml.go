package utils

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ParseYAML(path string, v any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, v)
}

func MarshalYAML(v any) ([]byte, error) {
	return yaml.Marshal(v)
}
