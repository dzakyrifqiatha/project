package server

import (
	"context"
)

type buku struct {
	writer ReadWriter
}

func NewBuku(writer ReadWriter) BukuService {
	return &buku{writer: writer}
}

//Methode pada interface MahasiswaService di service.go
func (c *buku) AddBukuService(ctx context.Context, buku Buku) error {
	//fmt.Println("mahasiswa")
	err := c.writer.AddBuku(buku)
	if err != nil {
		return err
	}

	return nil
}

func (c *buku) ReadBukuByKodeService(ctx context.Context, mob string) (Buku, error) {
	mhs, err := c.writer.ReadBukuByKode(mob)
	//fmt.Println(mhs)
	if err != nil {
		return mhs, err
	}
	return mhs, nil
}

func (c *buku) ReadBukuService(ctx context.Context) (Bukus, error) {
	mhs, err := c.writer.ReadBuku()
	//fmt.Println("mahasiswa", mhs)
	if err != nil {
		return mhs, err
	}
	return mhs, nil
}

func (c *buku) UpdateBukuService(ctx context.Context, mhs Buku) error {
	err := c.writer.UpdateBuku(mhs)
	if err != nil {
		return err
	}
	return nil
}

func (c *buku) ReadBukuByKeteranganService(ctx context.Context, keterangan string) (Bukus, error) {
	mhs, err := c.writer.ReadBukuByKeterangan(keterangan)
	//fmt.Println("mahasiswa:", mhs)
	if err != nil {
		return mhs, err
	}
	return mhs, nil
}
