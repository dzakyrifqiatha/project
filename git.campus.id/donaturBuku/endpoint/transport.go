package endpoint

import (
	"context"

	scv "MiniProject/git.campus.id/donaturBuku/server"

	pb "MiniProject/git.campus.id/donaturBuku/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcDonaturBukuServer struct {
	addDonaturBuku                  grpctransport.Handler
	readDonaturBukuByKd_DonaturBuku grpctransport.Handler
	readDonaturBuku                 grpctransport.Handler
	updateDonaturBuku               grpctransport.Handler
	readDonaturByNama_DonaturBuku   grpctransport.Handler
}

func NewGRPCDonaturBukuServer(endpoints DonaturBukuEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.DonaturBukuServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcDonaturBukuServer{
		addDonaturBuku: grpctransport.NewServer(endpoints.AddDonaturBukuEndpoint,
			decodeAddDonaturBukuRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddDonaturBuku", logger)))...),
		readDonaturBukuByKd_DonaturBuku: grpctransport.NewServer(endpoints.ReadDonaturBukuByKd_DonaturBukuEndpoint,
			decodeReadDonaturBukuByKd_DonaturBukuRequest,
			encodeReadDonaturBukuByKd_DonaturBukuResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadDonaturBukuByKd_DonaturBuku", logger)))...),
		readDonaturBuku: grpctransport.NewServer(endpoints.ReadDonaturBukuEndpoint,
			decodeReadDonaturBukuRequest,
			encodeReadDonaturBukuResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadDonaturBuku", logger)))...),
		updateDonaturBuku: grpctransport.NewServer(endpoints.UpdateDonaturBukuEndpoint,
			decodeUpdateDonaturBukuRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateDonaturBuku", logger)))...),
		readDonaturBukuByNama_DonaturBuku: grpctransport.NewServer(endpoints.ReadDonaturBukuByNama_DonaturBukuEndpoint,
			decodeReadDonaturBukuByNama_DonaturBukuRequest,
			encodeReadDonaturBukuByNama_DonaturBukuResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadDonaturBukuByNama_DonaturBuku", logger)))...),
	}
}

func decodeAddDonaturBukuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddDonaturBukuReq)
	return scv.DonaturBuku{Name: req.GetName(), CustomerType: req.GetCustomerType(),
		Kode: req.GetMobile(), Nama: req.GetEmail(), Status: req.GetGender(),
		CallbackPhone: req.GetCallbackPhone(), Status: req.GetStatus()}, nil
}

func decodeReadDonaturBukuByKd_DonaturBukuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadDonaturBukuByKd_DonaturBukuReq)
	return scv.DonaturBuku{Kode: req.Kode}, nil
}

func decodeReadDonaturBukuRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeUpdateDonaturBukuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateDonaturBukuReq)
	return scv.DonaturBuku{Kd_DonaturBuku: req.Kd_DonaturBuku, Nama_DonaturBuku: req.Nama_DonaturBuku, Status: req.Status}, nil
}

func decodeReadCustomerByEmailRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadCustomerByEmailReq)
	return scv.Customer{Email: req.Email}, nil

}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadDonaturBukuByKd_DonaturBukuResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.DonaturBuku)
	return &pb.ReadDonaturBukuByKd_DonaturBukuResp{Kd_DonaturBuku: resp.Kd_DonaturBuku, Nama_DonaturBuku: resp.Nama_DonaturBuku, Status: resp.Status}, nil
}

func encodeReadDonaturBukuResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.DonaturBuku)

	rsp := &pb.ReadDonaturBukuResp{}

	for _, v := range resp {
		itm := &pb.ReadDonaturBukuByKd_DonaturBukuResp{
			Kd_DonaturBuku:   v.Kd_DonaturBuku,
			Nama_DonaturBuku: v.Nama_DonaturBuku,
			Status:           v.Status,
		}
		rsp.AllDonaturBuku = append(rsp.AllDonaturBuku, itm)
	}
	return rsp, nil
}

func encodeReadDonaturBukuByNama_DonaturBukuResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.DonaturBuku)
	return &pb.ReadDonaturBukuByNama_DonaturBukuResp{Kd_DonaturBuku: resp.Kd_DonaturBuku, Nama_DonaturBuku: resp.Nama_DonaturBuku, Status: resp.Status}, nil
}

func (s *grpcDonaturBukuServer) AddDonaturBuku(ctx oldcontext.Context, donaturBk *pb.AddDonaturBukuReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addDonaturBuku.ServeGRPC(ctx, donaturBk)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcDonaturBukuServer) ReadDonaturBukuByKd_DonaturBuku(ctx oldcontext.Context, kode *pb.ReadDonaturBukuByKd_DonaturBukuReq) (*pb.ReadDonaturBukuByKd_DonaturBukuResp, error) {
	_, resp, err := s.readDonaturBukuByKd_DonaturBuku.ServeGRPC(ctx, kode)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadDonaturBukuByKd_DonaturBukuResp), nil
}

func (s *grpcDonaturBukuServer) ReadDonaturBuku(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadDonaturBukuResp, error) {
	_, resp, err := s.readDonaturBuku.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadDonaturBukuResp), nil
}

func (s *grpcDonaturBukuServer) UpdateDonaturBuku(ctx oldcontext.Context, dbk *pb.UpdateDonaturBukuReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateDonaturBuku.ServeGRPC(ctx, dbk)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcDonaturBukuServer) ReadDonaturBukuByNama_DonaturBuku(ctx oldcontext.Context, nama *pb.ReadDonaturBukuByNama_DonaturBukuReq) (*pb.ReadDonaturBukuByNama_DonaturBukuResp, error) {
	_, resp, err := s.readDonaturBukuByNama_DonaturBuku.ServeGRPC(ctx, nama)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadDonaturBukuByNama_DonaturBukuResp), nil
}
