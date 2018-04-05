package endpoint

import (
	"context"
	"time"

	svc "MiniProject/git.campus.id/buku2/server"

	pb "MiniProject/git.campus.id/buku2/grpc"

	util "MiniProject/git.campus.id/util/grpc"
	disc "MiniProject/git.campus.id/util/microservice"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	grpcName = "grpc.BukuService"
)

func NewGRPCBukuClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.BukuService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addBukuEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddBukuEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addBukuEp = retry
	}

	var readBukuByKodeEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadBukuByKodeEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readBukuByKodeEp = retry
	}

	var readBukuEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadBukuEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readBukuEp = retry
	}

	var updateBukuEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateBuku, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateBukuEp = retry
	}

	var readBukuByKeteranganEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadBukuByKeterangan, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readBukuByKeteranganEp = retry
	}
	return BukuEndpoint{AddBukuEndpoint: addBukuEp, ReadBukuByKodeEndpoint: readBukuByKodeEp,
		ReadBukuEndpoint: readBukuEp, UpdateBukuEndpoint: updateBukuEp,
		ReadBukuByKeteranganEndpoint: readBukuByKeteranganEp}, nil
}

func encodeAddBukuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Buku)
	return &pb.AddBukuReq{
		KodeBuku:   req.KodeBuku,
		NamaBuku:   req.NamaBuku,
		Status:     req.Status,
		Keterangan: req.Keterangan,
	}, nil
}

func encodeReadBukuByKodeRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Buku)
	return &pb.ReadBukuByKodeReq{KodeBuku: req.KodeBuku}, nil
}

func encodeReadBukuRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateBukuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Buku)
	return &pb.UpdateBukuReq{
		KodeBuku:   req.KodeBuku,
		NamaBuku:   req.NamaBuku,
		Status:     req.Status,
		Keterangan: req.Keterangan,
	}, nil
}

func encodeReadBukuByKeteranganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Buku)
	return &pb.ReadBukuByKeteranganReq{Keterangan: req.Keterangan}, nil
}

func decodeBukuResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadBukuByKodeRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadBukuByKodeResp)
	return svc.Buku{
		KodeBuku:   resp.KodeBuku,
		NamaBuku:   resp.NamaBuku,
		Status:     resp.Status,
		Keterangan: resp.Keterangan,
	}, nil
}

func decodeReadBukuResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadBukuResp)
	var rsp svc.Bukus

	for _, v := range resp.AllBuku {
		itm := svc.Buku{
			KodeBuku:   v.KodeBuku,
			NamaBuku:   v.NamaBuku,
			Status:     v.Status,
			Keterangan: v.Keterangan,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func decodeReadBukuByKeteranganRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadBukuByKeteranganResp)
	var rsp svc.Bukus

	for _, v := range resp.AllBukuKet {
		itm := svc.Buku{
			KodeBuku:   v.KodeBuku,
			NamaBuku:   v.NamaBuku,
			Status:     v.Status,
			Keterangan: v.Keterangan,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func makeClientAddBukuEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddBuku",
		encodeAddBukuRequest,
		decodeBukuResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddBuku")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddBuku",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadBukuByKodeEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadBukuByKode",
		encodeReadBukuByKodeRequest,
		decodeReadBukuByKodeRespones,
		pb.ReadBukuByKodeResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadBukuByKode")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadBukuByKode",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadBukuEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadBuku",
		encodeReadBukuRequest,
		decodeReadBukuResponse,
		pb.ReadBukuResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadBuku")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadBuku",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateBuku(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateBuku",
		encodeUpdateBukuRequest,
		decodeBukuResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateBuku")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateBuku",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadBukuByKeterangan(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadBukuByKeterangan",
		encodeReadBukuByKeteranganRequest,
		decodeReadBukuByKeteranganRespones,
		pb.ReadBukuByKeteranganResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadBukuByKeterangan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadBukuByKeterangan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
