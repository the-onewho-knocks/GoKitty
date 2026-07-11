package utils

import "encoding/json"

func MarshalJSON(v any) ([]byte, error) {
	return json.Marshal(v)
}

func MarshalIndentJSON(v any) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}

func UnmarshalJSON(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func PrettyPrintJSON(v any) string {
	data, err := MarshalIndentJSON(v)
	if err != nil {
		return err.Error()
	}
	return string(data)
}
