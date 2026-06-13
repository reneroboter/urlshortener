package infrastructure

import (
	"crypto/sha1"
	"encoding/hex"
)

type CodeGenerator interface {
	GenerateCode(s string) string
}

type SHA1CodeGenerator struct{}

func (g SHA1CodeGenerator) GenerateCode(s string) string {
	h := sha1.New()
	h.Write([]byte(s))

	generatedCode := hex.EncodeToString(h.Sum(nil))

	return generatedCode
}
