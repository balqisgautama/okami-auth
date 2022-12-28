package util

import (
	"bytes"
	"github.com/balqisgautama/okami-auth/constanta"
	"html/template"
	"net/smtp"
)

var auth smtp.Auth

func SendEmailGeneral(receiver []string, subject string, url string, token string, htmlAsset string) (result bool) {
	auth = smtp.PlainAuth("", constanta.EmailProject, constanta.EmailAppPassword, constanta.EmailHostGmail)
	templateData := struct {
		Name  string
		URL   string
		Token string
	}{
		Name:  receiver[0],
		URL:   url,
		Token: token,
	}
	r := newRequest(receiver, subject, "")
	err := r.parseTemplate(htmlAsset, templateData)
	if err = r.parseTemplate(htmlAsset, templateData); err == nil {
		ok, _ := r.sendEmail()
		result = ok
	}
	return result
}

type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

func newRequest(to []string, subject, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (r *Request) sendEmail() (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := constanta.EmailHostGmailWithPort

	if err := smtp.SendMail(addr, auth, constanta.EmailProject, r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (r *Request) parseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}

func SendEmailWithTemplate(receiver []string, subject string, htmlAsset string, title string,
	beforeButton string, button string, buttonUrl string, afterButton string) (result bool) {
	auth = smtp.PlainAuth("", constanta.EmailProject, constanta.EmailAppPassword, constanta.EmailHostGmail)

	templateData := struct {
		Title        string
		BeforeButton string
		Button       string
		ButtonUrl    string
		AfterButton  string
	}{
		Title:        title,
		BeforeButton: beforeButton,
		Button:       button,
		ButtonUrl:    buttonUrl,
		AfterButton:  afterButton,
	}

	r := newRequest(receiver, subject, "")
	err := r.parseTemplate(htmlAsset, templateData)
	if err = r.parseTemplate(htmlAsset, templateData); err == nil {
		ok, _ := r.sendEmail()
		result = ok
	}
	return result
}
