package server

import "context"

type Status int32

const (
	//ServiceID is dispatch service ID
	ServiceID        = "Campus.Buku.id"
	OnAdd     Status = 1
)

type Buku struct {
	KodeBuku   string
	NamaBuku   string
	Keterangan string
	Status     int32
	CreateBy   string
	UpdateBy   string
}
type Bukus []Buku

/*type Location struct {
	customerID   int64
	label        []int32
	locationType []int32
	name         []string
	street       string
	village      string
	district     string
	city         string
	province     string
	latitude     float64
	longitude    float64
}*/

type ReadWriter interface {
	AddBuku(Buku) error
	ReadBukuByKode(string) (Buku, error)
	ReadBuku() (Bukus, error)
	UpdateBuku(Buku) error
	ReadBukuByKeterangan(string) (Bukus, error)
}

type BukuService interface {
	AddBukuService(context.Context, Buku) error
	ReadBukuByKodeService(context.Context, string) (Buku, error)
	ReadBukuService(context.Context) (Bukus, error)
	UpdateBukuService(context.Context, Buku) error
	ReadBukuByKeteranganService(context.Context, string) (Bukus, error)
}
