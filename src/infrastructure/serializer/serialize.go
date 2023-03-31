package serializer

import "encoding/json"

func Marshal(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func Unmarshal(buffer []byte, pOutput interface{}) error {
	return json.Unmarshal(buffer, pOutput)
}
