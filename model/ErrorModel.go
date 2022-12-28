package model

import (
	"github.com/balqisgautama/okami-auth/config"
	"github.com/balqisgautama/okami-auth/dto/res"
	"strings"
)

var resourceID string

func GenerateDBServerError(fileName string, funcName string, tableName string, causedBy error) (output res.APIResponse) {
	resourceID = strings.ToUpper(config.ApplicationConfiguration.GetServerResourceID())
	output.Status.Success = false
	output.Status.Code = resourceID + "-370001-DB-SERVER"
	output.Status.Message = causedBy.Error()
	output.Status.Detail = []string{fileName, funcName, tableName}
	return
}

func GenerateValidationFailed(fileName string, funcName string, causedBy error) (output res.APIResponse) {
	resourceID = strings.ToUpper(config.ApplicationConfiguration.GetServerResourceID())
	output.Status.Success = false
	output.Status.Code = resourceID + "-370002-VALIDATION"
	output.Status.Message = causedBy.Error()
	output.Status.Detail = []string{fileName, funcName}
	return
}

func GenerateJWTError(fileName string, funcName string, causedBy error) (output res.APIResponse) {
	resourceID = strings.ToUpper(config.ApplicationConfiguration.GetServerResourceID())
	output.Status.Success = false
	output.Status.Code = resourceID + "-370003-JWT"
	output.Status.Message = causedBy.Error()
	output.Status.Detail = []string{fileName, funcName}
	return
}

func GenerateUnauthorizedError() (output res.APIResponse) {
	resourceID = strings.ToUpper(config.ApplicationConfiguration.GetServerResourceID())
	output.Status.Success = false
	output.Status.Code = resourceID + "-370004-UNAUTHORIZED"
	return
}

func GenerateMailError(fileName string, funcName string) (output res.APIResponse) {
	resourceID = strings.ToUpper(config.ApplicationConfiguration.GetServerResourceID())
	output.Status.Success = false
	output.Status.Code = resourceID + "-370005-MAIL"
	output.Status.Detail = []string{fileName, funcName}
	return
}
