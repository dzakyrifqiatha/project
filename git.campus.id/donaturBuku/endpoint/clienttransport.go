package endpoint

import (
	"context"
	"time"

	svc "MiniProject/git.campus.id/donaturBuku/server"

	pb "MiniProject/git.campus.id/donaturBuku/grpc"

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
	grpcName = "grpc.DonaturBukuService"
)

func NewGRPCDonaturBukuClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.DonaturBukuService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addDonaturBukuEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddDonaturBukuEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addDonaturBukuEp = retry
	}

	var readDonaturBukuByKd_DonaturBukuEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadDonaturBukuByKd_DonaturBukuEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readDonaturBukuByKd_DOnaturBukuEp = retry
	}

	var readDonaturBukuEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadDonaturBukuEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readDonaturBukuEp = retry
	}

	var updateDonaturBukuEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateDonaturBuku, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateDonaturBukuEp = retry
	}

	var readDonaturBukuByNama_DonaturBukuEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadDonaturBukuByNama_DonaturBuku, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readDonaturBukuByNam_DonaturBukuEp = retry
	}
	return DonaturBukuEndpoint{AddDonaturBukuEndpoint: addDonaturBukuEp, ReadDonaturBukuByKd_DonaturBukuEndpoint: readDonaturBukuByKd_DonaturBukuEp,
		ReadDonaturBukuEndpoint: readDonaturBukuEp, UpdateDonaturBukuEndpoint: updateDonaturBukuEp,
		ReadDonaturBukuByNama_DonaturBukuEndpoint: readDonaturBukuByNama_DonaturBukuEp}, nil
}

func encodeAddDonaturBukuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.DonaturBuku)
	return &pb.AddDonaturBukuReq{
		Kd_DonaturBuku:   req.Kd_DonaturBuku,
		Nama_DonaturBuku: req.Nama_DonaturBuku,
		Status:           req.Status,
	}, nil
}

func encodeReadDonaturBukuByKd_DonaturBukuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.DonaturBuku)
	return &pb.ReadDonaturBukuByKd_DonaturBukuReq{Kode: req.Kode}, nil
}

func encodeReadDonaturBukuRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateDonaturBukuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.DonaturBuku)
	return &pb.UpdateDonaturBukuReq{
		Kd_DonaturBuku:   req.Kd_DonaturBuku,
		Nama_DonaturBuku: req.Nama_DonaturBuku,
		Status:           req.Status,
	}, nil
}

func encodeReadDonaturBukuByNama_DonaturBukuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.DonaturBuku)
	return &pb.ReadDonaturBukuByNama_DonaturBukuReq{Nama: req.Nama}, nil
}

func decodeDonaturBukuResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadDonaturBukuByKd_DonaturBukuRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadDonaturBukuByKd_DonaturBukuResp)
	return svc.DonaturBuku{
		Kd_DonaturBuku:   resp.Kd_DonaturBuku,
		Nama_DonaturBuku: resp.Nama_DonaturBuku,
		Status:           resp.Status,
	}, nil
}

func decodeReadDonaturBukuResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadDonaturBukuResp)
	var rsp svc.DonaturBuku

	for _, v := range resp.AllDonaturBuku {
		itm := svc.DonaturBuku{
			Kd_DonaturBuku:   v.Kd_DonaturBuku,
			Nama_DonaturBuku: v.Nama_DonaturBuku,
			Status:           v.Status,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func decodeReadDonaturBukuByNama_DonaturBukuRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadDonaturBukuByNama_DonaturBukuResp)
	return svc.DonaturBuku{
		Kd_DonaturBuku:   resp.Kd_DonaturBuku,
		Nama_DonaturBuku: resp.Nama_DonaturBuku,
		Status:           resp.Status,
	}, nil
}

func makeClientAddDonaturBukuEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddDonaturBuku",
		encodeAddDonaturBukuRequest,
		decodeAddDonaturBukuResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddDonaturBuku")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddDonaturBuku",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadDonaturBukuByKd_DonaturBukuEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadDonaturBukuByKd_DonaturBuku",
		encodeReadDonaturBukuByKd_DonaturBukuRequest,
		decodeReadDonaturBukuByKd_DonaturBukuRespones,
		pb.ReadDonaturBukuByKd_DonaturBukuResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadDonaturBukuByKd_DonaturBuku")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadDonaturBukuByKd_DonaturBuku",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadDonaturBukuEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadDonaturBuku",
		encodeReadDonaturBukuRequest,
		decodeReadDonaturBukuResponse,
		pb.ReadDonaturBukuResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadDonaturBuku")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadDonaturBuku",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateDonaturBuku(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateDonaturBuku",
		encodeUpdateDonaturBukuRequest,
		decodeUpdateDonaturBukuResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateDonaturBuku")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateDonaturBuku",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadDonaturBukuByNama_DonaturBuku(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadDonaturBukuByNama_DonaturBuku",
		encodeReadDonaturBukuByNama_DonaturBukuRequest,
		decodeReadDonaturBukuByNama_DonaturBukuRespones,
		pb.ReadDonaturBukuByNama_DonaturBukuResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadDonaturBukuByNama_DonaturBuku")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadDonaturBukuByNama_DonaturBuku",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
