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
