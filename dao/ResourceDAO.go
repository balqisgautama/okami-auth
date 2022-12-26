package dao

import (
	"errors"
	"github.com/balqisgautama/okami-auth/config/server"
	"github.com/balqisgautama/okami-auth/constanta"
	"github.com/balqisgautama/okami-auth/dto/req"
	"github.com/balqisgautama/okami-auth/dto/res"
	"github.com/balqisgautama/okami-auth/model"
	"github.com/balqisgautama/okami-auth/util"
	"time"
)

type resourceDAO struct {
	AbstractDAO
}

var ResourceDAO = resourceDAO{}.New()

func (input resourceDAO) New() (output resourceDAO) {
	output.FileName = "ResourceDAO.go"
	output.TableName = "resources"
	return
}

func (input resourceDAO) ResourceInsert(resource req.ResourceForm, clientID string) (result model.ResourceModel, output res.APIResponse) {
	input.FuncName = "ResourceInsert"

	sqlStatement := `INSERT INTO ` + input.TableName + ` (surname, nickname, client_id, created_at, created_client) ` +
		`VALUES ($1, $2, $3, $4, $5) ` +
		`RETURNING *`

	db := server.ServerConfig.DBConnection
	row := db.QueryRow(sqlStatement, resource.Surname, resource.Nickname, util.GetUUID(), time.Now(), clientID)

	err := row.Scan(&result.ResourceID, &result.Surname, &result.Nickname, &result.ClientID,
		&result.CreatedAt, &result.CreateClient, &result.UpdatedAt, &result.UpdatedClient,
		&result.DeletedAt, &result.DeletedClient)

	if err != nil {
		output = model.GenerateDBServerError(input.FileName, input.FuncName, input.TableName, err)
		return
	}

	if result.ResourceID.Int64 == 0 {
		output = model.GenerateDBServerError(input.FileName, input.FuncName, input.TableName, errors.New(constanta.DescDataNotFound))
		return
	}

	return
}

func (input resourceDAO) UpdateResourceByID(id int64, surname string, nickname string, clientID string) (result model.ResourceModel, output res.APIResponse) {
	input.FuncName = "UpdateResourceByID"

	sqlStatement := `UPDATE ` + input.TableName + ` SET surname=$2, nickname=$3, ` +
		`updated_at=$4, updated_client=$5 ` +
		`WHERE resource_id=$1 ` +
		`RETURNING *`

	db := server.ServerConfig.DBConnection
	row := db.QueryRow(sqlStatement, id, surname, nickname, time.Now(), clientID)

	err := row.Scan(&result.ResourceID, &result.Surname, &result.Nickname, &result.ClientID,
		&result.CreatedAt, &result.CreateClient, &result.UpdatedAt, &result.UpdatedClient,
		&result.DeletedAt, &result.DeletedClient)

	if err != nil {
		output = model.GenerateDBServerError(input.FileName, input.FuncName, input.TableName, err)
		return
	}

	if result.ResourceID.Int64 == 0 {
		output = model.GenerateDBServerError(input.FileName, input.FuncName, input.TableName, errors.New(constanta.DescDataNotFound))
		return
	}

	return
}

func (input resourceDAO) DeleteResourceByID(id int64, clientID string, nickname string) (result model.ResourceModel, output res.APIResponse) {
	input.FuncName = "DeleteResourceByID"

	sqlStatement := `UPDATE ` + input.TableName + ` SET deleted_at=$2, deleted_client=$3, nickname=$4` +
		`WHERE resource_id=$1 ` +
		`RETURNING *`

	resourceDelete := nickname + "-" + util.GetUUID()

	db := server.ServerConfig.DBConnection
	row := db.QueryRow(sqlStatement, id, time.Now(), clientID, resourceDelete)

	err := row.Scan(&result.ResourceID, &result.Surname, &result.Nickname, &result.ClientID,
		&result.CreatedAt, &result.CreateClient, &result.UpdatedAt, &result.UpdatedClient,
		&result.DeletedAt, &result.DeletedClient)

	if err != nil {
		output = model.GenerateDBServerError(input.FileName, input.FuncName, input.TableName, err)
		return
	}

	if result.ResourceID.Int64 == 0 {
		output = model.GenerateDBServerError(input.FileName, input.FuncName, input.TableName, errors.New(constanta.DescDataNotFound))
		return
	}

	return
}

func (input resourceDAO) GetByResourceID(id int64) (result model.ResourceModel, output res.APIResponse) {
	input.FuncName = "GetByResourceID"

	sqlStatement := `SELECT * ` +
		`FROM ` + input.TableName +
		` WHERE resource_id=$1 AND EXTRACT(EPOCH FROM deleted_at) is NULL`

	db := server.ServerConfig.DBConnection
	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&result.ResourceID, &result.Surname, &result.Nickname, &result.ClientID,
		&result.CreatedAt, &result.CreateClient, &result.UpdatedAt, &result.UpdatedClient,
		&result.DeletedAt, &result.DeletedClient)

	if err != nil {
		output = model.GenerateDBServerError(input.FileName, input.FuncName, input.TableName, err)
		return
	}

	if result.ResourceID.Int64 == 0 {
		output = model.GenerateDBServerError(input.FileName, input.FuncName, input.TableName, errors.New(constanta.DescDataNotFound))
		return
	}

	return
}

func (input resourceDAO) GetResources() (resources []model.ResourceModel, output res.APIResponse) {
	input.FuncName = "GetResources"

	sqlStatement := `SELECT * ` +
		`FROM ` + input.TableName +
		` WHERE EXTRACT(EPOCH FROM deleted_at) is NULL`

	db := server.ServerConfig.DBConnection
	rows, err := db.Query(sqlStatement)

	for rows.Next() {
		var resource model.ResourceModel
		err = rows.Scan(&resource.ResourceID, &resource.Surname, &resource.Nickname, &resource.ClientID,
			&resource.CreatedAt, &resource.CreateClient, &resource.UpdatedAt, &resource.UpdatedClient,
			&resource.DeletedAt, &resource.DeletedClient)
		resources = append(resources, resource)
	}

	if err != nil {
		output = model.GenerateDBServerError(input.FileName, input.FuncName, input.TableName, err)
		return
	}

	if len(resources) <= 0 {
		output = model.GenerateDBServerError(input.FileName, input.FuncName, input.TableName, errors.New(constanta.DescDataNotFound))
		return
	}

	return
}
