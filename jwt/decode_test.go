package jwt

import (
	"testing"

	"github.com/designsbysm/server-go/database"
	"github.com/google/uuid"
)

func TestDecode(t *testing.T) {
	uuid := uuid.New()
	role := database.Role{
		ID:      1,
		IsAdmin: true,
		Name:    "Test",
	}
	token, _ := Encode(uuid, role)

	claims, err := Decode(token)

	if err != nil {
		t.Errorf("should be empty, got %v", err)
	} else if claims == nil {
		t.Errorf("should not be empty, got nil")
	}
}

func TestEmptyToken(t *testing.T) {
	claims, err := Decode("")

	if err == nil {
		t.Errorf("should not be empty, got nil")
	} else if claims != nil {
		t.Errorf("should be empty, got %v", claims)
	}
}

func TestInvalidToken(t *testing.T) {
	claims, err := Decode("eyJhbGciOiJQUzM4NCIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.W0vasx0Lxmn3hkfi3RjJSLDSZnZbqDRZy6fau8wvIm_sdtTCiSJAjMSO1eR8YbnUF_MD10JYwR5unI8JO_qB8uSrW5vg4OPCDuPSQQu0Pnf9Q2Cy3WqSK166lgidlwAqyijgPFp5ggOnJM20IY4F8W6HOqcGeXcRzsaM2DIBEnku32TM7Xb-aCJYdKtawEfvD1zEwE1of02BoGva3sf_RhijMZpOA3yIG3FCDll-3M1rILP9Bi4FPz_uOAbJkKvUirAdaMX-KNaw6T_0nnBwHaFc8M9GGVZV6bT4uaOQ1U0Ezi4SAcBWW8kkPKZNorpRC5EtO_X6uvIVazj5EHERuw")

	if err == nil {
		t.Error("empty err")
	} else if claims != nil {
		t.Errorf("claims should be empty: got %v", claims)
	}
}

// func TestEmptyClaims(t *testing.T) {
// 	claims, err := Decode("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.Et9HFtf9R3GEMA0IICOfFMVXY7kkTX1wr4qCyhIf58U")

// 	if err == nil {
// 		t.Error("empty err")
// 	} else if claims != nil {
// 		t.Errorf("claims should be empty: got %v", claims)
// 	}
// }
