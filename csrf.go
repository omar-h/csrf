package csrf

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
)

const (
	chars   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	idxBits = 6
	idxMask = 1<<idxBits - 1
)

func GenerateRandom(length int) string {
	result := make([]byte, length)
	bufferSize := int(float64(length) * 1.3)
	for i, j, randomBytes := 0, 0, []byte{}; i < length; j++ {
		if j%bufferSize == 0 {
			randomBytes = secureRandomBytes(bufferSize)
		}
		if idx := int(randomBytes[j%length] & idxMask); idx < len(chars) {
			result[i] = chars[idx]
			i++
		}
	}

	return string(result)
}

func GenerateToken(secret, salt string) string {
	return salt + hash(salt+"-"+secret)
}

func Verify(token, secret string, saltLen int) bool {
	salt := token[0:saltLen]
	return salt+hash(salt+"-"+secret) == token
}

func hash(str string) string {
	hash := sha256.New()
	io.WriteString(hash, str)

	hashedString := base64.RawURLEncoding.EncodeToString(hash.Sum(nil))
	return hashedString
}

func secureRandomBytes(length int) []byte {
	var randomBytes = make([]byte, length)
	rand.Read(randomBytes)
	return randomBytes
}
