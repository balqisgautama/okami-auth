package jwt

import (
	"encoding/json"
	"github.com/balqisgautama/okami-auth/config"
	"github.com/balqisgautama/okami-auth/constanta"
	"github.com/balqisgautama/okami-auth/dto/res"
	"github.com/balqisgautama/okami-auth/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type PayloadJWTSignIn struct {
	Code1 string `json:"code1"` // from step1
	Code2 string `json:"code2"` // sha512 from email
	jwt.StandardClaims
}

func GenerateJWTSignIn(clientID string, uuid string, sha512uuid string) (token string, output res.APIResponse) {
	tokenCode := PayloadJWTSignIn{
		Code1: uuid,
		Code2: sha512uuid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(constanta.Time30Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "server",
			Subject:   clientID,
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, tokenCode)
	token, err := jwtToken.SignedString([]byte(config.ApplicationConfiguration.GetJWTKey()))
	if err != nil {
		output = model.GenerateJWTError("SignInToken.go", "GenerateJWTSignIn", err)
		return
	}
	return
}

func ValidatorJWTSignIn(jwtToken string) (tokenData *PayloadJWTSignIn, output res.APIResponse) {

	claims := &PayloadJWTSignIn{}
	token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.ApplicationConfiguration.GetJWTKey()), nil
	})

	if err != nil {
		output = model.GenerateJWTError("SignInToken.go", "ValidatorJWTSignIn", err)
		return
	}

	tokenData = convertSignInToken(token.Claims.(*PayloadJWTSignIn))
	return
}

func convertSignInToken(input interface{}) *PayloadJWTSignIn {
	bolB, _ := json.Marshal(input)
	tokenData := PayloadJWTSignIn{}
	json.Unmarshal(bolB, &tokenData)
	return &tokenData
}
