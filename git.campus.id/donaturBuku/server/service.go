package server

import "context"

type Status int32

const (
	//serviceID adalah dispatch dari service ID
	ServiceID        = "donaturBuku.kampus.id"
	OnAdd     Status = 1
)

type DonaturBuku struct {
	Kd_DonaturBuku   string
	Nama_DonaturBuku string
	Status           int32
}
type DonaturBk []DonaturBuku

type ReadWriter interface {
	AddDonaturBk(DonaturBuku) error
	ReadDonaturBkByKd_DonaturBuku(string) (DonaturBuku, error)
	ReadMahasiswa() (DonaturBuku, error)
	UpdateDonaturBk(DonaturBuku) error
	ReadDonaturBkByNama_DonaturBuku(string) (DonaturBuku, error)
}

type DonaturBkService interface {
	AddDonaturBkService(context.Context, DonaturBuku) error
	ReadDOnaturBkByKd_DonaturBukuService(context.Context, string) (DonaturBuku, error)
	ReadDonaturBkService(context.Context, DonaturBuku) error
	ReadDonaturBkByNama_DonaturBukuService(context.Context, string) (DonaturBuku, error)
}
