package jwt

import (
	"encoding/json"
	"github.com/balqisgautama/okami-auth/config"
	"github.com/balqisgautama/okami-auth/constanta"
	"github.com/balqisgautama/okami-auth/dto/res"
	"github.com/balqisgautama/okami-auth/model"
	"github.com/balqisgautama/okami-auth/util"
	"github.com/balqisgautama/okami-auth/util/crypto"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type PayloadJWTActivation struct {
	Email    string `json:"email"`
	ClientID string `json:"client_id"`
	jwt.StandardClaims
}

func GenerateJWTActivation(email string, clientID string) (token string, output res.APIResponse) {
	subject := util.CheckSumWithMD5([]byte(clientID))
	encryptClientID := crypto.AESEncrypt(clientID)
	tokenCode := PayloadJWTActivation{
		Email:    email,
		ClientID: encryptClientID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(constanta.Time3Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "server-mail",
			Subject:   subject,
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, tokenCode)
	token, err := jwtToken.SignedString([]byte(config.ApplicationConfiguration.GetJWTKey()))
	if err != nil {
		output = model.GenerateJWTError("ActivationToken.go", "GenerateJWTActivation", err)
		return
	}
	return
}

func ValidatorJWTActivation(jwtToken string) (tokenData *PayloadJWTActivation, output res.APIResponse) {
	claims := &PayloadJWTActivation{}
	token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.ApplicationConfiguration.GetJWTKey()), nil
	})

	if err != nil {
		output = model.GenerateJWTError("ActivationToken.go", "ValidatorJWTActivation", err)
		return
	}

	tokenData = convertActivationToken(token.Claims.(*PayloadJWTActivation))

	if tokenData != nil {
		decryptClientID := crypto.AESDecrypt(tokenData.ClientID)
		if util.CheckSumWithMD5([]byte(decryptClientID)) != tokenData.Subject {
			output = model.GenerateJWTError("ActivationToken.go", "ValidatorJWTActivation", err)
			return
		}
	}

	return
}

func convertActivationToken(input interface{}) *PayloadJWTActivation {
	bolB, _ := json.Marshal(input)
	tokenData := PayloadJWTActivation{}
	json.Unmarshal(bolB, &tokenData)
	return &tokenData
}
