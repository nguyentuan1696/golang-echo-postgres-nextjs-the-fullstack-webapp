package util

import "encoding/json"

func ToJson(v any) string {
	b, err := json.Marshal(v)

	if err != nil {
		return ""
	}
	
	return string(b)
}
