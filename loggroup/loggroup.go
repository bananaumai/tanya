package loggroup

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs/cloudwatchlogsiface"
	"github.com/ieee0824/tanya/selector"
	"github.com/ieee0824/tanya/sess"
)

type LogGroupClient struct {
	svc cloudwatchlogsiface.CloudWatchLogsAPI
}

func NewClient() *LogGroupClient {
	return &LogGroupClient{
		cloudwatchlogs.New(sess.GetSession()),
	}
}

func (c *LogGroupClient) List() ([]string, error) {
	var ret = []string{}
	result, err := c.svc.DescribeLogGroups(&cloudwatchlogs.DescribeLogGroupsInput{})
	if err != nil {
		return nil, err
	}

	for _, logGroup := range result.LogGroups {
		ret = append(ret, *logGroup.LogGroupName)
	}

	return ret, nil
}

func (c LogGroupClient) String() string {
	list, err := c.List()
	if err != nil {
		return ""
	}

	return strings.Join(list, "\n")
}

func (c *LogGroupClient) ListCmd(args []string) error {
	list, err := c.List()
	if err != nil {
		return err
	}

	fmt.Println(strings.Join(list, "\n"))

	return nil
}

func (c *LogGroupClient) SelectorUI() (string, error) {
	list, err := c.List()
	if err != nil {
		return "", err
	}

	return selector.New(list).RunSelector()

	return "", nil
}
