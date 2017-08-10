package main

import (
	"os"

	"github.com/ieee0824/tanya/option"
	"github.com/ieee0824/tanya/sess"
)

func main() {
	args, err := sess.Initial(os.Args)
	if err != nil {
		panic(err)
	}

	o := option.ParseArgs(args)
	if o == nil {
		panic("options is nil")
	}

	if err := o.Exec(); err != nil {
		panic(err)
	}
}
