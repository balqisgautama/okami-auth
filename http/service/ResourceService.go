package service

import (
	"encoding/json"
	"github.com/balqisgautama/okami-auth/constanta"
	"github.com/balqisgautama/okami-auth/dao"
	"github.com/balqisgautama/okami-auth/dto"
	"github.com/balqisgautama/okami-auth/dto/req"
	"github.com/balqisgautama/okami-auth/dto/res"
	"github.com/balqisgautama/okami-auth/util"
	"github.com/balqisgautama/okami-auth/util/converter"
	"net/http"
	"strconv"
)

type resourceService struct {
	AbstractService
}

var ResourceService = resourceService{}.New()

func (input resourceService) New() (output resourceService) {
	output.FileName = "ResourceService.go"
	return
}

func (input resourceService) CreateResource(request *http.Request) (output res.APIResponse, header map[string]string, err error) {
	input.FuncName = "CreateResource"

	result, output_ := input.readBodyAndValidateResourceForm(request, req.ValidateResourceForm)
	if output_.Status.Code != "" {
		output = output_
		return
	}
	clientID := request.Header.Get(constanta.ClientIDHeaderNameConstanta)
	resourceInserted, output_ := dao.ResourceDAO.ResourceInsert(result, clientID)
	if output_.Status.Code != "" {
		output = output_
		return
	}

	output.Status.Success = true
	output.Data = converter.ModelToResponse(resourceInserted)
	return
}

func (input resourceService) EditResource(request *http.Request) (output res.APIResponse, header map[string]string, err error) {
	input.FuncName = "EditResource"

	result, output_ := input.readBodyAndValidateResourceForm(request, req.ValidateResourceForm)
	if output_.Status.Code != "" {
		output = output_
		return
	}

	clientID := request.Header.Get(constanta.ClientIDHeaderNameConstanta)
	resourceUpdated, output_ := dao.ResourceDAO.UpdateResourceByID(result.ResourceID, result.Surname, result.Nickname, clientID)
	if output_.Status.Code != "" {
		output = output_
		return
	}

	output.Status.Success = true
	output.Data = converter.ModelToResponse(resourceUpdated)
	return
}

func (input resourceService) DeleteResource(request *http.Request) (output res.APIResponse, header map[string]string, err error) {
	input.FuncName = "DeleteResource"

	params := util.GenerateQueryParam(request)
	temp := params[constanta.ParamSearchID]
	searchByID, _ := strconv.ParseInt(temp, 10, 64)

	clientID := request.Header.Get(constanta.ClientIDHeaderNameConstanta)

	resourceFound, output_ := dao.ResourceDAO.GetByResourceID(searchByID)
	if output_.Status.Code != "" {
		output = output_
		return
	}
	resourceDeleted, output_ := dao.ResourceDAO.DeleteResourceByID(searchByID, clientID, resourceFound.Nickname.String)
	if output_.Status.Code != "" {
		output = output_
		return
	}

	output.Status.Success = true
	output.Data = converter.ModelToResponse(resourceDeleted)
	return
}

func (input resourceService) GetResource(request *http.Request) (output res.APIResponse, header map[string]string, err error) {
	input.FuncName = "GetResource"

	params := util.GenerateQueryParam(request)
	temp := params[constanta.ParamSearchID]
	searchByID, _ := strconv.ParseInt(temp, 10, 64)

	if searchByID == 0 {
		resources, output_ := dao.ResourceDAO.GetResources()
		if output_.Status.Code != "" {
			output = output_
			return
		}
		output.Status.Success = true
		output.Data = converter.ArrayModelToResponse(resources)
		return
	}

	resourceFound, output_ := dao.ResourceDAO.GetByResourceID(searchByID)
	if output_.Status.Code != "" {
		output = output_
		return
	}

	output.Status.Success = true
	output.Data = converter.ModelToResponse(resourceFound)
	return
}

func (input resourceService) readBodyAndValidateResourceForm(request *http.Request, validation func(input *req.ResourceForm) (output res.APIResponse)) (inputStruct req.ResourceForm, output res.APIResponse) {
	input.FuncName = "readBodyAndValidateResourceForm"
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
