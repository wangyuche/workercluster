package workercluster

import (
	"context"
	"os"

	"github.com/wangyuche/workercluster/core/k8s"
	"github.com/wangyuche/workercluster/core/leader"
	clientset "k8s.io/client-go/kubernetes"
)

type WorkerCluster struct {
	k8s          *clientset.Clientset
	leader       *leader.LeaderElection
	leadercancel context.CancelFunc
}

func New() *WorkerCluster {
	instance := &WorkerCluster{}
	instance.connectk8s()

	return instance

}

func (this *WorkerCluster) connectk8s() {
	this.k8s = k8s.ConnectKubernetes(os.Getenv("K8SCONFIG"))
}

func (this *WorkerCluster) LeaderElection(name string, namespace string, id string) {
	ctx, cancel := context.WithCancel(context.Background())
	this.leadercancel = cancel
	l := leader.New(this.k8s, ctx, name, namespace, id, &LeaderCallBack{})
	this.leader = l
	go l.LeaderElection()
}

func (this *WorkerCluster) LeaderCancel() {
	this.LeaderCancel()
}

type LeaderCallBack struct {
}

func (this *LeaderCallBack) BeMaster() {

}

func (this *LeaderCallBack) MasterChange(master string) {

}
