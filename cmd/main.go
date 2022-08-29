package main

import (
	"github.com/robinlieb/validate-csv/foundation/helper"
)

var version = "develop"

func main() {
	if err := CmdInit(version); err != nil {
		helper.PrintErrorAndExit(err)
	}
}
