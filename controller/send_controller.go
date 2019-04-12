package controller

import (
	"net/http"

	"github.com/cloudmile/gae_alert_aws_dms_log/model"
	"google.golang.org/appengine"
)

// SendHandle is POST '/reset_send'
func SendHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/reset_send" {
		return
	}
	if r.Method != http.MethodPost {
		return
	}
	ctx := appengine.NewContext(r)
	var send model.Send
	send.Reset(ctx)
}

// InitHandle is POST '/init' for setup ignore chars into firestore
func InitHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/init" {
		return
	}
	if r.Method != http.MethodPost {
		return
	}
	ctx := appengine.NewContext(r)
	ignoreChars := model.IgnoreChars{
		Chars: []string{"MySQL server has gone away", "I:  Error executing data handler", "was detached because of recoverable error. Will try to reattach", "I:  Connection error"},
	}
	ignoreChars.AddAndPut(ctx)
}
