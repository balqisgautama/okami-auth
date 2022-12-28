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

type userDAO struct {
	AbstractDAO
}

var UserDAO = userDAO{}.New()

func (input userDAO) New() (output userDAO) {
	output.FileName = "UserDAO.go"
	output.TableName = "users"
	return
}

func (input userDAO) InsertUser(data req.RegistrationForm) (result model.UserModel, output res.APIResponse) {
	input.FuncName = "InsertUser"

	sqlStatement := `INSERT INTO ` + input.TableName + ` (email, password, client_id, created_at) ` +
		`VALUES ($1, $2, $3, $4) ` +
		`RETURNING user_id, email, password, client_id, status, sysadmin`

	db := server.ServerConfig.DBConnection
	passwordSHA256 := util.CheckSumWithSha256([]byte(data.Password))
	row := db.QueryRow(sqlStatement, data.Email, passwordSHA256, util.GetUUID(), time.Now())

	err := row.Scan(&result.UserID, &result.Email, &result.Password, &result.ClientID,
		&result.Status, &result.Sysadmin)

	if err != nil {
		output = model.GenerateDBServerError(input.FileName, input.FuncName, input.TableName, err)
		return
	}

	if result.UserID.Int64 == 0 {
		output = model.GenerateDBServerError(input.FileName, input.FuncName, input.TableName, errors.New(constanta.DescDataNotFound))
		return
	}

	return
}

func (input userDAO) GetUserByEmail(email string) (result model.UserModel, output res.APIResponse) {
	input.FuncName = "GetUserByEmail"

	sqlStatement := `SELECT user_id, email, password, client_id, status, sysadmin ` +
		`FROM ` + input.TableName +
		` WHERE email=$1 AND EXTRACT(EPOCH FROM deleted_at) is NULL`

	row := server.ServerConfig.DBConnection.QueryRow(sqlStatement, email)

	err := row.Scan(&result.UserID, &result.Email, &result.Password, &result.ClientID,
		&result.Status, &result.Sysadmin)

	if err != nil {
		output = model.GenerateDBServerError(input.FileName, input.FuncName, input.TableName, err)
		return
	}

	if result.UserID.Int64 == 0 {
		output = model.GenerateDBServerError(input.FileName, input.FuncName, input.TableName, errors.New(constanta.DescDataNotFound))
		return
	}

	return
}

func (input userDAO) GetUserByClientID(clientID string) (result model.UserModel, output res.APIResponse) {
	input.FuncName = "GetUserByClientID"

	sqlStatement := `SELECT user_id, email, password, client_id, status, sysadmin ` +
		`FROM ` + input.TableName +
		` WHERE client_id=$1 AND EXTRACT(EPOCH FROM deleted_at) is NULL`

	db := server.ServerConfig.DBConnection
	row := db.QueryRow(sqlStatement, clientID)

	err := row.Scan(&result.UserID, &result.Email, &result.Password, &result.ClientID,
		&result.Status, &result.Sysadmin)

	if err != nil {
		output = model.GenerateDBServerError(input.FileName, input.FuncName, input.TableName, err)
		return
	}

	if result.UserID.Int64 == 0 {
		output = model.GenerateDBServerError(input.FileName, input.FuncName, input.TableName, errors.New(constanta.DescDataNotFound))
		return
	}

	return
}

func (input userDAO) ActiveAccountByClientID(clientID string) (result model.UserModel, output res.APIResponse) {
	input.FuncName = "ActiveAccountByClientID"

	sqlStatement := `UPDATE ` + input.TableName + ` SET updated_at=$2, status=$3` +
		`WHERE client_id=$1 ` +
		`RETURNING user_id, email, client_id, status, sysadmin`

	row := server.ServerConfig.DBConnection.QueryRow(sqlStatement, clientID, time.Now(), constanta.UserActive)

	err := row.Scan(&result.UserID, &result.Email, &result.ClientID,
		&result.Status, &result.Sysadmin)

	if err != nil {
		output = model.GenerateDBServerError(input.FileName, input.FuncName, input.TableName, err)
		return
	}

	if result.UserID.Int64 == 0 {
		output = model.GenerateDBServerError(input.FileName, input.FuncName, input.TableName, errors.New(constanta.DescDataNotFound))
		return
	}

	return
}
