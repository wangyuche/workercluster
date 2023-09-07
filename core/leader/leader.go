package leader

import (
	"context"
	"os"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
)

func New(client *clientset.Clientset, ctx context.Context, name string, namespace string, id string, cb iLeaderCallBack) *LeaderElection {
	return &LeaderElection{client, ctx, name, namespace, id, cb}
}

type LeaderCallBack struct {
}

type iLeaderCallBack interface {
	BeMaster()
	MasterChange(string)
}

type LeaderElection struct {
	client    *clientset.Clientset
	ctx       context.Context
	name      string
	namespace string
	id        string
	cb        iLeaderCallBack
}

func (this *LeaderElection) LeaderElection() {

	lock := &resourcelock.LeaseLock{
		LeaseMeta: metav1.ObjectMeta{
			Name:      this.name,
			Namespace: this.namespace,
		},
		Client: this.client.CoordinationV1(),
		LockConfig: resourcelock.ResourceLockConfig{
			Identity: this.id,
		},
	}
	leaderelection.RunOrDie(this.ctx, leaderelection.LeaderElectionConfig{
		Lock:            lock,
		ReleaseOnCancel: true,
		LeaseDuration:   3 * time.Second,
		RenewDeadline:   2 * time.Second,
		RetryPeriod:     1 * time.Second,
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: func(ctx context.Context) {
				this.cb.BeMaster()
			},
			OnStoppedLeading: func() {
				os.Exit(0)
			},
			OnNewLeader: func(identity string) {
				if identity == this.id {
					return
				}
				this.cb.MasterChange(identity)
			},
		},
	})
}
