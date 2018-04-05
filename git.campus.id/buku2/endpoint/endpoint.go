package endpoint

import (
	"context"

	svc "MiniProject/git.campus.id/buku2/server"
	kit "github.com/go-kit/kit/endpoint"
)

type BukuEndpoint struct {
	AddBukuEndpoint              kit.Endpoint
	ReadBukuByKodeEndpoint       kit.Endpoint
	ReadBukuEndpoint             kit.Endpoint
	UpdateBukuEndpoint           kit.Endpoint
	ReadBukuByKeteranganEndpoint kit.Endpoint
}

func NewBukuEndpoint(service svc.BukuService) BukuEndpoint {
	addBukuEp := makeAddBukuEndpoint(service)
	readBukuByKodeEp := makeReadBukuByKodeEndpoint(service)
	readBukuEp := makeReadBukuEndpoint(service)
	updateBukuEp := makeUpdateBukuEndpoint(service)
	readBukuByKeteranganEp := makeReadBukuByKeteranganEndpoint(service)
	return BukuEndpoint{AddBukuEndpoint: addBukuEp,
		ReadBukuByKodeEndpoint:       readBukuByKodeEp,
		ReadBukuEndpoint:             readBukuEp,
		UpdateBukuEndpoint:           updateBukuEp,
		ReadBukuByKeteranganEndpoint: readBukuByKeteranganEp,
	}
}

func makeAddBukuEndpoint(service svc.BukuService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Buku)
		err := service.AddBukuService(ctx, req)
		return nil, err
	}
}

func makeReadBukuByKodeEndpoint(service svc.BukuService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Buku)
		result, err := service.ReadBukuByKodeService(ctx, req.KodeBuku)
		/*return svc.Customer{CustomerId: result.CustomerId, Name: result.Name,
		CustomerType: result.CustomerType, Mobile: result.Mobile, Email: result.Email,
		Gender: result.Gender, CallbackPhone: result.CallbackPhone, Status: result.Status}, err*/
		return result, err
	}
}

func makeReadBukuEndpoint(service svc.BukuService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadBukuService(ctx)
		return result, err
	}
}

func makeUpdateBukuEndpoint(service svc.BukuService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Buku)
		err := service.UpdateBukuService(ctx, req)
		return nil, err
	}
}

func makeReadBukuByKeteranganEndpoint(service svc.BukuService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Buku)
		result, err := service.ReadBukuByKeteranganService(ctx, req.Keterangan)
		return result, err
	}
}
