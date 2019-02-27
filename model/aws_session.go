package model

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// AWSSession is aws session
type AWSSession struct {
	Session *session.Session
}

// GetSession is get aws session
func (aws_session *AWSSession) GetSession() {
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("Region")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AKID"), os.Getenv("SecretKey"), ""),
	})
	aws_session.Session = sess
}
