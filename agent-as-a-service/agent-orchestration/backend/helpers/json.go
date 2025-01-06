package helpers

import "encoding/json"

func ConvertJsonString(data interface{}) string {
	if data == nil {
		return ""
	}
	b, _ := json.Marshal(data)
	return string(b)
}

func ConvertJsonObject(jsonStr string, data interface{}) error {
	return json.Unmarshal([]byte(jsonStr), data)
}
