package model

import (
	"context"
	"os"
	"strings"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/mail"
)

// SENDKEY is for queue to send mail
const SENDKEY = "SEND_KEY"

// Mail for send mail
type Mail struct {
	Ctx      context.Context
	HTMLBody string
}

// Send will send mail
func (m *Mail) Send() {
	ctx := m.Ctx
	msg := &mail.Message{
		Sender:   "noreply@" + appengine.AppID(ctx) + ".appspotmail.com",
		Bcc:      strings.Split(os.Getenv("TO"), ","),
		Subject:  os.Getenv("Subject"),
		HTMLBody: m.HTMLBody,
	}
	err := mail.Send(ctx, msg)
	if err != nil {
		log.Errorf(ctx, "sent mail failed: %s", err.Error())
	}
}
