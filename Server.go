package workercluster

import (
	"context"

	"github.com/wangyuche/workercluster/core/k8s"
	"github.com/wangyuche/workercluster/core/leader"
	"github.com/wangyuche/workercluster/core/worker"
	wk "github.com/wangyuche/workercluster/core/worker/protos"
	clientset "k8s.io/client-go/kubernetes"
)

type WorkerCluster struct {
	k8s            *clientset.Clientset
	leader         *leader.LeaderElection
	leadercancel   context.CancelFunc
	bemastercb     chan bool
	masterchangecb chan string
	wkserver       *worker.Server
	stream         chan wk.WorkerGRPC_StreamWorkerServer
	cb             iWorkerCallBack
	worker         *worker.Worker
}

type iWorkerCallBack interface {
	BeMaster()
	WorkerIdleEvent(*worker.WorkerInfo)
	MasterRecv([]byte)
	WorkerRecv([]byte)
}

func New(path string) *WorkerCluster {
	instance := &WorkerCluster{}
	instance.connectk8s(path)
	return instance
}

func (this *WorkerCluster) connectk8s(path string) {
	this.k8s = k8s.ConnectKubernetes(path)
}

func (this *WorkerCluster) Run(name string, namespace string, id string, port string, cb iWorkerCallBack) {
	ctx, cancel := context.WithCancel(context.Background())
	this.wkserver = worker.NewServer(port, cb)
	this.leadercancel = cancel
	this.bemastercb = make(chan bool)
	this.masterchangecb = make(chan string)
	this.cb = cb
	go this.leaderElectionEvent()
	l := leader.New(this.k8s, ctx, name, namespace, id, &LeaderCallBack{bemastercb: this.bemastercb, masterchangecb: this.masterchangecb})
	this.leader = l
	go l.LeaderElection()
}

func (this *WorkerCluster) leaderElectionEvent() {
	for {
		select {
		case <-this.bemastercb:
			this.cb.BeMaster()
			break
		case m := <-this.masterchangecb:
			this.worker = worker.NewWorker(m, this.cb)
			this.worker.Run()
			break
		default:
		}
	}
}

func (this *WorkerCluster) LeaderCancel() {
	this.LeaderCancel()
}

func (this *WorkerCluster) GetWorkers() map[*worker.WorkerInfo]wk.WorkerStatus {
	return this.wkserver.GetWorkers()
}

func (this *WorkerCluster) ChangeStatusIdle() {
	this.worker.ChangeStatusIdle()
}

func (this *WorkerCluster) ChangeStatusBusy() {
	this.worker.ChangeStatusBusy()
}

func (this *WorkerCluster) GetWorker() *worker.Worker {
	return this.worker
}

type LeaderCallBack struct {
	bemastercb     chan bool
	masterchangecb chan string
}

func (this *LeaderCallBack) BeMaster() {
	this.bemastercb <- true
}

func (this *LeaderCallBack) MasterChange(master string) {
	this.masterchangecb <- master
}
