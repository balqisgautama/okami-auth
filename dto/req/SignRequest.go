package req

import (
	"github.com/balqisgautama/okami-auth/dto"
	"github.com/balqisgautama/okami-auth/dto/res"
	"github.com/go-playground/validator/v10"
)

func init() {
	filename = "SignRequest.go"
}

type SignForm struct {
	Email    string `json:"email" validate:"required,min=5,email"`
	Password string `json:"password" validate:"required,min=3"`
}

func ValidateSign(inputStruct *SignForm) (output res.APIResponse) {
	funcName = "ValidateResourceForm"
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
