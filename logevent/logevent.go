package logevent

import (
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs/cloudwatchlogsiface"
	"github.com/ieee0824/tanya/sess"
)

type LogEventClient struct {
	svc cloudwatchlogsiface.CloudWatchLogsAPI
}

func NewClient(group string) *LogStreamClient {
	return &LogEventClient{
		cloudwatchlogs.New(sess.GetSession()),
	}
}
