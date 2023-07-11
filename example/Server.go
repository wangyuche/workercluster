package main

import (
	"os"

	"github.com/wangyuche/goutils/log"
	"github.com/wangyuche/workercluster"
)

func main() {
	log.New(log.LogType(os.Getenv("LogType")))
	w := workercluster.New()
	w.LeaderElection("aaa", "default", os.Getenv("HOSTNAME"))
	ch := make(chan os.Signal, 1)
	<-ch
}
