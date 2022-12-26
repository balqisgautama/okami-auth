package service

import (
	"github.com/balqisgautama/okami-auth/dto"
	"github.com/balqisgautama/okami-auth/dto/res"
	"github.com/balqisgautama/okami-auth/util"
	"net/http"
)

type AbstractService struct {
	FileName string
	FuncName string
}

func (input AbstractService) ReadBody(request *http.Request) (stringBody string, output res.APIResponse) {
	var errorS error
	input.FuncName = "ReadBody"
	input.FileName = "AbstractService.go"

	if request.Method != "GET" {
		stringBody, _, errorS = util.ReadBody(request)
		if errorS != nil {
			output = dto.GenerateInvalidRequestBody(errorS, input.FileName, input.FuncName)
			return
		}
	}

	return
}
