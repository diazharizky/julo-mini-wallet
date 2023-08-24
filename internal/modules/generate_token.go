package modules

import (
	"encoding/base64"

	"github.com/google/uuid"
)

type generateTokenModule struct {
}

func NewGenerateTokenModule() generateTokenModule {
	return generateTokenModule{}
}

func (generateTokenModule) Call(accountID uuid.UUID) string {
	data := []byte(accountID.String())
	return base64.StdEncoding.EncodeToString(data)
}
