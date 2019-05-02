package main

import (
	"flag"

	"github.com/jskcnsl/tower/command"
	_ "github.com/jskcnsl/tower/help"
)

func main() {
	flag.Parse()
	cmd := flag.Args()
	if len(cmd) == 0 {
		if err := command.Get("help").Handle(); err != nil {
			panic(err)
		}
	} else if c := command.Get(cmd[1]); c == nil {
		if err := command.Get("help").Handle(); err != nil {
			panic(err)
		}
	} else {
		if err := c.Handle(); err != nil {
			panic(err)
		}
	}
}
