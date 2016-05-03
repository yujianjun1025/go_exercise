package main

import "github.com/cihub/seelog"

import (
	"./selflogs"
)

func main() {
	defer seelog.Flush()
	defer selflogs.Logger.Flush()
	seelog.Info("Hello from Seelog!")

	seelog.Info("main begin")
	selflogs.Logger.Error("selflog xxxx")
	seelog.Info("main end")

}
