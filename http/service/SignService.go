package service

import (
	"encoding/json"
	"errors"
	"github.com/balqisgautama/okami-auth/constanta"
	"github.com/balqisgautama/okami-auth/dao"
	"github.com/balqisgautama/okami-auth/dto"
	"github.com/balqisgautama/okami-auth/dto/req"
	"github.com/balqisgautama/okami-auth/dto/res"
	"github.com/balqisgautama/okami-auth/model"
	"github.com/balqisgautama/okami-auth/util"
	"github.com/balqisgautama/okami-auth/util/converter"
	"github.com/balqisgautama/okami-auth/util/jwt"
	"net/http"
)

type signService struct {
	AbstractService
}

var SignService = signService{}.New()

func (input signService) New() (output signService) {
	output.FileName = "SignService.go"
	return
}

func (input signService) SignIn(request *http.Request) (output res.APIResponse, header map[string]string, err error) {
	input.FuncName = "SignIn"

	result, output_ := input.readBodyAndValidateSigning(request, req.ValidateSign)
	if output_.Status.Code != "" {
		output = output_
		return
	}

	userFound, output_ := dao.UserDAO.GetUserByEmail(result.Email)
	if output_.Status.Code != "" {
		output = output_
		return
	}

	if util.CheckSumWithSha256([]byte(result.Password)) != userFound.Password.String {
		output = model.GenerateValidationFailed(input.FileName, input.FuncName, errors.New(constanta.DescInvalidPassword))
		return
	}

	token, output_ := jwt.GenerateJWTUser(userFound.UserID.Int64, userFound.ClientID.String)
	if output_.Status.Code != "" {
		output = output_
		return
	}

	headerData := map[string]string{
		"user-token": token,
	}
	header = headerData

	output.Status.Success = true
	output.Data = converter.UserModelToSignRes(userFound)
	return
}

func (input signService) readBodyAndValidateSigning(request *http.Request, validation func(input *req.SignForm) (output res.APIResponse)) (inputStruct req.SignForm, output res.APIResponse) {
	input.FuncName = "readBodyAndValidateSigning"
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

	check, text, limit := util.IsPasswordStandardValid(inputStruct.Password)
	if !check {
		err := errors.New(text + "_" + limit)
		output = dto.GenerateInvalidRequestBody(err, input.FileName, input.FuncName)
		return
	}

	return
}
