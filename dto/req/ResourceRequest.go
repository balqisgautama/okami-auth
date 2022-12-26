package req

import (
	"github.com/balqisgautama/okami-auth/dto"
	"github.com/balqisgautama/okami-auth/dto/res"
	"github.com/go-playground/validator/v10"
)

func init() {
	filename = "ResourceRequest.go"
}

type ResourceForm struct {
	//Email    string `json:"email" validate:"required,email"`
	//Password string `json:"password" validate:"required"`
	Surname    string `json:"surname" validate:"required,min=3"`
	Nickname   string `json:"nickname" validate:"required,min=3,alpha,lowercase"`
	ResourceID int64  `json:"resource_id" validate:"omitempty,number,min=1"`
}

func ValidateResourceForm(inputStruct *ResourceForm) (output res.APIResponse) {
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
