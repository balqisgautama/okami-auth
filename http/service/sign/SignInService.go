package sign

import (
	"errors"
	"github.com/balqisgautama/okami-auth/constanta"
	"github.com/balqisgautama/okami-auth/dao"
	"github.com/balqisgautama/okami-auth/dto/req"
	"github.com/balqisgautama/okami-auth/dto/res"
	"github.com/balqisgautama/okami-auth/http/service"
	"github.com/balqisgautama/okami-auth/model"
	"github.com/balqisgautama/okami-auth/util"
	"github.com/balqisgautama/okami-auth/util/converter"
	"github.com/balqisgautama/okami-auth/util/jwt"
	"github.com/gorilla/mux"
	"net/http"
)

type signInService struct {
	abstractSignInService
}

var SignInService = signInService{}.New()

func (input signInService) New() (output signInService) {
	output.FileName = "SignInService.go"
	return
}

func (input signInService) Step1(request *http.Request) (output res.APIResponse, header map[string]string, err error) {
	input.FuncName = "Step1"

	result, output_ := input.abstractSignInService.readBodyAndValidateStep1(request, req.ValidateStep1)
	if output_.Status.Code != "" {
		output = output_
		return
	}

	token, output_ := jwt.GenerateJWTSignIn("", result.UUID, "")
	if output_.Status.Code != "" {
		output = output_
		return
	}

	headerData := map[string]string{
		constanta.SecretTokenHeaderName: token,
	}
	header = headerData

	output.Status.Success = true

	return
}

func (input signInService) Step2(request *http.Request) (output res.APIResponse, header map[string]string, err error) {
	input.FuncName = "Step2"

	tokenStep1 := request.Header.Get(constanta.SecretTokenHeaderName)

	result, output_ := input.abstractSignInService.readBodyAndValidateStep2(request, req.ValidateStep2)
	if output_.Status.Code != "" {
		output = output_
		return
	}

	tokenData, output_ := jwt.ValidatorJWTSignIn(tokenStep1)
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

	token, output_ := jwt.GenerateJWTSignIn(userFound.ClientID.String, tokenData.Code1, util.CheckSumWithSha512([]byte(tokenData.Code1)))
	if output_.Status.Code != "" {
		output = output_
		return
	}

	output_ = service.EmailService.SendLoginConfirmationCode(userFound.Email.String, token)
	if output_.Status.Code != "" {
		output = output_
		return
	}

	output.Status.Success = true

	return
}

func (input signInService) Step3(request *http.Request) (output res.APIResponse, header map[string]string, err error) {

	vars := mux.Vars(request)
	token := vars[constanta.VarToken]

	//output.Status.Code = constanta.ResponseTypeRedirect

	tokenData, output_ := jwt.ValidatorJWTSignIn(token)
	if output_.Status.Code != "" || tokenData == nil {
		//output.Status.Message = "https://www.google.com/"
		output = output_
		return
	}

	if util.CheckSumWithSha512([]byte(tokenData.Code1)) != tokenData.Code2 {
		//output.Status.Message = "https://www.google.com/"
		output = model.GenerateLoginError(input.FileName, input.FuncName)
		return
	}

	userFound, output_ := dao.UserDAO.GetUserByClientID(tokenData.Subject)
	if output_.Status.Code != "" {
		output = output_
		return
	}

	userToken, output_ := jwt.GenerateJWTUser(userFound.UserID.Int64, userFound.ClientID.String)
	if output_.Status.Code != "" {
		output = output_
		return
	}

	output.Status.Success = true
	//output.Status.Message = "https://www.google.com/maps"
	output.Data = converter.ToSignInResponse(userFound, userToken)

	return
}
