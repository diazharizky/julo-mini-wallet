package modules

import "encoding/base64"

type validateTokenModule struct{}

func NewValidateTokenModule() validateTokenModule {
	return validateTokenModule{}
}

func (m validateTokenModule) Call(token string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
