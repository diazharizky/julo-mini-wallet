package utils

import "encoding/base64"

func Encode(value string) string {
	data := []byte(value)
	return base64.StdEncoding.EncodeToString(data)
}

func Decode(hash string) (*string, error) {
	data, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		return nil, err
	}

	decoded := string(data)
	return &decoded, nil
}
