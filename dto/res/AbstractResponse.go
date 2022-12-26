package res

import (
	"github.com/balqisgautama/okami-auth/util"
)

type APIResponse struct {
	Timestamp int64          `json:"timestamp"`
	Status    StatusResponse `json:"status"`
	Data      interface{}
}

type StatusResponse struct {
	Success bool     `json:"success"`
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Detail  []string `json:"detail"`
}

func (input APIResponse) String() string {
	return util.StructToJSON(input)
}
