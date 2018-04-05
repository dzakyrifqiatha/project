package endpoint

import (
	"context"

	scv "MiniProject/git.campus.id/buku2/server"

	pb "MiniProject/git.campus.id/buku2/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcBukuServer struct {
	addBuku              grpctransport.Handler
	readBukuByKode       grpctransport.Handler
	readBuku             grpctransport.Handler
	updateBuku           grpctransport.Handler
	readBukuByKeterangan grpctransport.Handler
}

func NewGRPCBukuServer(endpoints BukuEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.BukuServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcBukuServer{
		addBuku: grpctransport.NewServer(endpoints.AddBukuEndpoint,
			decodeAddBukuRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddBuku", logger)))...),
		readBukuByKode: grpctransport.NewServer(endpoints.ReadBukuByKodeEndpoint,
			decodeReadBukuByKodeRequest,
			encodeReadBukuByKodeResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadBukuByKode", logger)))...),
		readBuku: grpctransport.NewServer(endpoints.ReadBukuEndpoint,
			decodeReadBukuRequest,
			encodeReadBukuResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadBuku", logger)))...),
		updateBuku: grpctransport.NewServer(endpoints.UpdateBukuEndpoint,
			decodeUpdateBukuRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateBuku", logger)))...),
		readBukuByKeterangan: grpctransport.NewServer(endpoints.ReadBukuByKeteranganEndpoint,
			decodeReadBukuByKeteranganRequest,
			encodeReadBukuByKeteranganResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadBukuByKeterangan", logger)))...),
	}
}

func decodeAddBukuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddBukuReq)
	return scv.Buku{KodeBuku: req.GetKodeBuku(), NamaBuku: req.GetNamaBuku(),
		Status: req.GetStatus(), Keterangan: req.GetKeterangan()}, nil
}

func decodeReadBukuByKodeRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadBukuByKodeReq)
	return scv.Buku{KodeBuku: req.KodeBuku}, nil
}

func decodeReadBukuRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeUpdateBukuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateBukuReq)
	return scv.Buku{KodeBuku: req.KodeBuku, NamaBuku: req.NamaBuku, Status: req.Status, Keterangan: req.Keterangan}, nil
}

func decodeReadBukuByKeteranganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadBukuByKeteranganReq)
	return scv.Buku{Keterangan: req.Keterangan}, nil

}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadBukuByKodeResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Buku)
	return &pb.ReadBukuByKodeResp{KodeBuku: resp.KodeBuku, NamaBuku: resp.NamaBuku, Status: resp.Status, Keterangan: resp.Keterangan}, nil
}

func encodeReadBukuResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Bukus)

	rsp := &pb.ReadBukuResp{}

	for _, v := range resp {
		itm := &pb.ReadBukuByKodeResp{
			KodeBuku:   v.KodeBuku,
			NamaBuku:   v.NamaBuku,
			Status:     v.Status,
			Keterangan: v.Keterangan,
		}
		rsp.AllBuku = append(rsp.AllBuku, itm)
	}
	return rsp, nil
}

func encodeReadBukuByKeteranganResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Bukus)

	rsp := &pb.ReadBukuByKeteranganResp{}

	for _, v := range resp {
		itm := &pb.ReadBukuByKodeResp{
			KodeBuku:   v.KodeBuku,
			NamaBuku:   v.NamaBuku,
			Status:     v.Status,
			Keterangan: v.Keterangan,
		}
		rsp.AllBukuKet = append(rsp.AllBukuKet, itm)
	}
	return rsp, nil
}

func (s *grpcBukuServer) AddBuku(ctx oldcontext.Context, buku *pb.AddBukuReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addBuku.ServeGRPC(ctx, buku)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcBukuServer) ReadBukuByKode(ctx oldcontext.Context, kode *pb.ReadBukuByKodeReq) (*pb.ReadBukuByKodeResp, error) {
	_, resp, err := s.readBukuByKode.ServeGRPC(ctx, kode)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadBukuByKodeResp), nil
}

func (s *grpcBukuServer) ReadBuku(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadBukuResp, error) {
	_, resp, err := s.readBuku.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadBukuResp), nil
}

func (s *grpcBukuServer) UpdateBuku(ctx oldcontext.Context, buku *pb.UpdateBukuReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateBuku.ServeGRPC(ctx, buku)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcBukuServer) ReadBukuByKeterangan(ctx oldcontext.Context, keterangan *pb.ReadBukuByKeteranganReq) (*pb.ReadBukuByKeteranganResp, error) {
	_, resp, err := s.readBukuByKeterangan.ServeGRPC(ctx, keterangan)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadBukuByKeteranganResp), nil
}
