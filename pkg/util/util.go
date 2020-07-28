package util

import (
	"encoding/json"
)

func ToJSON(i interface{}) (string, error) {
	b, err := json.Marshal(i)
	return string(b), err
}
