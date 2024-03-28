package handlers

import (
	"4th_Assignment/db"
	"4th_Assignment/ent"
	"context"
	"crypto/tls"
	"fmt"

	"github.com/gofiber/fiber/v2"
	gomail "gopkg.in/mail.v2"
)

func ContactSubmit(c *fiber.Ctx) error {
	p := new(ent.ContactSubmission)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	p, err := db.DBConn.ContactSubmission.Create().
		SetName(p.Name).
		SetPhone(p.Phone).
		SetEmail(p.Email).
		SetSubject(p.Subject).
		SetMessage(p.Message).
		Save(context.Background())
	if err != nil {
		return err
	}
	//send mail
	m := gomail.NewMessage()
	m.SetHeader("From", "hoangtrungkien2002@hotmail.com")
	m.SetHeader("To", "hoangtrungkien2002@gmail.com")
	m.SetHeader("Subject", p.Subject)
	m.SetBody("text/plain", "From:"+p.Email+"\n"+"Name:"+p.Name+"\n"+"Phone:"+p.Phone+"\n"+"Content:"+p.Message)
	d := gomail.NewDialer("smtp-mail.outlook.com", 587, "hoangtrungkien2002@hotmail.com", "kiendzpro25")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
	return c.JSON(p)
}
