package util

import (
	"github.com/google/uuid"
)

func GetUUID() (output string) {
	UUID, _ := uuid.NewRandom() // uuid v4
	output = UUID.String()
	//output = strings.Replace(output, "-", "", -1)
	return
}
