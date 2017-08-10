package option

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/ieee0824/tanya/loggroup"
	"github.com/ieee0824/tanya/tail"
)

type Option struct {
	Main      *string  `json:"main,omitempty"`
	SubOption []string `json:"sub_option,omitempty"`
}

func (o Option) String() string {
	bin, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return ""
	}
	return string(bin)
}

func ParseArgs(args []string) *Option {
	if len(args) == 1 {
		return &Option{}
	}

	if strings.HasPrefix(args[1], "-") {
		return &Option{
			Main:      nil,
			SubOption: args[1:],
		}
	}

	return &Option{
		Main:      &args[1],
		SubOption: args[2:],
	}
}

func (o *Option) Exec() error {
	if o.Main == nil {
		return tail.Tail(o.SubOption)
	}

	switch *o.Main {
	case "tail":
		return tail.Tail(o.SubOption)
	case "list":
		return loggroup.NewClient().ListCmd(o.SubOption)
	}

	return errors.New("no option")
}
