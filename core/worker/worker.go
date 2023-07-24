package worker

import (
	"context"
	"io"
	"net"
	"sync"

	"github.com/wangyuche/goutils/log"
	wk "github.com/wangyuche/workercluster/core/worker/protos"
	"google.golang.org/grpc"
)

type Server struct {
	wk.UnimplementedWorkerGRPCServer
	workers map[*WorkerInfo]wk.WorkerStatus
	m       *sync.RWMutex
	cb      iWorkerServerCallBack
}

type WorkerInfo struct {
	Worker wk.WorkerGRPC_StreamWorkerServer
}

type iWorkerServerCallBack interface {
	WorkerIdleEvent(*WorkerInfo)
	MasterRecv([]byte)
}

func NewServer(port string, cb iWorkerServerCallBack) *Server {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fail(err.Error())
	}
	grpcServer := grpc.NewServer()
	instace := &Server{workers: make(map[*WorkerInfo]wk.WorkerStatus), m: new(sync.RWMutex), cb: cb}
	wk.RegisterWorkerGRPCServer(grpcServer, instace)
	go func() {
		log.Info("Listen Port:" + port)
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Fail(err.Error())
		}
	}()
	return instace
}

func (s *Server) StreamWorker(stream wk.WorkerGRPC_StreamWorkerServer) error {
	s.m.RLock()
	c := &WorkerInfo{Worker: stream}
	s.workers[c] = wk.WorkerStatus_idle
	s.m.RUnlock()
	for {
		data, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			s.m.RLock()
			delete(s.workers, c)
			s.m.RUnlock()
			log.Error(err.Error())
			return err
		}

		switch data.Event {
		case wk.WorkerEvent_getstatus:
			s.workers[c] = data.Status
			if data.Status == wk.WorkerStatus_idle {
				s.cb.WorkerIdleEvent(c)
			}
			break
		case wk.WorkerEvent_dispatch:
			s.cb.MasterRecv(data.Data)
			break
		default:
			break
		}
	}
}

func (s *Server) GetWorkers() map[*WorkerInfo]wk.WorkerStatus {
	return s.workers
}

func (wi *WorkerInfo) Send(data []byte) {
	wi.Worker.Send(&wk.WorkerReqStruct{Event: wk.WorkerEvent_dispatch, Data: data})
}

type Worker struct {
	workercancel context.CancelFunc
	stream       wk.WorkerGRPC_StreamWorkerClient
	status       wk.WorkerStatus
	cb           iWorkerClientCallBack
}

type iWorkerClientCallBack interface {
	WorkerRecv([]byte)
}

func NewWorker(host string, cb iWorkerClientCallBack) *Worker {
	w := &Worker{}
	conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fail(err.Error())
	}
	wkclient := wk.NewWorkerGRPCClient(conn)
	stream, err := wkclient.StreamWorker(context.Background())
	if err != nil {
		log.Fail(err.Error())
	}
	w.cb = cb
	w.stream = stream
	return w
}

func (this *Worker) Run() {
	go func() {
		for {
			data, err := this.stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Error(err.Error())
				break
			}
			switch data.Event {
			case wk.WorkerEvent_getstatus:
				this.stream.Send(&wk.WorkerResStruct{Event: wk.WorkerEvent_getstatus, Status: this.status})
				break
			case wk.WorkerEvent_dispatch:
				this.cb.WorkerRecv(data.Data)
				break
			default:
				break
			}
		}
	}()
}

func (this *Worker) ChangeStatusIdle() {
	this.status = wk.WorkerStatus_idle
	this.stream.Send(&wk.WorkerResStruct{Event: wk.WorkerEvent_getstatus, Status: this.status})
}

func (this *Worker) ChangeStatusBusy() {
	this.status = wk.WorkerStatus_busy
}

func (this *Worker) Send(data []byte) {
	this.stream.Send(&wk.WorkerResStruct{Event: wk.WorkerEvent_dispatch, Data: data})
}
