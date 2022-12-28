package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"github.com/balqisgautama/okami-auth/config"
	"github.com/balqisgautama/okami-auth/util"
	"go.uber.org/zap"
	"io"
)

func AESEncrypt(text string) string {
	plaintext := []byte(text)

	block, err := aes.NewCipher([]byte(config.ApplicationConfiguration.GetCryptoKey()))
	if err != nil {
		util.Logger.Error("Failed to encrypt", zap.String("details", err.Error()))
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		util.Logger.Error("Failed to encrypt", zap.String("details", err.Error()))
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}
