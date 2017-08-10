package sess

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	s      *session.Session
	option = struct {
		region  *string
		profile *string
	}{}
)

func roundFlags(s []string) (o []string, region string, profile string) {
	for i := 0; i < len(s); i++ {
		if s[i] == "-region" {
			region = s[i+1]
			i++
		} else if s[i] == "-profile" {
			profile = s[i+1]
			i++
		} else if strings.HasPrefix(s[i], "-region=") {
			region = strings.TrimPrefix(s[i], "-region=")
		} else if strings.HasPrefix(s[i], "-profile=") {
			profile = strings.TrimPrefix(s[i], "-profile=")

		} else {
			o = append(o, s[i])
		}
	}
	return
}

func Initial(o []string) ([]string, error) {
	o, region, profile := roundFlags(o)
	option.region = &region
	option.profile = &profile

	s = session.Must(session.NewSessionWithOptions(session.Options{
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
		SharedConfigState:       session.SharedConfigEnable,
		Profile:                 *option.profile,
		Config: aws.Config{
			Region: option.region,
		},
	}))

	return o, nil
}

func GetSession() *session.Session {
	return s
}
