package main

import (
	"os"

	"github.com/wangyuche/goutils/log"
	"github.com/wangyuche/workercluster"
	"github.com/wangyuche/workercluster/core/worker"
)

func main() {
	log.New(log.LogType(os.Getenv("LogType")))
	w := workercluster.New(os.Getenv("K8SCONFIG"))
	w.Run("aaa", "default", os.Getenv("HOSTNAME"), "8888", &WorkerServerCallBack{})
	ch := make(chan os.Signal, 1)
	<-ch
}

type WorkerServerCallBack struct {
}

func (this *WorkerServerCallBack) WorkerRecv(data []byte) {

}

func (this *WorkerServerCallBack) WorkerIdleEvent(server *worker.WorkerInfo) {

}

func (this *WorkerServerCallBack) MasterRecv(data []byte) {

}

func (this *WorkerServerCallBack) BeMaster() {
}
