package main

import (
	"net/http"

	"github.com/cloudmile/gae_alert_aws_dms_log/controller"
	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/log", controller.LogHandle)
	http.HandleFunc("/queue/log", controller.QueueLogHandle)

	http.HandleFunc("/reset_send", controller.SendHandle)
	// http.HandleFunc("/init", controller.InitHandle)

	appengine.Main()
}
