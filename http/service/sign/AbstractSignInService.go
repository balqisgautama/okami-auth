package sign

import (
	"encoding/json"
	"github.com/balqisgautama/okami-auth/dto"
	"github.com/balqisgautama/okami-auth/dto/req"
	"github.com/balqisgautama/okami-auth/dto/res"
	"github.com/balqisgautama/okami-auth/http/service"
	"net/http"
)

type abstractSignInService struct {
	service.AbstractService
}

var AbstractSignInService = abstractSignInService{}.New()

func (input abstractSignInService) New() (output abstractSignInService) {
	output.FileName = "AbstractSignInService.go"
	return
}

func (input abstractSignInService) readBodyAndValidateStep1(request *http.Request,
	validation func(input *req.Step1) (output res.APIResponse)) (inputStruct req.Step1, output res.APIResponse) {
	input.FuncName = "readBodyAndValidateStep1"
	var stringBody string

	stringBody, output = input.ReadBody(request)
	if output.Status.Code != "" {
		return
	}

	if stringBody != "" {
		errorS := json.Unmarshal([]byte(stringBody), &inputStruct)
		if errorS != nil {
			output = dto.GenerateInvalidRequestBody(errorS, input.FileName, input.FuncName)
			return
		}
	}
	output = validation(&inputStruct)
	if output.Status.Code != "" {
		return
	}

	return
}

func (input abstractSignInService) readBodyAndValidateStep2(request *http.Request,
	validation func(input *req.Step2) (output res.APIResponse)) (inputStruct req.Step2, output res.APIResponse) {
	input.FuncName = "readBodyAndValidateStep2"
	var stringBody string

	stringBody, output = input.ReadBody(request)
	if output.Status.Code != "" {
		return
	}

	if stringBody != "" {
		errorS := json.Unmarshal([]byte(stringBody), &inputStruct)
		if errorS != nil {
			output = dto.GenerateInvalidRequestBody(errorS, input.FileName, input.FuncName)
			return
		}
	}
	output = validation(&inputStruct)
	if output.Status.Code != "" {
		return
	}

	//check, text, limit := util.IsPasswordStandardValid(inputStruct.Password)
	//if !check {
	//	err := errors.New(text + "_" + limit)
	//	output = dto.GenerateInvalidRequestBody(err, input.FileName, input.FuncName)
	//	return
	//}

	return
}
