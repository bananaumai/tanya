package tail

import (
	"flag"

	"github.com/ieee0824/tanya/help"
	"github.com/ieee0824/tanya/loggroup"
)

const (
	helpMsg = ""
)

func init() {
	help.AddMsg(helpMsg)
}

func Tail(args []string) error {
	flagSet := &flag.FlagSet{}

	h := flagSet.Bool("h", false, "view help")
	hl := flagSet.Bool("help", false, "view help")
	logGroupName := flagSet.String("loggroup", "", "set log group name")

	if err := flagSet.Parse(args); err != nil {
		return err
	}

	if *h || *hl {
		help.Help()
	}

	if *logGroupName == "" {
		name, err := loggroup.NewClient().SelectorUI()
		if err != nil {
			return err
		}
		logGroupName = &name
	}

	return nil
}
