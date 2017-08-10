package help

import (
	"fmt"
	"os"
)

var helpMsgs = []string{}

func AddMsg(s string) {
	helpMsgs = append(helpMsgs, s)
	helpMsgs = append(helpMsgs, "=====================")
}

func Help() {
	for _, msg := range helpMsgs {
		fmt.Println(msg)
	}
	os.Exit(1)
}
