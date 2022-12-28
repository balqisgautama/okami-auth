package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"github.com/balqisgautama/okami-auth/config"
	"github.com/balqisgautama/okami-auth/util"
	"go.uber.org/zap"
)

func AESDecrypt(text string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(text)

	block, err := aes.NewCipher([]byte(config.ApplicationConfiguration.GetCryptoKey()))
	if err != nil {
		util.Logger.Error("Failed to encrypt", zap.String("details", err.Error()))
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		util.Logger.Error("Failed to encrypt", zap.String("details", "ciphertext too short"))
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}
