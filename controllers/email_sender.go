package controllers

import (
	"crypto/tls"
	"io/ioutil"

	"github.com/gin-gonic/gin"

	gomail "gopkg.in/mail.v2"
)

type Email struct {
	to              string
	attachmentBytes []byte
	filename        string
}

func SendEmail(c *gin.Context) {

	defaultEmail := "eanderea1@gmail.com"
	defaultPassword := ""
	var email Email
	ioutil.WriteFile("images/"+email.filename, email.attachmentBytes, 0644)
	c.BindJSON(&email)
	m := gomail.NewMessage()
	m.SetHeader("From", defaultEmail)
	m.SetHeader("To", email.to)
	m.Attach("images/" + email.filename)

	d := gomail.NewDialer("smtp.gmail.com", 587, defaultEmail, defaultPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		c.JSON(400, gin.H{
			"error": true,
		})
	}

	c.JSON(200, gin.H{
		"data": true,
	})

}
