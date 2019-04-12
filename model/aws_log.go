package model

import (
	"strings"

	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

// AWSLog is aws cloud watch
type AWSLog struct {
	IgnoreChars      []string
	Cloudwatchlogs   *cloudwatchlogs.CloudWatchLogs
	Limit            *int64
	ConLogGroupName  *string
	ConLogStreamName *string
	StartTimestamp   *int64
	EndTimestamp     *int64
}

// ExecLogs is check log need to send or not
func (awsLog *AWSLog) ExecLogs() (htmlBody string, count int) {
	getLogEventsInput := cloudwatchlogs.GetLogEventsInput{
		LogGroupName:  awsLog.ConLogGroupName,
		LogStreamName: awsLog.ConLogStreamName,
		StartTime:     awsLog.StartTimestamp,
		EndTime:       awsLog.EndTimestamp,
	}
	out, _ := awsLog.Cloudwatchlogs.GetLogEvents(&getLogEventsInput)

	count = 0
	for _, event := range out.Events {
		if awsLog.isError(*event.Message) {
			if awsLog.isIgnore(*event.Message) {
				continue
			}
			htmlBody += *event.Message + "<br/>"
			count++
		}
	}
	return
}

func (awsLog *AWSLog) isError(eventLog string) bool {
	return strings.Index(eventLog, "E:  ") >= 0 || strings.Index(strings.ToLower(eventLog), "error") >= 0
}

func (awsLog *AWSLog) isIgnore(eventLog string) (ans bool) {
	for _, chars := range awsLog.IgnoreChars {
		ans = strings.Index(eventLog, chars) > 0
		if ans {
			break
		}
	}

	return
}
