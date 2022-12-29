package req

import (
	"github.com/balqisgautama/okami-auth/dto"
	"github.com/balqisgautama/okami-auth/dto/res"
	"github.com/go-playground/validator/v10"
)

func init() {
	filename = "SignInRequest.go"
}

type Step1 struct {
	UUID string `json:"uuid" validate:"required,uuid4_rfc4122"`
}

type Step2 struct {
	Email    string `json:"email" validate:"required,min=5,email"`
	Password string `json:"password" validate:"required,min=3"`
}

type Step3 struct {
	Token string `json:"token" validate:"required,sha512"` // from email
}

func ValidateStep1(inputStruct *Step1) (output res.APIResponse) {
	funcName = "ValidateStep1"
	validate = validator.New()
	err := validate.Struct(inputStruct)
	if err != nil {
		output = dto.GenerateValidationFailed(err, filename, funcName)
		if errV, ok := err.(*validator.InvalidValidationError); ok {
			output = dto.GenerateValidationFailed(errV, filename, funcName)
			return
		}
		return
	}
	output.Status.Success = true
	return
}

func ValidateStep2(inputStruct *Step2) (output res.APIResponse) {
	funcName = "ValidateStep2"
	validate = validator.New()
	err := validate.Struct(inputStruct)
	if err != nil {
		output = dto.GenerateValidationFailed(err, filename, funcName)
		if errV, ok := err.(*validator.InvalidValidationError); ok {
			output = dto.GenerateValidationFailed(errV, filename, funcName)
			return
		}
		return
	}
	output.Status.Success = true
	return
}

func ValidateStep3(inputStruct *Step3) (output res.APIResponse) {
	funcName = "ValidateStep3"
	validate = validator.New()
	err := validate.Struct(inputStruct)
	if err != nil {
		output = dto.GenerateValidationFailed(err, filename, funcName)
		if errV, ok := err.(*validator.InvalidValidationError); ok {
			output = dto.GenerateValidationFailed(errV, filename, funcName)
			return
		}
		return
	}
	output.Status.Success = true
	return
}
