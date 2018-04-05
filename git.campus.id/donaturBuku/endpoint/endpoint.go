package endpoint

import (
	"context"

	svc "MiniProject/git.campus.id/donaturBuku/server"
	kit "github.com/go-kit/kit/endpoint"
)

type DonaturBukuEndpoint struct {
	AddDonaturBukuEndpoint                    kit.Endpoint
	ReadDonaturBukuByKd_DonaturBukuEndpoint   kit.Endpoint
	ReadDonaturBukuEndpoint                   kit.Endpoint
	UpdateDonaturBukuEndpoint                 kit.Endpoint
	ReadDonaturBukuByNama_DonaturBukuEndpoint kit.Endpoint
}

func NewDonaturBukuEndpoint(service svc.CustomerService) CustomerEndpoint {
	addDonaturBukuEp := makeAddDonaturBukuEndpoint(service)
	readDonaturBukuByKd_DonaturBukuEp := makeReadDonaturBukuByKd_DonaturBukuEndpoint(service)
	readDonaturBukuEp := makeReadDonaturBukuEndpoint(service)
	updateDonaturBukuEp := makeUpdateDonaturBukuEndpoint(service)
	readDonaturBukuByNama_DonaturBukulEp := makeReadDonaturBukuByNama_DonaturBukuEndpoint(service)
	return DonaturBukuEndpoint{AddDonaturBukuEndpoint: addDonaturBukuEp,
		ReadDonaturBukuByKd_DonaturBukuEndpoint:   readDonaturBukuByKd_DonaturBukuEp,
		ReadDonaturBukuEndpoint:                   readDonaturBukuEp,
		UpdateDonaturBukuEndpoint:                 updateDonaturBukuEp,
		ReadDonaturBukuByNama_DonaturBukuEndpoint: readDonaturBukuByNama_DonaturBukuEp,
	}
}

func makeAddDonaturBukuEndpoint(service svc.DonaturBukuService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.DonaturBuku)
		err := service.AddDonaturBukuService(ctx, req)
		return nil, err
	}
}

func makeReadDonaturBukuByKd_DonaturBukuEndpoint(service svc.DonaturBukuService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.DonaturBuku)
		result, err := service.ReadDonaturBukuByKd_DonaturBukuService(ctx, req.Kode)
		/*return svc.Customer{CustomerId: result.CustomerId, Name: result.Name,
		CustomerType: result.CustomerType, Mobile: result.Mobile, Email: result.Email,
		Gender: result.Gender, CallbackPhone: result.CallbackPhone, Status: result.Status}, err*/
		return result, err
	}
}

func makeReadDonaturBukuEndpoint(service svc.DonaturBukuService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadDonaturBukuService(ctx)
		return result, err
	}
}

func makeUpdateDonaturBukuEndpoint(service svc.DonaturBukuService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.DonaturBuku)
		err := service.UpdateDonaturBukuService(ctx, req)
		return nil, err
	}
}

func makeReadDonaturBukuByNama_DonaturBukuEndpoint(service svc.DonaturBukuService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.DonaturBuku)
		result, err := service.ReadDonaturBukuByNama_DonaturBukuService(ctx, req.Nama)
		return result, err
	}
}
