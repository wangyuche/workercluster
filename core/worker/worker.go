package worker

import (
	"context"
	"io"
	"net"
	"sync"
	"time"

	"github.com/wangyuche/goutils/log"
	wk "github.com/wangyuche/workercluster/core/worker/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
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
	s.m.Lock()
	c := &WorkerInfo{Worker: stream}
	s.workers[c] = wk.WorkerStatus_idle
	s.m.Unlock()
	for {
		data, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			s.m.Lock()
			delete(s.workers, c)
			s.m.Unlock()
			log.Error(err.Error())
			return err
		}

		switch data.Event {
		case wk.WorkerEvent_getstatus:
			s.m.Lock()
			s.workers[c] = data.Status
			s.m.Unlock()
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
	s.m.RLock()
	defer s.m.RUnlock()
	var d map[*WorkerInfo]wk.WorkerStatus = make(map[*WorkerInfo]wk.WorkerStatus)
	for k, v := range s.workers {
		d[k] = v
	}
	return d
}

func (wi *WorkerInfo) Send(data []byte) {
	wi.Worker.Send(&wk.WorkerReqStruct{Event: wk.WorkerEvent_dispatch, Data: data})
}

type Worker struct {
	workercancel context.CancelFunc
	stream       wk.WorkerGRPC_StreamWorkerClient
	status       wk.WorkerStatus
	cb           iWorkerClientCallBack
	conn         *grpc.ClientConn
	done         chan bool
	reconnect    chan bool
}

type iWorkerClientCallBack interface {
	WorkerRecv([]byte)
}

func NewWorker(host string, cb iWorkerClientCallBack) *Worker {
	w := &Worker{}
	conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		log.Fail(err.Error())
	}
	wkclient := wk.NewWorkerGRPCClient(conn)
	stream, err := wkclient.StreamWorker(context.Background())
	if err != nil {
		log.Fail(err.Error())
	}
	w.conn = conn
	w.cb = cb
	w.stream = stream
	return w
}

func (this *Worker) waitUntilReady() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	return this.conn.WaitForStateChange(ctx, connectivity.Ready)
}

func (this *Worker) process() {
	go func() {
		for {
			data, err := this.stream.Recv()
			if err == io.EOF {
				this.done <- true
				return
			}
			if err != nil {
				log.Error(err.Error())
				this.reconnect <- true
				return
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

func (this *Worker) Run() {
	go this.process()
	go func() {
		select {
		case <-this.reconnect:
			if !this.waitUntilReady() {
				log.Fail("ReConnect Fail")
			}
			go this.process()
		case <-this.done:
			return
		}
	}()
}

func (this *Worker) ChangeStatusIdle() {
	this.status = wk.WorkerStatus_idle
	this.stream.Send(&wk.WorkerResStruct{Event: wk.WorkerEvent_getstatus, Status: this.status})
}

func (this *Worker) ChangeStatusBusy() {
	this.status = wk.WorkerStatus_busy
	this.stream.Send(&wk.WorkerResStruct{Event: wk.WorkerEvent_getstatus, Status: this.status})
}

func (this *Worker) Send(data []byte) {
	this.stream.Send(&wk.WorkerResStruct{Event: wk.WorkerEvent_dispatch, Data: data})
}
