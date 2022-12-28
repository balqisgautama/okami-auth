package util

import (
	"bytes"
	"github.com/balqisgautama/okami-auth/constanta"
	"html/template"
)

func ParseHTMLFileToString(templateFileName string, data interface{}) string {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return constanta.DescActivationFailed
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return constanta.DescActivationFailed
	}

	return buf.String()
}
