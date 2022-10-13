package api

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mailgun/mailgun-go/v4"
)

const (
	mailgunDomain          = "sandboxb6c2fb30d44b41e495272143b5d5c41f.mailgun.org"
	privateAPIKeyOfMailgun = "6bac0046ac203622eb58e7c7fe1ae5fe-381f2624-91da3481"
	sender                 = "postmaster@sandboxb6c2fb30d44b41e495272143b5d5c41f.mailgun.org"
	msgTemplate            = `
	user name: %s
	user email: %s
	user feedback: %s
	`
	feedbackEmail = "wcs399@proton.me"
)

func configGinContactRouter(router gin.IRoutes) {
	router.POST("/notifyContact", notifyContact)
}

func notifyContact(c *gin.Context) {
	// create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(mailgunDomain, privateAPIKeyOfMailgun)

	name, _ := c.GetQuery("name")
	email, _ := c.GetQuery("email")
	feedback, _ := c.GetQuery("feedback")
	if name == "" || email == "" || feedback == "" {
		c.JSON(400, gin.H{
			"message": "invalid request",
		})
		return
	}

	subject := "New Contact Message!"
	body := fmt.Sprintf(msgTemplate, name, email, feedback)

	message := mg.NewMessage(sender, subject, body, feedbackEmail)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// send the message with a 10 second timeout
	_, _, err := mg.Send(ctx, message)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed to send email,err" + err.Error(),
		})
	}

	c.JSON(200, gin.H{})
}
