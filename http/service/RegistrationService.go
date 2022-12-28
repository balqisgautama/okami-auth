package service

import (
	"encoding/json"
	"errors"
	"github.com/balqisgautama/okami-auth/config"
	"github.com/balqisgautama/okami-auth/constanta"
	"github.com/balqisgautama/okami-auth/dao"
	"github.com/balqisgautama/okami-auth/dto"
	"github.com/balqisgautama/okami-auth/dto/req"
	"github.com/balqisgautama/okami-auth/dto/res"
	"github.com/balqisgautama/okami-auth/util"
	"github.com/balqisgautama/okami-auth/util/crypto"
	"github.com/balqisgautama/okami-auth/util/jwt"
	"github.com/gorilla/mux"
	"net/http"
)

type registrationService struct {
	AbstractService
}

var RegistrationService = registrationService{}.New()

func (input registrationService) New() (output registrationService) {
	output.FileName = "RegistrationService.go"
	return
}

func (input registrationService) UserRegistration(request *http.Request) (output res.APIResponse, header map[string]string, err error) {
	input.FuncName = "UserRegistration"
	result, output_ := input.readBodyAndValidateRegister(request, req.ValidateRegistration)
	if output_.Status.Code != "" {
		output = output_
		return
	}

	userInserted, output_ := dao.UserDAO.InsertUser(result)
	if output_.Status.Code != "" {
		output = output_
		return
	}

	output_ = EmailService.SendActivationLink(result.Email, userInserted.ClientID.String)
	if output_.Status.Code != "" {
		output = output_
		return
	}

	output.Status.Success = true
	output.Status.Message = "Please check your mailbox"
	return
}

func (input registrationService) ActivationAccount(request *http.Request) (output res.APIResponse, header map[string]string, err error) {

	vars := mux.Vars(request)
	token := vars[constanta.VarToken]
	code := vars[constanta.VarCode]

	tokenData, output_ := jwt.ValidatorJWTActivation(token)
	if output_.Status.Code != "" {
		output = input.toHTMLResponse(false, false, false, code)
		return
	}

	userFound, output_ := dao.UserDAO.GetUserByEmail(tokenData.Email)
	if output_.Status.Code != "" {
		output = input.toHTMLResponse(false, false, false, code)
		return
	}

	_, output_ = dao.UserDAO.ActiveAccountByClientID(userFound.ClientID.String)
	if output_.Status.Code != "" {
		output = input.toHTMLResponse(false, true, false, code)
		return
	}

	output.Status.Success = true
	output = input.toHTMLResponse(true, true, false, "")
	return
}

func (input registrationService) ActivationAccountResend(request *http.Request) (output res.APIResponse, header map[string]string, err error) {

	vars := mux.Vars(request)
	token := vars[constanta.VarToken]
	code := vars[constanta.VarCode]

	_, output_ := jwt.ValidatorJWTActivation(token)
	if output_.Status.Code != "" {
		output = input.toHTMLResponse(false, false, true, code)
		return
	}

	decryptClientID := crypto.AESDecrypt(code)
	userFound, output_ := dao.UserDAO.GetUserByClientID(decryptClientID)
	if output_.Status.Code != "" {
		output = input.toHTMLResponse(false, false, true, code)
		return
	}
	output_ = EmailService.SendActivationLink(userFound.Email.String, userFound.ClientID.String)
	if output_.Status.Code != "" {
		output = input.toHTMLResponse(false, true, true, code)
		return
	}

	output.Status.Success = true
	output = input.toHTMLResponse(true, true, true, "")
	return
}

func (input registrationService) readBodyAndValidateRegister(request *http.Request, validation func(input *req.RegistrationForm) (output res.APIResponse)) (inputStruct req.RegistrationForm, output res.APIResponse) {
	input.FuncName = "readBodyAndValidateRegister"
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

func (input registrationService) toHTMLResponse(success bool, authorize bool, resend bool, clientID string) (output res.APIResponse) {

	var templateHTML res.DataHTMLFile
	templateHTML.Title = "Welcome,"
	templateHTML.BeforeButton = "Your account is active. Enjoy our live demo"
	templateHTML.Button = "Back To Home"
	templateHTML.ButtonUrl = config.ApplicationConfiguration.GetServerHost()

	if !success && !resend {
		decryptClientID := crypto.AESDecrypt(clientID)
		userFound, _ := dao.UserDAO.GetUserByClientID(decryptClientID)
		templateHTML.Title = "Sorry,"
		templateHTML.BeforeButton = "Failed to activate your account. Please click the following link for re-activation."
		templateHTML.Button = "Resend"
		templateHTML.ButtonUrl = EmailService.generateResendActivationURL(userFound.Email.String, userFound.ClientID.String)
	}

	if !success && resend {
		decryptClientID := crypto.AESDecrypt(clientID)
		userFound, _ := dao.UserDAO.GetUserByClientID(decryptClientID)
		templateHTML.Title = "Sorry,"
		templateHTML.BeforeButton = "Failed to send activation URL. Please click the following link to resend."
		templateHTML.Button = "Resend"
		templateHTML.ButtonUrl = EmailService.generateResendActivationURL(userFound.Email.String, userFound.ClientID.String)
	}

	if success && resend {
		templateHTML.Title = "Hi,"
		templateHTML.BeforeButton = "Activation email has been sent. Please check your mailbox."
	}

	if !authorize {
		templateHTML.Title = "Sorry,"
		templateHTML.BeforeButton = "Something wrong with your activation account. Please contact our administrator"
		templateHTML.Button = "WhatsApp Us"
		templateHTML.ButtonUrl = "https://wa.link/alq90l"
	}

	templateHTML.AfterButton = "Regards Okami"
	output.Status.Code = constanta.HeaderValueContentTypeHTML
	output.Status.Message = util.ParseHTMLFileToString(constanta.PathAssetResponseHTMLButton, templateHTML)
	return
}
