package converter

import (
	"github.com/balqisgautama/okami-auth/dto/res"
	"github.com/balqisgautama/okami-auth/model"
)

func UserModelToSignRes(user model.UserModel) (result res.SignIn) {
	result.UserID = user.UserID.Int64
	result.ClientID = user.ClientID.String
	result.Email = user.Email.String
	result.Status = user.Status.Int16
	result.Sysadmin = user.Sysadmin.Int16
	return
}
