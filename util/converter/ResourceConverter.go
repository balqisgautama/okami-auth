package converter

import (
	"github.com/balqisgautama/okami-auth/dto/res"
	"github.com/balqisgautama/okami-auth/model"
)

func ModelToResponse(data model.ResourceModel) (result res.ResourceDetail) {
	result.ResourceID = data.ResourceID.Int64
	result.ClientID = data.ClientID.String
	result.Surname = data.Surname.String
	result.Nickname = data.Nickname.String
	result.CreatedAt = data.CreatedAt.Time.Unix()
	result.CreateClient = data.CreateClient.String
	result.UpdatedAt = data.UpdatedAt.Time.Unix()
	if data.UpdatedAt.Time.Unix() < 0 {
		result.UpdatedAt = 0
	}
	result.UpdatedClient = data.UpdatedClient.String
	result.DeletedAt = data.DeletedAt.Time.Unix()
	if data.DeletedAt.Time.Unix() < 0 {
		result.DeletedAt = 0
	}
	result.DeletedClient = data.DeletedClient.String
	return
}

func ArrayModelToResponse(dataDB []model.ResourceModel) (result []res.ResourceDetail) {
	for _, value := range dataDB {
		temp := ModelToResponse(value)
		result = append(result, temp)
	}
	return
}
