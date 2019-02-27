package controller

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/cloudmile/gae_alert_aws_dms_log/model"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/taskqueue"
)

// LogHandle is GET '/log'
func LogHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/log" {
		return
	}
	ctx := appengine.NewContext(r)
	t := taskqueue.NewPOSTTask("/queue/log", nil)
	if _, err := taskqueue.Add(ctx, t, "get-aws-dms-log"); err != nil {
		return
	}
}

// QueueLogHandle is POST '/queue/log'
func QueueLogHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/queue/log" {
		return
	}
	ctx := appengine.NewContext(r)
	send, ok := needCheck(ctx)
	if !ok {
		return
	}

	logLimit, _ := strconv.Atoi(os.Getenv("LogLimit"))
	conLogGroupName := os.Getenv("LogGroupName")
	conLogStreamName := os.Getenv("LogStreamName")
	var limit int64
	limit = int64(logLimit)

	var awsSession model.AWSSession
	awsSession.GetSession()
	svc := cloudwatchlogs.New(awsSession.Session)
	getLogEventsInput := cloudwatchlogs.GetLogEventsInput{
		Limit:         &limit,
		LogGroupName:  &conLogGroupName,
		LogStreamName: &conLogStreamName,
	}
	out, _ := svc.GetLogEvents(&getLogEventsInput)

	count := 0
	var htmlBody string
	for _, event := range out.Events {
		if strings.Index(strings.ToLower(*event.Message), "error") >= 0 {
			htmlBody += *event.Message + "<br/>"
			count++
		}
	}
	if count > 0 {
		sendMail(ctx, htmlBody)
		send.AddAndPut(ctx)
	}
	log.Infof(ctx, "send count is %d", send.Count)
}

func needCheck(ctx context.Context) (model.Send, bool) {
	var send model.Send
	send.Get(ctx)
	countLimit, _ := strconv.Atoi(os.Getenv("CountLimit"))
	if send.Count > countLimit {
		return send, false
	}
	return send, true
}

func sendMail(ctx context.Context, bodyMessage string) {
	mail := model.Mail{Ctx: ctx, HTMLBody: bodyMessage}
	mail.Send()
}
