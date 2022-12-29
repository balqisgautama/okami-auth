package service

import (
	"github.com/balqisgautama/okami-auth/config"
	"github.com/balqisgautama/okami-auth/constanta"
	"github.com/balqisgautama/okami-auth/dto/res"
	"github.com/balqisgautama/okami-auth/model"
	"github.com/balqisgautama/okami-auth/util"
	"github.com/balqisgautama/okami-auth/util/crypto"
	"github.com/balqisgautama/okami-auth/util/jwt"
	"strconv"
)

type emailService struct {
	AbstractService
}

var EmailService = emailService{}.New()

func (input emailService) New() (output emailService) {
	output.FileName = "EmailService.go"
	return
}

func (input emailService) SendActivationLink(sendTo string, clientID string) (output res.APIResponse) {
	input.FuncName = "SendActivationLink"

	subject := "Activation Account"
	title := "Hi,"
	beforeButton := "Thank you for your registration. Please click the following link for activation."
	button := "Active My Account"
	buttonUrl := input.generateActivationURL(sendTo, clientID)
	afterButton := "Please do not reply to this email"

	send := util.SendEmailWithTemplate([]string{sendTo}, subject,
		constanta.PathAssetResponseHTMLButton, title, beforeButton, button, buttonUrl, afterButton)
	if !send {
		output = model.GenerateMailError(input.FileName, input.FuncName)
		return
	}

	return
}

func (input emailService) generateActivationURL(email string, clientID string) (url string) {
	token, output := jwt.GenerateJWTActivation(email, clientID)
	encrypt := crypto.AESEncrypt(clientID)

	if output.Status.Code != "" {
		return
	}
	url = config.ApplicationConfiguration.GetServerHost() + ":" +
		strconv.Itoa(config.ApplicationConfiguration.GetServerPort()) + "/" +
		config.ApplicationConfiguration.GetServerPrefixPath() + "/active/" +
		token + "/" + encrypt
	return
}

func (input emailService) generateResendActivationURL(email string, clientID string) (url string) {
	token, output := jwt.GenerateJWTActivation(email, clientID)
	encrypt := crypto.AESEncrypt(clientID)
	if output.Status.Code != "" {
		return
	}
	url = config.ApplicationConfiguration.GetServerHost() + ":" +
		strconv.Itoa(config.ApplicationConfiguration.GetServerPort()) + "/" +
		config.ApplicationConfiguration.GetServerPrefixPath() + "/active/resend/" +
		token + "/" + encrypt
	return
}

func (input emailService) SendLoginConfirmationLink(sendTo string, token string) (output res.APIResponse) {
	input.FuncName = "SendLoginConfirmationLink"

	subject := "Login Confirmation"
	title := "Hi, " + sendTo
	beforeButton := "To validate this login, please click on the confirmation link below"
	button := "Continue Login"
	buttonUrl := input.generateLoginConfirmationURL(token)
	afterButton := "Please do not reply to this email"

	send := util.SendEmailWithTemplate([]string{sendTo}, subject,
		constanta.PathAssetResponseHTMLButton, title, beforeButton, button, buttonUrl, afterButton)
	if !send {
		output = model.GenerateMailError(input.FileName, input.FuncName)
		return
	}

	return
}

func (input emailService) generateLoginConfirmationURL(token string) (url string) {
	url = config.ApplicationConfiguration.GetServerHost() + ":" +
		strconv.Itoa(config.ApplicationConfiguration.GetServerPort()) + "/" +
		config.ApplicationConfiguration.GetServerPrefixPath() + "/sign/in/step3/" + token
	return
}
