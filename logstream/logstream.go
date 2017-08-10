package logstream

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs/cloudwatchlogsiface"
	"github.com/ieee0824/tanya/sess"
)

var latestStreamCache string

type LogStreamClient struct {
	svc          cloudwatchlogsiface.CloudWatchLogsAPI
	LogGroupName string
}

func NewClient(group string) *LogStreamClient {
	return &LogStreamClient{
		cloudwatchlogs.New(sess.GetSession()),
		group,
	}
}

func (c *LogStreamClient) LatestStream() (string, error) {
	list, err := c.List()
	if err != nil {
		return "", err
	}
	return list[0], nil
}

func (c *LogStreamClient) List() ([]string, error) {
	var ret = []string{}

	result, err := c.svc.DescribeLogStreams(&cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName: &c.LogGroupName,
		OrderBy:      aws.String("LastEventTime"),
		Descending:   aws.Bool(true),
		Limit:        aws.Int64(5),
	})
	if err != nil {
		return nil, err
	}

	for _, stream := range result.LogStreams {
		ret = append(ret, *stream.LogStreamName)
	}

	return ret, nil
}

func (c *LogStreamClient) RangeList(to, from time.Time) ([]string, error) {
	params := &cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName: &c.LogGroupName,
	}
	ret := []string{}

	err := c.svc.DescribeLogStreamsPages(params, func(page *cloudwatchlogs.DescribeLogStreamsOutput, lastPage bool) bool {
		for _, stream := range page.LogStreams {
			if (aws.TimeUnixMilli(to) < *stream.FirstEventTimestamp && *stream.FirstEventTimestamp < aws.TimeUnixMilli(from)) || (aws.TimeUnixMilli(to) < *stream.LastEventTimestamp && *stream.LastEventTimestamp < aws.TimeUnixMilli(from)) {
				ret = append(ret, *stream.LogStreamName)
			}
		}
		return !lastPage
	})
	return ret, err
}
