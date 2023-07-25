package main

import (
	"os"
	"strconv"
	"time"

	"github.com/wangyuche/goutils/log"
	"github.com/wangyuche/workercluster"
	"github.com/wangyuche/workercluster/core/worker"
	wk "github.com/wangyuche/workercluster/core/worker/protos"
)

var workcluster *workercluster.WorkerCluster

func main() {
	log.New(log.LogType(os.Getenv("LogType")))
	workcluster = workercluster.New(os.Getenv("K8SCONFIG"))
	workcluster.Run("aaa", "default", os.Getenv("HOSTNAME")+":"+os.Getenv("PORT"), os.Getenv("PORT"), &WorkerServerCallBack{})
	time.Sleep(5 * time.Second)
	/*
		wks := workcluster.GetWorkers()
		for worker, status := range wks {
			if status == wk.WorkerStatus_idle {
				worker.Send([]byte(strconv.Itoa(count)))
				count = count + skip
				time.Sleep(2 * time.Second)
			}
		}
	*/
	ch := make(chan os.Signal, 1)
	<-ch
}

type WorkerServerCallBack struct {
}

func (this *WorkerServerCallBack) WorkerRecv(data []byte) {
	workcluster.ChangeStatusBusy()
	log.Debug(string(data))
	time.Sleep(5 * time.Second)
	workcluster.GetWorker().Send([]byte("OK~"))
	workcluster.ChangeStatusIdle()
}

func (this *WorkerServerCallBack) WorkerIdleEvent(server *worker.WorkerInfo) {
	/*
		log.Debug("idle")
		server.Send([]byte(strconv.Itoa(count)))
		count = count + skip
		time.Sleep(2 * time.Second)
	*/
}

func (this *WorkerServerCallBack) MasterRecv(data []byte) {
	log.Debug(string(data))
}

var count int = 0
var skip int = 5

func (this *WorkerServerCallBack) BeMaster() {
	time.Sleep(5 * time.Second)
	go func() {
		for {
			wks := workcluster.GetWorkers()
			if wks == nil {
				continue
			}
			for worker, status := range wks {
				if status == wk.WorkerStatus_idle {
					log.Debug(strconv.Itoa(count))
					worker.Send([]byte(strconv.Itoa(count)))
					count = count + skip
					time.Sleep(2 * time.Second)
				}
			}
		}
	}()
}
